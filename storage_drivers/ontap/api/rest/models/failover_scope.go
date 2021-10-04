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

// FailoverScope Defines where an interface may failover.
//
// swagger:model failover_scope
type FailoverScope string

func NewFailoverScope(value FailoverScope) *FailoverScope {
	v := value
	return &v
}

const (

	// FailoverScopeHomePortOnly captures enum value "home_port_only"
	FailoverScopeHomePortOnly FailoverScope = "home_port_only"

	// FailoverScopeDefault captures enum value "default"
	FailoverScopeDefault FailoverScope = "default"

	// FailoverScopeHomeNodeOnly captures enum value "home_node_only"
	FailoverScopeHomeNodeOnly FailoverScope = "home_node_only"

	// FailoverScopeSfoPartnersOnly captures enum value "sfo_partners_only"
	FailoverScopeSfoPartnersOnly FailoverScope = "sfo_partners_only"

	// FailoverScopeBroadcastDomainOnly captures enum value "broadcast_domain_only"
	FailoverScopeBroadcastDomainOnly FailoverScope = "broadcast_domain_only"
)

// for schema
var failoverScopeEnum []interface{}

func init() {
	var res []FailoverScope
	if err := json.Unmarshal([]byte(`["home_port_only","default","home_node_only","sfo_partners_only","broadcast_domain_only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		failoverScopeEnum = append(failoverScopeEnum, v)
	}
}

func (m FailoverScope) validateFailoverScopeEnum(path, location string, value FailoverScope) error {
	if err := validate.EnumCase(path, location, value, failoverScopeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this failover scope
func (m FailoverScope) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFailoverScopeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this failover scope based on context it is used
func (m FailoverScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
