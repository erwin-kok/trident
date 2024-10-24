// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// AccountTotpDeleteCollectionReader is a Reader for the AccountTotpDeleteCollection structure.
type AccountTotpDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountTotpDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountTotpDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAccountTotpDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountTotpDeleteCollectionOK creates a AccountTotpDeleteCollectionOK with default headers values
func NewAccountTotpDeleteCollectionOK() *AccountTotpDeleteCollectionOK {
	return &AccountTotpDeleteCollectionOK{}
}

/*
AccountTotpDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type AccountTotpDeleteCollectionOK struct {
}

// IsSuccess returns true when this account totp delete collection o k response has a 2xx status code
func (o *AccountTotpDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account totp delete collection o k response has a 3xx status code
func (o *AccountTotpDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account totp delete collection o k response has a 4xx status code
func (o *AccountTotpDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account totp delete collection o k response has a 5xx status code
func (o *AccountTotpDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account totp delete collection o k response a status code equal to that given
func (o *AccountTotpDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account totp delete collection o k response
func (o *AccountTotpDeleteCollectionOK) Code() int {
	return 200
}

func (o *AccountTotpDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /security/login/totps][%d] accountTotpDeleteCollectionOK", 200)
}

func (o *AccountTotpDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /security/login/totps][%d] accountTotpDeleteCollectionOK", 200)
}

func (o *AccountTotpDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountTotpDeleteCollectionDefault creates a AccountTotpDeleteCollectionDefault with default headers values
func NewAccountTotpDeleteCollectionDefault(code int) *AccountTotpDeleteCollectionDefault {
	return &AccountTotpDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
AccountTotpDeleteCollectionDefault describes a response with status code -1, with default header values.

Error
*/
type AccountTotpDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this account totp delete collection default response has a 2xx status code
func (o *AccountTotpDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this account totp delete collection default response has a 3xx status code
func (o *AccountTotpDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this account totp delete collection default response has a 4xx status code
func (o *AccountTotpDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this account totp delete collection default response has a 5xx status code
func (o *AccountTotpDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this account totp delete collection default response a status code equal to that given
func (o *AccountTotpDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the account totp delete collection default response
func (o *AccountTotpDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *AccountTotpDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/login/totps][%d] account_totp_delete_collection default %s", o._statusCode, payload)
}

func (o *AccountTotpDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/login/totps][%d] account_totp_delete_collection default %s", o._statusCode, payload)
}

func (o *AccountTotpDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AccountTotpDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AccountTotpDeleteCollectionBody account totp delete collection body
swagger:model AccountTotpDeleteCollectionBody
*/
type AccountTotpDeleteCollectionBody struct {

	// totp response inline records
	TotpResponseInlineRecords []*models.Totp `json:"records,omitempty"`
}

// Validate validates this account totp delete collection body
func (o *AccountTotpDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTotpResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccountTotpDeleteCollectionBody) validateTotpResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.TotpResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.TotpResponseInlineRecords); i++ {
		if swag.IsZero(o.TotpResponseInlineRecords[i]) { // not required
			continue
		}

		if o.TotpResponseInlineRecords[i] != nil {
			if err := o.TotpResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this account totp delete collection body based on the context it is used
func (o *AccountTotpDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateTotpResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccountTotpDeleteCollectionBody) contextValidateTotpResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.TotpResponseInlineRecords); i++ {

		if o.TotpResponseInlineRecords[i] != nil {
			if err := o.TotpResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AccountTotpDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccountTotpDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res AccountTotpDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
