package value_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-setting/value"
)

func TestFromValueUnsupportedType(t *testing.T) {
	i := int32(2)
	u := uint32(2)
	f := float32(2.5)
	si := []int32{2}
	su := []uint32{2}
	sf := []float32{2.5}

	testFromValueFail(t, i)
	testFromValueFail(t, u)
	testFromValueFail(t, f)
	testFromValueFail(t, si)
	testFromValueFail(t, su)
	testFromValueFail(t, sf)

	testFromValueFail(t, &i)
	testFromValueFail(t, &u)
	testFromValueFail(t, &f)
	testFromValueFail(t, &si)
	testFromValueFail(t, &su)
	testFromValueFail(t, &sf)
}

func TestFromValue(t *testing.T) {
	testFromValueSingle(t, int64(2), value.TypeInt)
	testFromValueSingle(t, uint64(2), value.TypeUInt)
	testFromValueSingle(t, 2.5, value.TypeFloat)
	testFromValueSingle(t, true, value.TypeBool)
	testFromValueSingle(t, "abc", value.TypeString)

	testFromValueSlice(t, []int64{2}, value.TypeInt)
	testFromValueSlice(t, []uint64{2}, value.TypeUInt)
	testFromValueSlice(t, []float64{2.5}, value.TypeFloat)
	testFromValueSlice(t, []bool{true}, value.TypeBool)
	testFromValueSlice(t, []string{"abc"}, value.TypeString)
}

func TestFromValuePtr(t *testing.T) {
	i := int64(2)
	u := uint64(2)
	f := 2.5
	b := true
	s := "abc"

	si := []int64{2}
	su := []uint64{2}
	sf := []float64{2.5}
	sb := []bool{true}
	ss := []string{"abc"}

	testFromValueSingle(t, &i, value.TypeInt)
	testFromValueSingle(t, &u, value.TypeUInt)
	testFromValueSingle(t, &f, value.TypeFloat)
	testFromValueSingle(t, &b, value.TypeBool)
	testFromValueSingle(t, &s, value.TypeString)

	testFromValueSlice(t, &si, value.TypeInt)
	testFromValueSlice(t, &su, value.TypeUInt)
	testFromValueSlice(t, &sf, value.TypeFloat)
	testFromValueSlice(t, &sb, value.TypeBool)
	testFromValueSlice(t, &ss, value.TypeString)
}

func testFromValueFail(t *testing.T, val interface{}) {
	v := value.FromValue(reflect.ValueOf(val))

	assert.Nil(t, v)
}

func testFromValueSingle(t *testing.T, val interface{}, expectedType value.Type) {
	v := value.FromValue(reflect.ValueOf(val))

	if !assert.NotNil(t, v) {
		return
	}

	assert.Equal(t, expectedType, v.Type())
	assert.False(t, v.IsSlice())

	s := v.(value.Single)

	if !assert.NotNil(t, s) {
		return
	}
}

func testFromValueSlice(t *testing.T, val interface{}, expectedType value.Type) {
	v := value.FromValue(reflect.ValueOf(val))

	if !assert.NotNil(t, v) {
		return
	}

	assert.Equal(t, expectedType, v.Type())
	assert.True(t, v.IsSlice())

	s := v.(value.Slice)

	if !assert.NotNil(t, s) {
		return
	}
}
