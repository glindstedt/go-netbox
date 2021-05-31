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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewDcimDeviceRolesBulkPartialUpdateParams creates a new DcimDeviceRolesBulkPartialUpdateParams object
// with the default values initialized.
func NewDcimDeviceRolesBulkPartialUpdateParams() *DcimDeviceRolesBulkPartialUpdateParams {
	var ()
	return &DcimDeviceRolesBulkPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceRolesBulkPartialUpdateParamsWithTimeout creates a new DcimDeviceRolesBulkPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceRolesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimDeviceRolesBulkPartialUpdateParams {
	var ()
	return &DcimDeviceRolesBulkPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimDeviceRolesBulkPartialUpdateParamsWithContext creates a new DcimDeviceRolesBulkPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceRolesBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimDeviceRolesBulkPartialUpdateParams {
	var ()
	return &DcimDeviceRolesBulkPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimDeviceRolesBulkPartialUpdateParamsWithHTTPClient creates a new DcimDeviceRolesBulkPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceRolesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimDeviceRolesBulkPartialUpdateParams {
	var ()
	return &DcimDeviceRolesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimDeviceRolesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim device roles bulk partial update operation typically these are written to a http.Request
*/
type DcimDeviceRolesBulkPartialUpdateParams struct {

	/*Data*/
	Data *models.DeviceRole

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device roles bulk partial update params
func (o *DcimDeviceRolesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimDeviceRolesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device roles bulk partial update params
func (o *DcimDeviceRolesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device roles bulk partial update params
func (o *DcimDeviceRolesBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimDeviceRolesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device roles bulk partial update params
func (o *DcimDeviceRolesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device roles bulk partial update params
func (o *DcimDeviceRolesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimDeviceRolesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device roles bulk partial update params
func (o *DcimDeviceRolesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device roles bulk partial update params
func (o *DcimDeviceRolesBulkPartialUpdateParams) WithData(data *models.DeviceRole) *DcimDeviceRolesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device roles bulk partial update params
func (o *DcimDeviceRolesBulkPartialUpdateParams) SetData(data *models.DeviceRole) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceRolesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
