package constraint

import "github.com/jamestunnell/go-setting/value"

// Greater restricts a value to be greater than the parameter
type Greater struct {
	val value.Single
}

// NewGreater makes a new Greater constraint
func NewGreater(val value.Single) *Greater {
	return &Greater{val: val}
}

// Type returns the constraint type.
func (c *Greater) Type() Type { return TypeGreater }

// Param returns the constraint parameter.
func (c *Greater) Param() value.Value { return c.val }

// CompatibleWith returns true if the given constraint is compatible with the current one.
// Returns a non-nil error in case of failure.
func (c *Greater) CompatibleWith(c2 Constraint) (bool, error) {
	switch c2.Type() {
	case TypeGreater, TypeGreaterEqual, TypeOneOf:
		return false, nil
	case TypeLess, TypeLessEqual:
		return c.val.Less(c2.Param().(value.Single))
	}

	return true, nil
}
