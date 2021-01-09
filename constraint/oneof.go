package constraint

import "github.com/jamestunnell/go-setting/value"

// OneOf is a restricts a value to one of those in the slice parameter
type OneOf struct {
	val value.Slice
}

// NewOneOf makes a new OneOf constraint
func NewOneOf(val value.Slice) *OneOf {
	return &OneOf{val: val}
}

// Type returns the constraint type.
func (c *OneOf) Type() Type { return TypeOneOf }

// Param returns the constraint parameter.
func (c *OneOf) Param() value.Value { return c.val }

// CompatibleWith returns true if the given constraint is compatible with the current one.
// Returns a non-nil error in case of failure.
func (c *OneOf) CompatibleWith(c2 Constraint) (bool, error) {
	switch c2.Type() {
	case TypeOneOf, TypeGreater, TypeGreaterEqual, TypeLess, TypeLessEqual:
		return false, nil
	}

	return true, nil
}
