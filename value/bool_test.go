package value_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

type compareValueFunc func(value.Single) (bool, error)

func TestBool(t *testing.T) {
	v := value.NewBool(false)

	assert.Equal(t, value.TypeBool, v.Type())
	assert.False(t, v.IsSlice())
	assert.Equal(t, false, v.Value())

	ptr := v.ValuePointer().(*bool)
	*ptr = true

	assert.Equal(t, true, v.Value())

	v.Set(false)

	assert.Equal(t, false, v.Value())
}

func TestBoolFromPtr(t *testing.T) {
	val := false
	v := value.NewBoolFromPtr(&val)

	assert.Equal(t, false, v.Value())

	val = true

	assert.Equal(t, true, v.Value())

	ptr := v.ValuePointer().(*bool)

	*ptr = false

	assert.Equal(t, false, v.Value())
}

func TestBoolOperations(t *testing.T) {
	v := value.NewBool(false)
	vF := value.NewBool(false)
	vT := value.NewBool(true)

	verifyCompare(t, v.Equal, vF, true)
	verifyCompare(t, v.Equal, vT, false)
	verifyCompare(t, v.Greater, vF, false)
	verifyCompare(t, v.Greater, vT, false)
	verifyCompare(t, v.GreaterEqual, vF, true)
	verifyCompare(t, v.GreaterEqual, vT, false)
	verifyCompare(t, v.Less, vF, false)
	verifyCompare(t, v.Less, vT, true)
	verifyCompare(t, v.LessEqual, vF, true)
	verifyCompare(t, v.LessEqual, vT, true)

	v.Set(true)

	verifyCompare(t, v.Equal, vF, false)
	verifyCompare(t, v.Equal, vT, true)
	verifyCompare(t, v.Greater, vF, true)
	verifyCompare(t, v.Greater, vT, false)
	verifyCompare(t, v.GreaterEqual, vF, true)
	verifyCompare(t, v.GreaterEqual, vT, true)
	verifyCompare(t, v.Less, vF, false)
	verifyCompare(t, v.Less, vT, false)
	verifyCompare(t, v.LessEqual, vF, false)
	verifyCompare(t, v.LessEqual, vT, true)
}

func TestBoolOperationsWrongType(t *testing.T) {
	v := value.NewBool(false)
	v2 := value.NewFloat(0.0)

	verifyCompareWrongType(t, v.Equal, v2)
	verifyCompareWrongType(t, v.Greater, v2)
	verifyCompareWrongType(t, v.GreaterEqual, v2)
	verifyCompareWrongType(t, v.Less, v2)
	verifyCompareWrongType(t, v.LessEqual, v2)
}

func TestBoolOneOf(t *testing.T) {
	testBoolOneOf(t, false, []bool{true}, false)
	testBoolOneOf(t, false, []bool{false}, true)
	testBoolOneOf(t, false, []bool{true, false}, true)
	testBoolOneOf(t, true, []bool{true, false}, true)
	testBoolOneOf(t, true, []bool{true}, true)
	testBoolOneOf(t, true, []bool{false}, false)
}

func TestBoolClone(t *testing.T) {
	const val = true

	v1 := value.NewBool(val)
	v2 := v1.Clone()

	assert.Equal(t, v1.Value(), v2.(value.Single).Value())

	// make sure the values are independent
	v1.Set(false)

	assert.Equal(t, val, v2.(value.Single).Value())
}

func TestBoolParse(t *testing.T) {
	v := value.NewBool(false)

	assert.Error(t, v.Parse("2.5"))
	assert.NoError(t, v.Parse("true"))
	assert.Equal(t, true, v.Value())
}

func testBoolOneOf(t *testing.T, bVal bool, bVals []bool, expected bool) {
	v := value.NewBool(bVal)
	s := value.NewBoolSlice(bVals...)

	result, err := v.OneOf(s)

	if assert.NoError(t, err) {
		assert.Equal(t, expected, result)
	}
}

func verifyCompare(t *testing.T, f compareValueFunc, val2 value.Single, expected interface{}) {
	result, err := f(val2)

	if assert.NoError(t, err) {
		assert.Equal(t, expected, result)
	}
}

func verifyCompareWrongType(t *testing.T, f compareValueFunc, val2 value.Single) {
	_, err := f(val2)

	assert.Error(t, err)
}
