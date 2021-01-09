package constraint_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/value"

	"github.com/jamestunnell/go-setting/constraint"
	"github.com/stretchr/testify/assert"
)

var allSingleVals = []value.Single{
	value.NewUInt(0), value.NewInt(0), value.NewFloat(0.0),
	value.NewString(""), value.NewBool(false),
}

var allSliceVals = []value.Slice{
	value.NewUIntSlice(), value.NewIntSlice(), value.NewFloatSlice(),
	value.NewStringSlice(), value.NewBoolSlice(),
}

func TestKnownTypes(t *testing.T) {
	for _, typ := range constraint.AllTypes() {
		str := typ.String()

		assert.NotEmpty(t, str)
		assert.True(t, typ.Valid())
	}
}

func TestUnknownType(t *testing.T) {
	typ := constraint.Type(-1)

	assert.Empty(t, typ.String())
	assert.False(t, typ.Valid())
}

func TestApplicableToInvalidType(t *testing.T) {
	assert.False(t, constraint.Type(-1).ApplicableTo(value.NewFloat(0.0)))
}

func TestApplicableToWithCompareConstraints(t *testing.T) {
	cTypes := []constraint.Type{
		constraint.TypeGreater, constraint.TypeGreaterEqual,
		constraint.TypeLess, constraint.TypeLessEqual,
	}

	for _, cType := range cTypes {
		for _, val := range allSingleVals {
			assert.True(t, cType.ApplicableTo(val))
		}

		for _, val := range allSliceVals {
			assert.True(t, cType.ApplicableTo(val))
		}
	}
}

func TestApplicableToWithOneOf(t *testing.T) {
	for _, val := range allSingleVals {
		assert.True(t, constraint.TypeOneOf.ApplicableTo(val))
	}

	for _, val := range allSliceVals {
		assert.False(t, constraint.TypeOneOf.ApplicableTo(val))
	}
}

func TestApplicableToWithLen(t *testing.T) {
	allSingleValsExceptString := []value.Single{
		value.NewUInt(0), value.NewInt(0), value.NewFloat(0.0), value.NewBool(false),
	}
	allSliceValsPlusSingleString := []value.Value{
		value.NewUIntSlice(), value.NewIntSlice(), value.NewFloatSlice(),
		value.NewStringSlice(), value.NewBoolSlice(), value.NewString(""),
	}

	for _, val := range allSingleValsExceptString {
		assert.False(t, constraint.TypeMinLen.ApplicableTo(val))
		assert.False(t, constraint.TypeMaxLen.ApplicableTo(val))
	}

	for _, val := range allSliceValsPlusSingleString {
		assert.True(t, constraint.TypeMinLen.ApplicableTo(val))
		assert.True(t, constraint.TypeMaxLen.ApplicableTo(val))
	}
}
