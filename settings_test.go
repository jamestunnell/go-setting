package settings_test

import (
	"testing"

	"github.com/jamestunnell/go-settings"
	"github.com/stretchr/testify/assert"
)

func TestSettingsTwoElementsNoSubsettings(t *testing.T) {
	type MyStruct struct {
		A float64 `name:"a" default:"7.2"`
		B int64   `name:"b" greater:"20"`
	}

	x := &MyStruct{}

	s, err := settings.FromStructPtr("ABC", x)

	if !assert.NoError(t, err) || !assert.NotNil(t, s) {
		return
	}

	assert.Empty(t, s.Subsettings())
	assert.Len(t, s.Elements(), 2)
	assert.NotNil(t, s.Element("a"))
	assert.NotNil(t, s.Element("b"))
	assert.Nil(t, s.Element("xyz"))
	assert.Nil(t, s.Subsetting("xyz"))
}

func TestSettingsTwoSubsettings(t *testing.T) {
	type X struct {
		A float64 `name:"a" default:"7.2"`
		B int64   `name:"b" greater:"20"`
	}
	type MyStruct struct {
		Y *X
		Z *X
	}

	x := &MyStruct{}

	s, err := settings.FromStructPtr("ABC", x)

	if !assert.NoError(t, err) || !assert.NotNil(t, s) {
		return
	}

	assert.Empty(t, s.Elements())
	assert.Len(t, s.Subsettings(), 2)

	y := s.Subsetting("Y")
	assert.NotNil(t, y)

	z := s.Subsetting("Z")
	assert.NotNil(t, z)
}
