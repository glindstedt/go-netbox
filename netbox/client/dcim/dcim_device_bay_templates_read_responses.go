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

// DcimDeviceBayTemplatesReadReader is a Reader for the DcimDeviceBayTemplatesRead structure.
type DcimDeviceBayTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBayTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceBayTemplatesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceBayTemplatesReadOK creates a DcimDeviceBayTemplatesReadOK with default headers values
func NewDcimDeviceBayTemplatesReadOK() *DcimDeviceBayTemplatesReadOK {
	return &DcimDeviceBayTemplatesReadOK{}
}

/*DcimDeviceBayTemplatesReadOK handles this case with default header values.

DcimDeviceBayTemplatesReadOK dcim device bay templates read o k
*/
type DcimDeviceBayTemplatesReadOK struct {
	Payload *models.DeviceBayTemplate
}

func (o *DcimDeviceBayTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBayTemplatesReadOK) GetPayload() *models.DeviceBayTemplate {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceBayTemplatesReadDefault creates a DcimDeviceBayTemplatesReadDefault with default headers values
func NewDcimDeviceBayTemplatesReadDefault(code int) *DcimDeviceBayTemplatesReadDefault {
	return &DcimDeviceBayTemplatesReadDefault{
		_statusCode: code,
	}
}

/*DcimDeviceBayTemplatesReadDefault handles this case with default header values.

DcimDeviceBayTemplatesReadDefault dcim device bay templates read default
*/
type DcimDeviceBayTemplatesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device bay templates read default response
func (o *DcimDeviceBayTemplatesReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimDeviceBayTemplatesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/device-bay-templates/{id}/][%d] dcim_device-bay-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBayTemplatesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
