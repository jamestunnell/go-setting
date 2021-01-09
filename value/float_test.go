package value_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestFloat(t *testing.T) {
	v := value.NewFloat(0.0)

	assert.Equal(t, value.TypeFloat, v.Type())
	assert.False(t, v.IsSlice())
	assert.Equal(t, float64(0), v.Value())

	ptr := v.ValuePointer().(*float64)
	*ptr = 25.2

	assert.Equal(t, 25.2, v.Value())

	v.Set(31.5)

	assert.Equal(t, 31.5, v.Value())
}

func TestFloatFromPtr(t *testing.T) {
	val := 2.0
	v := value.NewFloatFromPtr(&val)

	assert.Equal(t, 2.0, v.Value())

	val = 2.5

	assert.Equal(t, 2.5, v.Value())

	ptr := v.ValuePointer().(*float64)

	*ptr = 3.0

	assert.Equal(t, 3.0, v.Value())
}

func TestFloatOperations(t *testing.T) {
	v := value.NewFloat(3.7)
	vEq := value.NewFloat(3.7)
	vLt := value.NewFloat(3.6)
	vGt := value.NewFloat(3.8)

	verifyCompares(t, v, vEq, vLt, vGt)
}

func TestFloatOperationsWrongType(t *testing.T) {
	v := value.NewFloat(0.0)
	v2 := value.NewBool(false)

	verifyCompareWrongType(t, v.Equal, v2)
	verifyCompareWrongType(t, v.Greater, v2)
	verifyCompareWrongType(t, v.GreaterEqual, v2)
	verifyCompareWrongType(t, v.Less, v2)
	verifyCompareWrongType(t, v.LessEqual, v2)
}

func verifyCompares(t *testing.T, v, vEq, vLt, vGt value.Single) {
	verifyCompare(t, v.Equal, vEq, true)
	verifyCompare(t, v.Equal, vLt, false)
	verifyCompare(t, v.Equal, vGt, false)
	verifyCompare(t, v.Greater, vEq, false)
	verifyCompare(t, v.Greater, vLt, true)
	verifyCompare(t, v.Greater, vGt, false)
	verifyCompare(t, v.GreaterEqual, vEq, true)
	verifyCompare(t, v.GreaterEqual, vLt, true)
	verifyCompare(t, v.GreaterEqual, vGt, false)
	verifyCompare(t, v.Less, vEq, false)
	verifyCompare(t, v.Less, vLt, false)
	verifyCompare(t, v.Less, vGt, true)
	verifyCompare(t, v.LessEqual, vEq, true)
	verifyCompare(t, v.LessEqual, vLt, false)
	verifyCompare(t, v.LessEqual, vGt, true)
}

func TestFloatOneOf(t *testing.T) {
	testFloatOneOf(t, 0.0, []float64{}, false)
	testFloatOneOf(t, 0.0, []float64{0.0}, true)
	testFloatOneOf(t, 0.0, []float64{1.1}, false)
	testFloatOneOf(t, 0.0, []float64{1.1, 0.0}, true)
}

func TestFloatClone(t *testing.T) {
	const val = float64(2.9)

	v1 := value.NewFloat(val)
	v2 := v1.Clone()

	assert.Equal(t, v1.Value(), v2.(value.Single).Value())

	// make sure the values are independent
	v1.Set(4)

	assert.Equal(t, val, v2.(value.Single).Value())
}

func TestFloatParse(t *testing.T) {
	v := value.NewFloat(0.0)

	assert.Error(t, v.Parse("abc"))
	assert.NoError(t, v.Parse("2.5e-1"))
	assert.Equal(t, float64(0.25), v.Value())
}

func testFloatOneOf(t *testing.T, fVal float64, fVals []float64, expected bool) {
	v := value.NewFloat(fVal)
	s := value.NewFloatSlice(fVals...)

	result, err := v.OneOf(s)

	if assert.NoError(t, err) {
		assert.Equal(t, expected, result)
	}
}
