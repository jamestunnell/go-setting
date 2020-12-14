package compare

import "fmt"

// Less returns true if the first value is less than the second.
// Returns non-nil error if the value types do not match.
// Returns non-nil error if the values are not int64, uint64, or float64.
func Less(v1, v2 interface{}) (bool, error) {
	switch val1 := v1.(type) {
	case int64:
		switch val2 := v2.(type) {
		case int64:
			return val1 < val2, nil
		default:
			return false, typeMismatchErr(v1, v2)
		}
	case uint64:
		switch val2 := v2.(type) {
		case uint64:
			return val1 < val2, nil
		default:
			return false, typeMismatchErr(v1, v2)
		}
	case float64:
		switch val2 := v2.(type) {
		case float64:
			return val1 < val2, nil
		default:
			return false, typeMismatchErr(v1, v2)
		}
	}

	return false, fmt.Errorf("type for %v does not support <", v1)
}
