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

// DcimVirtualChassisBulkUpdateReader is a Reader for the DcimVirtualChassisBulkUpdate structure.
type DcimVirtualChassisBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimVirtualChassisBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimVirtualChassisBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimVirtualChassisBulkUpdateOK creates a DcimVirtualChassisBulkUpdateOK with default headers values
func NewDcimVirtualChassisBulkUpdateOK() *DcimVirtualChassisBulkUpdateOK {
	return &DcimVirtualChassisBulkUpdateOK{}
}

/*DcimVirtualChassisBulkUpdateOK handles this case with default header values.

DcimVirtualChassisBulkUpdateOK dcim virtual chassis bulk update o k
*/
type DcimVirtualChassisBulkUpdateOK struct {
	Payload *models.VirtualChassis
}

func (o *DcimVirtualChassisBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/virtual-chassis/][%d] dcimVirtualChassisBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimVirtualChassisBulkUpdateOK) GetPayload() *models.VirtualChassis {
	return o.Payload
}

func (o *DcimVirtualChassisBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimVirtualChassisBulkUpdateDefault creates a DcimVirtualChassisBulkUpdateDefault with default headers values
func NewDcimVirtualChassisBulkUpdateDefault(code int) *DcimVirtualChassisBulkUpdateDefault {
	return &DcimVirtualChassisBulkUpdateDefault{
		_statusCode: code,
	}
}

/*DcimVirtualChassisBulkUpdateDefault handles this case with default header values.

DcimVirtualChassisBulkUpdateDefault dcim virtual chassis bulk update default
*/
type DcimVirtualChassisBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim virtual chassis bulk update default response
func (o *DcimVirtualChassisBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimVirtualChassisBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/virtual-chassis/][%d] dcim_virtual-chassis_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualChassisBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimVirtualChassisBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
