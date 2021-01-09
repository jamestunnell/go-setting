package value_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestStringSlice(t *testing.T) {
	s := value.NewStringSlice()

	assert.Equal(t, value.TypeString, s.Type())
	assert.True(t, s.IsSlice())
	assert.Equal(t, []string{}, s.Slice())
	assert.Equal(t, 0, s.Len())

	ptr := s.SlicePointer().(*[]string)
	*ptr = append(*ptr, "aaa")

	assert.Equal(t, []string{"aaa"}, s.Slice())
	assert.Equal(t, 1, s.Len())

	*ptr = []string{"bbb", "ccc"}

	assert.Equal(t, []string{"bbb", "ccc"}, s.Slice())
	assert.Equal(t, 2, s.Len())

	s.Set([]string{"ddd", "eee"})

	assert.Equal(t, []string{"ddd", "eee"}, s.Slice())
	assert.Equal(t, 2, s.Len())
}

func TestStringSliceFromPtr(t *testing.T) {
	vals := []string{"a"}
	v := value.NewStringSliceFromPtr(&vals)

	assert.Equal(t, []string{"a"}, v.Slice())

	vals[0] = "b"

	assert.Equal(t, []string{"b"}, v.Slice())

	vals = append(vals, "c")

	assert.Equal(t, []string{"b", "c"}, v.Slice())

	ptr := v.SlicePointer().(*[]string)

	*ptr = []string{"d"}

	assert.Equal(t, []string{"d"}, v.Slice())
}

func TestStringSliceComparesEmpty(t *testing.T) {
	s := value.NewStringSlice()
	v := value.NewString("")

	testSliceComparesEmpty(t, s, v)
}

func TestStringSliceComparesWrongType(t *testing.T) {
	s := value.NewStringSlice("abc")
	v := value.NewBool(false)

	testSliceComparesWrongType(t, s, v)
}

func TestStringSliceCompares(t *testing.T) {
	const (
		valBefore2 = "ddd"
		valBefore  = "eee"
		val        = "fff"
		valAfter   = "ggg"
		valAfter2  = "hhh"
	)

	s := value.NewStringSlice(val)
	v := value.NewString(val)

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, true)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, true)

	s.Set([]string{val, valAfter})

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, true)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, false)

	s.Set([]string{val, valBefore})

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, false)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, true)

	s.Set([]string{valAfter2, valAfter})

	testSliceCompare(t, s.Greater, v, true)
	testSliceCompare(t, s.GreaterEqual, v, true)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, false)

	s.Set([]string{valBefore2, valBefore})

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, false)
	testSliceCompare(t, s.Less, v, true)
	testSliceCompare(t, s.LessEqual, v, true)
}

func TestStringSliceEqual(t *testing.T) {
	testStringSliceEqual(t, []string{}, []string{}, true)
	testStringSliceEqual(t, []string{"a"}, []string{}, false)
	testStringSliceEqual(t, []string{"a"}, []string{"a"}, true)
	testStringSliceEqual(t, []string{"a"}, []string{"b"}, false)
	testStringSliceEqual(t, []string{"a", "b"}, []string{"a", "b"}, true)
}

func TestStringSliceEqualWrongType(t *testing.T) {
	s1 := value.NewStringSlice()
	s2 := value.NewBoolSlice()

	_, err := s1.Equal(s2)

	assert.Error(t, err)
}

func TestStringSliceContainsWrongType(t *testing.T) {
	s := value.NewStringSlice("xyz")
	v := value.NewBool(false)

	_, err := s.Contains(v)

	assert.Error(t, err)
}

func TestStringSliceClone(t *testing.T) {
	vals := []string{"a", "b"}

	v1 := value.NewStringSlice(vals...)
	v2 := v1.Clone()

	assert.Equal(t, v1.Slice(), v2.(value.Slice).Slice())

	// make sure the values are independent
	v1.Set([]string{"c"})

	assert.Equal(t, vals, v2.(value.Slice).Slice())
}

func TestStringSliceParse(t *testing.T) {
	v := value.NewStringSlice()

	assert.NoError(t, v.Parse("true"))
	assert.Equal(t, []string{"true"}, v.Slice())

	assert.NoError(t, v.Parse("cat, dog"))
	assert.Equal(t, []string{"cat", "dog"}, v.Slice())
}

func testStringSliceEqual(t *testing.T, vals1, vals2 []string, expected bool) {
	s1 := value.NewStringSlice()
	s2 := value.NewStringSlice()

	s1.Set(vals1)
	s2.Set(vals2)

	testSliceEqual(t, s1, s2, expected)
}
