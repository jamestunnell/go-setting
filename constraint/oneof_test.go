package constraint_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/constraint"
	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestOneOf(t *testing.T) {
	c := constraint.NewOneOf(value.NewFloatSlice(3.3, 2.5))
	p := c.Param()

	assert.Equal(t, c.Type(), constraint.TypeOneOf)
	assert.Equal(t, p.Type(), value.TypeFloat)
	assert.True(t, p.IsSlice())

	compatible := []constraint.Constraint{
		constraint.NewMaxLen(0),
		constraint.NewMinLen(0),
	}
	incompatible := []constraint.Constraint{
		constraint.NewOneOf(value.NewFloatSlice(3.3)),
		constraint.NewGreater(value.NewFloat(0.0)),
		constraint.NewGreaterEqual(value.NewFloat(0.0)),
		constraint.NewLess(value.NewFloat(0.0)),
		constraint.NewLessEqual(value.NewFloat(0.0)),
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
