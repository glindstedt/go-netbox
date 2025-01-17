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

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimLocationsDeleteReader is a Reader for the DcimLocationsDelete structure.
type DcimLocationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimLocationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimLocationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimLocationsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimLocationsDeleteNoContent creates a DcimLocationsDeleteNoContent with default headers values
func NewDcimLocationsDeleteNoContent() *DcimLocationsDeleteNoContent {
	return &DcimLocationsDeleteNoContent{}
}

/* DcimLocationsDeleteNoContent describes a response with status code 204, with default header values.

DcimLocationsDeleteNoContent dcim locations delete no content
*/
type DcimLocationsDeleteNoContent struct {
}

func (o *DcimLocationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/locations/{id}/][%d] dcimLocationsDeleteNoContent ", 204)
}

func (o *DcimLocationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimLocationsDeleteDefault creates a DcimLocationsDeleteDefault with default headers values
func NewDcimLocationsDeleteDefault(code int) *DcimLocationsDeleteDefault {
	return &DcimLocationsDeleteDefault{
		_statusCode: code,
	}
}

/* DcimLocationsDeleteDefault describes a response with status code -1, with default header values.

DcimLocationsDeleteDefault dcim locations delete default
*/
type DcimLocationsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim locations delete default response
func (o *DcimLocationsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimLocationsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/locations/{id}/][%d] dcim_locations_delete default  %+v", o._statusCode, o.Payload)
}
func (o *DcimLocationsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimLocationsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
