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

// DcimFrontPortTemplatesListReader is a Reader for the DcimFrontPortTemplatesList structure.
type DcimFrontPortTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimFrontPortTemplatesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortTemplatesListOK creates a DcimFrontPortTemplatesListOK with default headers values
func NewDcimFrontPortTemplatesListOK() *DcimFrontPortTemplatesListOK {
	return &DcimFrontPortTemplatesListOK{}
}

/*DcimFrontPortTemplatesListOK handles this case with default header values.

DcimFrontPortTemplatesListOK dcim front port templates list o k
*/
type DcimFrontPortTemplatesListOK struct {
	Payload *DcimFrontPortTemplatesListOKBody
}

func (o *DcimFrontPortTemplatesListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/front-port-templates/][%d] dcimFrontPortTemplatesListOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortTemplatesListOK) GetPayload() *DcimFrontPortTemplatesListOKBody {
	return o.Payload
}

func (o *DcimFrontPortTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimFrontPortTemplatesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortTemplatesListDefault creates a DcimFrontPortTemplatesListDefault with default headers values
func NewDcimFrontPortTemplatesListDefault(code int) *DcimFrontPortTemplatesListDefault {
	return &DcimFrontPortTemplatesListDefault{
		_statusCode: code,
	}
}

/*DcimFrontPortTemplatesListDefault handles this case with default header values.

DcimFrontPortTemplatesListDefault dcim front port templates list default
*/
type DcimFrontPortTemplatesListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim front port templates list default response
func (o *DcimFrontPortTemplatesListDefault) Code() int {
	return o._statusCode
}

func (o *DcimFrontPortTemplatesListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/front-port-templates/][%d] dcim_front-port-templates_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortTemplatesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortTemplatesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DcimFrontPortTemplatesListOKBody dcim front port templates list o k body
swagger:model DcimFrontPortTemplatesListOKBody
*/
type DcimFrontPortTemplatesListOKBody struct {

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
	Results []*models.FrontPortTemplate `json:"results"`
}

// Validate validates this dcim front port templates list o k body
func (o *DcimFrontPortTemplatesListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DcimFrontPortTemplatesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimFrontPortTemplatesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimFrontPortTemplatesListOKBody) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimFrontPortTemplatesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimFrontPortTemplatesListOKBody) validatePrevious(formats strfmt.Registry) error {

	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimFrontPortTemplatesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimFrontPortTemplatesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimFrontPortTemplatesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimFrontPortTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimFrontPortTemplatesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimFrontPortTemplatesListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimFrontPortTemplatesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
