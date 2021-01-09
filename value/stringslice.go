package value

import "strings"

type compareStringFunc func(a, b string) bool

// StringSlice holds a slice of string values
type StringSlice struct {
	valsPtr *[]string
}

// NewStringSlice makes a new StringSlice with the given string values.
func NewStringSlice(vals ...string) *StringSlice {
	slice := make([]string, len(vals))

	copy(slice, vals)

	return &StringSlice{valsPtr: &slice}
}

// NewStringSliceFromPtr makes a new StringSlice with the given pointer to string values.
func NewStringSliceFromPtr(valsPtr *[]string) *StringSlice {
	return &StringSlice{valsPtr: valsPtr}
}

// Set changes the string values.
func (v *StringSlice) Set(vals []string) { *v.valsPtr = vals }

// Type return TypeString.
func (v *StringSlice) Type() Type { return TypeString }

// IsSlice returns true.
func (v *StringSlice) IsSlice() bool { return true }

// Clone produce a clone that is identical except for the backing pointer.
func (v *StringSlice) Clone() Value { return NewStringSlice(*v.valsPtr...) }

// Parse sets the values from the given string.
func (v *StringSlice) Parse(str string) error {
	substrings := strings.Split(str, ",")
	trimmed := make([]string, len(substrings))

	for i := 0; i < len(substrings); i++ {
		trimmed[i] = strings.TrimSpace(substrings[i])
	}

	*v.valsPtr = trimmed

	return nil
}

// SlicePointer returns the pointer for storage of slice values.
func (v *StringSlice) SlicePointer() interface{} { return v.valsPtr }

// Slice returns the string slice values.
func (v *StringSlice) Slice() interface{} { return *v.valsPtr }

// Len returns the number of slice elements.
func (v *StringSlice) Len() int { return len(*v.valsPtr) }

// Equal checks if length and values of given slice equal the current.
// Returns a non-nil error if types do not match.
func (v *StringSlice) Equal(v2 Slice) (bool, error) {
	if err := CheckType(TypeString, v2.Type()); err != nil {
		return false, err
	}

	vals1 := *v.valsPtr
	vals2 := v2.Slice().([]string)

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
func (v *StringSlice) Greater(v2 Single) (bool, error) {
	return compareStrings(*v.valsPtr, v2, stringGreater)
}

// GreaterEqual checks if all values of the current slice are greater or equal
// to the given single.
// Returns a non-nil error if types do not match.
func (v *StringSlice) GreaterEqual(v2 Single) (bool, error) {
	return compareStrings(*v.valsPtr, v2, stringGreaterEqual)
}

// Less checks if all values of the current slice are less than that of
// the given single.
// Returns a non-nil error if types do not match.
func (v *StringSlice) Less(v2 Single) (bool, error) {
	return compareStrings(*v.valsPtr, v2, stringLess)
}

// LessEqual checks if all values of the current slice are less or equal
// to the given single.
// Returns a non-nil error if types do not match.
func (v *StringSlice) LessEqual(v2 Single) (bool, error) {
	return compareStrings(*v.valsPtr, v2, stringLessEqual)
}

// Contains checks if the given single value is equal to one of the
// current slice values.
// Returns a non-nil error if types do not match.
func (v *StringSlice) Contains(v2 Single) (bool, error) {
	if err := CheckType(TypeString, v2.Type()); err != nil {
		return false, err
	}

	vals := *v.valsPtr
	val2 := v2.Value().(string)

	for _, val1 := range vals {
		if val1 == val2 {
			return true, nil
		}
	}

	return false, nil
}

func compareStrings(vals []string, v2 Single, f compareStringFunc) (bool, error) {
	if err := CheckType(TypeString, v2.Type()); err != nil {
		return false, err
	}

	if len(vals) == 0 {
		return false, nil
	}

	val2 := v2.Value().(string)

	for _, val1 := range vals {
		if !f(val1, val2) {
			return false, nil
		}
	}
	return true, nil
}

func stringGreater(a, b string) bool {
	return a > b
}

func stringGreaterEqual(a, b string) bool {
	return a >= b
}

func stringLess(a, b string) bool {
	return a < b
}

func stringLessEqual(a, b string) bool {
	return a <= b
}
