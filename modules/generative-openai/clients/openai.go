//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/entities/moduletools"
	"github.com/weaviate/weaviate/modules/generative-openai/config"
	"github.com/weaviate/weaviate/modules/generative-openai/ent"
)

var compile, _ = regexp.Compile(`{([\w\s]*?)}`)

type openai struct {
	apiKey     string
	host       string
	path       string
	httpClient *http.Client
	logger     logrus.FieldLogger
}

func New(apiKey string, logger logrus.FieldLogger) *openai {
	return &openai{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
		host:   "https://api.openai.com",
		path:   "/v1/completions",
		logger: logger,
	}
}

func (v *openai) GenerateSingleResult(ctx context.Context, textProperties map[string]string, prompt string, cfg moduletools.ClassConfig) (*ent.GenerateResult, error) {
	forPrompt, err := v.generateForPrompt(textProperties, prompt)
	if err != nil {
		return nil, err
	}
	return v.Generate(ctx, cfg, forPrompt)
}

func (v *openai) GenerateAllResults(ctx context.Context, textProperties []map[string]string, task string, cfg moduletools.ClassConfig) (*ent.GenerateResult, error) {
	forTask, err := v.generatePromptForTask(textProperties, task)
	if err != nil {
		return nil, err
	}
	return v.Generate(ctx, cfg, forTask)
}

func (v *openai) Generate(ctx context.Context, cfg moduletools.ClassConfig, prompt string) (*ent.GenerateResult, error) {
	settings := config.NewClassSettings(cfg)

	body, err := json.Marshal(generateInput{
		Prompt:           prompt,
		Model:            settings.Model(),
		MaxTokens:        settings.MaxTokens(),
		Temperature:      settings.Temperature(),
		FrequencyPenalty: settings.FrequencyPenalty(),
		PresencePenalty:  settings.PresencePenalty(),
		TopP:             settings.TopP(),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "marshal body")
	}

	oaiUrl, err := url.JoinPath(v.host, v.path)
	if err != nil {
		return nil, errors.Wrap(err, "join OpenAI API host and path")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", oaiUrl,
		bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "create POST request")
	}
	apiKey, err := v.getApiKey(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "OpenAI API Key")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := v.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "send POST request")
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}

	var resBody generateResponse
	if err := json.Unmarshal(bodyBytes, &resBody); err != nil {
		return nil, errors.Wrap(err, "unmarshal response body")
	}

	if res.StatusCode > 399 {
		if resBody.Error != nil {
			return nil, errors.Errorf("failed with status: %d error: %v", res.StatusCode, resBody.Error.Message)
		}
		return nil, errors.Errorf("failed with status: %d", res.StatusCode)
	}
	textResponse := resBody.Choices[0].Text
	if len(resBody.Choices) > 0 && textResponse != "" {
		replaceAll := strings.ReplaceAll(textResponse, "\n", "")
		return &ent.GenerateResult{
			Result: &replaceAll,
		}, nil
	}
	return &ent.GenerateResult{
		Result: nil,
	}, nil
}

func (v *openai) generatePromptForTask(textProperties []map[string]string, task string) (string, error) {
	marshal, err := json.Marshal(textProperties)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`'%v:
%v`, task, string(marshal)), nil
}

func (v *openai) generateForPrompt(textProperties map[string]string, prompt string) (string, error) {
	all := compile.FindAll([]byte(prompt), -1)
	for _, match := range all {
		originalProperty := string(match)
		replacedProperty := compile.FindStringSubmatch(originalProperty)[1]
		replacedProperty = strings.TrimSpace(replacedProperty)
		value := textProperties[replacedProperty]
		if value == "" {
			return "", errors.Errorf("Following property has empty value: '%v'. Make sure you spell the property name correctly, verify that the property exists and has a value", replacedProperty)
		}
		prompt = strings.ReplaceAll(prompt, originalProperty, value)
	}
	return prompt, nil
}

func (v *openai) getApiKey(ctx context.Context) (string, error) {
	if len(v.apiKey) > 0 {
		return v.apiKey, nil
	}
	apiKey := ctx.Value("X-Openai-Api-Key")
	if apiKeyHeader, ok := apiKey.([]string); ok &&
		len(apiKeyHeader) > 0 && len(apiKeyHeader[0]) > 0 {
		return apiKeyHeader[0], nil
	}
	return "", errors.New("no api key found " +
		"neither in request header: X-OpenAI-Api-Key " +
		"nor in environment variable under OPENAI_APIKEY")
}

type generateInput struct {
	Prompt           string   `json:"prompt"`
	Model            string   `json:"model"`
	MaxTokens        float64  `json:"max_tokens"`
	Temperature      float64  `json:"temperature"`
	Stop             []string `json:"stop"`
	FrequencyPenalty float64  `json:"frequency_penalty"`
	PresencePenalty  float64  `json:"presence_penalty"`
	TopP             float64  `json:"top_p"`
}

type generateResponse struct {
	Choices []choice
	Error   *openAIApiError `json:"error,omitempty"`
}

type choice struct {
	FinishReason string
	Index        float32
	Logprobs     string
	Text         string
}

type openAIApiError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    string `json:"code"`
}
