// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewAddAccUpdaterParams creates a new AddAccUpdaterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddAccUpdaterParams() *AddAccUpdaterParams {
	return &AddAccUpdaterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddAccUpdaterParamsWithTimeout creates a new AddAccUpdaterParams object
// with the ability to set a timeout on a request.
func NewAddAccUpdaterParamsWithTimeout(timeout time.Duration) *AddAccUpdaterParams {
	return &AddAccUpdaterParams{
		timeout: timeout,
	}
}

// NewAddAccUpdaterParamsWithContext creates a new AddAccUpdaterParams object
// with the ability to set a context for a request.
func NewAddAccUpdaterParamsWithContext(ctx context.Context) *AddAccUpdaterParams {
	return &AddAccUpdaterParams{
		Context: ctx,
	}
}

// NewAddAccUpdaterParamsWithHTTPClient creates a new AddAccUpdaterParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddAccUpdaterParamsWithHTTPClient(client *http.Client) *AddAccUpdaterParams {
	return &AddAccUpdaterParams{
		HTTPClient: client,
	}
}

/*
AddAccUpdaterParams contains all the parameters to send to the API endpoint

	for the add acc updater operation.

	Typically these are written to a http.Request.
*/
type AddAccUpdaterParams struct {

	// XRequestID.
	XRequestID string

	// AccountID.
	//
	// Format: int64
	AccountID int64

	// ClientID.
	//
	// Format: int64
	ClientID *int64

	/* Force.

	   if set, account will be updated without cache from beginning of time
	*/
	Force *bool

	// UserID.
	//
	// Format: int64
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add acc updater params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAccUpdaterParams) WithDefaults() *AddAccUpdaterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add acc updater params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAccUpdaterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add acc updater params
func (o *AddAccUpdaterParams) WithTimeout(timeout time.Duration) *AddAccUpdaterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add acc updater params
func (o *AddAccUpdaterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add acc updater params
func (o *AddAccUpdaterParams) WithContext(ctx context.Context) *AddAccUpdaterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add acc updater params
func (o *AddAccUpdaterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add acc updater params
func (o *AddAccUpdaterParams) WithHTTPClient(client *http.Client) *AddAccUpdaterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add acc updater params
func (o *AddAccUpdaterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the add acc updater params
func (o *AddAccUpdaterParams) WithXRequestID(xRequestID string) *AddAccUpdaterParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the add acc updater params
func (o *AddAccUpdaterParams) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithAccountID adds the accountID to the add acc updater params
func (o *AddAccUpdaterParams) WithAccountID(accountID int64) *AddAccUpdaterParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the add acc updater params
func (o *AddAccUpdaterParams) SetAccountID(accountID int64) {
	o.AccountID = accountID
}

// WithClientID adds the clientID to the add acc updater params
func (o *AddAccUpdaterParams) WithClientID(clientID *int64) *AddAccUpdaterParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the add acc updater params
func (o *AddAccUpdaterParams) SetClientID(clientID *int64) {
	o.ClientID = clientID
}

// WithForce adds the force to the add acc updater params
func (o *AddAccUpdaterParams) WithForce(force *bool) *AddAccUpdaterParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the add acc updater params
func (o *AddAccUpdaterParams) SetForce(force *bool) {
	o.Force = force
}

// WithUserID adds the userID to the add acc updater params
func (o *AddAccUpdaterParams) WithUserID(userID int64) *AddAccUpdaterParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the add acc updater params
func (o *AddAccUpdaterParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AddAccUpdaterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-RequestID
	if err := r.SetHeaderParam("X-RequestID", o.XRequestID); err != nil {
		return err
	}

	// query param account_id
	qrAccountID := o.AccountID
	qAccountID := swag.FormatInt64(qrAccountID)
	if qAccountID != "" {

		if err := r.SetQueryParam("account_id", qAccountID); err != nil {
			return err
		}
	}

	if o.ClientID != nil {

		// query param client_id
		var qrClientID int64

		if o.ClientID != nil {
			qrClientID = *o.ClientID
		}
		qClientID := swag.FormatInt64(qrClientID)
		if qClientID != "" {

			if err := r.SetQueryParam("client_id", qClientID); err != nil {
				return err
			}
		}
	}

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	// query param user_id
	qrUserID := o.UserID
	qUserID := swag.FormatInt64(qrUserID)
	if qUserID != "" {

		if err := r.SetQueryParam("user_id", qUserID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
