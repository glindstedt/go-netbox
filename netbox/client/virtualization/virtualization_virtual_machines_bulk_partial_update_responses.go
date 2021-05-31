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

// VirtualizationVirtualMachinesBulkPartialUpdateReader is a Reader for the VirtualizationVirtualMachinesBulkPartialUpdate structure.
type VirtualizationVirtualMachinesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualMachinesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationVirtualMachinesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationVirtualMachinesBulkPartialUpdateOK creates a VirtualizationVirtualMachinesBulkPartialUpdateOK with default headers values
func NewVirtualizationVirtualMachinesBulkPartialUpdateOK() *VirtualizationVirtualMachinesBulkPartialUpdateOK {
	return &VirtualizationVirtualMachinesBulkPartialUpdateOK{}
}

/*VirtualizationVirtualMachinesBulkPartialUpdateOK handles this case with default header values.

VirtualizationVirtualMachinesBulkPartialUpdateOK virtualization virtual machines bulk partial update o k
*/
type VirtualizationVirtualMachinesBulkPartialUpdateOK struct {
	Payload *models.VirtualMachineWithConfigContext
}

func (o *VirtualizationVirtualMachinesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/virtual-machines/][%d] virtualizationVirtualMachinesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationVirtualMachinesBulkPartialUpdateOK) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationVirtualMachinesBulkPartialUpdateDefault creates a VirtualizationVirtualMachinesBulkPartialUpdateDefault with default headers values
func NewVirtualizationVirtualMachinesBulkPartialUpdateDefault(code int) *VirtualizationVirtualMachinesBulkPartialUpdateDefault {
	return &VirtualizationVirtualMachinesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*VirtualizationVirtualMachinesBulkPartialUpdateDefault handles this case with default header values.

VirtualizationVirtualMachinesBulkPartialUpdateDefault virtualization virtual machines bulk partial update default
*/
type VirtualizationVirtualMachinesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization virtual machines bulk partial update default response
func (o *VirtualizationVirtualMachinesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationVirtualMachinesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/virtual-machines/][%d] virtualization_virtual-machines_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationVirtualMachinesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
