package constraint_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/constraint"
	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestMaxLen(t *testing.T) {
	c := constraint.NewMaxLen(7)

	assert.Equal(t, c.Type(), constraint.TypeMaxLen)

	compatible := []constraint.Constraint{
		constraint.NewMinLen(7),
		constraint.NewMinLen(6),
		constraint.NewGreater(value.NewFloat(0.0)),
		constraint.NewGreaterEqual(value.NewFloat(0.0)),
		constraint.NewLess(value.NewFloat(0.0)),
		constraint.NewLessEqual(value.NewFloat(0.0)),
	}
	incompatible := []constraint.Constraint{
		constraint.NewMinLen(8),
		constraint.NewMaxLen(8),
		constraint.NewOneOf(value.NewFloatSlice(0.0)),
	}

	for _, c2 := range compatible {
		result, err := c.CompatibleWith(c2)
		assert.NoError(t, err)
		assert.True(t, result)
	}

	for _, c2 := range incompatible {
		result, err := c.CompatibleWith(c2)
		assert.NoError(t, err)
		assert.False(t, result)
	}
}
