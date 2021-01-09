package value

import (
	"strconv"
	"strings"
)

type compareIntFunc func(a, b int64) bool

// IntSlice holds a slice of int64 values
type IntSlice struct {
	valsPtr *[]int64
}

// NewIntSlice makes a new IntSlice with the given int64 values.
func NewIntSlice(vals ...int64) *IntSlice {
	slice := make([]int64, len(vals))

	copy(slice, vals)

	return &IntSlice{valsPtr: &slice}
}

// NewIntSliceFromPtr makes a new IntSlice with the given pointer to int64 values.
func NewIntSliceFromPtr(valsPtr *[]int64) *IntSlice {
	return &IntSlice{valsPtr: valsPtr}
}

// Set changes the int64 values.
func (v *IntSlice) Set(vals []int64) { *v.valsPtr = vals }

// Type return TypeInt.
func (v *IntSlice) Type() Type { return TypeInt }

// IsSlice returns true.
func (v *IntSlice) IsSlice() bool { return true }

// Clone produce a clone that is identical except for the backing pointer.
func (v *IntSlice) Clone() Value { return NewIntSlice(*v.valsPtr...) }

// Parse sets the values from the given string.
func (v *IntSlice) Parse(str string) error {
	substrings := strings.Split(str, ",")
	vals := make([]int64, len(substrings))

	for i := 0; i < len(substrings); i++ {
		substr := strings.TrimSpace(substrings[i])

		val, err := strconv.ParseInt(substr, 10, 64)
		if err != nil {
			return err
		}

		vals[i] = val
	}

	*v.valsPtr = vals

	return nil
}

// SlicePointer returns the pointer for storage of slice values.
func (v *IntSlice) SlicePointer() interface{} { return v.valsPtr }

// Slice returns the int64 slice values.
func (v *IntSlice) Slice() interface{} { return *v.valsPtr }

// Len returns the number of slice elements.
func (v *IntSlice) Len() int { return len(*v.valsPtr) }

// Equal checks if length and values of given slice equal the current.
// Returns a non-nil error if types do not match.
func (v *IntSlice) Equal(v2 Slice) (bool, error) {
	if err := CheckType(TypeInt, v2.Type()); err != nil {
		return false, err
	}

	vals1 := *v.valsPtr
	vals2 := v2.Slice().([]int64)

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
func (v *IntSlice) Greater(v2 Single) (bool, error) {
	return compareInts(*v.valsPtr, v2, intGreater)
}

// GreaterEqual checks if all values of the current slice are greater or equal
// to the given single.
// Returns a non-nil error if types do not match.
func (v *IntSlice) GreaterEqual(v2 Single) (bool, error) {
	return compareInts(*v.valsPtr, v2, intGreaterEqual)
}

// Less checks if all values of the current slice are less than that of
// the given single.
// Returns a non-nil error if types do not match.
func (v *IntSlice) Less(v2 Single) (bool, error) {
	return compareInts(*v.valsPtr, v2, intLess)
}

// LessEqual checks if all values of the current slice are less or equal
// to the given single.
// Returns a non-nil error if types do not match.
func (v *IntSlice) LessEqual(v2 Single) (bool, error) {
	return compareInts(*v.valsPtr, v2, intLessEqual)
}

// Contains checks if the given single value is equal to one of the
// current slice values.
// Returns a non-nil error if types do not match.
func (v *IntSlice) Contains(v2 Single) (bool, error) {
	if err := CheckType(TypeInt, v2.Type()); err != nil {
		return false, err
	}

	vals := *v.valsPtr
	val2 := v2.Value().(int64)

	for _, val1 := range vals {
		if val1 == val2 {
			return true, nil
		}
	}

	return false, nil
}

func compareInts(vals []int64, v2 Single, f compareIntFunc) (bool, error) {
	if err := CheckType(TypeInt, v2.Type()); err != nil {
		return false, err
	}

	if len(vals) == 0 {
		return false, nil
	}

	val2 := v2.Value().(int64)

	for _, val1 := range vals {
		if !f(val1, val2) {
			return false, nil
		}
	}
	return true, nil
}

func intGreater(a, b int64) bool {
	return a > b
}

func intGreaterEqual(a, b int64) bool {
	return a >= b
}

func intLess(a, b int64) bool {
	return a < b
}

func intLessEqual(a, b int64) bool {
	return a <= b
}
