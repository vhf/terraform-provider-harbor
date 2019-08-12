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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sandhose/terraform-provider-harbor/api/models"
)

// NewPostLdapUsersImportParams creates a new PostLdapUsersImportParams object
// with the default values initialized.
func NewPostLdapUsersImportParams() *PostLdapUsersImportParams {
	var ()
	return &PostLdapUsersImportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLdapUsersImportParamsWithTimeout creates a new PostLdapUsersImportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLdapUsersImportParamsWithTimeout(timeout time.Duration) *PostLdapUsersImportParams {
	var ()
	return &PostLdapUsersImportParams{

		timeout: timeout,
	}
}

// NewPostLdapUsersImportParamsWithContext creates a new PostLdapUsersImportParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLdapUsersImportParamsWithContext(ctx context.Context) *PostLdapUsersImportParams {
	var ()
	return &PostLdapUsersImportParams{

		Context: ctx,
	}
}

// NewPostLdapUsersImportParamsWithHTTPClient creates a new PostLdapUsersImportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLdapUsersImportParamsWithHTTPClient(client *http.Client) *PostLdapUsersImportParams {
	var ()
	return &PostLdapUsersImportParams{
		HTTPClient: client,
	}
}

/*PostLdapUsersImportParams contains all the parameters to send to the API endpoint
for the post ldap users import operation typically these are written to a http.Request
*/
type PostLdapUsersImportParams struct {

	/*UIDList
	  The uid listed for importing. This list will check users validity of ldap service based on configuration from the system.

	*/
	UIDList *models.LdapImportUsers

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ldap users import params
func (o *PostLdapUsersImportParams) WithTimeout(timeout time.Duration) *PostLdapUsersImportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ldap users import params
func (o *PostLdapUsersImportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ldap users import params
func (o *PostLdapUsersImportParams) WithContext(ctx context.Context) *PostLdapUsersImportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ldap users import params
func (o *PostLdapUsersImportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ldap users import params
func (o *PostLdapUsersImportParams) WithHTTPClient(client *http.Client) *PostLdapUsersImportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ldap users import params
func (o *PostLdapUsersImportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUIDList adds the uIDList to the post ldap users import params
func (o *PostLdapUsersImportParams) WithUIDList(uIDList *models.LdapImportUsers) *PostLdapUsersImportParams {
	o.SetUIDList(uIDList)
	return o
}

// SetUIDList adds the uidList to the post ldap users import params
func (o *PostLdapUsersImportParams) SetUIDList(uIDList *models.LdapImportUsers) {
	o.UIDList = uIDList
}

// WriteToRequest writes these params to a swagger request
func (o *PostLdapUsersImportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UIDList != nil {
		if err := r.SetBodyParam(o.UIDList); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
