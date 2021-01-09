package value_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestStringValue(t *testing.T) {
	v := value.NewString("")

	assert.Equal(t, value.TypeString, v.Type())
	assert.False(t, v.IsSlice())
	assert.Equal(t, "", v.Value())

	ptr := v.ValuePointer().(*string)
	*ptr = "abc"

	assert.Equal(t, "abc", v.Value())

	v.Set("xyz")

	assert.Equal(t, "xyz", v.Value())
}

func TestStringFromPtr(t *testing.T) {
	val := "abc"
	v := value.NewStringFromPtr(&val)

	assert.Equal(t, "abc", v.Value())

	val = "xyz"

	assert.Equal(t, "xyz", v.Value())

	ptr := v.ValuePointer().(*string)

	*ptr = "dog"

	assert.Equal(t, "dog", v.Value())
}

func TestStringOperations(t *testing.T) {
	v := value.NewString("cat")
	vEq := value.NewString("cat")
	vLt := value.NewString("bat")
	vGt := value.NewString("dog")

	verifyCompares(t, v, vEq, vLt, vGt)
}

func TestStringOperationsWrongType(t *testing.T) {
	v := value.NewString("")
	v2 := value.NewFloat(0.0)

	verifyCompareWrongType(t, v.Equal, v2)
	verifyCompareWrongType(t, v.Greater, v2)
	verifyCompareWrongType(t, v.GreaterEqual, v2)
	verifyCompareWrongType(t, v.Less, v2)
	verifyCompareWrongType(t, v.LessEqual, v2)
}

func TestStringOneOf(t *testing.T) {
	testStringOneOf(t, "abc", []string{}, false)
	testStringOneOf(t, "abc", []string{"abc"}, true)
	testStringOneOf(t, "abc", []string{"xyz"}, false)
	testStringOneOf(t, "abc", []string{"xyz", "abc"}, true)
}

func TestStringClone(t *testing.T) {
	const val = "abc"

	v1 := value.NewString(val)
	v2 := v1.Clone()

	assert.Equal(t, v1.Value(), v2.(value.Single).Value())

	// make sure the values are independent
	v1.Set("xyz")

	assert.Equal(t, val, v2.(value.Single).Value())
}

func TestStringParse(t *testing.T) {
	v := value.NewString("")

	assert.NoError(t, v.Parse("-7"))
	assert.Equal(t, "-7", v.Value())
}

func testStringOneOf(t *testing.T, sVal string, sVals []string, expected bool) {
	v := value.NewString(sVal)
	s := value.NewStringSlice(sVals...)

	result, err := v.OneOf(s)

	if assert.NoError(t, err) {
		assert.Equal(t, expected, result)
	}
}
