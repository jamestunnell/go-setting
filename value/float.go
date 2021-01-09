package value

import "strconv"

// Float holds a single float64 value.
type Float struct {
	valPtr *float64
}

// NewFloat makes a new Float with the given float64 value.
func NewFloat(val float64) *Float {
	valPtr := new(float64)
	*valPtr = val

	return &Float{valPtr: valPtr}
}

// NewFloatFromPtr makes a new Float with the given pointer to float64 value.
func NewFloatFromPtr(valPtr *float64) *Float {
	return &Float{valPtr: valPtr}
}

// Set changes the float64 value.
func (v *Float) Set(val float64) { *v.valPtr = val }

// Type return TypeFloat.
func (v *Float) Type() Type { return TypeFloat }

// IsSlice returns false.
func (v *Float) IsSlice() bool { return false }

// Clone produce a clone that is identical except for the backing pointer.
func (v *Float) Clone() Value { return NewFloat(*v.valPtr) }

// Parse sets the value from the given string.
func (v *Float) Parse(str string) error {
	f, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return err
	}

	*v.valPtr = f

	return nil
}

// ValuePointer returns the pointer for value storage.
func (v *Float) ValuePointer() interface{} { return v.valPtr }

// Value returns the float64 value.
func (v *Float) Value() interface{} { return *v.valPtr }

// Equal returns checks if type and value of the given single are equal.
func (v *Float) Equal(v2 Single) (bool, error) {
	if err := CheckType(TypeFloat, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr == v2.Value().(float64), nil
}

// Greater checks if the current value is greater than the given.
// Returns non-nil error if types do not match.
func (v *Float) Greater(v2 Single) (bool, error) {
	if err := CheckType(TypeFloat, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr > v2.Value().(float64), nil
}

// GreaterEqual checks if the current value is greater or equal to the given.
// Returns non-nil error if types do not match.
func (v *Float) GreaterEqual(v2 Single) (bool, error) {
	if err := CheckType(TypeFloat, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr >= v2.Value().(float64), nil
}

// Less checks if the current value is less than the given.
// Returns non-nil error if types do not match.
func (v *Float) Less(v2 Single) (bool, error) {
	if err := CheckType(TypeFloat, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr < v2.Value().(float64), nil
}

// LessEqual checks if the current value is less or equal to the given.
// Returns non-nil error if types do not match.
func (v *Float) LessEqual(v2 Single) (bool, error) {
	if err := CheckType(TypeFloat, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr <= v2.Value().(float64), nil
}

// OneOf checks if the current value is one of the given.
// Returns non-nil error if types do not match.
func (v *Float) OneOf(v2 Slice) (bool, error) {
	return v2.Contains(v)
}
