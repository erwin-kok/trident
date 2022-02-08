// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// S3AuditEventSelector S3 audit event selector per SVM per bucket.  Use to set access and permission type for S3 event audit.
//
// swagger:model s3_audit_event_selector
type S3AuditEventSelector struct {

	// Specifies read and write access types.
	//
	// Enum: [read write all]
	Access *string `json:"access,omitempty"`

	// Specifies the name of the bucket. Bucket name is a string that can only contain the following combination of ASCII-range alphanumeric characters 0-9, a-z, ".", and "-".
	// Example: bucket1
	// Max Length: 63
	// Min Length: 3
	Bucket string `json:"bucket,omitempty"`

	// Specifies allow and deny permission types.
	//
	// Enum: [deny allow all]
	Permission *string `json:"permission,omitempty"`

	// svm
	Svm *S3AuditEventSelectorSvm `json:"svm,omitempty"`
}

// Validate validates this s3 audit event selector
func (m *S3AuditEventSelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBucket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var s3AuditEventSelectorTypeAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["read","write","all"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		s3AuditEventSelectorTypeAccessPropEnum = append(s3AuditEventSelectorTypeAccessPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// s3_audit_event_selector
	// S3AuditEventSelector
	// access
	// Access
	// read
	// END DEBUGGING
	// S3AuditEventSelectorAccessRead captures enum value "read"
	S3AuditEventSelectorAccessRead string = "read"

	// BEGIN DEBUGGING
	// s3_audit_event_selector
	// S3AuditEventSelector
	// access
	// Access
	// write
	// END DEBUGGING
	// S3AuditEventSelectorAccessWrite captures enum value "write"
	S3AuditEventSelectorAccessWrite string = "write"

	// BEGIN DEBUGGING
	// s3_audit_event_selector
	// S3AuditEventSelector
	// access
	// Access
	// all
	// END DEBUGGING
	// S3AuditEventSelectorAccessAll captures enum value "all"
	S3AuditEventSelectorAccessAll string = "all"
)

// prop value enum
func (m *S3AuditEventSelector) validateAccessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, s3AuditEventSelectorTypeAccessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *S3AuditEventSelector) validateAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.Access) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessEnum("access", "body", *m.Access); err != nil {
		return err
	}

	return nil
}

func (m *S3AuditEventSelector) validateBucket(formats strfmt.Registry) error {
	if swag.IsZero(m.Bucket) { // not required
		return nil
	}

	if err := validate.MinLength("bucket", "body", m.Bucket, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("bucket", "body", m.Bucket, 63); err != nil {
		return err
	}

	return nil
}

var s3AuditEventSelectorTypePermissionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["deny","allow","all"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		s3AuditEventSelectorTypePermissionPropEnum = append(s3AuditEventSelectorTypePermissionPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// s3_audit_event_selector
	// S3AuditEventSelector
	// permission
	// Permission
	// deny
	// END DEBUGGING
	// S3AuditEventSelectorPermissionDeny captures enum value "deny"
	S3AuditEventSelectorPermissionDeny string = "deny"

	// BEGIN DEBUGGING
	// s3_audit_event_selector
	// S3AuditEventSelector
	// permission
	// Permission
	// allow
	// END DEBUGGING
	// S3AuditEventSelectorPermissionAllow captures enum value "allow"
	S3AuditEventSelectorPermissionAllow string = "allow"

	// BEGIN DEBUGGING
	// s3_audit_event_selector
	// S3AuditEventSelector
	// permission
	// Permission
	// all
	// END DEBUGGING
	// S3AuditEventSelectorPermissionAll captures enum value "all"
	S3AuditEventSelectorPermissionAll string = "all"
)

// prop value enum
func (m *S3AuditEventSelector) validatePermissionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, s3AuditEventSelectorTypePermissionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *S3AuditEventSelector) validatePermission(formats strfmt.Registry) error {
	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	// value enum
	if err := m.validatePermissionEnum("permission", "body", *m.Permission); err != nil {
		return err
	}

	return nil
}

func (m *S3AuditEventSelector) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 audit event selector based on the context it is used
func (m *S3AuditEventSelector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3AuditEventSelector) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3AuditEventSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3AuditEventSelector) UnmarshalBinary(b []byte) error {
	var res S3AuditEventSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3AuditEventSelectorSvm s3 audit event selector svm
//
// swagger:model S3AuditEventSelectorSvm
type S3AuditEventSelectorSvm struct {

	// links
	Links *S3AuditEventSelectorSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this s3 audit event selector svm
func (m *S3AuditEventSelectorSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3AuditEventSelectorSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 audit event selector svm based on the context it is used
func (m *S3AuditEventSelectorSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3AuditEventSelectorSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3AuditEventSelectorSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3AuditEventSelectorSvm) UnmarshalBinary(b []byte) error {
	var res S3AuditEventSelectorSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3AuditEventSelectorSvmLinks s3 audit event selector svm links
//
// swagger:model S3AuditEventSelectorSvmLinks
type S3AuditEventSelectorSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this s3 audit event selector svm links
func (m *S3AuditEventSelectorSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3AuditEventSelectorSvmLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 audit event selector svm links based on the context it is used
func (m *S3AuditEventSelectorSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3AuditEventSelectorSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3AuditEventSelectorSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3AuditEventSelectorSvmLinks) UnmarshalBinary(b []byte) error {
	var res S3AuditEventSelectorSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}