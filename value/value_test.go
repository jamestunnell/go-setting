package value_test

import (
	"testing"

	"github.com/jamestunnell/go-settings/value"
	"github.com/stretchr/testify/assert"
)

func TestStringKnownTypes(t *testing.T) {
	strings := []string{
		value.Int64.String(),
		value.Int64s.String(),
		value.UInt64.String(),
		value.UInt64s.String(),
		value.Float64.String(),
		value.Float64s.String(),
		value.Bool.String(),
		value.Bools.String(),
		value.String.String(),
		value.Strings.String(),
	}

	for i := 0; i < len(strings); i++ {
		assert.NotEmpty(t, strings[i])

		for j := 0; j < len(strings); j++ {
			if i != j {
				assert.NotEqual(t, strings[j], strings[i])
			}
		}
	}
}

func TestStringUnknownType(t *testing.T) {
	assert.Empty(t, value.Type(-1).String())
}
