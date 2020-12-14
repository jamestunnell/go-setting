package option_test

import (
	"testing"

	"github.com/jamestunnell/go-settings/option"
	"github.com/stretchr/testify/assert"
)

func TestCompatibleWithMinMaxLen(t *testing.T) {
	minLen := option.New(option.MinLen, uint64(10))

	incompatible := []*option.Option{
		option.New(option.MinLen, uint64(10)),
		option.New(option.MaxLen, uint64(9)),
	}

	verifyCompatiblility(t, minLen, incompatible, false)

	compatible := []*option.Option{
		option.New(option.MaxLen, uint64(10)),
		option.New(option.MaxLen, uint64(11)),
		option.New(option.Greater, 2.5),
		option.New(option.GreaterEqual, -5),
		option.New(option.Less, uint64(0)),
		option.New(option.LessEqual, 7.7),
		option.New(option.OneOf, []float64{2.5, 5.5}),
	}

	verifyCompatiblility(t, minLen, compatible, true)
}

func TestCompatibleWithLess(t *testing.T) {
	less := option.New(option.Less, uint64(5))

	incompatible := []*option.Option{
		option.New(option.Less, uint64(5)),
		option.New(option.LessEqual, uint64(5)),
		option.New(option.Greater, uint64(5)),
		option.New(option.GreaterEqual, uint64(5)),
	}

	verifyCompatiblility(t, less, incompatible, false)

	compatible := []*option.Option{
		option.New(option.Greater, uint64(4)),
		option.New(option.GreaterEqual, uint64(4)),
	}

	verifyCompatiblility(t, less, compatible, true)
}

func TestCompatibleWithLessEqual(t *testing.T) {
	lessEqual := option.New(option.LessEqual, uint64(5))

	incompatible := []*option.Option{
		option.New(option.Less, uint64(5)),
		option.New(option.LessEqual, uint64(5)),
		option.New(option.Greater, uint64(5)),
	}

	verifyCompatiblility(t, lessEqual, incompatible, false)

	compatible := []*option.Option{
		option.New(option.Greater, uint64(4)),
		option.New(option.GreaterEqual, uint64(5)),
	}

	verifyCompatiblility(t, lessEqual, compatible, true)
}

func TestCompatibleWithOneOf(t *testing.T) {
	oneOf := option.New(option.OneOf, []float64{0.0, 2.5})

	incompatible := []*option.Option{
		option.New(option.OneOf, []float64{0.0, 2.5}),
		option.New(option.Less, 6.5),
		option.New(option.LessEqual, 6.5),
		option.New(option.Greater, -1.0),
		option.New(option.GreaterEqual, -1.0),
		option.New(option.Default, 2.0),
	}

	verifyCompatiblility(t, oneOf, incompatible, false)

	compatible := []*option.Option{
		option.New(option.MinLen, uint64(4)),
		option.New(option.MaxLen, uint64(4)),
		option.New(option.Default, 2.5),
	}

	verifyCompatiblility(t, oneOf, compatible, true)
}

func TestCompatibleWithDefault(t *testing.T) {
	def := option.New(option.Default, 2.5)

	incompatible := []*option.Option{
		option.New(option.OneOf, []float64{0.0, 2.7}),
		option.New(option.Less, 2.5),
		option.New(option.LessEqual, 2.4),
		option.New(option.Greater, 2.5),
		option.New(option.GreaterEqual, 2.6),
		option.New(option.Default, 2.5),
	}

	verifyCompatiblility(t, def, incompatible, false)

	compatible := []*option.Option{
		option.New(option.MinLen, uint64(4)),
		option.New(option.MaxLen, uint64(4)),
		option.New(option.OneOf, []float64{0.0, 2.5}),
		option.New(option.Less, 2.6),
		option.New(option.LessEqual, 2.5),
		option.New(option.Greater, 2.4),
		option.New(option.GreaterEqual, 2.5),
	}

	verifyCompatiblility(t, def, compatible, true)
}

func verifyCompatiblility(
	t *testing.T,
	o1 *option.Option,
	options []*option.Option,
	expected bool,
) {
	for _, o2 := range options {
		compatible, err := o2.CompatibleWith(o1)

		assert.NoError(t, err)
		assert.Equal(t, expected, compatible)

		compatible, err = o1.CompatibleWith(o2)

		assert.NoError(t, err)
		assert.Equal(t, expected, compatible)
	}
}
