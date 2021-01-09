package value

import (
	"strconv"
	"strings"
)

// BoolSlice holds a slice of boolean values
type BoolSlice struct {
	valsPtr *[]bool
}

// NewBoolSlice makes a new BoolSlice with the given boolean values.
func NewBoolSlice(vals ...bool) *BoolSlice {
	slice := make([]bool, len(vals))

	copy(slice, vals)

	return &BoolSlice{valsPtr: &slice}
}

// NewBoolSliceFromPtr makes a new BoolSlice with the given pointer to boolean values.
func NewBoolSliceFromPtr(valsPtr *[]bool) *BoolSlice {
	return &BoolSlice{valsPtr: valsPtr}
}

// Set changes the boolean values.
func (v *BoolSlice) Set(vals []bool) { *v.valsPtr = vals }

// Type return TypeBool.
func (v *BoolSlice) Type() Type { return TypeBool }

// IsSlice returns true.
func (v *BoolSlice) IsSlice() bool { return true }

// Clone produce a clone that is identical except for the backing pointer.
func (v *BoolSlice) Clone() Value { return NewBoolSlice(*v.valsPtr...) }

// Parse sets the values from the given string.
func (v *BoolSlice) Parse(str string) error {
	substrings := strings.Split(str, ",")
	vals := make([]bool, len(substrings))

	for i := 0; i < len(substrings); i++ {
		trimmed := strings.TrimSpace(substrings[i])

		val, err := strconv.ParseBool(trimmed)
		if err != nil {
			return err
		}

		vals[i] = val
	}

	*v.valsPtr = vals

	return nil
}

// SlicePointer returns the pointer for storage of slice values.
func (v *BoolSlice) SlicePointer() interface{} { return v.valsPtr }

// Slice returns the boolean slice values.
func (v *BoolSlice) Slice() interface{} { return *v.valsPtr }

// Len returns the number of slice elements.
func (v *BoolSlice) Len() int { return len(*v.valsPtr) }

// Equal checks if length and values of given slice equal the current.
// Returns a non-nil error if types do not match.
func (v *BoolSlice) Equal(v2 Slice) (bool, error) {
	if err := CheckType(TypeBool, v2.Type()); err != nil {
		return false, err
	}

	vals1 := *v.valsPtr
	vals2 := v2.Slice().([]bool)

	if len(vals1) != len(vals2) {
		return false, nil
	}

	for i, val1 := range vals1 {
		if val1 != vals2[i] {
			return false, nil
		}
	}

	return true, nil
}

// Greater checks if all values of the current slice are greater than that of
// the given single.
// Returns a non-nil error if types do not match.
func (v *BoolSlice) Greater(v2 Single) (bool, error) {
	return compareBools(*v.valsPtr, v2, BoolGreater)
}

// GreaterEqual checks if all values of the current slice are greater or equal
// to the given single.
// Returns a non-nil error if types do not match.
func (v *BoolSlice) GreaterEqual(v2 Single) (bool, error) {
	return compareBools(*v.valsPtr, v2, BoolGreaterEqual)
}

// Less checks if all values of the current slice are less than that of
// the given single.
// Returns a non-nil error if types do not match.
func (v *BoolSlice) Less(v2 Single) (bool, error) {
	return compareBools(*v.valsPtr, v2, BoolLess)
}

// LessEqual checks if all values of the current slice are less or equal
// to the given single.
// Returns a non-nil error if types do not match.
func (v *BoolSlice) LessEqual(v2 Single) (bool, error) {
	return compareBools(*v.valsPtr, v2, BoolLessEqual)
}

// Contains checks if the given single value is equal to one of the
// current slice values.
// Returns a non-nil error if types do not match.
func (v *BoolSlice) Contains(v2 Single) (bool, error) {
	if err := CheckType(TypeBool, v2.Type()); err != nil {
		return false, err
	}

	vals := *v.valsPtr
	val2 := v2.(Single).Value().(bool)

	for _, val1 := range vals {
		if val1 == val2 {
			return true, nil
		}
	}

	return false, nil
}

func compareBools(vals []bool, v2 Single, f CompareBoolFunc) (bool, error) {
	if err := CheckType(TypeBool, v2.Type()); err != nil {
		return false, err
	}

	if len(vals) == 0 {
		return false, nil
	}

	val2 := v2.Value().(bool)

	for _, val1 := range vals {
		if !f(val1, val2) {
			return false, nil
		}
	}
	return true, nil
}
