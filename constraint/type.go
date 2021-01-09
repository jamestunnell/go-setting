package constraint

import "github.com/jamestunnell/go-setting/value"

// Type is used to indicate option type.
type Type int

const (
	// TypeGreater indicates a minimum (non-inclusive) value
	TypeGreater Type = iota
	// TypeGreaterEqual indicates a minimum value
	TypeGreaterEqual
	// TypeLess indicates a maximum (non-inclusive) value
	TypeLess
	// TypeLessEqual indicates a maximum (non-inclusive) value
	TypeLessEqual
	// TypeOneOf indicates an enumerated value
	TypeOneOf
	// TypeMinLen indicates a minimum length for array value types
	TypeMinLen
	// TypeMaxLen indicates a maximum length for array value types
	TypeMaxLen

	// DefaultStr represents an optional default value
	DefaultStr = "default"
	// MinLenStr represents a minimum length for array value types
	MinLenStr = "minLen"
	// MaxLenStr represents a maximum length for array value types
	MaxLenStr = "maxLen"
	// GreaterStr represents a minimum (non-inclusive) value
	GreaterStr = "greater"
	// GreaterEqualStr represents a minimum value
	GreaterEqualStr = "greaterEqual"
	// LessStr represents a maximum (non-inclusive) value
	LessStr = "less"
	// LessEqualStr represents a maximum (non-inclusive) value
	LessEqualStr = "lessEqual"
	// OneOfStr represents an enumerated value
	OneOfStr = "oneOf"
)

// AllTypes returns all of the option types.
func AllTypes() []Type {
	return []Type{
		TypeGreater, TypeGreaterEqual, TypeLess, TypeLessEqual, TypeOneOf, TypeMinLen, TypeMaxLen}
}

// Valid returns if the current type is one of AllTypes
func (t Type) Valid() bool {
	for _, typ := range AllTypes() {
		if t == typ {
			return true
		}
	}

	return false
}

// String returns a string representation of the option type.
func (t Type) String() string {
	switch t {
	case TypeMinLen:
		return MinLenStr
	case TypeMaxLen:
		return MaxLenStr
	case TypeGreater:
		return GreaterStr
	case TypeGreaterEqual:
		return GreaterEqualStr
	case TypeLess:
		return LessStr
	case TypeLessEqual:
		return LessEqualStr
	case TypeOneOf:
		return OneOfStr
	}

	return ""
}

// ApplicableTo returns true if the current option is applicable to the given value type.
func (t Type) ApplicableTo(v value.Value) bool {
	switch t {
	case TypeGreater, TypeGreaterEqual, TypeLess, TypeLessEqual:
		return true
	case TypeMinLen, TypeMaxLen:
		return v.IsSlice() || (v.Type() == value.TypeString)
	case TypeOneOf:
		return !v.IsSlice()
	}

	return false
}
