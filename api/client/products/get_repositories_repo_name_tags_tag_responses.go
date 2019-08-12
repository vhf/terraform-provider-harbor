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

// GetRepositoriesRepoNameTagsTagReader is a Reader for the GetRepositoriesRepoNameTagsTag structure.
type GetRepositoriesRepoNameTagsTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesRepoNameTagsTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRepositoriesRepoNameTagsTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetRepositoriesRepoNameTagsTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRepositoriesRepoNameTagsTagOK creates a GetRepositoriesRepoNameTagsTagOK with default headers values
func NewGetRepositoriesRepoNameTagsTagOK() *GetRepositoriesRepoNameTagsTagOK {
	return &GetRepositoriesRepoNameTagsTagOK{}
}

/*GetRepositoriesRepoNameTagsTagOK handles this case with default header values.

Get tag successfully.
*/
type GetRepositoriesRepoNameTagsTagOK struct {
	Payload *models.DetailedTag
}

func (o *GetRepositoriesRepoNameTagsTagOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{repo_name}/tags/{tag}][%d] getRepositoriesRepoNameTagsTagOK  %+v", 200, o.Payload)
}

func (o *GetRepositoriesRepoNameTagsTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DetailedTag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesRepoNameTagsTagInternalServerError creates a GetRepositoriesRepoNameTagsTagInternalServerError with default headers values
func NewGetRepositoriesRepoNameTagsTagInternalServerError() *GetRepositoriesRepoNameTagsTagInternalServerError {
	return &GetRepositoriesRepoNameTagsTagInternalServerError{}
}

/*GetRepositoriesRepoNameTagsTagInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetRepositoriesRepoNameTagsTagInternalServerError struct {
}

func (o *GetRepositoriesRepoNameTagsTagInternalServerError) Error() string {
	return fmt.Sprintf("[GET /repositories/{repo_name}/tags/{tag}][%d] getRepositoriesRepoNameTagsTagInternalServerError ", 500)
}

func (o *GetRepositoriesRepoNameTagsTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
