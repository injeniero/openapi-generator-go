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

// UpdatableEntity is an object. An entity that can be updated.
type UpdatableEntity struct {
	// Id:
	Id string `json:"id" mapstructure:"id"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// UpdateTs:
	UpdateTs int32 `json:"update_ts" mapstructure:"update_ts"`
}

// Validate implements basic validation for this model
func (m UpdatableEntity) Validate() error {
	return validation.Errors{
		"id": validation.Validate(
			m.Id, validation.Required, is.UUID,
		),
	}.Filter()
}

// GetId returns the Id property
func (m UpdatableEntity) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m *UpdatableEntity) SetId(val string) {
	m.Id = val
}

// GetName returns the Name property
func (m UpdatableEntity) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *UpdatableEntity) SetName(val string) {
	m.Name = val
}

// GetUpdateTs returns the UpdateTs property
func (m UpdatableEntity) GetUpdateTs() int32 {
	return m.UpdateTs
}

// SetUpdateTs sets the UpdateTs property
func (m *UpdatableEntity) SetUpdateTs(val int32) {
	m.UpdateTs = val
}
