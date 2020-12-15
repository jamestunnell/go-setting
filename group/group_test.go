package group_test

import (
	"testing"

	"github.com/jamestunnell/go-settings/group"
	"github.com/stretchr/testify/assert"
)

func TestGroupTwoElementsNoSubsettings(t *testing.T) {
	type MyStruct struct {
		A float64 `less:"10.0" default:"7.2"`
		B int64   `lessEqual:"25" greater:"20"`
	}

	x := &MyStruct{}

	s, err := group.FromStructPtr("ABC", x)

	if !assert.NoError(t, err) || !assert.NotNil(t, s) {
		return
	}

	assert.Empty(t, s.Groups())
	assert.Len(t, s.Elements(), 2)
	assert.NotNil(t, s.Element("A"))
	assert.NotNil(t, s.Element("B"))
	assert.Nil(t, s.Element("xyz"))
	assert.Nil(t, s.Group("xyz"))
}

func TestGroupTwoSubsettings(t *testing.T) {
	type X struct {
		A float64 `less:"10.0" default:"7.2"`
		B int64   `lessEqual:"25" greater:"20"`
	}
	type MyStruct struct {
		Y *X
		Z *X
	}

	x := &MyStruct{}

	s, err := group.FromStructPtr("ABC", x)

	if !assert.NoError(t, err) || !assert.NotNil(t, s) {
		return
	}

	assert.Empty(t, s.Elements())
	assert.Len(t, s.Groups(), 2)

	y := s.Group("Y")
	assert.NotNil(t, y)

	z := s.Group("Z")
	assert.NotNil(t, z)
}
