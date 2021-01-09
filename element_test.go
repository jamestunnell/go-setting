package setting_test

import (
	"testing"

	"github.com/jamestunnell/go-setting"
	"github.com/jamestunnell/go-setting/constraint"
	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestNewNoOptions(t *testing.T) {
	e := setting.NewElement(value.NewFloatSlice(1.5))

	assert.Equal(t, value.TypeFloat, e.Value.Type())
	assert.True(t, e.Value.IsSlice())
	assert.Len(t, e.Constraints, 0)
	assert.Nil(t, e.Constraint(constraint.TypeGreater))
}

func TestNewWithCompatibleConstraints(t *testing.T) {
	startVal := value.NewFloat(1.0)
	lt := constraint.NewLess(value.NewFloat(10.0))
	ge := constraint.NewGreaterEqual(value.NewFloat(1.0))
	e := setting.NewElement(startVal, lt, ge)

	assert.NoError(t, e.CheckConstraints())
	assert.Len(t, e.Constraints, 2)
	assert.NotNil(t, e.Constraint(constraint.TypeLess))
	assert.NotNil(t, e.Constraint(constraint.TypeGreaterEqual))
}

func TestElementWithIncompatibleConstraints(t *testing.T) {
	startVal := value.NewFloat(1.0)
	lt := constraint.NewLess(value.NewFloat(5.0))
	le := constraint.NewLessEqual(value.NewFloat(5.0))
	e := setting.NewElement(startVal, lt, le)

	assert.Error(t, e.CheckConstraints())
}

func TestElementWithInapplicableConstraints(t *testing.T) {
	startVal := value.NewFloat(1.0)
	minlen := constraint.NewMinLen(5)
	e := setting.NewElement(startVal, minlen)

	assert.Error(t, e.CheckConstraints())
}
