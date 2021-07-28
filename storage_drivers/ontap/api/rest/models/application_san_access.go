// Code generated by go-swagger; DO NOT EDIT.

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

// ApplicationSanAccess application san access
//
// swagger:model application_san_access
type ApplicationSanAccess struct {

	// backing storage
	BackingStorage *ApplicationSanAccessBackingStorage `json:"backing_storage,omitempty"`

	// Clone
	// Read Only: true
	IsClone *bool `json:"is_clone,omitempty"`

	// lun mappings
	LunMappings []*ApplicationLunMappingObject `json:"lun_mappings,omitempty"`

	// LUN serial number
	// Read Only: true
	// Max Length: 12
	// Min Length: 12
	SerialNumber string `json:"serial_number,omitempty"`
}

// Validate validates this application san access
func (m *ApplicationSanAccess) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackingStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLunMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerialNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSanAccess) validateBackingStorage(formats strfmt.Registry) error {
	if swag.IsZero(m.BackingStorage) { // not required
		return nil
	}

	if m.BackingStorage != nil {
		if err := m.BackingStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backing_storage")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationSanAccess) validateLunMappings(formats strfmt.Registry) error {
	if swag.IsZero(m.LunMappings) { // not required
		return nil
	}

	for i := 0; i < len(m.LunMappings); i++ {
		if swag.IsZero(m.LunMappings[i]) { // not required
			continue
		}

		if m.LunMappings[i] != nil {
			if err := m.LunMappings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lun_mappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationSanAccess) validateSerialNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.SerialNumber) { // not required
		return nil
	}

	if err := validate.MinLength("serial_number", "body", m.SerialNumber, 12); err != nil {
		return err
	}

	if err := validate.MaxLength("serial_number", "body", m.SerialNumber, 12); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this application san access based on the context it is used
func (m *ApplicationSanAccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackingStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsClone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLunMappings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSerialNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSanAccess) contextValidateBackingStorage(ctx context.Context, formats strfmt.Registry) error {

	if m.BackingStorage != nil {
		if err := m.BackingStorage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backing_storage")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationSanAccess) contextValidateIsClone(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_clone", "body", m.IsClone); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationSanAccess) contextValidateLunMappings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LunMappings); i++ {

		if m.LunMappings[i] != nil {
			if err := m.LunMappings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lun_mappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationSanAccess) contextValidateSerialNumber(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "serial_number", "body", string(m.SerialNumber)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationSanAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSanAccess) UnmarshalBinary(b []byte) error {
	var res ApplicationSanAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationSanAccessBackingStorage application san access backing storage
//
// swagger:model ApplicationSanAccessBackingStorage
type ApplicationSanAccessBackingStorage struct {

	// Backing storage type
	// Read Only: true
	// Enum: [lun]
	Type string `json:"type,omitempty"`

	// Backing storage UUID
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this application san access backing storage
func (m *ApplicationSanAccessBackingStorage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var applicationSanAccessBackingStorageTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["lun"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationSanAccessBackingStorageTypeTypePropEnum = append(applicationSanAccessBackingStorageTypeTypePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// ApplicationSanAccessBackingStorage
	// ApplicationSanAccessBackingStorage
	// type
	// Type
	// lun
	// END RIPPY DEBUGGING
	// ApplicationSanAccessBackingStorageTypeLun captures enum value "lun"
	ApplicationSanAccessBackingStorageTypeLun string = "lun"
)

// prop value enum
func (m *ApplicationSanAccessBackingStorage) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, applicationSanAccessBackingStorageTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationSanAccessBackingStorage) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("backing_storage"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this application san access backing storage based on the context it is used
func (m *ApplicationSanAccessBackingStorage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSanAccessBackingStorage) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "backing_storage"+"."+"type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationSanAccessBackingStorage) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "backing_storage"+"."+"uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationSanAccessBackingStorage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSanAccessBackingStorage) UnmarshalBinary(b []byte) error {
	var res ApplicationSanAccessBackingStorage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY