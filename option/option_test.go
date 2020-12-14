package option_test

import (
	"testing"

	"github.com/jamestunnell/go-settings/option"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	o := option.New(option.Default, 2.5)

	assert.Equal(t, "default(2.5)", o.String())
}

func TestTypeStringKnownTypes(t *testing.T) {
	strings := []string{
		option.Default.String(),
		option.OneOf.String(),
		option.MinLen.String(),
		option.MaxLen.String(),
		option.Greater.String(),
		option.GreaterEqual.String(),
		option.Less.String(),
		option.LessEqual.String(),
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

func TestTypeStringUnknownType(t *testing.T) {
	assert.Empty(t, option.Type(-1).String())
}
