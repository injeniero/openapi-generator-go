// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// OrgEntity is an object. An organization
type OrgEntity struct {
	// Email:
	Email string `json:"email" mapstructure:"email"`
	// Id:
	Id string `json:"id" mapstructure:"id"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// UpdateTs:
	UpdateTs int32 `json:"update_ts" mapstructure:"update_ts"`
}

// Validate implements basic validation for this model
func (m OrgEntity) Validate() error {
	return validation.Errors{
		"email": validation.Validate(
			m.Email, validation.Required, is.EmailFormat,
		),
		"id": validation.Validate(
			m.Id, validation.Required, is.UUID,
		),
	}.Filter()
}

// GetEmail returns the Email property
func (m OrgEntity) GetEmail() string {
	return m.Email
}

// SetEmail sets the Email property
func (m *OrgEntity) SetEmail(val string) {
	m.Email = val
}

// GetId returns the Id property
func (m OrgEntity) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m *OrgEntity) SetId(val string) {
	m.Id = val
}

// GetName returns the Name property
func (m OrgEntity) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *OrgEntity) SetName(val string) {
	m.Name = val
}

// GetUpdateTs returns the UpdateTs property
func (m OrgEntity) GetUpdateTs() int32 {
	return m.UpdateTs
}

// SetUpdateTs sets the UpdateTs property
func (m *OrgEntity) SetUpdateTs(val int32) {
	m.UpdateTs = val
}
