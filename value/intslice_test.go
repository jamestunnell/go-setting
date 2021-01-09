package value_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestIntSlice(t *testing.T) {
	s := value.NewIntSlice()

	assert.Equal(t, value.TypeInt, s.Type())
	assert.True(t, s.IsSlice())
	assert.Equal(t, []int64{}, s.Slice())
	assert.Equal(t, 0, s.Len())

	ptr := s.SlicePointer().(*[]int64)
	*ptr = append(*ptr, 2)

	assert.Equal(t, []int64{2}, s.Slice())
	assert.Equal(t, 1, s.Len())

	*ptr = []int64{7, 13}

	assert.Equal(t, []int64{7, 13}, s.Slice())
	assert.Equal(t, 2, s.Len())

	s.Set([]int64{27, 88})

	assert.Equal(t, []int64{27, 88}, s.Slice())
	assert.Equal(t, 2, s.Len())
}

func TestIntSliceFromPtr(t *testing.T) {
	vals := []int64{0}
	v := value.NewIntSliceFromPtr(&vals)

	assert.Equal(t, []int64{0}, v.Slice())

	vals[0] = 7

	assert.Equal(t, []int64{7}, v.Slice())

	vals = append(vals, 9)

	assert.Equal(t, []int64{7, 9}, v.Slice())

	ptr := v.SlicePointer().(*[]int64)

	*ptr = []int64{11}

	assert.Equal(t, []int64{11}, v.Slice())
}

func TestIntSliceComparesEmpty(t *testing.T) {
	s := value.NewIntSlice()
	v := value.NewInt(0)

	testSliceComparesEmpty(t, s, v)
}

func TestIntSliceComparesWrongType(t *testing.T) {
	s := value.NewIntSlice(0)
	v := value.NewBool(false)

	testSliceComparesWrongType(t, s, v)
}

func TestIntSliceCompares(t *testing.T) {
	const val = int64(5)

	s := value.NewIntSlice(val)
	v := value.NewInt(val)

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, true)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, true)

	s.Set([]int64{val, val + 1})

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, true)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, false)

	s.Set([]int64{val, val - 1})

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, false)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, true)

	s.Set([]int64{val + 2, val + 1})

	testSliceCompare(t, s.Greater, v, true)
	testSliceCompare(t, s.GreaterEqual, v, true)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, false)

	s.Set([]int64{val - 2, val - 1})

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, false)
	testSliceCompare(t, s.Less, v, true)
	testSliceCompare(t, s.LessEqual, v, true)
}

func TestIntSliceEqual(t *testing.T) {
	testIntSliceEqual(t, []int64{}, []int64{}, true)
	testIntSliceEqual(t, []int64{22}, []int64{}, false)
	testIntSliceEqual(t, []int64{22}, []int64{22}, true)
	testIntSliceEqual(t, []int64{22}, []int64{33}, false)
	testIntSliceEqual(t, []int64{22, 33}, []int64{22, 33}, true)
}

func TestIntSliceEqualWrongType(t *testing.T) {
	s1 := value.NewIntSlice()
	s2 := value.NewBoolSlice()

	_, err := s1.Equal(s2)

	assert.Error(t, err)
}

func TestIntSliceContainsWrongType(t *testing.T) {
	s := value.NewIntSlice(0)
	v := value.NewBool(false)

	s.Set([]int64{7})

	_, err := s.Contains(v)

	assert.Error(t, err)
}

func TestIntSliceClone(t *testing.T) {
	vals := []int64{-88}

	v1 := value.NewIntSlice(vals...)
	v2 := v1.Clone()

	assert.Equal(t, v1.Slice(), v2.(value.Slice).Slice())

	// make sure the values are independent
	v1.Set([]int64{22, 5})

	assert.Equal(t, vals, v2.(value.Slice).Slice())
}

func TestIntSliceParse(t *testing.T) {
	v := value.NewIntSlice()

	assert.Error(t, v.Parse("-7.7"))
	assert.Error(t, v.Parse("-7, 2.2"))

	assert.NoError(t, v.Parse("-7, 2"))
	assert.Equal(t, []int64{-7, 2}, v.Slice())
}

func testIntSliceEqual(t *testing.T, vals1, vals2 []int64, expected bool) {
	s1 := value.NewIntSlice()
	s2 := value.NewIntSlice()

	s1.Set(vals1)
	s2.Set(vals2)

	testSliceEqual(t, s1, s2, expected)
}
