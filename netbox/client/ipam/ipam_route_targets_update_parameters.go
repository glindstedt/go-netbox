// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
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
//

package ipam

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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewIpamRouteTargetsUpdateParams creates a new IpamRouteTargetsUpdateParams object
// with the default values initialized.
func NewIpamRouteTargetsUpdateParams() *IpamRouteTargetsUpdateParams {
	var ()
	return &IpamRouteTargetsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRouteTargetsUpdateParamsWithTimeout creates a new IpamRouteTargetsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamRouteTargetsUpdateParamsWithTimeout(timeout time.Duration) *IpamRouteTargetsUpdateParams {
	var ()
	return &IpamRouteTargetsUpdateParams{

		timeout: timeout,
	}
}

// NewIpamRouteTargetsUpdateParamsWithContext creates a new IpamRouteTargetsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamRouteTargetsUpdateParamsWithContext(ctx context.Context) *IpamRouteTargetsUpdateParams {
	var ()
	return &IpamRouteTargetsUpdateParams{

		Context: ctx,
	}
}

// NewIpamRouteTargetsUpdateParamsWithHTTPClient creates a new IpamRouteTargetsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamRouteTargetsUpdateParamsWithHTTPClient(client *http.Client) *IpamRouteTargetsUpdateParams {
	var ()
	return &IpamRouteTargetsUpdateParams{
		HTTPClient: client,
	}
}

/*IpamRouteTargetsUpdateParams contains all the parameters to send to the API endpoint
for the ipam route targets update operation typically these are written to a http.Request
*/
type IpamRouteTargetsUpdateParams struct {

	/*Data*/
	Data *models.WritableRouteTarget
	/*ID
	  A unique integer value identifying this route target.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) WithTimeout(timeout time.Duration) *IpamRouteTargetsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) WithContext(ctx context.Context) *IpamRouteTargetsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) WithHTTPClient(client *http.Client) *IpamRouteTargetsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) WithData(data *models.WritableRouteTarget) *IpamRouteTargetsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) SetData(data *models.WritableRouteTarget) {
	o.Data = data
}

// WithID adds the id to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) WithID(id int64) *IpamRouteTargetsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam route targets update params
func (o *IpamRouteTargetsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRouteTargetsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
