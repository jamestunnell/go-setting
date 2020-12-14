package option

import "github.com/jamestunnell/go-settings/compare"

// CompatibleWith checks the current option with the given option to see if
// they are compatible.
// Options with the same type are incompatible.
// Options with less and lessEqual type are incompatible.
// Options with greater and greaterEqual type are incompatible.
// Options with oneOf and any other compare type (<, <=, >, >=) are incompatible.
// For the remaining cases, compatibility is determined by comparing the option
// param values. As an example less(2.0) and greater(0.0) are compatible, but
// less(2.0) and greater(2.0) are not.
// Returns a non-nil error in case of failure (e.g., for param values that cannot be compared).
func (o1 *Option) CompatibleWith(o2 *Option) (bool, error) {
	// duplicates are not allowed
	if o1.Type == o2.Type {
		return false, nil
	}

	switch o1.Type {
	case Default:
		switch o2.Type {
		case Greater:
			return compare.Greater(o1.Param, o2.Param)
		case GreaterEqual:
			return compare.GreaterEqual(o1.Param, o2.Param)
		case Less:
			return compare.Less(o1.Param, o2.Param)
		case LessEqual:
			return compare.LessEqual(o1.Param, o2.Param)
		case OneOf:
			return compare.OneOf(o1.Param, o2.Param)
		}
	case MinLen:
		if o2.Type == MaxLen {
			return compare.LessEqual(o1.Param, o2.Param)
		}
	case MaxLen:
		if o2.Type == MinLen {
			return compare.LessEqual(o2.Param, o1.Param)
		}
	case Greater:
		switch o2.Type {
		case GreaterEqual, OneOf:
			return false, nil
		case Less, LessEqual:
			return compare.Less(o1.Param, o2.Param)
		case Default:
			return compare.Greater(o2.Param, o1.Param)
		}
	case GreaterEqual:
		switch o2.Type {
		case Greater, OneOf:
			return false, nil
		case Less:
			return compare.Less(o1.Param, o2.Param)
		case LessEqual:
			return compare.LessEqual(o1.Param, o2.Param)
		case Default:
			return compare.GreaterEqual(o2.Param, o1.Param)
		}
	case Less:
		switch o2.Type {
		case LessEqual, OneOf:
			return false, nil
		case Greater, GreaterEqual:
			return compare.Less(o2.Param, o1.Param)
		case Default:
			return compare.Less(o2.Param, o1.Param)
		}
	case LessEqual:
		switch o2.Type {
		case Less, OneOf:
			return false, nil
		case Greater:
			return compare.Less(o2.Param, o1.Param)
		case GreaterEqual:
			return compare.LessEqual(o2.Param, o1.Param)
		case Default:
			return compare.LessEqual(o2.Param, o1.Param)
		}
	case OneOf:
		switch o2.Type {
		case Default:
			return compare.OneOf(o2.Param, o1.Param)
		case Greater, GreaterEqual, Less, LessEqual:
			return false, nil
		}
	}

	return true, nil
}
