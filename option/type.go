package option

// Type is used to indicate option type.
type Type int

const (
	// Default indicates an optional default value
	Default Type = iota
	// MinLen indicates a minimum length for array value types
	MinLen
	// MaxLen indicates a maximum length for array value types
	MaxLen
	// Greater indicates a minimum (non-inclusive) value
	Greater
	// GreaterEqual indicates a minimum value
	GreaterEqual
	// Less indicates a maximum (non-inclusive) value
	Less
	// LessEqual indicates a maximum (non-inclusive) value
	LessEqual
	// OneOf indicates an enumerated value
	OneOf

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
		Default, MinLen, MaxLen, Greater, GreaterEqual, Less, LessEqual, OneOf}
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
	case Default:
		return DefaultStr
	case MinLen:
		return MinLenStr
	case MaxLen:
		return MaxLenStr
	case Greater:
		return GreaterStr
	case GreaterEqual:
		return GreaterEqualStr
	case Less:
		return LessStr
	case LessEqual:
		return LessEqualStr
	case OneOf:
		return OneOfStr
	}

	return ""
}
