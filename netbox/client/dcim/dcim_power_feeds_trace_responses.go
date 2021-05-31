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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimPowerFeedsTraceReader is a Reader for the DcimPowerFeedsTrace structure.
type DcimPowerFeedsTraceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsTraceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerFeedsTraceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerFeedsTraceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerFeedsTraceOK creates a DcimPowerFeedsTraceOK with default headers values
func NewDcimPowerFeedsTraceOK() *DcimPowerFeedsTraceOK {
	return &DcimPowerFeedsTraceOK{}
}

/*DcimPowerFeedsTraceOK handles this case with default header values.

DcimPowerFeedsTraceOK dcim power feeds trace o k
*/
type DcimPowerFeedsTraceOK struct {
	Payload *models.PowerFeed
}

func (o *DcimPowerFeedsTraceOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-feeds/{id}/trace/][%d] dcimPowerFeedsTraceOK  %+v", 200, o.Payload)
}

func (o *DcimPowerFeedsTraceOK) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsTraceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerFeedsTraceDefault creates a DcimPowerFeedsTraceDefault with default headers values
func NewDcimPowerFeedsTraceDefault(code int) *DcimPowerFeedsTraceDefault {
	return &DcimPowerFeedsTraceDefault{
		_statusCode: code,
	}
}

/*DcimPowerFeedsTraceDefault handles this case with default header values.

DcimPowerFeedsTraceDefault dcim power feeds trace default
*/
type DcimPowerFeedsTraceDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power feeds trace default response
func (o *DcimPowerFeedsTraceDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerFeedsTraceDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/power-feeds/{id}/trace/][%d] dcim_power-feeds_trace default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerFeedsTraceDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerFeedsTraceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
