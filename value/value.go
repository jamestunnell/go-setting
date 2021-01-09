package value

// Value holds a single value or slice of values that have one of a fixed
// number of types
type Value interface {
	// Type returns the value type. For a slice, this returns type of each
	// slice element.
	Type() Type
	// IsSlice returns true if it is a slice
	IsSlice() bool
	// Clone produce a clone that is identical except for the backing pointer.
	Clone() Value
	// Parse sets the value from the given string
	Parse(string) error
	// Greater returns true if the value (or all values for a slice) is greater than
	// the given single value.
	// Returns non-nil error in case of type mismatch.
	Greater(Single) (bool, error)
	// GreaterEqual returns true if the value (or all values for a slice) is greater or equal to
	// the given single value.
	// Returns non-nil error in case of type mismatch.
	GreaterEqual(Single) (bool, error)
	// Less returns true if the value (or all values for a slice) is less than
	// the given single value.
	// Returns non-nil error in case of type mismatch.
	Less(Single) (bool, error)
	// LessEqual returns true if the value (or all values for a slice) is less or equal to
	// the given single value.
	// Returns non-nil error in case of type mismatch.
	LessEqual(Single) (bool, error)
}

// Single holds a single value.
type Single interface {
	Value
	ValuePointer() interface{}
	Value() interface{}
	Equal(Single) (bool, error)
	OneOf(Slice) (bool, error)
}

// Slice holds a slice of values.
type Slice interface {
	Value
	SlicePointer() interface{}
	Slice() interface{}
	Len() int
	Equal(Slice) (bool, error)
	Contains(Single) (bool, error)
}
