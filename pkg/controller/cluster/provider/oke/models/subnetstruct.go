// Code generated by go-swagger; DO NOT EDIT.

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Subnetstruct subnetstruct
// swagger:model Subnetstruct
type Subnetstruct struct {

	// e tag
	ETag string `json:"ETag,omitempty"`

	// availability domain
	AvailabilityDomain string `json:"availabilityDomain,omitempty"`

	// cidr block
	CidrBlock string `json:"cidrBlock,omitempty"`

	// compartment Id
	CompartmentID string `json:"compartmentId,omitempty"`

	// dhcp options Id
	DhcpOptionsID string `json:"dhcpOptionsId,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// dns label
	DNSLabel string `json:"dnsLabel,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// lifecycle state
	LifecycleState string `json:"lifecycleState,omitempty"`

	// route table Id
	RouteTableID string `json:"routeTableId,omitempty"`

	// security list ids
	SecurityListIds []string `json:"securityListIds"`

	// subnet domain name
	SubnetDomainName string `json:"subnetDomainName,omitempty"`

	// time created
	TimeCreated string `json:"timeCreated,omitempty"`

	// vcn Id
	VcnID string `json:"vcnId,omitempty"`

	// virtual router Ip
	VirtualRouterIP string `json:"virtualRouterIp,omitempty"`

	// virtual router mac
	VirtualRouterMac string `json:"virtualRouterMac,omitempty"`
}

// Validate validates this subnetstruct
func (m *Subnetstruct) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecurityListIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subnetstruct) validateSecurityListIds(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityListIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Subnetstruct) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subnetstruct) UnmarshalBinary(b []byte) error {
	var res Subnetstruct
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
