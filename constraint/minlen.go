package constraint

import "github.com/jamestunnell/go-setting/value"

// MinLen resticts the length of a slice or string.
type MinLen struct {
	val value.Single
}

// NewMinLen makes a new MinLen constraint
func NewMinLen(len uint64) *MinLen {
	return &MinLen{val: value.NewUInt(len)}
}

// Type returns the constraint type.
func (c *MinLen) Type() Type { return TypeMinLen }

// Param returns the constraint parameter.
func (c *MinLen) Param() value.Value { return c.val }

// CompatibleWith returns true if the given constraint is compatible with the current one.
// Returns a non-nil error in case of failure.
func (c *MinLen) CompatibleWith(c2 Constraint) (bool, error) {
	switch c2.Type() {
	case TypeMinLen, TypeOneOf:
		return false, nil
	case TypeMaxLen:
		return c.val.LessEqual(c2.Param().(value.Single))
	}

	return true, nil
}
