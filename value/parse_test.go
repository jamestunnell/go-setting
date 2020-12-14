package value_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-settings/value"
)

func TestParseValUnsupportedType(t *testing.T) {
	_, err := value.Parse("abc", value.Type(-1))

	assert.Error(t, err)
}

func TestParseValString(t *testing.T) {
	val, err := value.Parse("abc", value.String)

	assert.NoError(t, err)
	assert.Equal(t, "abc", val)
}

func TestParseValUInt(t *testing.T) {
	val, err := value.Parse("50", value.UInt64)

	assert.NoError(t, err)
	assert.Equal(t, uint64(50), val)

	val, err = value.Parse("-50", value.UInt64)

	assert.Error(t, err)
}

func TestParseValUInts(t *testing.T) {
	val, err := value.Parse("50,100", value.UInt64s)

	assert.NoError(t, err)
	assert.Equal(t, []uint64{50, 100}, val)

	val, err = value.Parse("25,-50", value.UInt64s)

	assert.Error(t, err)
}

func TestParseValInt(t *testing.T) {
	val, err := value.Parse("-50", value.Int64)

	assert.NoError(t, err)
	assert.Equal(t, int64(-50), val)

	val, err = value.Parse("-50.5", value.Int64)

	assert.Error(t, err)
}

func TestParseValInts(t *testing.T) {
	val, err := value.Parse("7,8,-50", value.Int64s)

	assert.NoError(t, err)
	assert.Equal(t, []int64{7, 8, -50}, val)

	val, err = value.Parse("7,8,-50.5", value.Int64s)

	assert.Error(t, err)
}

func TestParseValFloat(t *testing.T) {
	val, err := value.Parse("-50.0", value.Float64)

	assert.NoError(t, err)
	assert.Equal(t, -50.0, val)

	val, err = value.Parse("true", value.Float64)

	assert.Error(t, err)
}

func TestParseValFloats(t *testing.T) {
	val, err := value.Parse("-50.0, 2.5", value.Float64s)

	assert.NoError(t, err)
	assert.Equal(t, []float64{-50.0, 2.5}, val)

	val, err = value.Parse("2.2,true", value.Float64s)

	assert.Error(t, err)
}

func TestParseValBool(t *testing.T) {
	val, err := value.Parse("true", value.Bool)

	assert.NoError(t, err)
	assert.Equal(t, true, val)

	val, err = value.Parse("twue", value.Bool)

	assert.Error(t, err)
}

func TestParseValBools(t *testing.T) {
	val, err := value.Parse("true", value.Bools)

	assert.NoError(t, err)
	assert.Equal(t, []bool{true}, val)

	val, err = value.Parse("false, true", value.Bools)

	assert.NoError(t, err)
	assert.Equal(t, []bool{false, true}, val)

	val, err = value.Parse("false, twue", value.Bool)

	assert.Error(t, err)
}

func TestParseValStrings(t *testing.T) {
	val, err := value.Parse(" abc, def ", value.Strings)

	assert.NoError(t, err)
	assert.Equal(t, []string{"abc", "def"}, val)
}
