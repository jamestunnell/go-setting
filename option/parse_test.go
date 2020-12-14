package option_test

import (
	"fmt"
	"testing"

	"github.com/jamestunnell/go-settings/option"
	"github.com/jamestunnell/go-settings/value"
	"github.com/stretchr/testify/assert"
)

func TestParseSuccesses(t *testing.T) {
	testParseSuccess(t, value.Float64, option.Default, "2.0", 2.0)
	testParseSuccess(t, value.String, option.MinLen, "2", uint64(2))
	testParseSuccess(t, value.String, option.MaxLen, "10", uint64(10))
	testParseSuccess(t, value.Int64, option.Greater, "-2", int64(-2))
	testParseSuccess(t, value.Float64, option.GreaterEqual, "2.5", 2.5)
	testParseSuccess(t, value.UInt64, option.Less, "77", uint64(77))
	testParseSuccess(t, value.Float64, option.LessEqual, "27.5", 27.5)
	testParseSuccess(t, value.UInt64, option.OneOf, "2,4,6", []uint64{2, 4, 6})
	testParseSuccess(t, value.Int64, option.OneOf, "2,4,6", []int64{2, 4, 6})
	testParseSuccess(t, value.Float64, option.OneOf, "2,4,6", []float64{2, 4, 6})
	testParseSuccess(t, value.String, option.OneOf, "2,4,6", []string{"2", "4", "6"})
}

func testParseSuccess(
	t *testing.T,
	valType value.Type,
	optType option.Type,
	paramStr string,
	expectedParamVal interface{},
) {
	name := fmt.Sprintf("%s:\"%s\"", optType.String(), paramStr)
	t.Run(name, func(t *testing.T) {
		o, err := option.Parse(valType, optType, paramStr)

		if !assert.NoError(t, err) || !assert.NotNil(t, o) {
			return
		}

		assert.Equal(t, optType, o.Type)
		assert.Equal(t, expectedParamVal, o.Param)
	})
}

func TestParseFailUnknownValueType(t *testing.T) {
	o, err := option.Parse(value.Type(-1), option.Default, "2.0")

	assert.Error(t, err)
	assert.Nil(t, o)
}

func TestParseFailUnknownOptionType(t *testing.T) {
	o, err := option.Parse(value.Float64, option.Type(-1), "2.0")

	assert.Error(t, err)
	assert.Nil(t, o)
}

func TestParseFailBadLen(t *testing.T) {
	o, err := option.Parse(value.Float64, option.MinLen, "2.5")

	assert.Error(t, err)
	assert.Nil(t, o)
}

func TestParseFailBadVal(t *testing.T) {
	o, err := option.Parse(value.Bool, option.Default, "twue")

	assert.Error(t, err)
	assert.Nil(t, o)
}
