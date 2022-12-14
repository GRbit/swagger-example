// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewAddAccUpdaterParams creates a new AddAccUpdaterParams object
//
// There are no default values defined in the spec.
func NewAddAccUpdaterParams() AddAccUpdaterParams {

	return AddAccUpdaterParams{}
}

// AddAccUpdaterParams contains all the bound params for the add acc updater operation
// typically these are obtained from a http.Request
//
// swagger:parameters addAccUpdater
type AddAccUpdaterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  Min Length: 30
	  In: header
	*/
	XRequestID string
	/*
	  Required: true
	  In: query
	*/
	AccountID int64
	/*
	  In: query
	*/
	ClientID *int64
	/*if set, account will be updated without cache from beginning of time
	  In: query
	*/
	Force *bool
	/*
	  Required: true
	  In: query
	*/
	UserID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddAccUpdaterParams() beforehand.
func (o *AddAccUpdaterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXRequestID(r.Header[http.CanonicalHeaderKey("X-RequestID")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qAccountID, qhkAccountID, _ := qs.GetOK("account_id")
	if err := o.bindAccountID(qAccountID, qhkAccountID, route.Formats); err != nil {
		res = append(res, err)
	}

	qClientID, qhkClientID, _ := qs.GetOK("client_id")
	if err := o.bindClientID(qClientID, qhkClientID, route.Formats); err != nil {
		res = append(res, err)
	}

	qForce, qhkForce, _ := qs.GetOK("force")
	if err := o.bindForce(qForce, qhkForce, route.Formats); err != nil {
		res = append(res, err)
	}

	qUserID, qhkUserID, _ := qs.GetOK("user_id")
	if err := o.bindUserID(qUserID, qhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXRequestID binds and validates parameter XRequestID from header.
func (o *AddAccUpdaterParams) bindXRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-RequestID", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-RequestID", "header", raw); err != nil {
		return err
	}
	o.XRequestID = raw

	if err := o.validateXRequestID(formats); err != nil {
		return err
	}

	return nil
}

// validateXRequestID carries on validations for parameter XRequestID
func (o *AddAccUpdaterParams) validateXRequestID(formats strfmt.Registry) error {

	if err := validate.MinLength("X-RequestID", "header", o.XRequestID, 30); err != nil {
		return err
	}

	return nil
}

// bindAccountID binds and validates parameter AccountID from query.
func (o *AddAccUpdaterParams) bindAccountID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("account_id", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("account_id", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("account_id", "query", "int64", raw)
	}
	o.AccountID = value

	return nil
}

// bindClientID binds and validates parameter ClientID from query.
func (o *AddAccUpdaterParams) bindClientID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("client_id", "query", "int64", raw)
	}
	o.ClientID = &value

	return nil
}

// bindForce binds and validates parameter Force from query.
func (o *AddAccUpdaterParams) bindForce(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("force", "query", "bool", raw)
	}
	o.Force = &value

	return nil
}

// bindUserID binds and validates parameter UserID from query.
func (o *AddAccUpdaterParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("user_id", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("user_id", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("user_id", "query", "int64", raw)
	}
	o.UserID = value

	return nil
}
