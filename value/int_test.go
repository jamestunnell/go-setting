package value_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestIntValue(t *testing.T) {
	v := value.NewInt(0)

	assert.Equal(t, value.TypeInt, v.Type())
	assert.False(t, v.IsSlice())
	assert.Equal(t, int64(0), v.Value())

	ptr := v.ValuePointer().(*int64)
	*ptr = int64(25)

	assert.Equal(t, int64(25), v.Value())

	v.Set(315)

	assert.Equal(t, int64(315), v.Value())
}

func TestIntFromPtr(t *testing.T) {
	val := int64(2)
	v := value.NewIntFromPtr(&val)

	assert.Equal(t, int64(2), v.Value())

	val = 3

	assert.Equal(t, int64(3), v.Value())

	ptr := v.ValuePointer().(*int64)

	*ptr = 4

	assert.Equal(t, int64(4), v.Value())
}

func TestIntOperations(t *testing.T) {
	v := value.NewInt(37)
	vEq := value.NewInt(37)
	vLt := value.NewInt(36)
	vGt := value.NewInt(38)

	verifyCompares(t, v, vEq, vLt, vGt)
}

func TestIntOperationsWrongType(t *testing.T) {
	v := value.NewInt(0)
	v2 := value.NewFloat(0.0)

	verifyCompareWrongType(t, v.Equal, v2)
	verifyCompareWrongType(t, v.Greater, v2)
	verifyCompareWrongType(t, v.GreaterEqual, v2)
	verifyCompareWrongType(t, v.Less, v2)
	verifyCompareWrongType(t, v.LessEqual, v2)
}

func TestIntOneOf(t *testing.T) {
	testIntOneOf(t, 0, []int64{}, false)
	testIntOneOf(t, 0, []int64{0}, true)
	testIntOneOf(t, 0, []int64{1}, false)
	testIntOneOf(t, 0, []int64{1, 0}, true)
}

func TestIntClone(t *testing.T) {
	const val = int64(-4)

	v1 := value.NewInt(val)
	v2 := v1.Clone()

	assert.Equal(t, v1.Value(), v2.(value.Single).Value())

	// make sure the values are independent
	v1.Set(4)

	assert.Equal(t, val, v2.(value.Single).Value())
}

func TestIntParse(t *testing.T) {
	v := value.NewInt(0)

	assert.Error(t, v.Parse("2.5"))
	assert.NoError(t, v.Parse("-7"))
	assert.Equal(t, int64(-7), v.Value())
}

func testIntOneOf(t *testing.T, iVal int64, iVals []int64, expected bool) {
	v := value.NewInt(0)
	s := value.NewIntSlice(0)

	v.Set(iVal)
	s.Set(iVals)

	result, err := v.OneOf(s)

	if assert.NoError(t, err) {
		assert.Equal(t, expected, result)
	}
}
