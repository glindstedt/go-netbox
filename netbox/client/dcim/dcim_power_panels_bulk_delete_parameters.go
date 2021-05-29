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

// NewDcimPowerPanelsBulkDeleteParams creates a new DcimPowerPanelsBulkDeleteParams object
// with the default values initialized.
func NewDcimPowerPanelsBulkDeleteParams() *DcimPowerPanelsBulkDeleteParams {

	return &DcimPowerPanelsBulkDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPanelsBulkDeleteParamsWithTimeout creates a new DcimPowerPanelsBulkDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerPanelsBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimPowerPanelsBulkDeleteParams {

	return &DcimPowerPanelsBulkDeleteParams{

		timeout: timeout,
	}
}

// NewDcimPowerPanelsBulkDeleteParamsWithContext creates a new DcimPowerPanelsBulkDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerPanelsBulkDeleteParamsWithContext(ctx context.Context) *DcimPowerPanelsBulkDeleteParams {

	return &DcimPowerPanelsBulkDeleteParams{

		Context: ctx,
	}
}

// NewDcimPowerPanelsBulkDeleteParamsWithHTTPClient creates a new DcimPowerPanelsBulkDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerPanelsBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimPowerPanelsBulkDeleteParams {

	return &DcimPowerPanelsBulkDeleteParams{
		HTTPClient: client,
	}
}

/*DcimPowerPanelsBulkDeleteParams contains all the parameters to send to the API endpoint
for the dcim power panels bulk delete operation typically these are written to a http.Request
*/
type DcimPowerPanelsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power panels bulk delete params
func (o *DcimPowerPanelsBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimPowerPanelsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power panels bulk delete params
func (o *DcimPowerPanelsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power panels bulk delete params
func (o *DcimPowerPanelsBulkDeleteParams) WithContext(ctx context.Context) *DcimPowerPanelsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power panels bulk delete params
func (o *DcimPowerPanelsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power panels bulk delete params
func (o *DcimPowerPanelsBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimPowerPanelsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power panels bulk delete params
func (o *DcimPowerPanelsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPanelsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}