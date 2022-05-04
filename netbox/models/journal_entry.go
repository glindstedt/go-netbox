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

// JournalEntry journal entry
//
// swagger:model JournalEntry
type JournalEntry struct {

	// Assigned object
	// Read Only: true
	AssignedObject map[string]*string `json:"assigned_object,omitempty"`

	// Assigned object id
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	AssignedObjectID *int64 `json:"assigned_object_id"`

	// Assigned object type
	// Required: true
	AssignedObjectType *string `json:"assigned_object_type"`

	// Comments
	// Required: true
	// Min Length: 1
	Comments *string `json:"comments"`

	// Created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Created by
	CreatedBy *int64 `json:"created_by,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// kind
	Kind *JournalEntryKind `json:"kind,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this journal entry
func (m *JournalEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignedObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignedObjectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

func (m *JournalEntry) validateAssignedObjectID(formats strfmt.Registry) error {

	if err := validate.Required("assigned_object_id", "body", m.AssignedObjectID); err != nil {
		return err
	}

	if err := validate.MinimumInt("assigned_object_id", "body", *m.AssignedObjectID, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("assigned_object_id", "body", *m.AssignedObjectID, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *JournalEntry) validateAssignedObjectType(formats strfmt.Registry) error {

	if err := validate.Required("assigned_object_type", "body", m.AssignedObjectType); err != nil {
		return err
	}

	return nil
}

func (m *JournalEntry) validateComments(formats strfmt.Registry) error {

	if err := validate.Required("comments", "body", m.Comments); err != nil {
		return err
	}

	if err := validate.MinLength("comments", "body", *m.Comments, 1); err != nil {
		return err
	}

	return nil
}

func (m *JournalEntry) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JournalEntry) validateKind(formats strfmt.Registry) error {
	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	if m.Kind != nil {
		if err := m.Kind.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kind")
			}
			return err
		}
	}

	return nil
}

func (m *JournalEntry) validateTags(formats strfmt.Registry) error {
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

func (m *JournalEntry) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this journal entry based on the context it is used
func (m *JournalEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssignedObject(ctx, formats); err != nil {
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

	if err := m.contextValidateKind(ctx, formats); err != nil {
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

func (m *JournalEntry) contextValidateAssignedObject(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *JournalEntry) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *JournalEntry) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *JournalEntry) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *JournalEntry) contextValidateKind(ctx context.Context, formats strfmt.Registry) error {

	if m.Kind != nil {
		if err := m.Kind.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kind")
			}
			return err
		}
	}

	return nil
}

func (m *JournalEntry) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *JournalEntry) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JournalEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JournalEntry) UnmarshalBinary(b []byte) error {
	var res JournalEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// JournalEntryKind Kind
//
// swagger:model JournalEntryKind
type JournalEntryKind struct {

	// label
	// Required: true
	// Enum: [Info Success Warning Danger]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [info success warning danger]
	Value *string `json:"value"`
}

// Validate validates this journal entry kind
func (m *JournalEntryKind) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var journalEntryKindTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Info","Success","Warning","Danger"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		journalEntryKindTypeLabelPropEnum = append(journalEntryKindTypeLabelPropEnum, v)
	}
}

const (

	// JournalEntryKindLabelInfo captures enum value "Info"
	JournalEntryKindLabelInfo string = "Info"

	// JournalEntryKindLabelSuccess captures enum value "Success"
	JournalEntryKindLabelSuccess string = "Success"

	// JournalEntryKindLabelWarning captures enum value "Warning"
	JournalEntryKindLabelWarning string = "Warning"

	// JournalEntryKindLabelDanger captures enum value "Danger"
	JournalEntryKindLabelDanger string = "Danger"
)

// prop value enum
func (m *JournalEntryKind) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, journalEntryKindTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JournalEntryKind) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("kind"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("kind"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var journalEntryKindTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["info","success","warning","danger"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		journalEntryKindTypeValuePropEnum = append(journalEntryKindTypeValuePropEnum, v)
	}
}

const (

	// JournalEntryKindValueInfo captures enum value "info"
	JournalEntryKindValueInfo string = "info"

	// JournalEntryKindValueSuccess captures enum value "success"
	JournalEntryKindValueSuccess string = "success"

	// JournalEntryKindValueWarning captures enum value "warning"
	JournalEntryKindValueWarning string = "warning"

	// JournalEntryKindValueDanger captures enum value "danger"
	JournalEntryKindValueDanger string = "danger"
)

// prop value enum
func (m *JournalEntryKind) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, journalEntryKindTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JournalEntryKind) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("kind"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("kind"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this journal entry kind based on context it is used
func (m *JournalEntryKind) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JournalEntryKind) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JournalEntryKind) UnmarshalBinary(b []byte) error {
	var res JournalEntryKind
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
