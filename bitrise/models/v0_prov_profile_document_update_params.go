// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V0ProvProfileDocumentUpdateParams v0 prov profile document update params
// swagger:model v0.ProvProfileDocumentUpdateParams
type V0ProvProfileDocumentUpdateParams struct {

	// is expose
	IsExpose bool `json:"is_expose,omitempty"`

	// is protected
	IsProtected bool `json:"is_protected,omitempty"`

	// processed
	Processed bool `json:"processed,omitempty"`
}

// Validate validates this v0 prov profile document update params
func (m *V0ProvProfileDocumentUpdateParams) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0ProvProfileDocumentUpdateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0ProvProfileDocumentUpdateParams) UnmarshalBinary(b []byte) error {
	var res V0ProvProfileDocumentUpdateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
