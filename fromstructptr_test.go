package settings_test

import (
	"testing"

	"github.com/jamestunnell/go-settings"
	"github.com/jamestunnell/go-settings/element"
	"github.com/jamestunnell/go-settings/option"
	"github.com/jamestunnell/go-settings/value"
	"github.com/stretchr/testify/assert"
)

func TestFromStructGivenStructNonPointer(t *testing.T) {
	type MyStruct struct{}

	s, err := settings.FromStructPtr(MyStruct{})

	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestFromStructPtrEmpty(t *testing.T) {
	type MyStruct struct{}

	x := &MyStruct{}
	s, err := settings.FromStructPtr(x)

	if !assert.NoError(t, err) || !assert.NotNil(t, s) {
		return
	}

	assert.Equal(t, "MyStruct", s.Name())
	assert.Len(t, s.Elements(), 0)
	assert.Equal(t, x, s.StructPtr())
}

func TestFromStructPtrFieldsForAllSupportedValueTypes(t *testing.T) {
	type MyStruct struct {
		A float64   `name:"a" default:"7.2"`
		B int64     `name:"b" greater:"20"`
		C uint64    `name:"c" less:"100"`
		D string    `name:"d" minLen:"10"`
		E bool      `name:"e"`
		F []float64 `name:"f" default:"7.2,4.4"`
		G []int64   `name:"g" greater:"20"`
		H []uint64  `name:"h" less:"100"`
		I []string  `name:"i" minLen:"10"`
		J []bool    `name:"j" maxLen:"7"`
	}

	expectedElems := []*element.Element{
		element.New("a", value.Float64, option.New(option.Name, "a"), option.New(option.Default, 7.2)),
		element.New("b", value.Int64, option.New(option.Name, "b"), option.New(option.Greater, int64(20))),
		element.New("c", value.UInt64, option.New(option.Name, "c"), option.New(option.Less, uint64(100))),
		element.New("d", value.String, option.New(option.Name, "d"), option.New(option.MinLen, uint64(10))),
		element.New("e", value.Bool, option.New(option.Name, "e")),
		element.New("f", value.Float64s, option.New(option.Name, "f"), option.New(option.Default, []float64{7.2, 4.4})),
		element.New("g", value.Int64s, option.New(option.Name, "g"), option.New(option.Greater, int64(20))),
		element.New("h", value.UInt64s, option.New(option.Name, "h"), option.New(option.Less, uint64(100))),
		element.New("i", value.Strings, option.New(option.Name, "i"), option.New(option.MinLen, uint64(10))),
		element.New("j", value.Bools, option.New(option.Name, "j"), option.New(option.MaxLen, uint64(7))),
	}

	x := &MyStruct{}
	s, err := settings.FromStructPtr(x)

	if !assert.NoError(t, err) || !assert.NotNil(t, s) {
		return
	}

	assert.Equal(t, "MyStruct", s.Name())
	assert.Len(t, s.Elements(), len(expectedElems))
	assert.Equal(t, x, s.StructPtr())

	for _, expectedElem := range expectedElems {
		elem := s.Element(expectedElem.Name())

		if !assert.NotNil(t, elem) {
			return
		}

		assert.Equal(t, expectedElem.Type(), elem.Type())
		assert.Equal(t, expectedElem.Required(), elem.Required())
		assert.Equal(t, expectedElem.DefaultVal(), elem.DefaultVal())
		assert.Len(t, elem.Options(), len(expectedElem.Options()))

		for _, expectedOpt := range expectedElem.Options() {
			opt := elem.Option(expectedOpt.Type)

			if !assert.NotNil(t, opt) {
				return
			}

			assert.Equal(t, expectedOpt.Param, opt.Param)
		}
	}
}

func TestFromStructPtrDuplicateElemNames(t *testing.T) {
	type MyStruct struct {
		A float64 `name:"a" default:"7.2"`
		B int64   `name:"a" greater:"20"`
	}

	x := &MyStruct{}
	s, err := settings.FromStructPtr(x)

	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestFromStructPtrInvalidOptionTag(t *testing.T) {
	type MyStruct struct {
		X int `name=5`
	}

	x := &MyStruct{}
	s, err := settings.FromStructPtr(x)

	assert.Nil(t, s)
	assert.Error(t, err)
}

func TestFromStructPtrMissingNameOption(t *testing.T) {
	type MyStruct struct {
		X int64 `json:"x"`
	}

	x := &MyStruct{}
	s, err := settings.FromStructPtr(x)

	assert.NotNil(t, s)
	assert.NoError(t, err)
	assert.Len(t, s.Elements(), 1)

	// Field name is used by default
	assert.Equal(t, "X", s.Elements()[0].Name())
}

func TestFromStructPtrEmptyNameOption(t *testing.T) {
	type MyStruct struct {
		X int64 `name:""`
	}

	x := &MyStruct{}
	s, err := settings.FromStructPtr(x)

	// Empty option values are ignored
	assert.NotNil(t, s)
	assert.NoError(t, err)
}

func TestFromStructPtrBadElemName(t *testing.T) {
	type MyStruct struct {
		X int64 `name:"x..."`
	}

	x := &MyStruct{}
	s, err := settings.FromStructPtr(x)

	assert.Nil(t, s)
	assert.Error(t, err)
}

func TestFromStructPtrBadOptionValue(t *testing.T) {
	type MyStruct struct {
		X int64 `less:"abc"`
	}

	x := &MyStruct{}
	s, err := settings.FromStructPtr(x)

	assert.Nil(t, s)
	assert.Error(t, err)
}
