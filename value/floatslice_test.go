package value_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestFloatSlice(t *testing.T) {
	s := value.NewFloatSlice()

	assert.Equal(t, value.TypeFloat, s.Type())
	assert.True(t, s.IsSlice())
	assert.Equal(t, []float64{}, s.Slice())
	assert.Equal(t, 0, s.Len())

	ptr := s.SlicePointer().(*[]float64)
	*ptr = append(*ptr, 2.2)

	assert.Equal(t, []float64{2.2}, s.Slice())
	assert.Equal(t, 1, s.Len())

	*ptr = []float64{7.5, 13.5}

	assert.Equal(t, []float64{7.5, 13.5}, s.Slice())
	assert.Equal(t, 2, s.Len())

	s.Set([]float64{27.55, 88.88})

	assert.Equal(t, []float64{27.55, 88.88}, s.Slice())
	assert.Equal(t, 2, s.Len())
}

func TestFloatSliceFromPtr(t *testing.T) {
	vals := []float64{2.0}
	v := value.NewFloatSliceFromPtr(&vals)

	assert.Equal(t, []float64{2.0}, v.Slice())

	vals[0] = 2.5

	assert.Equal(t, []float64{2.5}, v.Slice())

	vals = append(vals, 3.0)

	assert.Equal(t, []float64{2.5, 3.0}, v.Slice())

	ptr := v.SlicePointer().(*[]float64)

	*ptr = []float64{5.0}

	assert.Equal(t, []float64{5.0}, v.Slice())
}

func TestFloatSliceComparesEmpty(t *testing.T) {
	s := value.NewFloatSlice()
	v := value.NewFloat(0.0)

	testSliceComparesEmpty(t, s, v)
}

func TestFloatSliceComparesWrongType(t *testing.T) {
	s := value.NewFloatSlice()
	v := value.NewBool(false)

	testSliceComparesWrongType(t, s, v)
}

func TestFloatSliceCompares(t *testing.T) {
	const val = 2.5

	s := value.NewFloatSlice(val)
	v := value.NewFloat(val)

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, true)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, true)

	s.Set([]float64{val, val + 1})

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, true)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, false)

	s.Set([]float64{val, val - 1})

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, false)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, true)

	s.Set([]float64{val + 2, val + 1})

	testSliceCompare(t, s.Greater, v, true)
	testSliceCompare(t, s.GreaterEqual, v, true)
	testSliceCompare(t, s.Less, v, false)
	testSliceCompare(t, s.LessEqual, v, false)

	s.Set([]float64{val - 2, val - 1})

	testSliceCompare(t, s.Greater, v, false)
	testSliceCompare(t, s.GreaterEqual, v, false)
	testSliceCompare(t, s.Less, v, true)
	testSliceCompare(t, s.LessEqual, v, true)
}

func TestFloatSliceEqual(t *testing.T) {
	testFloatSliceEqual(t, []float64{}, []float64{}, true)
	testFloatSliceEqual(t, []float64{22.2}, []float64{}, false)
	testFloatSliceEqual(t, []float64{22.2}, []float64{22.2}, true)
	testFloatSliceEqual(t, []float64{22.2}, []float64{33.3}, false)
	testFloatSliceEqual(t, []float64{22.2, 33.3}, []float64{22.2, 33.3}, true)
}

func TestFloatSliceEqualWrongType(t *testing.T) {
	s1 := value.NewFloatSlice()
	s2 := value.NewBoolSlice()

	_, err := s1.Equal(s2)

	assert.Error(t, err)
}

func TestFloatSliceContainsWrongType(t *testing.T) {
	s := value.NewFloatSlice()
	v := value.NewBool(false)

	s.Set([]float64{7.7})

	_, err := s.Contains(v)

	assert.Error(t, err)
}

func TestFloatSliceClone(t *testing.T) {
	vals := []float64{-88.5}

	v1 := value.NewFloatSlice(vals...)
	v2 := v1.Clone()

	assert.Equal(t, v1.Slice(), v2.(value.Slice).Slice())

	// make sure the values are independent
	v1.Set([]float64{22.1, -5.5})

	assert.Equal(t, vals, v2.(value.Slice).Slice())
}

func TestFloatSliceParse(t *testing.T) {
	v := value.NewFloatSlice()

	assert.Error(t, v.Parse("abc"))
	assert.Error(t, v.Parse("-7.5, abc"))

	assert.NoError(t, v.Parse("-7.2, 2.2"))
	assert.Equal(t, []float64{-7.2, 2.2}, v.Slice())
}

func testFloatSliceEqual(t *testing.T, vals1, vals2 []float64, expected bool) {
	s1 := value.NewFloatSlice(vals1...)
	s2 := value.NewFloatSlice(vals2...)

	testSliceEqual(t, s1, s2, expected)
}
