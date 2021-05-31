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

package dcim

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
)

// NewDcimPowerOutletsBulkDeleteParams creates a new DcimPowerOutletsBulkDeleteParams object
// with the default values initialized.
func NewDcimPowerOutletsBulkDeleteParams() *DcimPowerOutletsBulkDeleteParams {

	return &DcimPowerOutletsBulkDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletsBulkDeleteParamsWithTimeout creates a new DcimPowerOutletsBulkDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerOutletsBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimPowerOutletsBulkDeleteParams {

	return &DcimPowerOutletsBulkDeleteParams{

		timeout: timeout,
	}
}

// NewDcimPowerOutletsBulkDeleteParamsWithContext creates a new DcimPowerOutletsBulkDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerOutletsBulkDeleteParamsWithContext(ctx context.Context) *DcimPowerOutletsBulkDeleteParams {

	return &DcimPowerOutletsBulkDeleteParams{

		Context: ctx,
	}
}

// NewDcimPowerOutletsBulkDeleteParamsWithHTTPClient creates a new DcimPowerOutletsBulkDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerOutletsBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimPowerOutletsBulkDeleteParams {

	return &DcimPowerOutletsBulkDeleteParams{
		HTTPClient: client,
	}
}

/*DcimPowerOutletsBulkDeleteParams contains all the parameters to send to the API endpoint
for the dcim power outlets bulk delete operation typically these are written to a http.Request
*/
type DcimPowerOutletsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power outlets bulk delete params
func (o *DcimPowerOutletsBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimPowerOutletsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlets bulk delete params
func (o *DcimPowerOutletsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlets bulk delete params
func (o *DcimPowerOutletsBulkDeleteParams) WithContext(ctx context.Context) *DcimPowerOutletsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlets bulk delete params
func (o *DcimPowerOutletsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlets bulk delete params
func (o *DcimPowerOutletsBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimPowerOutletsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlets bulk delete params
func (o *DcimPowerOutletsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
