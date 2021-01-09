package value

import "strconv"

// UInt holds a single uint64 value.
type UInt struct {
	valPtr *uint64
}

// NewUInt makes a new UInt with the given uint64 value.
func NewUInt(val uint64) *UInt {
	valPtr := new(uint64)
	*valPtr = val

	return &UInt{valPtr: valPtr}
}

// NewUIntFromPtr makes a new UInt with the given pointer to uint64 value.
func NewUIntFromPtr(valPtr *uint64) *UInt {
	return &UInt{valPtr: valPtr}
}

// Set changes the uint64 value.
func (v *UInt) Set(val uint64) { *v.valPtr = val }

// Type return TypeUInt.
func (v *UInt) Type() Type { return TypeUInt }

// IsSlice returns false.
func (v *UInt) IsSlice() bool { return false }

// Clone produce a clone that is identical except for the backing pointer.
func (v *UInt) Clone() Value { return NewUInt(*v.valPtr) }

// Parse sets the value from the given string.
func (v *UInt) Parse(str string) error {
	u, err := strconv.ParseUint(str, 10, 64)

	if err != nil {
		return err
	}

	*v.valPtr = u

	return nil
}

// ValuePointer returns the pointer for value storage.
func (v *UInt) ValuePointer() interface{} { return v.valPtr }

// Value returns the uint64 value.
func (v *UInt) Value() interface{} { return *v.valPtr }

// Equal returns checks if type and value of the given single are equal.
func (v *UInt) Equal(v2 Single) (bool, error) {
	if err := CheckType(TypeUInt, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr == v2.Value().(uint64), nil
}

// Greater checks if the current value is greater than the given.
// Returns non-nil error if types do not match.
func (v *UInt) Greater(v2 Single) (bool, error) {
	if err := CheckType(TypeUInt, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr > v2.Value().(uint64), nil
}

// GreaterEqual checks if the current value is greater or equal to the given.
// Returns non-nil error if types do not match.
func (v *UInt) GreaterEqual(v2 Single) (bool, error) {
	if err := CheckType(TypeUInt, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr >= v2.Value().(uint64), nil
}

// Less checks if the current value is less than the given.
// Returns non-nil error if types do not match.
func (v *UInt) Less(v2 Single) (bool, error) {
	if err := CheckType(TypeUInt, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr < v2.Value().(uint64), nil
}

// LessEqual checks if the current value is less or equal to the given.
// Returns non-nil error if types do not match.
func (v *UInt) LessEqual(v2 Single) (bool, error) {
	if err := CheckType(TypeUInt, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr <= v2.Value().(uint64), nil
}

// OneOf checks if the current value is one of the given.
// Returns non-nil error if types do not match.
func (v *UInt) OneOf(v2 Slice) (bool, error) {
	return v2.Contains(v)
}
