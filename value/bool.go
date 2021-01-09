package value

import "strconv"

// Bool holds a single boolean value.
type Bool struct {
	valPtr *bool
}

// NewBool makes a new Bool with the given boolean value.
func NewBool(val bool) *Bool {
	valPtr := new(bool)
	*valPtr = val

	return &Bool{valPtr: valPtr}
}

// NewBoolFromPtr makes a new Bool with the given pointer to boolean value.
func NewBoolFromPtr(valPtr *bool) *Bool {
	return &Bool{valPtr: valPtr}
}

// Set changes the boolean value.
func (v *Bool) Set(val bool) { *v.valPtr = val }

// Type return TypeBool.
func (v *Bool) Type() Type { return TypeBool }

// IsSlice returns false.
func (v *Bool) IsSlice() bool { return false }

// Clone produce a clone that is identical except for the backing pointer.
func (v *Bool) Clone() Value { return NewBool(*v.valPtr) }

// Parse sets the value from the given string.
func (v *Bool) Parse(str string) error {
	b, err := strconv.ParseBool(str)

	if err != nil {
		return err
	}

	*v.valPtr = b

	return nil
}

// ValuePointer returns the pointer for value storage.
func (v *Bool) ValuePointer() interface{} { return v.valPtr }

// Value returns the boolean value.
func (v *Bool) Value() interface{} { return *v.valPtr }

// Equal returns checks if type and value of the given single are equal.
func (v *Bool) Equal(v2 Single) (bool, error) {
	return compareBool(*v.valPtr, v2, BoolEqual)
}

// Greater checks if the current value is greater than the given.
// Returns non-nil error if types do not match.
func (v *Bool) Greater(v2 Single) (bool, error) {
	return compareBool(*v.valPtr, v2, BoolGreater)
}

// GreaterEqual checks if the current value is greater or equal to the given.
// Returns non-nil error if types do not match.
func (v *Bool) GreaterEqual(v2 Single) (bool, error) {
	return compareBool(*v.valPtr, v2, BoolGreaterEqual)
}

// Less checks if the current value is less than the given.
// Returns non-nil error if types do not match.
func (v *Bool) Less(v2 Single) (bool, error) {
	return compareBool(*v.valPtr, v2, BoolLess)
}

// LessEqual checks if the current value is less or equal to the given.
// Returns non-nil error if types do not match.
func (v *Bool) LessEqual(v2 Single) (bool, error) {
	return compareBool(*v.valPtr, v2, BoolLessEqual)
}

// OneOf checks if the current value is one of the given.
// Returns non-nil error if types do not match.
func (v *Bool) OneOf(v2 Slice) (bool, error) { return v2.Contains(v) }

func compareBool(val bool, v2 Single, f CompareBoolFunc) (bool, error) {
	if err := CheckType(TypeBool, v2.Type()); err != nil {
		return false, err
	}

	return f(val, v2.Value().(bool)), nil
}
