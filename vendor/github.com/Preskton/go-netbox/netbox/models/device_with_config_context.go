// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceWithConfigContext device with config context
// swagger:model DeviceWithConfigContext
type DeviceWithConfigContext struct {

	// Asset tag
	//
	// A unique tag used to identify this device
	// Max Length: 50
	AssetTag string `json:"asset_tag,omitempty"`

	// cluster
	Cluster *NestedCluster `json:"cluster,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Config context
	// Read Only: true
	ConfigContext string `json:"config_context,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// device role
	// Required: true
	DeviceRole *NestedDeviceRole `json:"device_role"`

	// device type
	// Required: true
	DeviceType *NestedDeviceType `json:"device_type"`

	// Display name
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// face
	Face *DeviceWithConfigContextFace `json:"face,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Local context data
	LocalContextData string `json:"local_context_data,omitempty"`

	// Name
	// Max Length: 64
	Name string `json:"name,omitempty"`

	// Parent device
	// Read Only: true
	ParentDevice string `json:"parent_device,omitempty"`

	// platform
	Platform *NestedPlatform `json:"platform,omitempty"`

	// Position (U)
	//
	// The lowest-numbered unit occupied by the device
	// Maximum: 32767
	// Minimum: 1
	Position int64 `json:"position,omitempty"`

	// primary ip
	PrimaryIP *DeviceIPAddress `json:"primary_ip,omitempty"`

	// primary ip4
	PrimaryIp4 *DeviceIPAddress `json:"primary_ip4,omitempty"`

	// primary ip6
	PrimaryIp6 *DeviceIPAddress `json:"primary_ip6,omitempty"`

	// rack
	Rack *NestedRack `json:"rack,omitempty"`

	// Serial number
	// Max Length: 50
	Serial string `json:"serial,omitempty"`

	// site
	// Required: true
	Site *NestedSite `json:"site"`

	// status
	Status *DeviceWithConfigContextStatus `json:"status,omitempty"`

	// Tags
	Tags []string `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// Vc position
	// Maximum: 255
	// Minimum: 0
	VcPosition *int64 `json:"vc_position,omitempty"`

	// Vc priority
	// Maximum: 255
	// Minimum: 0
	VcPriority *int64 `json:"vc_priority,omitempty"`

	// virtual chassis
	VirtualChassis *DeviceVirtualChassis `json:"virtual_chassis,omitempty"`
}

// Validate validates this device with config context
func (m *DeviceWithConfigContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcPriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualChassis(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceWithConfigContext) validateAssetTag(formats strfmt.Registry) error {

	if swag.IsZero(m.AssetTag) { // not required
		return nil
	}

	if err := validate.MaxLength("asset_tag", "body", string(m.AssetTag), 50); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateDeviceRole(formats strfmt.Registry) error {

	if err := validate.Required("device_role", "body", m.DeviceRole); err != nil {
		return err
	}

	if m.DeviceRole != nil {
		if err := m.DeviceRole.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_role")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	if m.DeviceType != nil {
		if err := m.DeviceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateFace(formats strfmt.Registry) error {

	if swag.IsZero(m.Face) { // not required
		return nil
	}

	if m.Face != nil {
		if err := m.Face.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("face")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validatePlatform(formats strfmt.Registry) error {

	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	if m.Platform != nil {
		if err := m.Platform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.MinimumInt("position", "body", int64(m.Position), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("position", "body", int64(m.Position), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validatePrimaryIP(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryIP) { // not required
		return nil
	}

	if m.PrimaryIP != nil {
		if err := m.PrimaryIP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validatePrimaryIp4(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryIp4) { // not required
		return nil
	}

	if m.PrimaryIp4 != nil {
		if err := m.PrimaryIp4.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip4")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validatePrimaryIp6(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryIp6) { // not required
		return nil
	}

	if m.PrimaryIp6 != nil {
		if err := m.PrimaryIp6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip6")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateRack(formats strfmt.Registry) error {

	if swag.IsZero(m.Rack) { // not required
		return nil
	}

	if m.Rack != nil {
		if err := m.Rack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rack")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateSerial(formats strfmt.Registry) error {

	if swag.IsZero(m.Serial) { // not required
		return nil
	}

	if err := validate.MaxLength("serial", "body", string(m.Serial), 50); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateSite(formats strfmt.Registry) error {

	if err := validate.Required("site", "body", m.Site); err != nil {
		return err
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateTenant(formats strfmt.Registry) error {

	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateVcPosition(formats strfmt.Registry) error {

	if swag.IsZero(m.VcPosition) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_position", "body", int64(*m.VcPosition), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_position", "body", int64(*m.VcPosition), 255, false); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateVcPriority(formats strfmt.Registry) error {

	if swag.IsZero(m.VcPriority) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_priority", "body", int64(*m.VcPriority), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_priority", "body", int64(*m.VcPriority), 255, false); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateVirtualChassis(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualChassis) { // not required
		return nil
	}

	if m.VirtualChassis != nil {
		if err := m.VirtualChassis.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_chassis")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceWithConfigContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceWithConfigContext) UnmarshalBinary(b []byte) error {
	var res DeviceWithConfigContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DeviceWithConfigContextFace Face
// swagger:model DeviceWithConfigContextFace
type DeviceWithConfigContextFace struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *int64 `json:"value"`
}

// Validate validates this device with config context face
func (m *DeviceWithConfigContextFace) Validate(formats strfmt.Registry) error {
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

func (m *DeviceWithConfigContextFace) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("face"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContextFace) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("face"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceWithConfigContextFace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceWithConfigContextFace) UnmarshalBinary(b []byte) error {
	var res DeviceWithConfigContextFace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DeviceWithConfigContextStatus Status
// swagger:model DeviceWithConfigContextStatus
type DeviceWithConfigContextStatus struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *int64 `json:"value"`
}

// Validate validates this device with config context status
func (m *DeviceWithConfigContextStatus) Validate(formats strfmt.Registry) error {
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

func (m *DeviceWithConfigContextStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContextStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceWithConfigContextStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceWithConfigContextStatus) UnmarshalBinary(b []byte) error {
	var res DeviceWithConfigContextStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}