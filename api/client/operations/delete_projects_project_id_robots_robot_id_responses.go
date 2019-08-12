// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteProjectsProjectIDRobotsRobotIDReader is a Reader for the DeleteProjectsProjectIDRobotsRobotID structure.
type DeleteProjectsProjectIDRobotsRobotIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectsProjectIDRobotsRobotIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteProjectsProjectIDRobotsRobotIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteProjectsProjectIDRobotsRobotIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteProjectsProjectIDRobotsRobotIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteProjectsProjectIDRobotsRobotIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteProjectsProjectIDRobotsRobotIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProjectsProjectIDRobotsRobotIDOK creates a DeleteProjectsProjectIDRobotsRobotIDOK with default headers values
func NewDeleteProjectsProjectIDRobotsRobotIDOK() *DeleteProjectsProjectIDRobotsRobotIDOK {
	return &DeleteProjectsProjectIDRobotsRobotIDOK{}
}

/*DeleteProjectsProjectIDRobotsRobotIDOK handles this case with default header values.

The specified robot account is successfully deleted.
*/
type DeleteProjectsProjectIDRobotsRobotIDOK struct {
}

func (o *DeleteProjectsProjectIDRobotsRobotIDOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}/robots/{robot_id}][%d] deleteProjectsProjectIdRobotsRobotIdOK ", 200)
}

func (o *DeleteProjectsProjectIDRobotsRobotIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsProjectIDRobotsRobotIDUnauthorized creates a DeleteProjectsProjectIDRobotsRobotIDUnauthorized with default headers values
func NewDeleteProjectsProjectIDRobotsRobotIDUnauthorized() *DeleteProjectsProjectIDRobotsRobotIDUnauthorized {
	return &DeleteProjectsProjectIDRobotsRobotIDUnauthorized{}
}

/*DeleteProjectsProjectIDRobotsRobotIDUnauthorized handles this case with default header values.

User need to log in first.
*/
type DeleteProjectsProjectIDRobotsRobotIDUnauthorized struct {
}

func (o *DeleteProjectsProjectIDRobotsRobotIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}/robots/{robot_id}][%d] deleteProjectsProjectIdRobotsRobotIdUnauthorized ", 401)
}

func (o *DeleteProjectsProjectIDRobotsRobotIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsProjectIDRobotsRobotIDForbidden creates a DeleteProjectsProjectIDRobotsRobotIDForbidden with default headers values
func NewDeleteProjectsProjectIDRobotsRobotIDForbidden() *DeleteProjectsProjectIDRobotsRobotIDForbidden {
	return &DeleteProjectsProjectIDRobotsRobotIDForbidden{}
}

/*DeleteProjectsProjectIDRobotsRobotIDForbidden handles this case with default header values.

User in session does not have permission to the project.
*/
type DeleteProjectsProjectIDRobotsRobotIDForbidden struct {
}

func (o *DeleteProjectsProjectIDRobotsRobotIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}/robots/{robot_id}][%d] deleteProjectsProjectIdRobotsRobotIdForbidden ", 403)
}

func (o *DeleteProjectsProjectIDRobotsRobotIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsProjectIDRobotsRobotIDNotFound creates a DeleteProjectsProjectIDRobotsRobotIDNotFound with default headers values
func NewDeleteProjectsProjectIDRobotsRobotIDNotFound() *DeleteProjectsProjectIDRobotsRobotIDNotFound {
	return &DeleteProjectsProjectIDRobotsRobotIDNotFound{}
}

/*DeleteProjectsProjectIDRobotsRobotIDNotFound handles this case with default header values.

The robot account is not found.
*/
type DeleteProjectsProjectIDRobotsRobotIDNotFound struct {
}

func (o *DeleteProjectsProjectIDRobotsRobotIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}/robots/{robot_id}][%d] deleteProjectsProjectIdRobotsRobotIdNotFound ", 404)
}

func (o *DeleteProjectsProjectIDRobotsRobotIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsProjectIDRobotsRobotIDInternalServerError creates a DeleteProjectsProjectIDRobotsRobotIDInternalServerError with default headers values
func NewDeleteProjectsProjectIDRobotsRobotIDInternalServerError() *DeleteProjectsProjectIDRobotsRobotIDInternalServerError {
	return &DeleteProjectsProjectIDRobotsRobotIDInternalServerError{}
}

/*DeleteProjectsProjectIDRobotsRobotIDInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type DeleteProjectsProjectIDRobotsRobotIDInternalServerError struct {
}

func (o *DeleteProjectsProjectIDRobotsRobotIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}/robots/{robot_id}][%d] deleteProjectsProjectIdRobotsRobotIdInternalServerError ", 500)
}

func (o *DeleteProjectsProjectIDRobotsRobotIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}