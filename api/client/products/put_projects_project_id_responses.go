// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutProjectsProjectIDReader is a Reader for the PutProjectsProjectID structure.
type PutProjectsProjectIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutProjectsProjectIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutProjectsProjectIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutProjectsProjectIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPutProjectsProjectIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutProjectsProjectIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutProjectsProjectIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutProjectsProjectIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutProjectsProjectIDOK creates a PutProjectsProjectIDOK with default headers values
func NewPutProjectsProjectIDOK() *PutProjectsProjectIDOK {
	return &PutProjectsProjectIDOK{}
}

/*PutProjectsProjectIDOK handles this case with default header values.

Updated project properties successfully.
*/
type PutProjectsProjectIDOK struct {
}

func (o *PutProjectsProjectIDOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}][%d] putProjectsProjectIdOK ", 200)
}

func (o *PutProjectsProjectIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDBadRequest creates a PutProjectsProjectIDBadRequest with default headers values
func NewPutProjectsProjectIDBadRequest() *PutProjectsProjectIDBadRequest {
	return &PutProjectsProjectIDBadRequest{}
}

/*PutProjectsProjectIDBadRequest handles this case with default header values.

Illegal format of provided ID value.
*/
type PutProjectsProjectIDBadRequest struct {
}

func (o *PutProjectsProjectIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}][%d] putProjectsProjectIdBadRequest ", 400)
}

func (o *PutProjectsProjectIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDUnauthorized creates a PutProjectsProjectIDUnauthorized with default headers values
func NewPutProjectsProjectIDUnauthorized() *PutProjectsProjectIDUnauthorized {
	return &PutProjectsProjectIDUnauthorized{}
}

/*PutProjectsProjectIDUnauthorized handles this case with default header values.

User need to log in first.
*/
type PutProjectsProjectIDUnauthorized struct {
}

func (o *PutProjectsProjectIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}][%d] putProjectsProjectIdUnauthorized ", 401)
}

func (o *PutProjectsProjectIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDForbidden creates a PutProjectsProjectIDForbidden with default headers values
func NewPutProjectsProjectIDForbidden() *PutProjectsProjectIDForbidden {
	return &PutProjectsProjectIDForbidden{}
}

/*PutProjectsProjectIDForbidden handles this case with default header values.

User does not have permission to the project.
*/
type PutProjectsProjectIDForbidden struct {
}

func (o *PutProjectsProjectIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}][%d] putProjectsProjectIdForbidden ", 403)
}

func (o *PutProjectsProjectIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDNotFound creates a PutProjectsProjectIDNotFound with default headers values
func NewPutProjectsProjectIDNotFound() *PutProjectsProjectIDNotFound {
	return &PutProjectsProjectIDNotFound{}
}

/*PutProjectsProjectIDNotFound handles this case with default header values.

Project ID does not exist.
*/
type PutProjectsProjectIDNotFound struct {
}

func (o *PutProjectsProjectIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}][%d] putProjectsProjectIdNotFound ", 404)
}

func (o *PutProjectsProjectIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDInternalServerError creates a PutProjectsProjectIDInternalServerError with default headers values
func NewPutProjectsProjectIDInternalServerError() *PutProjectsProjectIDInternalServerError {
	return &PutProjectsProjectIDInternalServerError{}
}

/*PutProjectsProjectIDInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type PutProjectsProjectIDInternalServerError struct {
}

func (o *PutProjectsProjectIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}][%d] putProjectsProjectIdInternalServerError ", 500)
}

func (o *PutProjectsProjectIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
