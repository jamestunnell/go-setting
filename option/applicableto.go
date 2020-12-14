package option

import "github.com/jamestunnell/go-settings/value"

// ApplicableTo returns true if the current option is applicable to the given value type.
func (t Type) ApplicableTo(valType value.Type) bool {
	switch t {
	case Name, Default:
		return true
	case MinLen, MaxLen:
		switch valType {
		case value.Int64s, value.UInt64s, value.Float64s, value.String, value.Strings, value.Bools:
			return true
		}
	case Greater, GreaterEqual, Less, LessEqual:
		switch valType {
		case value.Int64, value.Int64s, value.UInt64, value.UInt64s, value.Float64, value.Float64s:
			return true
		}
	case OneOf:
		switch valType {
		case value.Int64, value.UInt64, value.Float64, value.String:
			return true
		}
	}

	return false
}
