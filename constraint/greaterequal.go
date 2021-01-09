package constraint

import "github.com/jamestunnell/go-setting/value"

// GreaterEqual restricts a value to be greater than or equal to the parameter
type GreaterEqual struct {
	val value.Single
}

// NewGreaterEqual makes a new GreaterEqual constraint
func NewGreaterEqual(val value.Single) *GreaterEqual {
	return &GreaterEqual{val: val}
}

// Type returns the constraint type.
func (c *GreaterEqual) Type() Type { return TypeGreaterEqual }

// Param returns the constraint parameter.
func (c *GreaterEqual) Param() value.Value { return c.val }

// CompatibleWith returns true if the given constraint is compatible with the current one.
// Returns a non-nil error in case of failure.
func (c *GreaterEqual) CompatibleWith(c2 Constraint) (bool, error) {
	switch c2.Type() {
	case TypeGreater, TypeGreaterEqual, TypeOneOf:
		return false, nil
	case TypeLess:
		return c.val.Less(c2.Param().(value.Single))
	case TypeLessEqual:
		return c.val.LessEqual(c2.Param().(value.Single))
	}

	return true, nil
}
