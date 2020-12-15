package group_test

import (
	"testing"

	"github.com/jamestunnell/go-settings/element"
	"github.com/jamestunnell/go-settings/group"
	"github.com/jamestunnell/go-settings/option"
	"github.com/jamestunnell/go-settings/value"
	"github.com/stretchr/testify/assert"
)

func TestFromStructGivenStructNonPointer(t *testing.T) {
	type MyStruct struct{}

	s, err := group.FromStructPtr("ABC", MyStruct{})

	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestFromStructPtrEmpty(t *testing.T) {
	type MyStruct struct{}

	x := &MyStruct{}
	s, err := group.FromStructPtr("ABC", x)

	if !assert.NoError(t, err) || !assert.NotNil(t, s) {
		return
	}

	assert.Equal(t, "ABC", s.Name())
	assert.Len(t, s.Elements(), 0)
	assert.Equal(t, x, s.StructPtr())
}

func TestFromStructPtrFieldsForAllSupportedValueTypes(t *testing.T) {
	type MyStruct struct {
		A float64 `default:"7.2"`
		B int64   `greater:"20"`
		C uint64  `less:"100"`
		D string  `minLen:"10"`
		E bool
		F []float64 `default:"7.2,4.4"`
		G []int64   `greater:"20"`
		H []uint64  `less:"100"`
		I []string  `minLen:"10"`
		J []bool    `maxLen:"7"`
	}

	expectedElems := []*element.Element{
		element.New("A", value.Float64, option.New(option.Default, 7.2)),
		element.New("B", value.Int64, option.New(option.Greater, int64(20))),
		element.New("C", value.UInt64, option.New(option.Less, uint64(100))),
		element.New("D", value.String, option.New(option.MinLen, uint64(10))),
		element.New("E", value.Bool),
		element.New("F", value.Float64s, option.New(option.Default, []float64{7.2, 4.4})),
		element.New("G", value.Int64s, option.New(option.Greater, int64(20))),
		element.New("H", value.UInt64s, option.New(option.Less, uint64(100))),
		element.New("I", value.Strings, option.New(option.MinLen, uint64(10))),
		element.New("J", value.Bools, option.New(option.MaxLen, uint64(7))),
	}

	x := &MyStruct{}
	s, err := group.FromStructPtr("XYZ", x)

	if !assert.NoError(t, err) || !assert.NotNil(t, s) {
		return
	}

	assert.Equal(t, "XYZ", s.Name())
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

func TestFromStructPtrInvalidOptionTag(t *testing.T) {
	type MyStruct struct {
		X int `default=5`
	}

	x := &MyStruct{}
	s, err := group.FromStructPtr("ABC", x)

	assert.Nil(t, s)
	assert.Error(t, err)
}

func TestFromStructPtrMissingNameOption(t *testing.T) {
	type MyStruct struct {
		X int64 `json:"x"`
	}

	x := &MyStruct{}
	s, err := group.FromStructPtr("ABC", x)

	assert.NotNil(t, s)
	assert.NoError(t, err)
	assert.Len(t, s.Elements(), 1)

	// Field name is used by default
	assert.Equal(t, "X", s.Elements()[0].Name())
}

func TestFromStructPtrBadOptionValue(t *testing.T) {
	type MyStruct struct {
		X int64 `less:"abc"`
	}

	x := &MyStruct{}
	s, err := group.FromStructPtr("ABC", x)

	assert.Nil(t, s)
	assert.Error(t, err)
}

func TestFromStructWithSubsetting(t *testing.T) {
	type A struct {
		X int64 `less:"2"`
	}

	type B struct {
		A *A
	}

	b := &B{}
	s, err := group.FromStructPtr("ABC", b)

	assert.NotNil(t, s)
	assert.NoError(t, err)
	assert.Len(t, s.Groups(), 1)
}

func TestFromStructFailToMakeSubsetting(t *testing.T) {
	type A struct {
		X int64 `less:"true"`
	}

	type B struct {
		A *A
	}

	b := &B{}
	s, err := group.FromStructPtr("ABC", b)

	assert.Nil(t, s)
	assert.Error(t, err)
}
