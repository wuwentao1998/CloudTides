// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResourceInfoItem resource info item
//
// swagger:model ResourceInfoItem
type ResourceInfoItem struct {

	// CPU percent
	CPUPercent float64 `json:"CPUPercent,omitempty"`

	// RAM percent
	RAMPercent float64 `json:"RAMPercent,omitempty"`

	// active v ms
	ActiveVMs int64 `json:"activeVMs,omitempty"`

	// cluster
	Cluster string `json:"cluster,omitempty"`

	// current CPU
	CurrentCPU float64 `json:"currentCPU,omitempty"`

	// current RAM
	CurrentRAM float64 `json:"currentRAM,omitempty"`

	// datacenter
	Datacenter string `json:"datacenter,omitempty"`

	// date added
	DateAdded string `json:"dateAdded,omitempty"`

	// host address
	HostAddress string `json:"hostAddress,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// is active
	IsActive bool `json:"isActive,omitempty"`

	// job completed
	JobCompleted int64 `json:"jobCompleted,omitempty"`

	// last deployed
	LastDeployed string `json:"lastDeployed,omitempty"`

	// monitored
	Monitored bool `json:"monitored,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// platform type
	// Enum: [vsphere kvm hyper-v]
	PlatformType string `json:"platformType,omitempty"`

	// policy name
	PolicyName string `json:"policyName,omitempty"`

	// status
	// Enum: [idle normal busy unknown]
	Status string `json:"status,omitempty"`

	// total CPU
	TotalCPU float64 `json:"totalCPU,omitempty"`

	// total jobs
	TotalJobs int64 `json:"totalJobs,omitempty"`

	// total RAM
	TotalRAM float64 `json:"totalRAM,omitempty"`

	// total v ms
	TotalVMs int64 `json:"totalVMs,omitempty"`
}

// Validate validates this resource info item
func (m *ResourceInfoItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlatformType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var resourceInfoItemTypePlatformTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["vsphere","kvm","hyper-v"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceInfoItemTypePlatformTypePropEnum = append(resourceInfoItemTypePlatformTypePropEnum, v)
	}
}

const (

	// ResourceInfoItemPlatformTypeVsphere captures enum value "vsphere"
	ResourceInfoItemPlatformTypeVsphere string = "vsphere"

	// ResourceInfoItemPlatformTypeKvm captures enum value "kvm"
	ResourceInfoItemPlatformTypeKvm string = "kvm"

	// ResourceInfoItemPlatformTypeHyperv captures enum value "hyper-v"
	ResourceInfoItemPlatformTypeHyperv string = "hyper-v"
)

// prop value enum
func (m *ResourceInfoItem) validatePlatformTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceInfoItemTypePlatformTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ResourceInfoItem) validatePlatformType(formats strfmt.Registry) error {

	if swag.IsZero(m.PlatformType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePlatformTypeEnum("platformType", "body", m.PlatformType); err != nil {
		return err
	}

	return nil
}

var resourceInfoItemTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["idle","normal","busy","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceInfoItemTypeStatusPropEnum = append(resourceInfoItemTypeStatusPropEnum, v)
	}
}

const (

	// ResourceInfoItemStatusIdle captures enum value "idle"
	ResourceInfoItemStatusIdle string = "idle"

	// ResourceInfoItemStatusNormal captures enum value "normal"
	ResourceInfoItemStatusNormal string = "normal"

	// ResourceInfoItemStatusBusy captures enum value "busy"
	ResourceInfoItemStatusBusy string = "busy"

	// ResourceInfoItemStatusUnknown captures enum value "unknown"
	ResourceInfoItemStatusUnknown string = "unknown"
)

// prop value enum
func (m *ResourceInfoItem) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceInfoItemTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ResourceInfoItem) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceInfoItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceInfoItem) UnmarshalBinary(b []byte) error {
	var res ResourceInfoItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
