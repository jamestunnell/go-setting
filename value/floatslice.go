package value

import (
	"strconv"
	"strings"
)

type compareFloatFunc func(a, b float64) bool

// FloatSlice holds a slice of float64 values
type FloatSlice struct {
	valsPtr *[]float64
}

// NewFloatSlice makes a new FloatSlice with the given float64 values.
func NewFloatSlice(vals ...float64) *FloatSlice {
	slice := make([]float64, len(vals))

	copy(slice, vals)

	return &FloatSlice{valsPtr: &slice}
}

// NewFloatSliceFromPtr makes a new FloatSlice with the given pointer to float64 values.
func NewFloatSliceFromPtr(valsPtr *[]float64) *FloatSlice {
	return &FloatSlice{valsPtr: valsPtr}
}

// Set changes the float64 values.
func (v *FloatSlice) Set(vals []float64) { *v.valsPtr = vals }

// Type return TypeFloat.
func (v *FloatSlice) Type() Type { return TypeFloat }

// IsSlice returns true.
func (v *FloatSlice) IsSlice() bool { return true }

// Clone produce a clone that is identical except for the backing pointer.
func (v *FloatSlice) Clone() Value { return NewFloatSlice(*v.valsPtr...) }

// Parse sets the values from the given string.
func (v *FloatSlice) Parse(str string) error {
	substrings := strings.Split(str, ",")
	vals := make([]float64, len(substrings))

	for i := 0; i < len(substrings); i++ {
		substr := strings.TrimSpace(substrings[i])

		val, err := strconv.ParseFloat(substr, 64)
		if err != nil {
			return err
		}

		vals[i] = val
	}

	*v.valsPtr = vals

	return nil
}

// SlicePointer returns the pointer for storage of slice values.
func (v *FloatSlice) SlicePointer() interface{} { return v.valsPtr }

// Slice returns the float64 slice values.
func (v *FloatSlice) Slice() interface{} { return *v.valsPtr }

// Len returns the number of slice elements.
func (v *FloatSlice) Len() int { return len(*v.valsPtr) }

// Equal checks if length and values of given slice equal the current.
// Returns a non-nil error if types do not match.
func (v *FloatSlice) Equal(v2 Slice) (bool, error) {
	if err := CheckType(TypeFloat, v2.Type()); err != nil {
		return false, err
	}

	vals1 := *v.valsPtr
	vals2 := v2.Slice().([]float64)

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
func (v *FloatSlice) Greater(v2 Single) (bool, error) {
	return compareFloats(*v.valsPtr, v2, floatGreater)
}

// GreaterEqual checks if all values of the current slice are greater or equal
// to the given single.
// Returns a non-nil error if types do not match.
func (v *FloatSlice) GreaterEqual(v2 Single) (bool, error) {
	return compareFloats(*v.valsPtr, v2, floatGreaterEqual)
}

// Less checks if all values of the current slice are less than that of
// the given single.
// Returns a non-nil error if types do not match.
func (v *FloatSlice) Less(v2 Single) (bool, error) {
	return compareFloats(*v.valsPtr, v2, floatLess)
}

// LessEqual checks if all values of the current slice are less or equal
// to the given single.
// Returns a non-nil error if types do not match.
func (v *FloatSlice) LessEqual(v2 Single) (bool, error) {
	return compareFloats(*v.valsPtr, v2, floatLessEqual)
}

// Contains checks if the given single value is equal to one of the
// current slice values.
// Returns a non-nil error if types do not match.
func (v *FloatSlice) Contains(v2 Single) (bool, error) {
	if err := CheckType(TypeFloat, v2.Type()); err != nil {
		return false, err
	}

	vals := *v.valsPtr
	val2 := v2.Value().(float64)

	for _, val1 := range vals {
		if val1 == val2 {
			return true, nil
		}
	}

	return false, nil
}

func compareFloats(vals []float64, v2 Single, f compareFloatFunc) (bool, error) {
	if err := CheckType(TypeFloat, v2.Type()); err != nil {
		return false, err
	}

	if len(vals) == 0 {
		return false, nil
	}

	val2 := v2.Value().(float64)

	for _, val1 := range vals {
		if !f(val1, val2) {
			return false, nil
		}
	}
	return true, nil
}

func floatGreater(a, b float64) bool {
	return a > b
}

func floatGreaterEqual(a, b float64) bool {
	return a >= b
}

func floatLess(a, b float64) bool {
	return a < b
}

func floatLessEqual(a, b float64) bool {
	return a <= b
}
