// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DiscoveryIgnitionParams discovery ignition params
//
// swagger:model discovery-ignition-params
type DiscoveryIgnitionParams struct {

	// config
	Config string `json:"config,omitempty"`
}

// Validate validates this discovery ignition params
func (m *DiscoveryIgnitionParams) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DiscoveryIgnitionParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiscoveryIgnitionParams) UnmarshalBinary(b []byte) error {
	var res DiscoveryIgnitionParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
