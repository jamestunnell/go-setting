package element_test

import (
	"testing"

	"github.com/jamestunnell/go-settings/element"
	"github.com/jamestunnell/go-settings/option"
	"github.com/jamestunnell/go-settings/value"
	"github.com/stretchr/testify/assert"
)

func TestNewNoOptions(t *testing.T) {
	e := element.New("abc", value.Float64)

	assert.Equal(t, "abc", e.Name())
	assert.Equal(t, value.Float64, e.Type())
	assert.Nil(t, e.DefaultVal())
	assert.True(t, e.Required())
	assert.Len(t, e.Options(), 0)
	assert.Nil(t, e.Option(option.Default))
}

func TestNewWithDefault(t *testing.T) {
	def := option.New(option.Default, 2.5)
	e := element.New("abc", value.Float64, def)

	assert.Equal(t, 2.5, e.DefaultVal())
	assert.False(t, e.Required())
	assert.Len(t, e.Options(), 1)
	assert.NotNil(t, e.Option(option.Default))
}

func TestNewWithTwoOptions(t *testing.T) {
	def := option.New(option.Default, 2.5)
	less := option.New(option.Less, 5.0)
	e := element.New("abc", value.Float64, def, less)

	assert.Len(t, e.Options(), 2)
	assert.NotNil(t, e.Option(option.Default))
	assert.NotNil(t, e.Option(option.Less))
}
