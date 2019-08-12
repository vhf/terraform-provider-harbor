// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sandhose/terraform-provider-harbor/api/models"
)

// GetProjectsProjectIDWebhookLasttriggerReader is a Reader for the GetProjectsProjectIDWebhookLasttrigger structure.
type GetProjectsProjectIDWebhookLasttriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDWebhookLasttriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProjectsProjectIDWebhookLasttriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetProjectsProjectIDWebhookLasttriggerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetProjectsProjectIDWebhookLasttriggerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetProjectsProjectIDWebhookLasttriggerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetProjectsProjectIDWebhookLasttriggerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProjectsProjectIDWebhookLasttriggerOK creates a GetProjectsProjectIDWebhookLasttriggerOK with default headers values
func NewGetProjectsProjectIDWebhookLasttriggerOK() *GetProjectsProjectIDWebhookLasttriggerOK {
	return &GetProjectsProjectIDWebhookLasttriggerOK{}
}

/*GetProjectsProjectIDWebhookLasttriggerOK handles this case with default header values.

Test webhook connection successfully.
*/
type GetProjectsProjectIDWebhookLasttriggerOK struct {
	Payload []*models.WebhookLastTrigger
}

func (o *GetProjectsProjectIDWebhookLasttriggerOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/lasttrigger][%d] getProjectsProjectIdWebhookLasttriggerOK  %+v", 200, o.Payload)
}

func (o *GetProjectsProjectIDWebhookLasttriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDWebhookLasttriggerBadRequest creates a GetProjectsProjectIDWebhookLasttriggerBadRequest with default headers values
func NewGetProjectsProjectIDWebhookLasttriggerBadRequest() *GetProjectsProjectIDWebhookLasttriggerBadRequest {
	return &GetProjectsProjectIDWebhookLasttriggerBadRequest{}
}

/*GetProjectsProjectIDWebhookLasttriggerBadRequest handles this case with default header values.

Illegal format of provided ID value.
*/
type GetProjectsProjectIDWebhookLasttriggerBadRequest struct {
}

func (o *GetProjectsProjectIDWebhookLasttriggerBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/lasttrigger][%d] getProjectsProjectIdWebhookLasttriggerBadRequest ", 400)
}

func (o *GetProjectsProjectIDWebhookLasttriggerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDWebhookLasttriggerUnauthorized creates a GetProjectsProjectIDWebhookLasttriggerUnauthorized with default headers values
func NewGetProjectsProjectIDWebhookLasttriggerUnauthorized() *GetProjectsProjectIDWebhookLasttriggerUnauthorized {
	return &GetProjectsProjectIDWebhookLasttriggerUnauthorized{}
}

/*GetProjectsProjectIDWebhookLasttriggerUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetProjectsProjectIDWebhookLasttriggerUnauthorized struct {
}

func (o *GetProjectsProjectIDWebhookLasttriggerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/lasttrigger][%d] getProjectsProjectIdWebhookLasttriggerUnauthorized ", 401)
}

func (o *GetProjectsProjectIDWebhookLasttriggerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDWebhookLasttriggerForbidden creates a GetProjectsProjectIDWebhookLasttriggerForbidden with default headers values
func NewGetProjectsProjectIDWebhookLasttriggerForbidden() *GetProjectsProjectIDWebhookLasttriggerForbidden {
	return &GetProjectsProjectIDWebhookLasttriggerForbidden{}
}

/*GetProjectsProjectIDWebhookLasttriggerForbidden handles this case with default header values.

User have no permission to get webhook policy of the project.
*/
type GetProjectsProjectIDWebhookLasttriggerForbidden struct {
}

func (o *GetProjectsProjectIDWebhookLasttriggerForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/lasttrigger][%d] getProjectsProjectIdWebhookLasttriggerForbidden ", 403)
}

func (o *GetProjectsProjectIDWebhookLasttriggerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDWebhookLasttriggerInternalServerError creates a GetProjectsProjectIDWebhookLasttriggerInternalServerError with default headers values
func NewGetProjectsProjectIDWebhookLasttriggerInternalServerError() *GetProjectsProjectIDWebhookLasttriggerInternalServerError {
	return &GetProjectsProjectIDWebhookLasttriggerInternalServerError{}
}

/*GetProjectsProjectIDWebhookLasttriggerInternalServerError handles this case with default header values.

Internal server errors.
*/
type GetProjectsProjectIDWebhookLasttriggerInternalServerError struct {
}

func (o *GetProjectsProjectIDWebhookLasttriggerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/lasttrigger][%d] getProjectsProjectIdWebhookLasttriggerInternalServerError ", 500)
}

func (o *GetProjectsProjectIDWebhookLasttriggerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
