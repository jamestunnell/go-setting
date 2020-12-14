package compare_test

import (
	"testing"

	"github.com/jamestunnell/go-settings/compare"
	"github.com/stretchr/testify/assert"
)

type val1val2 struct {
	Val1 interface{}
	Val2 interface{}
}

func newVal1Val2(v1, v2 interface{}) *val1val2 {
	return &val1val2{v1, v2}
}

func TestCompareFails(t *testing.T) {
	cases := []*val1val2{
		newVal1Val2(int(-2), int64(-2)), // int type is not supported
		newVal1Val2(int64(4), true),     // can't compare int64 to bool
		newVal1Val2(uint64(4), 2.5),     // can't compare uint64 to float64
		newVal1Val2(2.0, int64(2)),      // can't compare float64 to int64
	}

	for _, v1v2 := range cases {
		_, err := compare.Greater(v1v2.Val1, v1v2.Val2)

		assert.Error(t, err)

		_, err = compare.GreaterEqual(v1v2.Val1, v1v2.Val2)

		assert.Error(t, err)

		_, err = compare.Less(v1v2.Val1, v1v2.Val2)

		assert.Error(t, err)

		_, err = compare.LessEqual(v1v2.Val1, v1v2.Val2)

		assert.Error(t, err)
	}
}
