// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// WeaviateThingsDeleteReader is a Reader for the WeaviateThingsDelete structure.
type WeaviateThingsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateThingsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewWeaviateThingsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateThingsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateThingsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWeaviateThingsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateThingsDeleteNoContent creates a WeaviateThingsDeleteNoContent with default headers values
func NewWeaviateThingsDeleteNoContent() *WeaviateThingsDeleteNoContent {
	return &WeaviateThingsDeleteNoContent{}
}

/*WeaviateThingsDeleteNoContent handles this case with default header values.

Successful deleted.
*/
type WeaviateThingsDeleteNoContent struct {
}

func (o *WeaviateThingsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /things/{thingId}][%d] weaviateThingsDeleteNoContent ", 204)
}

func (o *WeaviateThingsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsDeleteUnauthorized creates a WeaviateThingsDeleteUnauthorized with default headers values
func NewWeaviateThingsDeleteUnauthorized() *WeaviateThingsDeleteUnauthorized {
	return &WeaviateThingsDeleteUnauthorized{}
}

/*WeaviateThingsDeleteUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateThingsDeleteUnauthorized struct {
}

func (o *WeaviateThingsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /things/{thingId}][%d] weaviateThingsDeleteUnauthorized ", 401)
}

func (o *WeaviateThingsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsDeleteForbidden creates a WeaviateThingsDeleteForbidden with default headers values
func NewWeaviateThingsDeleteForbidden() *WeaviateThingsDeleteForbidden {
	return &WeaviateThingsDeleteForbidden{}
}

/*WeaviateThingsDeleteForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateThingsDeleteForbidden struct {
}

func (o *WeaviateThingsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /things/{thingId}][%d] weaviateThingsDeleteForbidden ", 403)
}

func (o *WeaviateThingsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsDeleteNotFound creates a WeaviateThingsDeleteNotFound with default headers values
func NewWeaviateThingsDeleteNotFound() *WeaviateThingsDeleteNotFound {
	return &WeaviateThingsDeleteNotFound{}
}

/*WeaviateThingsDeleteNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type WeaviateThingsDeleteNotFound struct {
}

func (o *WeaviateThingsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /things/{thingId}][%d] weaviateThingsDeleteNotFound ", 404)
}

func (o *WeaviateThingsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}