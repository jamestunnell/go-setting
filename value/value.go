package value

// Type is the value data type
type Type int

const (
	// Int64 indicates int64 value
	Int64 Type = iota
	// Int64s indicates []int64 value
	Int64s
	// UInt64 indicates uint64 value
	UInt64
	// UInt64s indicates []uint64 value
	UInt64s
	// Float64 indicates float64 value
	Float64
	// Float64s indicates []float64 value
	Float64s
	// Bool indicates bool value
	Bool
	// Bools indicates []bool value
	Bools
	// String indicates string value
	String
	// Strings indicates []string value
	Strings
)

// AllTypes returns all of the value types.
func AllTypes() []Type {
	return []Type{
		Int64, Int64s, UInt64, UInt64s, Float64, Float64s, Bool, Bools, String, Strings}
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
	case Int64:
		return "int64"
	case Int64s:
		return "[]int64"
	case UInt64:
		return "uint64"
	case UInt64s:
		return "[]uint64"
	case Float64:
		return "float64"
	case Float64s:
		return "[]float64"
	case Bool:
		return "bool"
	case Bools:
		return "[]bool"
	case String:
		return "[]string"
	case Strings:
		return "[]strings"
	}

	return ""
}
