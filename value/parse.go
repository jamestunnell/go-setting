package value

import (
	"fmt"
	"strconv"
	"strings"
)

// Parse attempts to parse the value string as the given type.
// Returns non-nil error in case of failure.
// Failure can be due to an unsupported value type or a parse failure.
func Parse(str string, typ Type) (interface{}, error) {
	switch typ {
	case Int64:
		i, err := strconv.ParseInt(str, 10, 64)
		return i, err
	case Int64s:
		substrings := strings.Split(str, ",")
		vals := make([]int64, len(substrings))

		for i := 0; i < len(substrings); i++ {
			substr := strings.TrimSpace(substrings[i])

			val, err := strconv.ParseInt(substr, 10, 64)
			if err != nil {
				return []int64{}, err
			}

			vals[i] = val
		}

		return vals, nil
	case UInt64:
		u, err := strconv.ParseUint(str, 10, 64)
		return u, err
	case UInt64s:
		substrings := strings.Split(str, ",")
		vals := make([]uint64, len(substrings))

		for i := 0; i < len(substrings); i++ {
			substr := strings.TrimSpace(substrings[i])

			val, err := strconv.ParseUint(substr, 10, 64)
			if err != nil {
				return []uint64{}, err
			}

			vals[i] = val
		}

		return vals, nil
	case Float64:
		f, err := strconv.ParseFloat(str, 64)
		return f, err
	case Float64s:
		substrings := strings.Split(str, ",")
		vals := make([]float64, len(substrings))

		for i := 0; i < len(substrings); i++ {
			substr := strings.TrimSpace(substrings[i])

			val, err := strconv.ParseFloat(substr, 64)
			if err != nil {
				return []float64{}, err
			}

			vals[i] = val
		}

		return vals, nil
	case Bool:
		b, err := strconv.ParseBool(str)
		return b, err
	case Bools:
		substrings := strings.Split(str, ",")
		vals := make([]bool, len(substrings))

		for i := 0; i < len(substrings); i++ {
			substr := strings.TrimSpace(substrings[i])

			val, err := strconv.ParseBool(substr)
			if err != nil {
				return []bool{}, err
			}

			vals[i] = val
		}

		return vals, nil
	case String:
		return str, nil
	case Strings:
		substrings := strings.Split(str, ",")
		trimmed := make([]string, len(substrings))

		for i := 0; i < len(substrings); i++ {
			trimmed[i] = strings.TrimSpace(substrings[i])
		}

		return trimmed, nil
	}

	return nil, fmt.Errorf("unsupported value type %d", typ)
}
