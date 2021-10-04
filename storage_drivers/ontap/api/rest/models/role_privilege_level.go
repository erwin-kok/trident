// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// RolePrivilegeLevel Access level for the REST endpoint.
// Example: readonly
//
// swagger:model role_privilege_level
type RolePrivilegeLevel string

func NewRolePrivilegeLevel(value RolePrivilegeLevel) *RolePrivilegeLevel {
	v := value
	return &v
}

const (

	// RolePrivilegeLevelNone captures enum value "none"
	RolePrivilegeLevelNone RolePrivilegeLevel = "none"

	// RolePrivilegeLevelReadonly captures enum value "readonly"
	RolePrivilegeLevelReadonly RolePrivilegeLevel = "readonly"

	// RolePrivilegeLevelAll captures enum value "all"
	RolePrivilegeLevelAll RolePrivilegeLevel = "all"
)

// for schema
var rolePrivilegeLevelEnum []interface{}

func init() {
	var res []RolePrivilegeLevel
	if err := json.Unmarshal([]byte(`["none","readonly","all"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rolePrivilegeLevelEnum = append(rolePrivilegeLevelEnum, v)
	}
}

func (m RolePrivilegeLevel) validateRolePrivilegeLevelEnum(path, location string, value RolePrivilegeLevel) error {
	if err := validate.EnumCase(path, location, value, rolePrivilegeLevelEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this role privilege level
func (m RolePrivilegeLevel) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRolePrivilegeLevelEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this role privilege level based on context it is used
func (m RolePrivilegeLevel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
