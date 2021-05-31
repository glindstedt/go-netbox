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

package extras

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

// NewExtrasTagsBulkUpdateParams creates a new ExtrasTagsBulkUpdateParams object
// with the default values initialized.
func NewExtrasTagsBulkUpdateParams() *ExtrasTagsBulkUpdateParams {
	var ()
	return &ExtrasTagsBulkUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTagsBulkUpdateParamsWithTimeout creates a new ExtrasTagsBulkUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasTagsBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasTagsBulkUpdateParams {
	var ()
	return &ExtrasTagsBulkUpdateParams{

		timeout: timeout,
	}
}

// NewExtrasTagsBulkUpdateParamsWithContext creates a new ExtrasTagsBulkUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasTagsBulkUpdateParamsWithContext(ctx context.Context) *ExtrasTagsBulkUpdateParams {
	var ()
	return &ExtrasTagsBulkUpdateParams{

		Context: ctx,
	}
}

// NewExtrasTagsBulkUpdateParamsWithHTTPClient creates a new ExtrasTagsBulkUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasTagsBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasTagsBulkUpdateParams {
	var ()
	return &ExtrasTagsBulkUpdateParams{
		HTTPClient: client,
	}
}

/*ExtrasTagsBulkUpdateParams contains all the parameters to send to the API endpoint
for the extras tags bulk update operation typically these are written to a http.Request
*/
type ExtrasTagsBulkUpdateParams struct {

	/*Data*/
	Data *models.Tag

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasTagsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) WithContext(ctx context.Context) *ExtrasTagsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasTagsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) WithData(data *models.Tag) *ExtrasTagsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) SetData(data *models.Tag) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTagsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
