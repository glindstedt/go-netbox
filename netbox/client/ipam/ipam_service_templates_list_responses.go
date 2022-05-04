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
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpamServiceTemplatesListReader is a Reader for the IpamServiceTemplatesList structure.
type IpamServiceTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServiceTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamServiceTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamServiceTemplatesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamServiceTemplatesListOK creates a IpamServiceTemplatesListOK with default headers values
func NewIpamServiceTemplatesListOK() *IpamServiceTemplatesListOK {
	return &IpamServiceTemplatesListOK{}
}

/* IpamServiceTemplatesListOK describes a response with status code 200, with default header values.

IpamServiceTemplatesListOK ipam service templates list o k
*/
type IpamServiceTemplatesListOK struct {
	Payload *IpamServiceTemplatesListOKBody
}

func (o *IpamServiceTemplatesListOK) Error() string {
	return fmt.Sprintf("[GET /ipam/service-templates/][%d] ipamServiceTemplatesListOK  %+v", 200, o.Payload)
}
func (o *IpamServiceTemplatesListOK) GetPayload() *IpamServiceTemplatesListOKBody {
	return o.Payload
}

func (o *IpamServiceTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(IpamServiceTemplatesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamServiceTemplatesListDefault creates a IpamServiceTemplatesListDefault with default headers values
func NewIpamServiceTemplatesListDefault(code int) *IpamServiceTemplatesListDefault {
	return &IpamServiceTemplatesListDefault{
		_statusCode: code,
	}
}

/* IpamServiceTemplatesListDefault describes a response with status code -1, with default header values.

IpamServiceTemplatesListDefault ipam service templates list default
*/
type IpamServiceTemplatesListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam service templates list default response
func (o *IpamServiceTemplatesListDefault) Code() int {
	return o._statusCode
}

func (o *IpamServiceTemplatesListDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/service-templates/][%d] ipam_service-templates_list default  %+v", o._statusCode, o.Payload)
}
func (o *IpamServiceTemplatesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamServiceTemplatesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*IpamServiceTemplatesListOKBody ipam service templates list o k body
swagger:model IpamServiceTemplatesListOKBody
*/
type IpamServiceTemplatesListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.ServiceTemplate `json:"results"`
}

// Validate validates this ipam service templates list o k body
func (o *IpamServiceTemplatesListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IpamServiceTemplatesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("ipamServiceTemplatesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *IpamServiceTemplatesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamServiceTemplatesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IpamServiceTemplatesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamServiceTemplatesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IpamServiceTemplatesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("ipamServiceTemplatesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipamServiceTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ipam service templates list o k body based on the context it is used
func (o *IpamServiceTemplatesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IpamServiceTemplatesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipamServiceTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *IpamServiceTemplatesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IpamServiceTemplatesListOKBody) UnmarshalBinary(b []byte) error {
	var res IpamServiceTemplatesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
