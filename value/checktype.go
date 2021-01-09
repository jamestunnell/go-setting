package value

import "fmt"

// CheckType returns a non-nil error if the types are not equal.
func CheckType(expected, actual Type) error {
	if expected != actual {
		return fmt.Errorf(
			"unexpected type %s, wanted %s", actual.String(), expected.String())
	}

	return nil
}
