package option

import (
	"fmt"
)

// Option combines type with a parameter value, as in less(2.0) or default(0.0)
type Option struct {
	Type  Type
	Param interface{}
}

// New makes a new option with the given type and value.
func New(typ Type, param interface{}) *Option {
	return &Option{Type: typ, Param: param}
}

// String a string representation of the option with the form <type>(<param>)
func (o *Option) String() string {
	return fmt.Sprintf("%s(%v)", o.Type.String(), o.Param)
}
