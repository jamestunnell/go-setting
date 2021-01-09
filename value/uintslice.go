package value

import (
	"strconv"
	"strings"
)

type compareUIntFunc func(a, b uint64) bool

// UIntSlice holds a slice of uint64 values
type UIntSlice struct {
	valsPtr *[]uint64
}

// NewUIntSlice makes a new UIntSlice with the given uint64 values.
func NewUIntSlice(vals ...uint64) *UIntSlice {
	slice := make([]uint64, len(vals))

	copy(slice, vals)

	return &UIntSlice{valsPtr: &slice}
}

// NewUIntSliceFromPtr makes a new UIntSlice with the given pointer to uint64 values.
func NewUIntSliceFromPtr(valsPtr *[]uint64) *UIntSlice {
	return &UIntSlice{valsPtr: valsPtr}
}

// Set changes the uint64 values.
func (v *UIntSlice) Set(vals []uint64) { *v.valsPtr = vals }

// Type return TypeUInt.
func (v *UIntSlice) Type() Type { return TypeUInt }

// IsSlice returns true.
func (v *UIntSlice) IsSlice() bool { return true }

// Clone produce a clone that is identical except for the backing pointer.
func (v *UIntSlice) Clone() Value { return NewUIntSlice(*v.valsPtr...) }

// Parse sets the values from the given string.
func (v *UIntSlice) Parse(str string) error {
	substrings := strings.Split(str, ",")
	vals := make([]uint64, len(substrings))

	for i := 0; i < len(substrings); i++ {
		substr := strings.TrimSpace(substrings[i])

		val, err := strconv.ParseUint(substr, 10, 64)
		if err != nil {
			return err
		}

		vals[i] = val
	}

	*v.valsPtr = vals

	return nil
}

// SlicePointer returns the pointer for storage of slice values.
func (v *UIntSlice) SlicePointer() interface{} { return v.valsPtr }

// Slice returns the uint64 slice values.
func (v *UIntSlice) Slice() interface{} { return *v.valsPtr }

// Len returns the number of slice elements.
func (v *UIntSlice) Len() int { return len(*v.valsPtr) }

// Equal checks if length and values of given slice equal the current.
// Returns a non-nil error if types do not match.
func (v *UIntSlice) Equal(v2 Slice) (bool, error) {
	if err := CheckType(TypeUInt, v2.Type()); err != nil {
		return false, err
	}

	vals1 := *v.valsPtr
	vals2 := v2.Slice().([]uint64)

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
func (v *UIntSlice) Greater(v2 Single) (bool, error) {
	return compareUInts(*v.valsPtr, v2, uintGreater)
}

// GreaterEqual checks if all values of the current slice are greater or equal
// to the given single.
// Returns a non-nil error if types do not match.
func (v *UIntSlice) GreaterEqual(v2 Single) (bool, error) {
	return compareUInts(*v.valsPtr, v2, uintGreaterEqual)
}

// Less checks if all values of the current slice are less than that of
// the given single.
// Returns a non-nil error if types do not match.
func (v *UIntSlice) Less(v2 Single) (bool, error) {
	return compareUInts(*v.valsPtr, v2, uintLess)
}

// LessEqual checks if all values of the current slice are less or equal
// to the given single.
// Returns a non-nil error if types do not match.
func (v *UIntSlice) LessEqual(v2 Single) (bool, error) {
	return compareUInts(*v.valsPtr, v2, uintLessEqual)
}

// Contains checks if the given single value is equal to one of the
// current slice values.
// Returns a non-nil error if types do not match.
func (v *UIntSlice) Contains(v2 Single) (bool, error) {
	if err := CheckType(TypeUInt, v2.Type()); err != nil {
		return false, err
	}

	vals := *v.valsPtr
	val2 := v2.Value().(uint64)

	for _, val1 := range vals {
		if val1 == val2 {
			return true, nil
		}
	}

	return false, nil
}

func compareUInts(vals []uint64, v2 Single, f compareUIntFunc) (bool, error) {
	if err := CheckType(TypeUInt, v2.Type()); err != nil {
		return false, err
	}

	if len(vals) == 0 {
		return false, nil
	}

	val2 := v2.Value().(uint64)

	for _, val1 := range vals {
		if !f(val1, val2) {
			return false, nil
		}
	}
	return true, nil
}

func uintGreater(a, b uint64) bool {
	return a > b
}

func uintGreaterEqual(a, b uint64) bool {
	return a >= b
}

func uintLess(a, b uint64) bool {
	return a < b
}

func uintLessEqual(a, b uint64) bool {
	return a <= b
}
