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

// SvmPeerPermissionApplications Applications for an SVM peer permission.
//
// swagger:model svm_peer_permission_applications
type SvmPeerPermissionApplications string

func NewSvmPeerPermissionApplications(value SvmPeerPermissionApplications) *SvmPeerPermissionApplications {
	v := value
	return &v
}

const (

	// SvmPeerPermissionApplicationsSnapmirror captures enum value "snapmirror"
	SvmPeerPermissionApplicationsSnapmirror SvmPeerPermissionApplications = "snapmirror"

	// SvmPeerPermissionApplicationsFlexcache captures enum value "flexcache"
	SvmPeerPermissionApplicationsFlexcache SvmPeerPermissionApplications = "flexcache"
)

// for schema
var svmPeerPermissionApplicationsEnum []interface{}

func init() {
	var res []SvmPeerPermissionApplications
	if err := json.Unmarshal([]byte(`["snapmirror","flexcache"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		svmPeerPermissionApplicationsEnum = append(svmPeerPermissionApplicationsEnum, v)
	}
}

func (m SvmPeerPermissionApplications) validateSvmPeerPermissionApplicationsEnum(path, location string, value SvmPeerPermissionApplications) error {
	if err := validate.EnumCase(path, location, value, svmPeerPermissionApplicationsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this svm peer permission applications
func (m SvmPeerPermissionApplications) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSvmPeerPermissionApplicationsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this svm peer permission applications based on context it is used
func (m SvmPeerPermissionApplications) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
