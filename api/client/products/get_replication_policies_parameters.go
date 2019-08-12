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

// NewGetReplicationPoliciesParams creates a new GetReplicationPoliciesParams object
// with the default values initialized.
func NewGetReplicationPoliciesParams() *GetReplicationPoliciesParams {
	var ()
	return &GetReplicationPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReplicationPoliciesParamsWithTimeout creates a new GetReplicationPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReplicationPoliciesParamsWithTimeout(timeout time.Duration) *GetReplicationPoliciesParams {
	var ()
	return &GetReplicationPoliciesParams{

		timeout: timeout,
	}
}

// NewGetReplicationPoliciesParamsWithContext creates a new GetReplicationPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReplicationPoliciesParamsWithContext(ctx context.Context) *GetReplicationPoliciesParams {
	var ()
	return &GetReplicationPoliciesParams{

		Context: ctx,
	}
}

// NewGetReplicationPoliciesParamsWithHTTPClient creates a new GetReplicationPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReplicationPoliciesParamsWithHTTPClient(client *http.Client) *GetReplicationPoliciesParams {
	var ()
	return &GetReplicationPoliciesParams{
		HTTPClient: client,
	}
}

/*GetReplicationPoliciesParams contains all the parameters to send to the API endpoint
for the get replication policies operation typically these are written to a http.Request
*/
type GetReplicationPoliciesParams struct {

	/*Name
	  The replication policy name.

	*/
	Name *string
	/*Page
	  The page nubmer.

	*/
	Page *int32
	/*PageSize
	  The size of per page.

	*/
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get replication policies params
func (o *GetReplicationPoliciesParams) WithTimeout(timeout time.Duration) *GetReplicationPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get replication policies params
func (o *GetReplicationPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get replication policies params
func (o *GetReplicationPoliciesParams) WithContext(ctx context.Context) *GetReplicationPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get replication policies params
func (o *GetReplicationPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get replication policies params
func (o *GetReplicationPoliciesParams) WithHTTPClient(client *http.Client) *GetReplicationPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get replication policies params
func (o *GetReplicationPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get replication policies params
func (o *GetReplicationPoliciesParams) WithName(name *string) *GetReplicationPoliciesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get replication policies params
func (o *GetReplicationPoliciesParams) SetName(name *string) {
	o.Name = name
}

// WithPage adds the page to the get replication policies params
func (o *GetReplicationPoliciesParams) WithPage(page *int32) *GetReplicationPoliciesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get replication policies params
func (o *GetReplicationPoliciesParams) SetPage(page *int32) {
	o.Page = page
}

// WithPageSize adds the pageSize to the get replication policies params
func (o *GetReplicationPoliciesParams) WithPageSize(pageSize *int32) *GetReplicationPoliciesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get replication policies params
func (o *GetReplicationPoliciesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetReplicationPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
