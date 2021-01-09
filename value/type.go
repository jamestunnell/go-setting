package value

// Type is the value data type
type Type int

const (
	// TypeInt indicates int64 value
	TypeInt Type = iota
	// TypeUInt indicates uint64 value
	TypeUInt
	// TypeFloat indicates float64 value
	TypeFloat
	// TypeBool indicates bool value
	TypeBool
	// TypeString indicates string value
	TypeString
)

// AllTypes returns all of the value types.
func AllTypes() []Type {
	return []Type{TypeInt, TypeUInt, TypeFloat, TypeBool, TypeString}
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

func (t Type) String() string {
	switch t {
	case TypeInt:
		return "int64"
	case TypeUInt:
		return "uint64"
	case TypeFloat:
		return "float64"
	case TypeBool:
		return "bool"
	case TypeString:
		return "string"
	}

	return ""
}
