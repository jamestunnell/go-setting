package value_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestUIntSlice(t *testing.T) {
	s := value.NewUIntSlice()

	assert.Equal(t, value.TypeUInt, s.Type())
	assert.True(t, s.IsSlice())
	assert.Equal(t, []uint64{}, s.Slice())
	assert.Equal(t, 0, s.Len())

	ptr := s.SlicePointer().(*[]uint64)
	*ptr = append(*ptr, 2)

	assert.Equal(t, []uint64{2}, s.Slice())
	assert.Equal(t, 1, s.Len())

	*ptr = []uint64{7, 13}

	assert.Equal(t, []uint64{7, 13}, s.Slice())
	assert.Equal(t, 2, s.Len())

	s.Set([]uint64{27, 88})

	assert.Equal(t, []uint64{27, 88}, s.Slice())
	assert.Equal(t, 2, s.Len())
}

func TestUIntSliceFromPtr(t *testing.T) {
	vals := []uint64{0}
	v := value.NewUIntSliceFromPtr(&vals)

	assert.Equal(t, []uint64{0}, v.Slice())

	vals[0] = 7

	assert.Equal(t, []uint64{7}, v.Slice())

	vals = append(vals, 9)

	assert.Equal(t, []uint64{7, 9}, v.Slice())

	ptr := v.SlicePointer().(*[]uint64)

	*ptr = []uint64{11}

	assert.Equal(t, []uint64{11}, v.Slice())
}

func TestUIntSliceComparesEmpty(t *testing.T) {
	s := value.NewUIntSlice()
	v := value.NewUInt(0)

	testSliceComparesEmpty(t, s, v)
}

func TestUIntSliceComparesWrongType(t *testing.T) {
	s := value.NewUIntSlice()
	v := value.NewBool(false)

	testSliceComparesWrongType(t, s, v)
}

func TestUIntSliceCompares(t *testing.T) {
	const val = uint64(5)

	s := value.NewUIntSlice(val)
	v := value.NewUInt(val)

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, true)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, true)

	s.Set([]uint64{val, val + 1})

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, true)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, false)

	s.Set([]uint64{val, val - 1})

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, false)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, true)

	s.Set([]uint64{val + 2, val + 1})

	testSliceCompare(t, s.Greater, v, true)
	testSliceCompare(t, s.GreaterEqual, v, true)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, false)

	s.Set([]uint64{val - 2, val - 1})

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, false)
	testSliceCompare(t, s.Less, v, true)
	testSliceCompare(t, s.LessEqual, v, true)
}

func TestUIntSliceEqual(t *testing.T) {
	testUIntSliceEqual(t, []uint64{}, []uint64{}, true)
	testUIntSliceEqual(t, []uint64{22}, []uint64{}, false)
	testUIntSliceEqual(t, []uint64{22}, []uint64{22}, true)
	testUIntSliceEqual(t, []uint64{22}, []uint64{33}, false)
	testUIntSliceEqual(t, []uint64{22, 33}, []uint64{22, 33}, true)
}

func TestUIntSliceEqualWrongType(t *testing.T) {
	s1 := value.NewUIntSlice()
	s2 := value.NewBoolSlice()

	_, err := s1.Equal(s2)

	assert.Error(t, err)
}

func TestUIntSliceContainsWrongType(t *testing.T) {
	s := value.NewUIntSlice(7)
	v := value.NewBool(false)

	_, err := s.Contains(v)

	assert.Error(t, err)
}

func TestUIntSliceClone(t *testing.T) {
	vals := []uint64{88}

	v1 := value.NewUIntSlice(vals...)
	v2 := v1.Clone()

	assert.Equal(t, v1.Slice(), v2.(value.Slice).Slice())

	// make sure the values are independent
	v1.Set([]uint64{22, 5})

	assert.Equal(t, vals, v2.(value.Slice).Slice())
}

func TestUIntSliceParse(t *testing.T) {
	v := value.NewUIntSlice()

	assert.Error(t, v.Parse("7.7"))
	assert.Error(t, v.Parse("17, 2.2"))

	assert.NoError(t, v.Parse("17, 2"))
	assert.Equal(t, []uint64{17, 2}, v.Slice())
}

func testUIntSliceEqual(t *testing.T, vals1, vals2 []uint64, expected bool) {
	s1 := value.NewUIntSlice()
	s2 := value.NewUIntSlice()

	s1.Set(vals1)
	s2.Set(vals2)

	testSliceEqual(t, s1, s2, expected)
}
