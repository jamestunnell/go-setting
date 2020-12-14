package option_test

import (
	"testing"

	"github.com/jamestunnell/go-settings/option"
	"github.com/jamestunnell/go-settings/value"
	"github.com/stretchr/testify/assert"
)

func TestApplicableToNumber(t *testing.T) {
	testApplicableToNumber(t, value.Int64)
	testApplicableToNumber(t, value.UInt64)
	testApplicableToNumber(t, value.Float64)
}

func TestApplicableToNumberSlice(t *testing.T) {
	testApplicableToNumberSlice(t, value.Int64s)
	testApplicableToNumberSlice(t, value.UInt64s)
	testApplicableToNumberSlice(t, value.Float64s)
}

func TestApplicableToString(t *testing.T) {
	assert.True(t, option.Name.ApplicableTo(value.String))
	assert.True(t, option.Default.ApplicableTo(value.String))
	assert.True(t, option.OneOf.ApplicableTo(value.String))
	assert.False(t, option.Greater.ApplicableTo(value.String))
	assert.False(t, option.GreaterEqual.ApplicableTo(value.String))
	assert.False(t, option.Less.ApplicableTo(value.String))
	assert.False(t, option.LessEqual.ApplicableTo(value.String))
	assert.True(t, option.MinLen.ApplicableTo(value.String))
	assert.True(t, option.MaxLen.ApplicableTo(value.String))
}

func TestApplicableToBool(t *testing.T) {
	assert.True(t, option.Name.ApplicableTo(value.Bool))
	assert.True(t, option.Default.ApplicableTo(value.Bool))
	assert.False(t, option.OneOf.ApplicableTo(value.Bool))
	assert.False(t, option.Greater.ApplicableTo(value.Bool))
	assert.False(t, option.GreaterEqual.ApplicableTo(value.Bool))
	assert.False(t, option.Less.ApplicableTo(value.Bool))
	assert.False(t, option.LessEqual.ApplicableTo(value.Bool))
	assert.False(t, option.MinLen.ApplicableTo(value.Bool))
	assert.False(t, option.MaxLen.ApplicableTo(value.Bool))
}

func TestApplicableToNonNumberSlice(t *testing.T) {
	testApplicableToNonNumberSlice(t, value.Strings)
	testApplicableToNonNumberSlice(t, value.Bools)
}

func testApplicableToNumber(t *testing.T, valType value.Type) {
	assert.True(t, option.Name.ApplicableTo(valType))
	assert.True(t, option.Default.ApplicableTo(valType))
	assert.True(t, option.OneOf.ApplicableTo(valType))
	assert.True(t, option.Greater.ApplicableTo(valType))
	assert.True(t, option.GreaterEqual.ApplicableTo(valType))
	assert.True(t, option.Less.ApplicableTo(valType))
	assert.True(t, option.LessEqual.ApplicableTo(valType))
	assert.False(t, option.MinLen.ApplicableTo(valType))
	assert.False(t, option.MaxLen.ApplicableTo(valType))
}

func testApplicableToNumberSlice(t *testing.T, valType value.Type) {
	assert.True(t, option.Name.ApplicableTo(valType))
	assert.True(t, option.Default.ApplicableTo(valType))
	assert.False(t, option.OneOf.ApplicableTo(valType))
	assert.True(t, option.Greater.ApplicableTo(valType))
	assert.True(t, option.GreaterEqual.ApplicableTo(valType))
	assert.True(t, option.Less.ApplicableTo(valType))
	assert.True(t, option.LessEqual.ApplicableTo(valType))
	assert.True(t, option.MinLen.ApplicableTo(valType))
	assert.True(t, option.MaxLen.ApplicableTo(valType))
}

func testApplicableToNonNumberSlice(t *testing.T, valType value.Type) {
	assert.True(t, option.Name.ApplicableTo(valType))
	assert.True(t, option.Default.ApplicableTo(valType))
	assert.False(t, option.OneOf.ApplicableTo(valType))
	assert.False(t, option.Greater.ApplicableTo(valType))
	assert.False(t, option.GreaterEqual.ApplicableTo(valType))
	assert.False(t, option.Less.ApplicableTo(valType))
	assert.False(t, option.LessEqual.ApplicableTo(valType))
	assert.True(t, option.MinLen.ApplicableTo(valType))
	assert.True(t, option.MaxLen.ApplicableTo(valType))
}
