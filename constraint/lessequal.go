package constraint

import "github.com/jamestunnell/go-setting/value"

// LessEqual restricts a value to be less than or equal to the parameter
type LessEqual struct {
	val value.Single
}

// NewLessEqual makes a new LessEqual constraint
func NewLessEqual(val value.Single) *LessEqual {
	return &LessEqual{val: val}
}

// Type returns the constraint type.
func (c *LessEqual) Type() Type { return TypeLessEqual }

// Param returns the constraint parameter.
func (c *LessEqual) Param() value.Value { return c.val }

// CompatibleWith returns true if the given constraint is compatible with the current one.
// Returns a non-nil error in case of failure.
func (c *LessEqual) CompatibleWith(c2 Constraint) (bool, error) {
	switch c2.Type() {
	case TypeLess, TypeLessEqual, TypeOneOf:
		return false, nil
	case TypeGreater:
		return c.val.Greater(c2.Param().(value.Single))
	case TypeGreaterEqual:
		return c.val.GreaterEqual(c2.Param().(value.Single))
	}

	return true, nil
}
