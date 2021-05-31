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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamRolesBulkDeleteReader is a Reader for the IpamRolesBulkDelete structure.
type IpamRolesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamRolesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRolesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRolesBulkDeleteNoContent creates a IpamRolesBulkDeleteNoContent with default headers values
func NewIpamRolesBulkDeleteNoContent() *IpamRolesBulkDeleteNoContent {
	return &IpamRolesBulkDeleteNoContent{}
}

/*IpamRolesBulkDeleteNoContent handles this case with default header values.

IpamRolesBulkDeleteNoContent ipam roles bulk delete no content
*/
type IpamRolesBulkDeleteNoContent struct {
}

func (o *IpamRolesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/roles/][%d] ipamRolesBulkDeleteNoContent ", 204)
}

func (o *IpamRolesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamRolesBulkDeleteDefault creates a IpamRolesBulkDeleteDefault with default headers values
func NewIpamRolesBulkDeleteDefault(code int) *IpamRolesBulkDeleteDefault {
	return &IpamRolesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*IpamRolesBulkDeleteDefault handles this case with default header values.

IpamRolesBulkDeleteDefault ipam roles bulk delete default
*/
type IpamRolesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam roles bulk delete default response
func (o *IpamRolesBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpamRolesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/roles/][%d] ipam_roles_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRolesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRolesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
