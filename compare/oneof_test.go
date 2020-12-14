package compare_test

import (
	"testing"

	"github.com/jamestunnell/go-settings/compare"
	"github.com/stretchr/testify/assert"
)

func TestOneOfFails(t *testing.T) {
	cases := []*val1val2{
		newVal1Val2(true, []bool{false, true}), // bool is not supported
		newVal1Val2(int64(2), []int{2, 4}),     // type mismatch
		newVal1Val2(uint64(4), []uint{0, 4}),   // type mismatch
		newVal1Val2(2.0, []float32{2.0}),       // type mismatch
		newVal1Val2("abc", []float32{2.0}),     // type mismatch
	}

	for _, v1v2 := range cases {
		_, err := compare.OneOf(v1v2.Val1, v1v2.Val2)

		assert.Error(t, err)
	}
}

func TestOneOfInt64(t *testing.T) {
	result, err := compare.OneOf(int64(2), []int64{0, 2, 4})

	assert.NoError(t, err)
	assert.True(t, result)

	result, err = compare.OneOf(int64(2), []int64{1, 3, 5})

	assert.NoError(t, err)
	assert.False(t, result)
}

func TestOneOfUInt64(t *testing.T) {
	result, err := compare.OneOf(uint64(2), []uint64{0, 2, 4})

	assert.NoError(t, err)
	assert.True(t, result)

	result, err = compare.OneOf(uint64(2), []uint64{1, 3, 5})

	assert.NoError(t, err)
	assert.False(t, result)
}

func TestOneOfFloat64(t *testing.T) {
	result, err := compare.OneOf(2.0, []float64{0.0, 2.0, 4.0})

	assert.NoError(t, err)
	assert.True(t, result)

	result, err = compare.OneOf(2.0, []float64{1.0, 3.0, 5.0})

	assert.NoError(t, err)
	assert.False(t, result)
}

func TestOneOfString(t *testing.T) {
	result, err := compare.OneOf("abc", []string{"xyz", "abc"})

	assert.NoError(t, err)
	assert.True(t, result)

	result, err = compare.OneOf("abc", []string{"zyx", "cba"})

	assert.NoError(t, err)
	assert.False(t, result)
}
