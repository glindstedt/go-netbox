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

// IpamPrefixesBulkDeleteReader is a Reader for the IpamPrefixesBulkDelete structure.
type IpamPrefixesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamPrefixesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamPrefixesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamPrefixesBulkDeleteNoContent creates a IpamPrefixesBulkDeleteNoContent with default headers values
func NewIpamPrefixesBulkDeleteNoContent() *IpamPrefixesBulkDeleteNoContent {
	return &IpamPrefixesBulkDeleteNoContent{}
}

/*IpamPrefixesBulkDeleteNoContent handles this case with default header values.

IpamPrefixesBulkDeleteNoContent ipam prefixes bulk delete no content
*/
type IpamPrefixesBulkDeleteNoContent struct {
}

func (o *IpamPrefixesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/prefixes/][%d] ipamPrefixesBulkDeleteNoContent ", 204)
}

func (o *IpamPrefixesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamPrefixesBulkDeleteDefault creates a IpamPrefixesBulkDeleteDefault with default headers values
func NewIpamPrefixesBulkDeleteDefault(code int) *IpamPrefixesBulkDeleteDefault {
	return &IpamPrefixesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*IpamPrefixesBulkDeleteDefault handles this case with default header values.

IpamPrefixesBulkDeleteDefault ipam prefixes bulk delete default
*/
type IpamPrefixesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam prefixes bulk delete default response
func (o *IpamPrefixesBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpamPrefixesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/prefixes/][%d] ipam_prefixes_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamPrefixesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamPrefixesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
