package value_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestBoolEqual(t *testing.T) {
	assert.True(t, value.BoolEqual(true, true))
	assert.True(t, value.BoolEqual(false, false))
	assert.False(t, value.BoolEqual(true, false))
	assert.False(t, value.BoolEqual(false, true))
}

func TestBoolGreater(t *testing.T) {
	assert.False(t, value.BoolGreater(true, true))
	assert.False(t, value.BoolGreater(false, false))
	assert.True(t, value.BoolGreater(true, false))
	assert.False(t, value.BoolGreater(false, true))
}

func TestBoolGreaterEqual(t *testing.T) {
	assert.True(t, value.BoolGreaterEqual(true, true))
	assert.True(t, value.BoolGreaterEqual(false, false))
	assert.True(t, value.BoolGreaterEqual(true, false))
	assert.False(t, value.BoolGreaterEqual(false, true))
}

func TestBoolLess(t *testing.T) {
	assert.False(t, value.BoolLess(true, true))
	assert.False(t, value.BoolLess(false, false))
	assert.False(t, value.BoolLess(true, false))
	assert.True(t, value.BoolLess(false, true))
}

func TestBoolLessEqual(t *testing.T) {
	assert.True(t, value.BoolLessEqual(true, true))
	assert.True(t, value.BoolLessEqual(false, false))
	assert.False(t, value.BoolLessEqual(true, false))
	assert.True(t, value.BoolLessEqual(false, true))
}
