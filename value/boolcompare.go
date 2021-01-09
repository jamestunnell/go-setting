package value

// CompareBoolFunc performs a comparison on the given bool values
type CompareBoolFunc func(a, b bool) bool

// BoolEqual returns true if the given bools are equal
func BoolEqual(a, b bool) bool {
	return a == b
}

// BoolGreater has truth table:
//  A | B | A > B
// ---|---|-------
//  F | F | F
//  F | T | F
//  T | F | T
//  T | T | F
func BoolGreater(a, b bool) bool {
	return a && !b
}

// BoolGreaterEqual has truth table:
//  A | B | A >= B
// ---|---|-------
//  F | F | T
//  F | T | F
//  T | F | T
//  T | T | T
func BoolGreaterEqual(a, b bool) bool {
	return a || !b
}

// BoolLess has truth table:
//  A | B | A < B
// ---|---|-------
//  F | F | F
//  F | T | T
//  T | F | F
//  T | T | F
func BoolLess(a, b bool) bool {
	return !a && b
}

// BoolLessEqual has truth table:
//  A | B | A <= B
// ---|---|-------
//  F | F | T
//  F | T | T
//  T | F | F
//  T | T | T
func BoolLessEqual(a, b bool) bool {
	return !a || b
}
