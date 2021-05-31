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

// DcimSitesReadReader is a Reader for the DcimSitesRead structure.
type DcimSitesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSitesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSitesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSitesReadOK creates a DcimSitesReadOK with default headers values
func NewDcimSitesReadOK() *DcimSitesReadOK {
	return &DcimSitesReadOK{}
}

/*DcimSitesReadOK handles this case with default header values.

DcimSitesReadOK dcim sites read o k
*/
type DcimSitesReadOK struct {
	Payload *models.Site
}

func (o *DcimSitesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/sites/{id}/][%d] dcimSitesReadOK  %+v", 200, o.Payload)
}

func (o *DcimSitesReadOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *DcimSitesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSitesReadDefault creates a DcimSitesReadDefault with default headers values
func NewDcimSitesReadDefault(code int) *DcimSitesReadDefault {
	return &DcimSitesReadDefault{
		_statusCode: code,
	}
}

/*DcimSitesReadDefault handles this case with default header values.

DcimSitesReadDefault dcim sites read default
*/
type DcimSitesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim sites read default response
func (o *DcimSitesReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimSitesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/sites/{id}/][%d] dcim_sites_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimSitesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSitesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
