package value_test

import (
	"testing"

	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestTypeValid(t *testing.T) {
	assert.True(t, value.TypeUInt.Valid())
	assert.False(t, value.Type(-1).Valid())
}

func TestTypeString(t *testing.T) {
	assert.NotEmpty(t, value.TypeUInt.String())
	assert.Empty(t, value.Type(-1).String())
}
