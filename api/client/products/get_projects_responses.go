// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sandhose/terraform-provider-harbor/api/models"
)

// GetProjectsReader is a Reader for the GetProjects structure.
type GetProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetProjectsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetProjectsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProjectsOK creates a GetProjectsOK with default headers values
func NewGetProjectsOK() *GetProjectsOK {
	return &GetProjectsOK{}
}

/*GetProjectsOK handles this case with default header values.

Return all matched projects.
*/
type GetProjectsOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of projects
	 */
	XTotalCount int64

	Payload []*models.Project
}

func (o *GetProjectsOK) Error() string {
	return fmt.Sprintf("[GET /projects][%d] getProjectsOK  %+v", 200, o.Payload)
}

func (o *GetProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsUnauthorized creates a GetProjectsUnauthorized with default headers values
func NewGetProjectsUnauthorized() *GetProjectsUnauthorized {
	return &GetProjectsUnauthorized{}
}

/*GetProjectsUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetProjectsUnauthorized struct {
}

func (o *GetProjectsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects][%d] getProjectsUnauthorized ", 401)
}

func (o *GetProjectsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsInternalServerError creates a GetProjectsInternalServerError with default headers values
func NewGetProjectsInternalServerError() *GetProjectsInternalServerError {
	return &GetProjectsInternalServerError{}
}

/*GetProjectsInternalServerError handles this case with default header values.

Internal errors.
*/
type GetProjectsInternalServerError struct {
}

func (o *GetProjectsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects][%d] getProjectsInternalServerError ", 500)
}

func (o *GetProjectsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
