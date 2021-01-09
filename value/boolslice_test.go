package value_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestBoolSlice(t *testing.T) {
	s := value.NewBoolSlice()

	assert.Equal(t, value.TypeBool, s.Type())
	assert.True(t, s.IsSlice())
	assert.Equal(t, []bool{}, s.Slice())
	assert.Equal(t, 0, s.Len())

	ptr := s.SlicePointer().(*[]bool)
	*ptr = append(*ptr, true)

	assert.Equal(t, []bool{true}, s.Slice())
	assert.Equal(t, 1, s.Len())

	*ptr = []bool{false, true}

	assert.Equal(t, []bool{false, true}, s.Slice())
	assert.Equal(t, 2, s.Len())

	s.Set([]bool{true, false})

	assert.Equal(t, []bool{true, false}, s.Slice())
	assert.Equal(t, 2, s.Len())
}

func TestBoolSliceFromPtr(t *testing.T) {
	vals := []bool{false}
	v := value.NewBoolSliceFromPtr(&vals)

	assert.Equal(t, []bool{false}, v.Slice())

	vals[0] = true

	assert.Equal(t, []bool{true}, v.Slice())

	vals = append(vals, true)

	assert.Equal(t, []bool{true, true}, v.Slice())

	ptr := v.SlicePointer().(*[]bool)

	*ptr = []bool{false}

	assert.Equal(t, []bool{false}, v.Slice())
}

func TestBoolSliceComparesEmpty(t *testing.T) {
	s := value.NewBoolSlice()
	v := value.NewBool(false)

	testSliceComparesEmpty(t, s, v)

	v.Set(true)

	testSliceComparesEmpty(t, s, v)
}

func TestBoolSliceComparesWrongType(t *testing.T) {
	s := value.NewBoolSlice(false)
	v := value.NewFloat(0.0)

	testSliceComparesWrongType(t, s, v)
}

func TestBoolSliceCompares(t *testing.T) {
	testBoolSliceCompares(t, false)
	testBoolSliceCompares(t, true)
}

func TestBoolSliceEqual(t *testing.T) {
	testBoolSliceEqual(t, []bool{}, []bool{}, true)
	testBoolSliceEqual(t, []bool{true}, []bool{}, false)
	testBoolSliceEqual(t, []bool{true}, []bool{true}, true)
	testBoolSliceEqual(t, []bool{false}, []bool{true}, false)
	testBoolSliceEqual(t, []bool{true, false}, []bool{true, false}, true)
}

func TestBoolSliceEqualWrongType(t *testing.T) {
	s1 := value.NewBoolSlice()
	s2 := value.NewFloatSlice()

	_, err := s1.Equal(s2)

	assert.Error(t, err)
}

func TestBoolSliceContainsWrongType(t *testing.T) {
	s := value.NewBoolSlice()
	v := value.NewFloat(0.0)

	s.Set([]bool{true})

	_, err := s.Contains(v)

	assert.Error(t, err)
}

func TestBoolSliceClone(t *testing.T) {
	vals := []bool{false, false}

	v1 := value.NewBoolSlice(vals...)
	v2 := v1.Clone()

	assert.Equal(t, v1.Slice(), v2.(value.Slice).Slice())

	// make sure the values are independent
	v1.Set([]bool{true})

	assert.Equal(t, vals, v2.(value.Slice).Slice())
}

func TestBoolSliceParse(t *testing.T) {
	v := value.NewBoolSlice()

	assert.Error(t, v.Parse("2"))
	assert.Error(t, v.Parse("true,twue"))

	assert.NoError(t, v.Parse("true, false"))
	assert.Equal(t, []bool{true, false}, v.Slice())
}

func testBoolSliceCompares(t *testing.T, bVal bool) {
	s := value.NewBoolSlice()
	v := value.NewBool(false)

	v.Set(bVal)
	s.Set([]bool{false})

	testSliceCompare(t, s.Greater, v, value.BoolGreater(false, bVal))
	testSliceCompare(t, s.GreaterEqual, v, value.BoolGreaterEqual(false, bVal))
	testSliceCompare(t, s.Less, v, value.BoolLess(false, bVal))
	testSliceCompare(t, s.LessEqual, v, value.BoolLessEqual(false, bVal))

	s.Set([]bool{true})

	testSliceCompare(t, s.Greater, v, value.BoolGreater(true, bVal))
	testSliceCompare(t, s.GreaterEqual, v, value.BoolGreaterEqual(true, bVal))
	testSliceCompare(t, s.Less, v, value.BoolLess(true, bVal))
	testSliceCompare(t, s.LessEqual, v, value.BoolLessEqual(true, bVal))

	s.Set([]bool{false, true})

	testSliceCompare(t, s.Greater, v, value.BoolGreater(false, bVal) && value.BoolGreater(true, bVal))
	testSliceCompare(t, s.GreaterEqual, v, value.BoolGreaterEqual(false, bVal) && value.BoolGreaterEqual(true, bVal))
	testSliceCompare(t, s.Less, v, value.BoolLess(false, bVal) && value.BoolLess(true, bVal))
	testSliceCompare(t, s.LessEqual, v, value.BoolLessEqual(false, bVal) && value.BoolLessEqual(true, bVal))
}

func testBoolSliceEqual(t *testing.T, vals1, vals2 []bool, expected bool) {
	s1 := value.NewBoolSlice(vals1...)
	s2 := value.NewBoolSlice(vals2...)

	testSliceEqual(t, s1, s2, expected)
}

func testSliceEqual(t *testing.T, s1, s2 value.Slice, expected bool) {
	result, err := s1.Equal(s2)

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, expected, result)

	result, err = s2.Equal(s1)

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, expected, result)
}

func testSliceComparesEmpty(t *testing.T, s value.Slice, v value.Single) {
	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, false)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, false)
}

func testSliceComparesWrongType(t *testing.T, s value.Slice, v value.Single) {
	testSliceCompareWrongType(t, s.Greater, v)
	testSliceCompareWrongType(t, s.GreaterEqual, v)
	testSliceCompareWrongType(t, s.Less, v)
	testSliceCompareWrongType(t, s.LessEqual, v)
}

func testSliceCompare(
	t *testing.T, f compareValueFunc, v value.Single, expected bool) {
	result, err := f(v)

	if assert.NoError(t, err) {
		assert.Equal(t, expected, result)
	}
}

func testSliceCompareWrongType(t *testing.T, f compareValueFunc, v value.Single) {
	_, err := f(v)

	assert.Error(t, err)
}
