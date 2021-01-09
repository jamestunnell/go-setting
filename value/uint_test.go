package value_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestUIntValue(t *testing.T) {
	v := value.NewUInt(0)

	assert.Equal(t, value.TypeUInt, v.Type())
	assert.False(t, v.IsSlice())
	assert.Equal(t, uint64(0), v.Value())

	ptr := v.ValuePointer().(*uint64)
	*ptr = uint64(25)

	assert.Equal(t, uint64(25), v.Value())

	v.Set(315)

	assert.Equal(t, uint64(315), v.Value())
}

func TestUIntFromPtr(t *testing.T) {
	val := uint64(2)
	v := value.NewUIntFromPtr(&val)

	assert.Equal(t, uint64(2), v.Value())

	val = 3

	assert.Equal(t, uint64(3), v.Value())

	ptr := v.ValuePointer().(*uint64)

	*ptr = 4

	assert.Equal(t, uint64(4), v.Value())
}

func TestUIntOperations(t *testing.T) {
	v := value.NewUInt(37)
	vEq := value.NewUInt(37)
	vLt := value.NewUInt(36)
	vGt := value.NewUInt(38)

	verifyCompares(t, v, vEq, vLt, vGt)
}

func TestUIntOperationsWrongType(t *testing.T) {
	v := value.NewUInt(0)
	v2 := value.NewFloat(0.0)

	verifyCompareWrongType(t, v.Equal, v2)
	verifyCompareWrongType(t, v.Greater, v2)
	verifyCompareWrongType(t, v.GreaterEqual, v2)
	verifyCompareWrongType(t, v.Less, v2)
	verifyCompareWrongType(t, v.LessEqual, v2)
}

func TestUIntOneOf(t *testing.T) {
	testUIntOneOf(t, 0, []uint64{}, false)
	testUIntOneOf(t, 0, []uint64{0}, true)
	testUIntOneOf(t, 0, []uint64{1}, false)
	testUIntOneOf(t, 0, []uint64{1, 0}, true)
}

func TestUIntClone(t *testing.T) {
	const val = uint64(14)

	v1 := value.NewUInt(val)
	v2 := v1.Clone()

	assert.Equal(t, v1.Value(), v2.(value.Single).Value())

	// make sure the values are independent
	v1.Set(4)

	assert.Equal(t, val, v2.(value.Single).Value())
}

func TestUIntParse(t *testing.T) {
	v := value.NewUInt(0)

	assert.Error(t, v.Parse("2.5"))
	assert.NoError(t, v.Parse("17"))
	assert.Equal(t, uint64(17), v.Value())
}

func testUIntOneOf(t *testing.T, uVal uint64, uVals []uint64, expected bool) {
	v := value.NewUInt(uVal)
	s := value.NewUIntSlice(uVals...)

	result, err := v.OneOf(s)

	if assert.NoError(t, err) {
		assert.Equal(t, expected, result)
	}
}
