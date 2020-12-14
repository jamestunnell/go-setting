package settings_test

import (
	"testing"

	"github.com/jamestunnell/go-settings"
	"github.com/stretchr/testify/assert"
)

func TestSettingsElement(t *testing.T) {
	type MyStruct struct {
		A float64 `name:"a" default:"7.2"`
		B int64   `name:"b" greater:"20"`
	}

	x := &MyStruct{}

	s, err := settings.FromStructPtr(x)

	if !assert.NoError(t, err) || !assert.NotNil(t, s) {
		return
	}

	assert.NotNil(t, s.Element("a"))
	assert.NotNil(t, s.Element("b"))
	assert.Nil(t, s.Element("xyz"))
}
