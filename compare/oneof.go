package compare

import "fmt"

// OneOf returns true if the first value is equal to an element of the second.
// Returns non-nil error if the first value is not int64, uint64, or float64.
// Returns non-nil error if the second value is not the sliced type of first value (e.g. int64 and []int64)
func OneOf(v1, v2 interface{}) (bool, error) {
	switch val1 := v1.(type) {
	case int64:
		switch val2 := v2.(type) {
		case []int64:
			for _, x := range val2 {
				if x == val1 {
					return true, nil
				}
			}

			return false, nil
		default:
			return false, fmt.Errorf("%v does not have expected type []int64", v2)
		}
	case uint64:
		switch val2 := v2.(type) {
		case []uint64:
			for _, x := range val2 {
				if x == val1 {
					return true, nil
				}
			}

			return false, nil
		default:
			return false, fmt.Errorf("%v does not have expected type []uint64", v2)
		}
	case float64:
		switch val2 := v2.(type) {
		case []float64:
			for _, x := range val2 {
				if x == val1 {
					return true, nil
				}
			}

			return false, nil
		default:
			return false, fmt.Errorf("%v does not have expected type []float64", v2)
		}
	case string:
		switch val2 := v2.(type) {
		case []string:
			for _, x := range val2 {
				if x == val1 {
					return true, nil
				}
			}

			return false, nil
		default:
			return false, fmt.Errorf("%v does not have expected type []string", v2)
		}
	}

	return false, fmt.Errorf("type for %v does not support oneOf", v1)
}
