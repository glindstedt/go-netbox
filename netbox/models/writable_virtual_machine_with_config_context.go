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

// WritableVirtualMachineWithConfigContext writable virtual machine with config context
//
// swagger:model WritableVirtualMachineWithConfigContext
type WritableVirtualMachineWithConfigContext struct {

	// Cluster
	// Required: true
	Cluster *int64 `json:"cluster"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Config context
	// Read Only: true
	ConfigContext map[string]interface{} `json:"config_context,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Disk (GB)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	Disk *int64 `json:"disk"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Local context data
	LocalContextData map[string]interface{} `json:"local_context_data"`

	// Memory (MB)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	Memory *int64 `json:"memory"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Platform
	Platform *int64 `json:"platform"`

	// Primary ip
	// Read Only: true
	PrimaryIP string `json:"primary_ip,omitempty"`

	// Primary IPv4
	PrimaryIp4 *int64 `json:"primary_ip4"`

	// Primary IPv6
	PrimaryIp6 *int64 `json:"primary_ip6"`

	// Role
	Role *int64 `json:"role"`

	// Site
	// Read Only: true
	Site string `json:"site,omitempty"`

	// Status
	// Enum: [offline active planned staged failed decommissioning]
	Status string `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Tenant
	Tenant *int64 `json:"tenant"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// VCPUs
	// Minimum: 0.01
	Vcpus *float64 `json:"vcpus"`
}

// Validate validates this writable virtual machine with config context
func (m *WritableVirtualMachineWithConfigContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcpus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateDisk(formats strfmt.Registry) error {
	if swag.IsZero(m.Disk) { // not required
		return nil
	}

	if err := validate.MinimumInt("disk", "body", *m.Disk, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("disk", "body", *m.Disk, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateMemory(formats strfmt.Registry) error {
	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if err := validate.MinimumInt("memory", "body", *m.Memory, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("memory", "body", *m.Memory, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateName(formats strfmt.Registry) error {

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

var writableVirtualMachineWithConfigContextTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["offline","active","planned","staged","failed","decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableVirtualMachineWithConfigContextTypeStatusPropEnum = append(writableVirtualMachineWithConfigContextTypeStatusPropEnum, v)
	}
}

const (

	// WritableVirtualMachineWithConfigContextStatusOffline captures enum value "offline"
	WritableVirtualMachineWithConfigContextStatusOffline string = "offline"

	// WritableVirtualMachineWithConfigContextStatusActive captures enum value "active"
	WritableVirtualMachineWithConfigContextStatusActive string = "active"

	// WritableVirtualMachineWithConfigContextStatusPlanned captures enum value "planned"
	WritableVirtualMachineWithConfigContextStatusPlanned string = "planned"

	// WritableVirtualMachineWithConfigContextStatusStaged captures enum value "staged"
	WritableVirtualMachineWithConfigContextStatusStaged string = "staged"

	// WritableVirtualMachineWithConfigContextStatusFailed captures enum value "failed"
	WritableVirtualMachineWithConfigContextStatusFailed string = "failed"

	// WritableVirtualMachineWithConfigContextStatusDecommissioning captures enum value "decommissioning"
	WritableVirtualMachineWithConfigContextStatusDecommissioning string = "decommissioning"
)

// prop value enum
func (m *WritableVirtualMachineWithConfigContext) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableVirtualMachineWithConfigContextTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateTags(formats strfmt.Registry) error {
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

func (m *WritableVirtualMachineWithConfigContext) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateVcpus(formats strfmt.Registry) error {
	if swag.IsZero(m.Vcpus) { // not required
		return nil
	}

	if err := validate.Minimum("vcpus", "body", *m.Vcpus, 0.01, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable virtual machine with config context based on the context it is used
func (m *WritableVirtualMachineWithConfigContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigContext(ctx, formats); err != nil {
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

	if err := m.contextValidatePrimaryIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSite(ctx, formats); err != nil {
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

func (m *WritableVirtualMachineWithConfigContext) contextValidateConfigContext(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) contextValidatePrimaryIP(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "primary_ip", "body", string(m.PrimaryIP)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "site", "body", string(m.Site)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WritableVirtualMachineWithConfigContext) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableVirtualMachineWithConfigContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableVirtualMachineWithConfigContext) UnmarshalBinary(b []byte) error {
	var res WritableVirtualMachineWithConfigContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
