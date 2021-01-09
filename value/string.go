package value

// String holds a single string value.
type String struct {
	valPtr *string
}

// NewString makes a new String with the given string value.
func NewString(val string) *String {
	valPtr := new(string)
	*valPtr = val

	return &String{valPtr: valPtr}
}

// NewStringFromPtr makes a new String with the given pointer to string value.
func NewStringFromPtr(valPtr *string) *String {
	return &String{valPtr: valPtr}
}

// Set changes the string value.
func (v *String) Set(val string) { *v.valPtr = val }

// Type return TypeString.
func (v *String) Type() Type { return TypeString }

// IsSlice returns false.
func (v *String) IsSlice() bool { return false }

// Clone produce a clone that is identical except for the backing pointer.
func (v *String) Clone() Value { return NewString(*v.valPtr) }

// Parse sets the value from the given string.
func (v *String) Parse(str string) error {
	*v.valPtr = str

	return nil
}

// ValuePointer returns the pointer for value storage.
func (v *String) ValuePointer() interface{} { return v.valPtr }

// Value returns the string value.
func (v *String) Value() interface{} {
	return *v.valPtr
}

// Equal returns checks if type and value of the given single are equal.
func (v *String) Equal(v2 Single) (bool, error) {
	if err := CheckType(TypeString, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr == v2.Value().(string), nil
}

// Greater checks if the current value is greater than the given.
// Returns non-nil error if types do not match.
func (v *String) Greater(v2 Single) (bool, error) {
	if err := CheckType(TypeString, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr > v2.Value().(string), nil
}

// GreaterEqual checks if the current value is greater or equal to the given.
// Returns non-nil error if types do not match.
func (v *String) GreaterEqual(v2 Single) (bool, error) {
	if err := CheckType(TypeString, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr >= v2.Value().(string), nil
}

// Less checks if the current value is less than the given.
// Returns non-nil error if types do not match.
func (v *String) Less(v2 Single) (bool, error) {
	if err := CheckType(TypeString, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr < v2.Value().(string), nil
}

// LessEqual checks if the current value is less or equal to the given.
// Returns non-nil error if types do not match.
func (v *String) LessEqual(v2 Single) (bool, error) {
	if err := CheckType(TypeString, v2.Type()); err != nil {
		return false, err
	}

	return *v.valPtr <= v2.Value().(string), nil
}

// OneOf checks if the current value is one of the given.
// Returns non-nil error if types do not match.
func (v *String) OneOf(v2 Slice) (bool, error) {
	return v2.Contains(v)
}
