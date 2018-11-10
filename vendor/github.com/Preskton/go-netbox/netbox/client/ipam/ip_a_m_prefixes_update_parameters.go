// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/Preskton/go-netbox/netbox/models"
)

// NewIPAMPrefixesUpdateParams creates a new IPAMPrefixesUpdateParams object
// with the default values initialized.
func NewIPAMPrefixesUpdateParams() *IPAMPrefixesUpdateParams {
	var ()
	return &IPAMPrefixesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMPrefixesUpdateParamsWithTimeout creates a new IPAMPrefixesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMPrefixesUpdateParamsWithTimeout(timeout time.Duration) *IPAMPrefixesUpdateParams {
	var ()
	return &IPAMPrefixesUpdateParams{

		timeout: timeout,
	}
}

// NewIPAMPrefixesUpdateParamsWithContext creates a new IPAMPrefixesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMPrefixesUpdateParamsWithContext(ctx context.Context) *IPAMPrefixesUpdateParams {
	var ()
	return &IPAMPrefixesUpdateParams{

		Context: ctx,
	}
}

// NewIPAMPrefixesUpdateParamsWithHTTPClient creates a new IPAMPrefixesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMPrefixesUpdateParamsWithHTTPClient(client *http.Client) *IPAMPrefixesUpdateParams {
	var ()
	return &IPAMPrefixesUpdateParams{
		HTTPClient: client,
	}
}

/*IPAMPrefixesUpdateParams contains all the parameters to send to the API endpoint
for the ipam prefixes update operation typically these are written to a http.Request
*/
type IPAMPrefixesUpdateParams struct {

	/*Data*/
	Data *models.PrefixCreateUpdate
	/*ID
	  A unique integer value identifying this prefix.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam prefixes update params
func (o *IPAMPrefixesUpdateParams) WithTimeout(timeout time.Duration) *IPAMPrefixesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes update params
func (o *IPAMPrefixesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes update params
func (o *IPAMPrefixesUpdateParams) WithContext(ctx context.Context) *IPAMPrefixesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes update params
func (o *IPAMPrefixesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes update params
func (o *IPAMPrefixesUpdateParams) WithHTTPClient(client *http.Client) *IPAMPrefixesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes update params
func (o *IPAMPrefixesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam prefixes update params
func (o *IPAMPrefixesUpdateParams) WithData(data *models.PrefixCreateUpdate) *IPAMPrefixesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam prefixes update params
func (o *IPAMPrefixesUpdateParams) SetData(data *models.PrefixCreateUpdate) {
	o.Data = data
}

// WithID adds the id to the ipam prefixes update params
func (o *IPAMPrefixesUpdateParams) WithID(id int64) *IPAMPrefixesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes update params
func (o *IPAMPrefixesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMPrefixesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
