package value

import "strconv"

// Int holds a single int64 value.
type Int struct {
	valPtr *int64
}

// NewInt makes a new Int with the given int64 value.
func NewInt(val int64) *Int {
	valPtr := new(int64)
	*valPtr = val

	return &Int{valPtr: valPtr}
}

// NewIntFromPtr makes a new Int with the given pointer to int64 value.
func NewIntFromPtr(valPtr *int64) *Int {
	return &Int{valPtr: valPtr}
}

// Set changes the int64 value.
func (v *Int) Set(val int64) { *v.valPtr = val }

// Type return TypeInt.
func (v *Int) Type() Type { return TypeInt }

// IsSlice returns false.
func (v *Int) IsSlice() bool { return false }

// Clone produce a clone that is identical except for the backing pointer.
func (v *Int) Clone() Value { return NewInt(*v.valPtr) }

// Parse sets the value from the given string.
func (v *Int) Parse(str string) error {
	i, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return err
	}

	*v.valPtr = i

	return nil
}

// ValuePointer returns the pointer for value storage.
func (v *Int) ValuePointer() interface{} { return v.valPtr }

// Value returns the int64 value.
func (v *Int) Value() interface{} { return *v.valPtr }

// Equal returns checks if type and value of the given single are equal.
func (v *Int) Equal(v2 Single) (bool, error) {
	if err := CheckType(TypeInt, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr == v2.Value().(int64), nil
}

// Greater checks if the current value is greater than the given.
// Returns non-nil error if types do not match.
func (v *Int) Greater(v2 Single) (bool, error) {
	if err := CheckType(TypeInt, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr > v2.Value().(int64), nil
}

// GreaterEqual checks if the current value is greater or equal to the given.
// Returns non-nil error if types do not match.
func (v *Int) GreaterEqual(v2 Single) (bool, error) {
	if err := CheckType(TypeInt, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr >= v2.Value().(int64), nil
}

// Less checks if the current value is less than the given.
// Returns non-nil error if types do not match.
func (v *Int) Less(v2 Single) (bool, error) {
	if err := CheckType(TypeInt, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr < v2.Value().(int64), nil
}

// LessEqual checks if the current value is less or equal to the given.
// Returns non-nil error if types do not match.
func (v *Int) LessEqual(v2 Single) (bool, error) {
	if err := CheckType(TypeInt, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr <= v2.Value().(int64), nil
}

// OneOf checks if the current value is one of the given.
// Returns non-nil error if types do not match.
func (v *Int) OneOf(v2 Slice) (bool, error) {
	return v2.Contains(v)
}
