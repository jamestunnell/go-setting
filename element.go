package setting

import (
	"fmt"

	"github.com/jamestunnell/go-setting/constraint"
	"github.com/jamestunnell/go-setting/value"
)

// Element is a setting group element, specifying value type and zero
// or more constraints. If the default constraint is present then the element
// is considered to be 'required'.
type Element struct {
	Value       value.Value
	Constraints []constraint.Constraint
}

// New makes a new element.
func NewElement(
	val value.Value,
	constraints ...constraint.Constraint,
) *Element {
	return &Element{Constraints: constraints, Value: val}
}

// CheckConstraints ensures that the constraints are all applicable to the element
// value, and that all constraints are compatible with each other.
// Returns non-nil error in case of failure.
func (e *Element) CheckConstraints() error {
	const (
		notApplicableFmt = "constraint type %s is not applicable to value %v"
		notCompatibleFmt = "constraint %v is not compatible with %v"
	)

	for i, c := range e.Constraints {
		cType := c.Type()

		if !cType.ApplicableTo(e.Value) {
			return fmt.Errorf(notApplicableFmt, cType, e.Value)
		}

		for j := i + 1; j < len(e.Constraints); j++ {
			c2 := e.Constraints[j]

			compatible, err := c.CompatibleWith(c2)
			if err != nil {
				return err
			}

			if !compatible {
				return fmt.Errorf(notCompatibleFmt, c2, c)
			}
		}
	}

	return nil
}

// Constraint returns the element constraint with the given type.
// Returns nil if not found.
func (e *Element) Constraint(cType constraint.Type) constraint.Constraint {
	for _, c := range e.Constraints {
		if c.Type() == cType {
			return c
		}
	}
	return nil
}
