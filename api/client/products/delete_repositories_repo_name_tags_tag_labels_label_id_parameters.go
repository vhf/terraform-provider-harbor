// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams creates a new DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams object
// with the default values initialized.
func NewDeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams() *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams {
	var ()
	return &DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRepositoriesRepoNameTagsTagLabelsLabelIDParamsWithTimeout creates a new DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRepositoriesRepoNameTagsTagLabelsLabelIDParamsWithTimeout(timeout time.Duration) *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams {
	var ()
	return &DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams{

		timeout: timeout,
	}
}

// NewDeleteRepositoriesRepoNameTagsTagLabelsLabelIDParamsWithContext creates a new DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRepositoriesRepoNameTagsTagLabelsLabelIDParamsWithContext(ctx context.Context) *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams {
	var ()
	return &DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams{

		Context: ctx,
	}
}

// NewDeleteRepositoriesRepoNameTagsTagLabelsLabelIDParamsWithHTTPClient creates a new DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRepositoriesRepoNameTagsTagLabelsLabelIDParamsWithHTTPClient(client *http.Client) *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams {
	var ()
	return &DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams{
		HTTPClient: client,
	}
}

/*DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams contains all the parameters to send to the API endpoint
for the delete repositories repo name tags tag labels label ID operation typically these are written to a http.Request
*/
type DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams struct {

	/*LabelID
	  The ID of label.

	*/
	LabelID int64
	/*RepoName
	  The name of repository.

	*/
	RepoName string
	/*Tag
	  The tag of the image.

	*/
	Tag string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete repositories repo name tags tag labels label ID params
func (o *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams) WithTimeout(timeout time.Duration) *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete repositories repo name tags tag labels label ID params
func (o *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete repositories repo name tags tag labels label ID params
func (o *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams) WithContext(ctx context.Context) *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete repositories repo name tags tag labels label ID params
func (o *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete repositories repo name tags tag labels label ID params
func (o *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams) WithHTTPClient(client *http.Client) *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete repositories repo name tags tag labels label ID params
func (o *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabelID adds the labelID to the delete repositories repo name tags tag labels label ID params
func (o *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams) WithLabelID(labelID int64) *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams {
	o.SetLabelID(labelID)
	return o
}

// SetLabelID adds the labelId to the delete repositories repo name tags tag labels label ID params
func (o *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams) SetLabelID(labelID int64) {
	o.LabelID = labelID
}

// WithRepoName adds the repoName to the delete repositories repo name tags tag labels label ID params
func (o *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams) WithRepoName(repoName string) *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the delete repositories repo name tags tag labels label ID params
func (o *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WithTag adds the tag to the delete repositories repo name tags tag labels label ID params
func (o *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams) WithTag(tag string) *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the delete repositories repo name tags tag labels label ID params
func (o *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams) SetTag(tag string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRepositoriesRepoNameTagsTagLabelsLabelIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param label_id
	if err := r.SetPathParam("label_id", swag.FormatInt64(o.LabelID)); err != nil {
		return err
	}

	// path param repo_name
	if err := r.SetPathParam("repo_name", o.RepoName); err != nil {
		return err
	}

	// path param tag
	if err := r.SetPathParam("tag", o.Tag); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
