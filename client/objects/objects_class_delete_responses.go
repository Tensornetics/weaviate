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

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsClassDeleteReader is a Reader for the ObjectsClassDelete structure.
type ObjectsClassDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ObjectsClassDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewObjectsClassDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewObjectsClassDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewObjectsClassDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewObjectsClassDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewObjectsClassDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewObjectsClassDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewObjectsClassDeleteNoContent creates a ObjectsClassDeleteNoContent with default headers values
func NewObjectsClassDeleteNoContent() *ObjectsClassDeleteNoContent {
	return &ObjectsClassDeleteNoContent{}
}

/*
ObjectsClassDeleteNoContent handles this case with default header values.

Successfully deleted.
*/
type ObjectsClassDeleteNoContent struct {
}

func (o *ObjectsClassDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /objects/{className}/{id}][%d] objectsClassDeleteNoContent ", 204)
}

func (o *ObjectsClassDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsClassDeleteBadRequest creates a ObjectsClassDeleteBadRequest with default headers values
func NewObjectsClassDeleteBadRequest() *ObjectsClassDeleteBadRequest {
	return &ObjectsClassDeleteBadRequest{}
}

/*
ObjectsClassDeleteBadRequest handles this case with default header values.

Malformed request.
*/
type ObjectsClassDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsClassDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /objects/{className}/{id}][%d] objectsClassDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ObjectsClassDeleteBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsClassDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsClassDeleteUnauthorized creates a ObjectsClassDeleteUnauthorized with default headers values
func NewObjectsClassDeleteUnauthorized() *ObjectsClassDeleteUnauthorized {
	return &ObjectsClassDeleteUnauthorized{}
}

/*
ObjectsClassDeleteUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ObjectsClassDeleteUnauthorized struct {
}

func (o *ObjectsClassDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /objects/{className}/{id}][%d] objectsClassDeleteUnauthorized ", 401)
}

func (o *ObjectsClassDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsClassDeleteForbidden creates a ObjectsClassDeleteForbidden with default headers values
func NewObjectsClassDeleteForbidden() *ObjectsClassDeleteForbidden {
	return &ObjectsClassDeleteForbidden{}
}

/*
ObjectsClassDeleteForbidden handles this case with default header values.

Forbidden
*/
type ObjectsClassDeleteForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsClassDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /objects/{className}/{id}][%d] objectsClassDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsClassDeleteForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsClassDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsClassDeleteNotFound creates a ObjectsClassDeleteNotFound with default headers values
func NewObjectsClassDeleteNotFound() *ObjectsClassDeleteNotFound {
	return &ObjectsClassDeleteNotFound{}
}

/*
ObjectsClassDeleteNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type ObjectsClassDeleteNotFound struct {
}

func (o *ObjectsClassDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /objects/{className}/{id}][%d] objectsClassDeleteNotFound ", 404)
}

func (o *ObjectsClassDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsClassDeleteInternalServerError creates a ObjectsClassDeleteInternalServerError with default headers values
func NewObjectsClassDeleteInternalServerError() *ObjectsClassDeleteInternalServerError {
	return &ObjectsClassDeleteInternalServerError{}
}

/*
ObjectsClassDeleteInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ObjectsClassDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsClassDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /objects/{className}/{id}][%d] objectsClassDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsClassDeleteInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsClassDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
