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

// DcimDeviceBaysBulkUpdateReader is a Reader for the DcimDeviceBaysBulkUpdate structure.
type DcimDeviceBaysBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBaysBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBaysBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceBaysBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceBaysBulkUpdateOK creates a DcimDeviceBaysBulkUpdateOK with default headers values
func NewDcimDeviceBaysBulkUpdateOK() *DcimDeviceBaysBulkUpdateOK {
	return &DcimDeviceBaysBulkUpdateOK{}
}

/* DcimDeviceBaysBulkUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBaysBulkUpdateOK dcim device bays bulk update o k
*/
type DcimDeviceBaysBulkUpdateOK struct {
	Payload *models.DeviceBay
}

func (o *DcimDeviceBaysBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-bays/][%d] dcimDeviceBaysBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceBaysBulkUpdateOK) GetPayload() *models.DeviceBay {
	return o.Payload
}

func (o *DcimDeviceBaysBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceBaysBulkUpdateDefault creates a DcimDeviceBaysBulkUpdateDefault with default headers values
func NewDcimDeviceBaysBulkUpdateDefault(code int) *DcimDeviceBaysBulkUpdateDefault {
	return &DcimDeviceBaysBulkUpdateDefault{
		_statusCode: code,
	}
}

/* DcimDeviceBaysBulkUpdateDefault describes a response with status code -1, with default header values.

DcimDeviceBaysBulkUpdateDefault dcim device bays bulk update default
*/
type DcimDeviceBaysBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device bays bulk update default response
func (o *DcimDeviceBaysBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimDeviceBaysBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-bays/][%d] dcim_device-bays_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimDeviceBaysBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceBaysBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
