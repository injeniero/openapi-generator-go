// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// IntegerEnumWithNull is an enum.
type IntegerEnumWithNull int32

// Validate implements basic validation for this model
func (m IntegerEnumWithNull) Validate() error {
	return InKnownIntegerEnumWithNull.Validate(m)
}

var (
	IntegerEnumWithNull1 IntegerEnumWithNull = 1
	IntegerEnumWithNull2 IntegerEnumWithNull = 2

	// KnownIntegerEnumWithNull is the list of valid IntegerEnumWithNull
	KnownIntegerEnumWithNull = []IntegerEnumWithNull{
		IntegerEnumWithNull1,
		IntegerEnumWithNull2,
	}
	// KnownIntegerEnumWithNullInt32 is the list of valid IntegerEnumWithNull as int32
	KnownIntegerEnumWithNullInt32 = []int32{
		int32(IntegerEnumWithNull1),
		int32(IntegerEnumWithNull2),
	}

	// InKnownIntegerEnumWithNull is an ozzo-validator for IntegerEnumWithNull
	InKnownIntegerEnumWithNull = validation.In(
		IntegerEnumWithNull1,
		IntegerEnumWithNull2,
	)
)
