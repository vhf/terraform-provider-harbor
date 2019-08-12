// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteChartrepoRepoChartsName deletes all the versions of the specified chart

Delete all the versions of the specified chart
*/
func (a *Client) DeleteChartrepoRepoChartsName(params *DeleteChartrepoRepoChartsNameParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteChartrepoRepoChartsNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteChartrepoRepoChartsNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteChartrepoRepoChartsName",
		Method:             "DELETE",
		PathPattern:        "/chartrepo/{repo}/charts/{name}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteChartrepoRepoChartsNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteChartrepoRepoChartsNameOK), nil

}

/*
DeleteChartrepoRepoChartsNameVersion deletes the specified chart version

Delete the specified chart version
*/
func (a *Client) DeleteChartrepoRepoChartsNameVersion(params *DeleteChartrepoRepoChartsNameVersionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteChartrepoRepoChartsNameVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteChartrepoRepoChartsNameVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteChartrepoRepoChartsNameVersion",
		Method:             "DELETE",
		PathPattern:        "/chartrepo/{repo}/charts/{name}/{version}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteChartrepoRepoChartsNameVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteChartrepoRepoChartsNameVersionOK), nil

}

/*
DeleteChartrepoRepoChartsNameVersionLabelsID removes label from chart

Remove label from the specified chart version.
*/
func (a *Client) DeleteChartrepoRepoChartsNameVersionLabelsID(params *DeleteChartrepoRepoChartsNameVersionLabelsIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteChartrepoRepoChartsNameVersionLabelsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteChartrepoRepoChartsNameVersionLabelsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteChartrepoRepoChartsNameVersionLabelsID",
		Method:             "DELETE",
		PathPattern:        "/chartrepo/{repo}/charts/{name}/{version}/labels/{id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteChartrepoRepoChartsNameVersionLabelsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteChartrepoRepoChartsNameVersionLabelsIDOK), nil

}

/*
DeleteProjectsProjectIDRobotsRobotID deletes the specified robot account

Delete the specified robot account
*/
func (a *Client) DeleteProjectsProjectIDRobotsRobotID(params *DeleteProjectsProjectIDRobotsRobotIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteProjectsProjectIDRobotsRobotIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProjectsProjectIDRobotsRobotIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteProjectsProjectIDRobotsRobotID",
		Method:             "DELETE",
		PathPattern:        "/projects/{project_id}/robots/{robot_id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteProjectsProjectIDRobotsRobotIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteProjectsProjectIDRobotsRobotIDOK), nil

}

/*
GetChartrepoHealth checks the health of chart repository service

Check the health of chart repository service.
*/
func (a *Client) GetChartrepoHealth(params *GetChartrepoHealthParams, authInfo runtime.ClientAuthInfoWriter) (*GetChartrepoHealthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChartrepoHealthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetChartrepoHealth",
		Method:             "GET",
		PathPattern:        "/chartrepo/health",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetChartrepoHealthReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetChartrepoHealthOK), nil

}

/*
GetChartrepoRepoCharts gets all the charts under the specified project

Get all the charts under the specified project
*/
func (a *Client) GetChartrepoRepoCharts(params *GetChartrepoRepoChartsParams, authInfo runtime.ClientAuthInfoWriter) (*GetChartrepoRepoChartsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChartrepoRepoChartsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetChartrepoRepoCharts",
		Method:             "GET",
		PathPattern:        "/chartrepo/{repo}/charts",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetChartrepoRepoChartsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetChartrepoRepoChartsOK), nil

}

/*
GetChartrepoRepoChartsName gets all the versions of the specified chart

Get all the versions of the specified chart
*/
func (a *Client) GetChartrepoRepoChartsName(params *GetChartrepoRepoChartsNameParams, authInfo runtime.ClientAuthInfoWriter) (*GetChartrepoRepoChartsNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChartrepoRepoChartsNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetChartrepoRepoChartsName",
		Method:             "GET",
		PathPattern:        "/chartrepo/{repo}/charts/{name}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetChartrepoRepoChartsNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetChartrepoRepoChartsNameOK), nil

}

/*
GetChartrepoRepoChartsNameVersion gets the specified chart version

Get the specified chart version
*/
func (a *Client) GetChartrepoRepoChartsNameVersion(params *GetChartrepoRepoChartsNameVersionParams, authInfo runtime.ClientAuthInfoWriter) (*GetChartrepoRepoChartsNameVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChartrepoRepoChartsNameVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetChartrepoRepoChartsNameVersion",
		Method:             "GET",
		PathPattern:        "/chartrepo/{repo}/charts/{name}/{version}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetChartrepoRepoChartsNameVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetChartrepoRepoChartsNameVersionOK), nil

}

/*
GetChartrepoRepoChartsNameVersionLabels returns the attahced labels of chart

Return the attahced labels of the specified chart version.
*/
func (a *Client) GetChartrepoRepoChartsNameVersionLabels(params *GetChartrepoRepoChartsNameVersionLabelsParams, authInfo runtime.ClientAuthInfoWriter) (*GetChartrepoRepoChartsNameVersionLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChartrepoRepoChartsNameVersionLabelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetChartrepoRepoChartsNameVersionLabels",
		Method:             "GET",
		PathPattern:        "/chartrepo/{repo}/charts/{name}/{version}/labels",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetChartrepoRepoChartsNameVersionLabelsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetChartrepoRepoChartsNameVersionLabelsOK), nil

}

/*
GetProjectsProjectIDRobots gets all robot accounts of specified project

Get all robot accounts of specified project
*/
func (a *Client) GetProjectsProjectIDRobots(params *GetProjectsProjectIDRobotsParams, authInfo runtime.ClientAuthInfoWriter) (*GetProjectsProjectIDRobotsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectsProjectIDRobotsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetProjectsProjectIDRobots",
		Method:             "GET",
		PathPattern:        "/projects/{project_id}/robots",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProjectsProjectIDRobotsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProjectsProjectIDRobotsOK), nil

}

/*
GetProjectsProjectIDRobotsRobotID returns the infor of the specified robot account

Return the infor of the specified robot account.
*/
func (a *Client) GetProjectsProjectIDRobotsRobotID(params *GetProjectsProjectIDRobotsRobotIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetProjectsProjectIDRobotsRobotIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectsProjectIDRobotsRobotIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetProjectsProjectIDRobotsRobotID",
		Method:             "GET",
		PathPattern:        "/projects/{project_id}/robots/{robot_id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProjectsProjectIDRobotsRobotIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProjectsProjectIDRobotsRobotIDOK), nil

}

/*
GetSystemCVEWhitelist gets the system level whitelist of c v e

Get the system level whitelist of CVE.  This API can be called by all authenticated users.
*/
func (a *Client) GetSystemCVEWhitelist(params *GetSystemCVEWhitelistParams, authInfo runtime.ClientAuthInfoWriter) (*GetSystemCVEWhitelistOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemCVEWhitelistParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSystemCVEWhitelist",
		Method:             "GET",
		PathPattern:        "/system/CVEWhitelist",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSystemCVEWhitelistReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSystemCVEWhitelistOK), nil

}

/*
PostChartrepoCharts uploads a chart file to the defult library project

Upload a chart file to the default 'library' project. Uploading together with the prov file at the same time is also supported.
*/
func (a *Client) PostChartrepoCharts(params *PostChartrepoChartsParams, authInfo runtime.ClientAuthInfoWriter) (*PostChartrepoChartsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostChartrepoChartsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostChartrepoCharts",
		Method:             "POST",
		PathPattern:        "/chartrepo/charts",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostChartrepoChartsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostChartrepoChartsCreated), nil

}

/*
PostChartrepoRepoCharts uploads a chart file to the specified project

Upload a chart file to the specified project. With this API, the corresponding provance file can be uploaded together with chart file at once.
*/
func (a *Client) PostChartrepoRepoCharts(params *PostChartrepoRepoChartsParams, authInfo runtime.ClientAuthInfoWriter) (*PostChartrepoRepoChartsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostChartrepoRepoChartsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostChartrepoRepoCharts",
		Method:             "POST",
		PathPattern:        "/chartrepo/{repo}/charts",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostChartrepoRepoChartsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostChartrepoRepoChartsCreated), nil

}

/*
PostChartrepoRepoChartsNameVersionLabels marks label to chart

Mark label to the specified chart version.
*/
func (a *Client) PostChartrepoRepoChartsNameVersionLabels(params *PostChartrepoRepoChartsNameVersionLabelsParams, authInfo runtime.ClientAuthInfoWriter) (*PostChartrepoRepoChartsNameVersionLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostChartrepoRepoChartsNameVersionLabelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostChartrepoRepoChartsNameVersionLabels",
		Method:             "POST",
		PathPattern:        "/chartrepo/{repo}/charts/{name}/{version}/labels",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostChartrepoRepoChartsNameVersionLabelsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostChartrepoRepoChartsNameVersionLabelsOK), nil

}

/*
PostChartrepoRepoProv uploads a provance file to the specified project

Upload a provance file to the specified project. The provance file should be targeted for an existing chart file.
*/
func (a *Client) PostChartrepoRepoProv(params *PostChartrepoRepoProvParams, authInfo runtime.ClientAuthInfoWriter) (*PostChartrepoRepoProvCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostChartrepoRepoProvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostChartrepoRepoProv",
		Method:             "POST",
		PathPattern:        "/chartrepo/{repo}/prov",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostChartrepoRepoProvReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostChartrepoRepoProvCreated), nil

}

/*
PostProjectsProjectIDRobots creates a robot account for project

Create a robot account for project
*/
func (a *Client) PostProjectsProjectIDRobots(params *PostProjectsProjectIDRobotsParams, authInfo runtime.ClientAuthInfoWriter) (*PostProjectsProjectIDRobotsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostProjectsProjectIDRobotsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostProjectsProjectIDRobots",
		Method:             "POST",
		PathPattern:        "/projects/{project_id}/robots",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostProjectsProjectIDRobotsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostProjectsProjectIDRobotsCreated), nil

}

/*
PostSystemOidcPing tests the o ID c endpoint

Test the OIDC endpoint, the setting of the endpoint is provided in the request.  This API can only be called by system admin.
*/
func (a *Client) PostSystemOidcPing(params *PostSystemOidcPingParams, authInfo runtime.ClientAuthInfoWriter) (*PostSystemOidcPingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSystemOidcPingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostSystemOidcPing",
		Method:             "POST",
		PathPattern:        "/system/oidc/ping",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostSystemOidcPingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSystemOidcPingOK), nil

}

/*
PutProjectsProjectIDRobotsRobotID updates status of robot account

Used to disable/enable a specified robot account.
*/
func (a *Client) PutProjectsProjectIDRobotsRobotID(params *PutProjectsProjectIDRobotsRobotIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutProjectsProjectIDRobotsRobotIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutProjectsProjectIDRobotsRobotIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutProjectsProjectIDRobotsRobotID",
		Method:             "PUT",
		PathPattern:        "/projects/{project_id}/robots/{robot_id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutProjectsProjectIDRobotsRobotIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutProjectsProjectIDRobotsRobotIDOK), nil

}

/*
PutSystemCVEWhitelist updates the system level whitelist of c v e

This API overwrites the system level whitelist of CVE with the list in request body.  Only system Admin has permission to call this API.
*/
func (a *Client) PutSystemCVEWhitelist(params *PutSystemCVEWhitelistParams, authInfo runtime.ClientAuthInfoWriter) (*PutSystemCVEWhitelistOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutSystemCVEWhitelistParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutSystemCVEWhitelist",
		Method:             "PUT",
		PathPattern:        "/system/CVEWhitelist",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutSystemCVEWhitelistReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutSystemCVEWhitelistOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
