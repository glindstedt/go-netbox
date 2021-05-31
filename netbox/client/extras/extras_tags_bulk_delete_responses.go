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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasTagsBulkDeleteReader is a Reader for the ExtrasTagsBulkDelete structure.
type ExtrasTagsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasTagsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasTagsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasTagsBulkDeleteNoContent creates a ExtrasTagsBulkDeleteNoContent with default headers values
func NewExtrasTagsBulkDeleteNoContent() *ExtrasTagsBulkDeleteNoContent {
	return &ExtrasTagsBulkDeleteNoContent{}
}

/*ExtrasTagsBulkDeleteNoContent handles this case with default header values.

ExtrasTagsBulkDeleteNoContent extras tags bulk delete no content
*/
type ExtrasTagsBulkDeleteNoContent struct {
}

func (o *ExtrasTagsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/tags/][%d] extrasTagsBulkDeleteNoContent ", 204)
}

func (o *ExtrasTagsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasTagsBulkDeleteDefault creates a ExtrasTagsBulkDeleteDefault with default headers values
func NewExtrasTagsBulkDeleteDefault(code int) *ExtrasTagsBulkDeleteDefault {
	return &ExtrasTagsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*ExtrasTagsBulkDeleteDefault handles this case with default header values.

ExtrasTagsBulkDeleteDefault extras tags bulk delete default
*/
type ExtrasTagsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras tags bulk delete default response
func (o *ExtrasTagsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasTagsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /extras/tags/][%d] extras_tags_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasTagsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasTagsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
