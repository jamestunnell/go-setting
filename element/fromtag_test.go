package element_test

import (
	"fmt"
	"testing"

	"github.com/jamestunnell/go-settings/element"
	"github.com/jamestunnell/go-settings/option"
	"github.com/jamestunnell/go-settings/value"
	"github.com/stretchr/testify/assert"
)

type structtag map[string]string

func (tag structtag) Get(key string) string {
	if val, found := tag[key]; found {
		return val
	}

	return ""
}

func TestFromTagUnsupportedOptionIsIgnored(t *testing.T) {
	tag := &structtag{
		"abc": "xyz",
	}

	e, err := element.FromTag("X", tag, value.Float64)

	assert.NoError(t, err)
	assert.NotNil(t, e)
}

func TestFromTagEmptyOptionValIsIgnored(t *testing.T) {
	tag := &structtag{
		"name": "",
	}

	e, err := element.FromTag("X", tag, value.String)

	if !assert.NoError(t, err) || !assert.NotNil(t, e) {
		return
	}

	assert.Equal(t, "X", e.Name())
}

func TestFromTagOptionNotAllowedWithValue(t *testing.T) {
	tag := &structtag{
		"greater": "f",
	}

	_, err := element.FromTag("X", tag, value.String)

	assert.Error(t, err)
}

func TestFromTagInvalidDefaultValue(t *testing.T) {
	testFromTagInvalidOptionVal(t, value.Int64, option.Default, "2.5")
	testFromTagInvalidOptionVal(t, value.Int64s, option.Default, "-7,2.5")
	testFromTagInvalidOptionVal(t, value.UInt64, option.Default, "2.5")
	testFromTagInvalidOptionVal(t, value.UInt64s, option.Default, "5,2.5")
	testFromTagInvalidOptionVal(t, value.Float64, option.Default, "true")
	testFromTagInvalidOptionVal(t, value.Float64s, option.Default, "2.2,true")
	testFromTagInvalidOptionVal(t, value.Bool, option.Default, "twue")
	testFromTagInvalidOptionVal(t, value.Bools, option.Default, "true,twue")
}

func TestFromTagValidDefaultValue(t *testing.T) {
	testFromTagValidOptionVal(t, value.Int64, option.Default, "-7")
	testFromTagValidOptionVal(t, value.Int64s, option.Default, "-7")
	testFromTagValidOptionVal(t, value.Int64s, option.Default, "-7,2")
	testFromTagValidOptionVal(t, value.UInt64, option.Default, "6")
	testFromTagValidOptionVal(t, value.UInt64s, option.Default, "6")
	testFromTagValidOptionVal(t, value.UInt64s, option.Default, "6,7")
	testFromTagValidOptionVal(t, value.Float64, option.Default, "2.2")
	testFromTagValidOptionVal(t, value.Float64s, option.Default, "2.2")
	testFromTagValidOptionVal(t, value.Float64s, option.Default, "2.2,2.3")
	testFromTagValidOptionVal(t, value.Bool, option.Default, "true")
	testFromTagValidOptionVal(t, value.Bool, option.Default, "false")
	testFromTagValidOptionVal(t, value.Bools, option.Default, "false")
	testFromTagValidOptionVal(t, value.Bools, option.Default, "true,true,false")
	testFromTagValidOptionVal(t, value.String, option.Default, "anything!")
	testFromTagValidOptionVal(t, value.Strings, option.Default, "anything,and more!")
}

func TestFromTagInvalidCompareValue(t *testing.T) {
	compareOptTypes := []option.Type{
		option.Less, option.LessEqual, option.Greater, option.GreaterEqual}
	for _, optType := range compareOptTypes {
		testFromTagInvalidOptionVal(t, value.Int64, optType, "2.5")
		testFromTagInvalidOptionVal(t, value.Int64s, optType, "2,4")
		testFromTagInvalidOptionVal(t, value.Int64s, optType, "2.5")
		testFromTagInvalidOptionVal(t, value.UInt64, optType, "2.5")
		testFromTagInvalidOptionVal(t, value.UInt64s, optType, "2,4")
		testFromTagInvalidOptionVal(t, value.UInt64s, optType, "2.5")
		testFromTagInvalidOptionVal(t, value.Float64, optType, "true")
		testFromTagInvalidOptionVal(t, value.Float64s, optType, "2.5,5.2")
		testFromTagInvalidOptionVal(t, value.Float64s, optType, "true")
	}
}

func TestFromTagValidCompareValue(t *testing.T) {
	compareOptTypes := []option.Type{
		option.Less, option.LessEqual, option.Greater, option.GreaterEqual}
	for _, optType := range compareOptTypes {
		testFromTagValidOptionVal(t, value.Int64, optType, "-2")
		testFromTagValidOptionVal(t, value.Int64s, optType, "-2")
		testFromTagValidOptionVal(t, value.UInt64, optType, "75")
		testFromTagValidOptionVal(t, value.UInt64s, optType, "75")
		testFromTagValidOptionVal(t, value.Float64, optType, "7.7")
		testFromTagValidOptionVal(t, value.Float64s, optType, "8.5")
	}
}

func TestFromTagInvalidOneOfValue(t *testing.T) {
	testFromTagInvalidOptionVal(t, value.Int64, option.OneOf, "47,22.5")
	testFromTagInvalidOptionVal(t, value.UInt64, option.OneOf, "44,-7")
	testFromTagInvalidOptionVal(t, value.Float64, option.OneOf, "2.2,true")
}

func TestFromTagValidOneOfValue(t *testing.T) {
	testFromTagValidOptionVal(t, value.Int64, option.OneOf, "47")
	testFromTagValidOptionVal(t, value.Int64, option.OneOf, "47,22")
	testFromTagValidOptionVal(t, value.UInt64, option.OneOf, "44")
	testFromTagValidOptionVal(t, value.UInt64, option.OneOf, "44,0")
	testFromTagValidOptionVal(t, value.Float64, option.OneOf, "2.2")
	testFromTagValidOptionVal(t, value.Float64, option.OneOf, "2.2,22.22")
	testFromTagValidOptionVal(t, value.String, option.OneOf, "2.2,anything")
}

// func TestFromTagInvalidVal(t *testing.T) {
// 	testCases := map[value.Type][]string{
// 		// value.UInt64:  []string{"-1", "7.7", "true"},
// 		// value.Int64:   []string{"true", "7.7"},
// 		// value.Float64: []string{"true"},
// 		value.Bool: []string{"twue", "-5", "7.7"},
// 	}

// 	for valType, invalidVals := range testCases {
// 		optTypes := applicableTypes(valType)

// 		testFromTagInvalidVals(t, valType, optTypes, invalidVals)
// 	}
// }

// func TestFromTagInvalidLen(t *testing.T) {
// 	invalidLens := []string{"-1", "7.7", "true"}

// 	testFromTagInvalidVals(t, value.String, []option.Type{option.MinLen, option.MaxLen}, invalidLens)
// }

// func TestFromTagValidLen(t *testing.T) {
// 	validLens := []string{"0", "77", "123456"}

// 	testFromTagValidVals(t, value.String, []option.Type{option.MinLen, option.MaxLen}, validLens)
// }

// func TestFromTagValidVal(t *testing.T) {
// 	testCases := map[value.Type][]string{
// 		value.UInt64:  []string{"0", "77", "123456"},
// 		value.Int64:   []string{"-1", "0", "77", "123456"},
// 		value.Float64: []string{"-50", "10", "0.0", "2.7", "1.7e-25"},
// 		value.Bool:    []string{"true", "false", "0", "1"},
// 	}

// 	for valType, validVals := range testCases {
// 		optTypes := applicableTypes(valType)

// 		testFromTagValidVals(t, valType, optTypes, validVals)
// 	}

// 	validStrVals := []string{"!anything~*except@@parenthesis", "-50", "1.7e-25"}

// 	optTypesString := []option.Type{option.OneOf, option.Default}
// 	testFromTagValidVals(t, value.String, optTypesString, validStrVals)

// 	optTypesNumerSlice := []option.Type{option.Less, option.LessEqual, option.Greater, option.GreaterEqual}
// 	testFromTagValidVals(t, value.Int64s, optTypesNumerSlice, []string{"-1", "0", "77"})
// 	testFromTagValidVals(t, value.UInt64s, optTypesNumerSlice, []string{"0", "77"})
// 	testFromTagValidVals(t, value.Float64s, optTypesNumerSlice, []string{"-50", "10", "2.7"})
// }

func TestFromTagMultipleOptions(t *testing.T) {
	tag := structtag{
		"default": "2.5",
		"greater": "0.0",
	}

	e, err := element.FromTag("X", tag, value.Float64)

	assert.NoError(t, err)
	assert.NotNil(t, e)

	if !assert.Len(t, e.Options(), 2) {
		return
	}

	oDefault := e.Option(option.Default)
	oGreater := e.Option(option.Greater)

	if !assert.NotNil(t, oDefault) || !assert.NotNil(t, oGreater) {
		return
	}

	assert.Equal(t, 2.5, oDefault.Param)
	assert.Equal(t, 0.0, oGreater.Param)
}

func TestFromTagMultipleOptionsOneInvalid(t *testing.T) {
	tag := structtag{
		"default": "notfloat",
		"greater": "0.0",
	}

	e, err := element.FromTag("X", tag, value.Float64)

	assert.Error(t, err)
	assert.Nil(t, e)
}

func TestFromTagIncompatibleOptions(t *testing.T) {
	tag := structtag{
		"greater": "0.5",
		"less":    "0.5",
	}

	e, err := element.FromTag("X", tag, value.Float64)

	assert.Error(t, err)
	assert.Nil(t, e)
}

func testFromTagInvalidOptionVal(t *testing.T, valType value.Type, optType option.Type, valStr string) {
	tName := fmt.Sprintf("invalid %s %s in option %s", valType.String(), valStr, optType.String())

	t.Run(tName, func(t *testing.T) {
		tag := makeTagWithOption(optType, valStr)
		e, err := element.FromTag("X", tag, valType)

		assert.Error(t, err)
		assert.Nil(t, e)
	})
}

func testFromTagValidOptionVal(t *testing.T, valType value.Type, optType option.Type, valStr string) {
	name := fmt.Sprintf("valid %s %s in option %s", valType.String(), valStr, optType.String())

	t.Run(name, func(t *testing.T) {
		tag := makeTagWithOption(optType, valStr)
		e, err := element.FromTag("X", tag, valType)

		if !assert.NoError(t, err) || assert.NotNil(t, e) {
			return
		}

		assert.Equal(t, "X", e.Name())
		assert.Equal(t, valType, e.Type())
		assert.NotNil(t, e.Option(optType))

		if !assert.Len(t, e.Options(), 1) {
			return
		}

		o := e.Options()[0]

		assert.Equal(t, optType, o.Type)

		val, err := value.Parse(valStr, valType)

		assert.NoError(t, err)
		assert.Equal(t, val, o.Param)
	})
}

func makeTagWithOption(optType option.Type, valStr string) structtag {
	return structtag{optType.String(): valStr}
}

func applicableTypes(valType value.Type) []option.Type {
	allTypes := option.AllTypes()
	applicable := []option.Type{}

	for _, t := range allTypes {
		if t.ApplicableTo(valType) {
			applicable = append(applicable, t)
		}
	}

	return applicable
}
