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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// VirtualizationInterfacesBulkUpdateReader is a Reader for the VirtualizationInterfacesBulkUpdate structure.
type VirtualizationInterfacesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationInterfacesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationInterfacesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationInterfacesBulkUpdateOK creates a VirtualizationInterfacesBulkUpdateOK with default headers values
func NewVirtualizationInterfacesBulkUpdateOK() *VirtualizationInterfacesBulkUpdateOK {
	return &VirtualizationInterfacesBulkUpdateOK{}
}

/*VirtualizationInterfacesBulkUpdateOK handles this case with default header values.

VirtualizationInterfacesBulkUpdateOK virtualization interfaces bulk update o k
*/
type VirtualizationInterfacesBulkUpdateOK struct {
	Payload *models.VMInterface
}

func (o *VirtualizationInterfacesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/interfaces/][%d] virtualizationInterfacesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationInterfacesBulkUpdateOK) GetPayload() *models.VMInterface {
	return o.Payload
}

func (o *VirtualizationInterfacesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationInterfacesBulkUpdateDefault creates a VirtualizationInterfacesBulkUpdateDefault with default headers values
func NewVirtualizationInterfacesBulkUpdateDefault(code int) *VirtualizationInterfacesBulkUpdateDefault {
	return &VirtualizationInterfacesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*VirtualizationInterfacesBulkUpdateDefault handles this case with default header values.

VirtualizationInterfacesBulkUpdateDefault virtualization interfaces bulk update default
*/
type VirtualizationInterfacesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization interfaces bulk update default response
func (o *VirtualizationInterfacesBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationInterfacesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /virtualization/interfaces/][%d] virtualization_interfaces_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationInterfacesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationInterfacesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
