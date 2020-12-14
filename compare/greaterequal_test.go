package compare_test

import (
	"testing"

	"github.com/jamestunnell/go-settings/compare"
	"github.com/stretchr/testify/assert"
)

func TestGreaterEqual(t *testing.T) {
	cases := map[*val1val2]bool{
		newVal1Val2(int64(-2), int64(2)):  false,
		newVal1Val2(uint64(4), uint64(0)): true,
		newVal1Val2(3.3, 3.3):             true,
	}

	for v1v2, expected := range cases {
		result, err := compare.GreaterEqual(v1v2.Val1, v1v2.Val2)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	}
}
