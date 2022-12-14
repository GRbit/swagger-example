// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddAccUpdater(params *AddAccUpdaterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddAccUpdaterOK, *AddAccUpdaterCreated, error)

	DelAccUpdater(params *DelAccUpdaterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DelAccUpdaterOK, error)

	GetAdsStats(params *GetAdsStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdsStatsOK, error)

	GetCampaignsStats(params *GetCampaignsStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCampaignsStatsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	AddAccUpdater Starts loading account ad and campaign statistic or raise account loading task in line

Response for new account is 201 CREATED
Response for raising in line is 200 OK
*/
func (a *Client) AddAccUpdater(params *AddAccUpdaterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddAccUpdaterOK, *AddAccUpdaterCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddAccUpdaterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addAccUpdater",
		Method:             "POST",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddAccUpdaterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *AddAccUpdaterOK:
		return value, nil, nil
	case *AddAccUpdaterCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for operations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	DelAccUpdater Delete account from loading queue and stop regularly updates.

Note: in --loader-acc-list-upd-timeout (default 600s) account list will be updated from cabinets service, if account presented in service, new update will be started.
*/
func (a *Client) DelAccUpdater(params *DelAccUpdaterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DelAccUpdaterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDelAccUpdaterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delAccUpdater",
		Method:             "DELETE",
		PathPattern:        "/{account_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DelAccUpdaterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DelAccUpdaterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delAccUpdater: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAdsStats Gives array of stats for set time period. If updatedSince set, then reutrns only stats updated after set time
*/
func (a *Client) GetAdsStats(params *GetAdsStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdsStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdsStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAdsStats",
		Method:             "GET",
		PathPattern:        "/{account_id}/ads",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAdsStatsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAdsStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAdsStats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCampaignsStats Gives array of stats for set time period. If updatedSince set, then reutrns only stats updated after set time
*/
func (a *Client) GetCampaignsStats(params *GetCampaignsStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCampaignsStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCampaignsStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCampaignsStats",
		Method:             "GET",
		PathPattern:        "/{account_id}/campaigns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCampaignsStatsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCampaignsStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCampaignsStats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
