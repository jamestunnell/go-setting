package constraint

import "github.com/jamestunnell/go-setting/value"

// Less restricts a value to be less than the parameter
type Less struct {
	val value.Single
}

// NewLess makes a new Less constraint
func NewLess(val value.Single) *Less {
	return &Less{val: val}
}

// Type returns the constraint type.
func (c *Less) Type() Type { return TypeLess }

// Param returns the constraint parameter.
func (c *Less) Param() value.Value { return c.val }

// CompatibleWith returns true if the given constraint is compatible with the current one.
// Returns a non-nil error in case of failure.
func (c *Less) CompatibleWith(c2 Constraint) (bool, error) {
	switch c2.Type() {
	case TypeLess, TypeLessEqual, TypeOneOf:
		return false, nil
	case TypeGreater, TypeGreaterEqual:
		return c.val.Greater(c2.Param().(value.Single))
	}

	return true, nil
}
