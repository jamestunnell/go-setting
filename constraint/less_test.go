package constraint_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/constraint"
	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestLess(t *testing.T) {
	c := constraint.NewLess(value.NewFloat(2.5))

	assert.Equal(t, c.Type(), constraint.TypeLess)

	compatible := []constraint.Constraint{
		constraint.NewGreater(value.NewFloat(2.0)),
		constraint.NewMaxLen(0),
		constraint.NewMinLen(0),
	}
	incompatible := []constraint.Constraint{
		constraint.NewOneOf(value.NewFloatSlice()),
		constraint.NewLess(value.NewFloat(5.0)),
		constraint.NewLessEqual(value.NewFloat(5.0)),
		constraint.NewGreater(value.NewFloat(2.5)),
		constraint.NewGreaterEqual(value.NewFloat(2.5)),
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
