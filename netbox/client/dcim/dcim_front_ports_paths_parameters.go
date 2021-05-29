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
	"github.com/go-openapi/swag"
)

// NewDcimFrontPortsPathsParams creates a new DcimFrontPortsPathsParams object
// with the default values initialized.
func NewDcimFrontPortsPathsParams() *DcimFrontPortsPathsParams {
	var ()
	return &DcimFrontPortsPathsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortsPathsParamsWithTimeout creates a new DcimFrontPortsPathsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimFrontPortsPathsParamsWithTimeout(timeout time.Duration) *DcimFrontPortsPathsParams {
	var ()
	return &DcimFrontPortsPathsParams{

		timeout: timeout,
	}
}

// NewDcimFrontPortsPathsParamsWithContext creates a new DcimFrontPortsPathsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimFrontPortsPathsParamsWithContext(ctx context.Context) *DcimFrontPortsPathsParams {
	var ()
	return &DcimFrontPortsPathsParams{

		Context: ctx,
	}
}

// NewDcimFrontPortsPathsParamsWithHTTPClient creates a new DcimFrontPortsPathsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimFrontPortsPathsParamsWithHTTPClient(client *http.Client) *DcimFrontPortsPathsParams {
	var ()
	return &DcimFrontPortsPathsParams{
		HTTPClient: client,
	}
}

/*DcimFrontPortsPathsParams contains all the parameters to send to the API endpoint
for the dcim front ports paths operation typically these are written to a http.Request
*/
type DcimFrontPortsPathsParams struct {

	/*ID
	  A unique integer value identifying this front port.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) WithTimeout(timeout time.Duration) *DcimFrontPortsPathsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) WithContext(ctx context.Context) *DcimFrontPortsPathsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) WithHTTPClient(client *http.Client) *DcimFrontPortsPathsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) WithID(id int64) *DcimFrontPortsPathsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortsPathsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}