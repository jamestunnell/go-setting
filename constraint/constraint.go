package constraint

import (
	"github.com/jamestunnell/go-setting/value"
)

// Constraint restricts a value according to the parameter
type Constraint interface {
	// Type returns the constraint type.
	Type() Type
	// Param returns the constraint parameter value.
	Param() value.Value
	// CompatibleWith returns true if the given constraint is compatible with the current one.
	// Returns a non-nil error in case of failure.
	CompatibleWith(Constraint) (bool, error)
}
