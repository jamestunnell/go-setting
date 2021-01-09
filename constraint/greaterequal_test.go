package constraint_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/constraint"
	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestGreaterEqual(t *testing.T) {
	c := constraint.NewGreaterEqual(value.NewFloat(2.5))

	assert.Equal(t, c.Type(), constraint.TypeGreaterEqual)

	compatible := []constraint.Constraint{
		constraint.NewLess(value.NewFloat(3.0)),
		constraint.NewLessEqual(value.NewFloat(2.5)),
		constraint.NewMaxLen(0),
		constraint.NewMinLen(0),
	}
	incompatible := []constraint.Constraint{
		constraint.NewOneOf(value.NewFloatSlice()),
		constraint.NewGreater(value.NewFloat(0.0)),
		constraint.NewGreaterEqual(value.NewFloat(0.0)),
		constraint.NewLess(value.NewFloat(2.5)),
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
