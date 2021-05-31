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

// DcimRearPortsBulkPartialUpdateReader is a Reader for the DcimRearPortsBulkPartialUpdate structure.
type DcimRearPortsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRearPortsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortsBulkPartialUpdateOK creates a DcimRearPortsBulkPartialUpdateOK with default headers values
func NewDcimRearPortsBulkPartialUpdateOK() *DcimRearPortsBulkPartialUpdateOK {
	return &DcimRearPortsBulkPartialUpdateOK{}
}

/*DcimRearPortsBulkPartialUpdateOK handles this case with default header values.

DcimRearPortsBulkPartialUpdateOK dcim rear ports bulk partial update o k
*/
type DcimRearPortsBulkPartialUpdateOK struct {
	Payload *models.RearPort
}

func (o *DcimRearPortsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rear-ports/][%d] dcimRearPortsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortsBulkPartialUpdateOK) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRearPortsBulkPartialUpdateDefault creates a DcimRearPortsBulkPartialUpdateDefault with default headers values
func NewDcimRearPortsBulkPartialUpdateDefault(code int) *DcimRearPortsBulkPartialUpdateDefault {
	return &DcimRearPortsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*DcimRearPortsBulkPartialUpdateDefault handles this case with default header values.

DcimRearPortsBulkPartialUpdateDefault dcim rear ports bulk partial update default
*/
type DcimRearPortsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rear ports bulk partial update default response
func (o *DcimRearPortsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRearPortsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rear-ports/][%d] dcim_rear-ports_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
