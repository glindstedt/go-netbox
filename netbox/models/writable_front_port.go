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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableFrontPort writable front port
//
// swagger:model WritableFrontPort
type WritableFrontPort struct {

	// occupied
	// Read Only: true
	Occupied *bool `json:"_occupied,omitempty"`

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Color
	// Max Length: 6
	// Pattern: ^[0-9a-f]{6}$
	Color string `json:"color,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Device
	// Required: true
	Device *int64 `json:"device"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Link peer
	//
	//
	// Return the appropriate serializer for the link termination model.
	//
	// Read Only: true
	LinkPeer map[string]*string `json:"link_peer,omitempty"`

	// Link peer type
	// Read Only: true
	LinkPeerType string `json:"link_peer_type,omitempty"`

	// Mark connected
	//
	// Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Rear port
	// Required: true
	RearPort *int64 `json:"rear_port"`

	// Rear port position
	// Maximum: 1024
	// Minimum: 1
	RearPortPosition int64 `json:"rear_port_position,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Type
	// Required: true
	// Enum: [8p8c 8p6c 8p4c 8p2c 6p6c 6p4c 6p2c 4p4c 4p2c gg45 tera-4p tera-2p tera-1p 110-punch bnc f n mrj21 fc lc lc-pc lc-upc lc-apc lsh lsh-pc lsh-upc lsh-apc mpo mtrj sc sc-pc sc-upc sc-apc st cs sn sma-905 sma-906 urm-p2 urm-p4 urm-p8 splice]
	Type *string `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable front port
func (m *WritableFrontPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearPortPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableFrontPort) validateCable(formats strfmt.Registry) error {
	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritableFrontPort) validateColor(formats strfmt.Registry) error {
	if swag.IsZero(m.Color) { // not required
		return nil
	}

	if err := validate.MaxLength("color", "body", m.Color, 6); err != nil {
		return err
	}

	if err := validate.Pattern("color", "body", m.Color, `^[0-9a-f]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateRearPort(formats strfmt.Registry) error {

	if err := validate.Required("rear_port", "body", m.RearPort); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateRearPortPosition(formats strfmt.Registry) error {
	if swag.IsZero(m.RearPortPosition) { // not required
		return nil
	}

	if err := validate.MinimumInt("rear_port_position", "body", m.RearPortPosition, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("rear_port_position", "body", m.RearPortPosition, 1024, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var writableFrontPortTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","8p6c","8p4c","8p2c","6p6c","6p4c","6p2c","4p4c","4p2c","gg45","tera-4p","tera-2p","tera-1p","110-punch","bnc","f","n","mrj21","fc","lc","lc-pc","lc-upc","lc-apc","lsh","lsh-pc","lsh-upc","lsh-apc","mpo","mtrj","sc","sc-pc","sc-upc","sc-apc","st","cs","sn","sma-905","sma-906","urm-p2","urm-p4","urm-p8","splice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableFrontPortTypeTypePropEnum = append(writableFrontPortTypeTypePropEnum, v)
	}
}

const (

	// WritableFrontPortTypeNr8p8c captures enum value "8p8c"
	WritableFrontPortTypeNr8p8c string = "8p8c"

	// WritableFrontPortTypeNr8p6c captures enum value "8p6c"
	WritableFrontPortTypeNr8p6c string = "8p6c"

	// WritableFrontPortTypeNr8p4c captures enum value "8p4c"
	WritableFrontPortTypeNr8p4c string = "8p4c"

	// WritableFrontPortTypeNr8p2c captures enum value "8p2c"
	WritableFrontPortTypeNr8p2c string = "8p2c"

	// WritableFrontPortTypeNr6p6c captures enum value "6p6c"
	WritableFrontPortTypeNr6p6c string = "6p6c"

	// WritableFrontPortTypeNr6p4c captures enum value "6p4c"
	WritableFrontPortTypeNr6p4c string = "6p4c"

	// WritableFrontPortTypeNr6p2c captures enum value "6p2c"
	WritableFrontPortTypeNr6p2c string = "6p2c"

	// WritableFrontPortTypeNr4p4c captures enum value "4p4c"
	WritableFrontPortTypeNr4p4c string = "4p4c"

	// WritableFrontPortTypeNr4p2c captures enum value "4p2c"
	WritableFrontPortTypeNr4p2c string = "4p2c"

	// WritableFrontPortTypeGg45 captures enum value "gg45"
	WritableFrontPortTypeGg45 string = "gg45"

	// WritableFrontPortTypeTeraDash4p captures enum value "tera-4p"
	WritableFrontPortTypeTeraDash4p string = "tera-4p"

	// WritableFrontPortTypeTeraDash2p captures enum value "tera-2p"
	WritableFrontPortTypeTeraDash2p string = "tera-2p"

	// WritableFrontPortTypeTeraDash1p captures enum value "tera-1p"
	WritableFrontPortTypeTeraDash1p string = "tera-1p"

	// WritableFrontPortTypeNr110DashPunch captures enum value "110-punch"
	WritableFrontPortTypeNr110DashPunch string = "110-punch"

	// WritableFrontPortTypeBnc captures enum value "bnc"
	WritableFrontPortTypeBnc string = "bnc"

	// WritableFrontPortTypeF captures enum value "f"
	WritableFrontPortTypeF string = "f"

	// WritableFrontPortTypeN captures enum value "n"
	WritableFrontPortTypeN string = "n"

	// WritableFrontPortTypeMrj21 captures enum value "mrj21"
	WritableFrontPortTypeMrj21 string = "mrj21"

	// WritableFrontPortTypeFc captures enum value "fc"
	WritableFrontPortTypeFc string = "fc"

	// WritableFrontPortTypeLc captures enum value "lc"
	WritableFrontPortTypeLc string = "lc"

	// WritableFrontPortTypeLcDashPc captures enum value "lc-pc"
	WritableFrontPortTypeLcDashPc string = "lc-pc"

	// WritableFrontPortTypeLcDashUpc captures enum value "lc-upc"
	WritableFrontPortTypeLcDashUpc string = "lc-upc"

	// WritableFrontPortTypeLcDashApc captures enum value "lc-apc"
	WritableFrontPortTypeLcDashApc string = "lc-apc"

	// WritableFrontPortTypeLsh captures enum value "lsh"
	WritableFrontPortTypeLsh string = "lsh"

	// WritableFrontPortTypeLshDashPc captures enum value "lsh-pc"
	WritableFrontPortTypeLshDashPc string = "lsh-pc"

	// WritableFrontPortTypeLshDashUpc captures enum value "lsh-upc"
	WritableFrontPortTypeLshDashUpc string = "lsh-upc"

	// WritableFrontPortTypeLshDashApc captures enum value "lsh-apc"
	WritableFrontPortTypeLshDashApc string = "lsh-apc"

	// WritableFrontPortTypeMpo captures enum value "mpo"
	WritableFrontPortTypeMpo string = "mpo"

	// WritableFrontPortTypeMtrj captures enum value "mtrj"
	WritableFrontPortTypeMtrj string = "mtrj"

	// WritableFrontPortTypeSc captures enum value "sc"
	WritableFrontPortTypeSc string = "sc"

	// WritableFrontPortTypeScDashPc captures enum value "sc-pc"
	WritableFrontPortTypeScDashPc string = "sc-pc"

	// WritableFrontPortTypeScDashUpc captures enum value "sc-upc"
	WritableFrontPortTypeScDashUpc string = "sc-upc"

	// WritableFrontPortTypeScDashApc captures enum value "sc-apc"
	WritableFrontPortTypeScDashApc string = "sc-apc"

	// WritableFrontPortTypeSt captures enum value "st"
	WritableFrontPortTypeSt string = "st"

	// WritableFrontPortTypeCs captures enum value "cs"
	WritableFrontPortTypeCs string = "cs"

	// WritableFrontPortTypeSn captures enum value "sn"
	WritableFrontPortTypeSn string = "sn"

	// WritableFrontPortTypeSmaDash905 captures enum value "sma-905"
	WritableFrontPortTypeSmaDash905 string = "sma-905"

	// WritableFrontPortTypeSmaDash906 captures enum value "sma-906"
	WritableFrontPortTypeSmaDash906 string = "sma-906"

	// WritableFrontPortTypeUrmDashP2 captures enum value "urm-p2"
	WritableFrontPortTypeUrmDashP2 string = "urm-p2"

	// WritableFrontPortTypeUrmDashP4 captures enum value "urm-p4"
	WritableFrontPortTypeUrmDashP4 string = "urm-p4"

	// WritableFrontPortTypeUrmDashP8 captures enum value "urm-p8"
	WritableFrontPortTypeUrmDashP8 string = "urm-p8"

	// WritableFrontPortTypeSplice captures enum value "splice"
	WritableFrontPortTypeSplice string = "splice"
)

// prop value enum
func (m *WritableFrontPort) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableFrontPortTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableFrontPort) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable front port based on the context it is used
func (m *WritableFrontPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOccupied(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinkPeer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinkPeerType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableFrontPort) contextValidateOccupied(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_occupied", "body", m.Occupied); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {
		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritableFrontPort) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) contextValidateLinkPeer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritableFrontPort) contextValidateLinkPeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "link_peer_type", "body", string(m.LinkPeerType)); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableFrontPort) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableFrontPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableFrontPort) UnmarshalBinary(b []byte) error {
	var res WritableFrontPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
