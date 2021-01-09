package constraint

import "github.com/jamestunnell/go-setting/value"

// MaxLen resticts the length of a slice or string.
type MaxLen struct {
	val value.Single
}

// NewMaxLen makes a new MaxLen constraint
func NewMaxLen(len uint64) *MaxLen {
	return &MaxLen{val: value.NewUInt(len)}
}

// Type returns the constraint type.
func (c *MaxLen) Type() Type { return TypeMaxLen }

// Param returns the constraint parameter.
func (c *MaxLen) Param() value.Value { return c.val }

// CompatibleWith returns true if the given constraint is compatible with the current one.
// Returns a non-nil error in case of failure.
func (c *MaxLen) CompatibleWith(c2 Constraint) (bool, error) {
	switch c2.Type() {
	case TypeMaxLen, TypeOneOf:
		return false, nil
	case TypeMinLen:
		return c.val.GreaterEqual(c2.Param().(value.Single))
	}

	return true, nil
}
