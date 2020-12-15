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

// String a string representation of the option with the form <type>:"<param>"
func (o *Option) String() string {
	return fmt.Sprintf("%s:\"%v\"", o.Type.String(), o.Param)
}

// NewDefault makes an option that specifies default value
func NewDefault(param interface{}) *Option {
	return &Option{Type: Default, Param: param}
}

// NewMinLen makes an option that specifies min string/slice length
func NewMinLen(param interface{}) *Option {
	return &Option{Type: MinLen, Param: param}
}

// NewMaxLen makes an option that specifies max string/slice length
func NewMaxLen(param interface{}) *Option {
	return &Option{Type: MaxLen, Param: param}
}

// NewGreater makes an option that specifies a minimum (non-inclusive) value
func NewGreater(param interface{}) *Option {
	return &Option{Type: Greater, Param: param}
}

// NewGreaterEqual makes an option that specifies a minimum (inclusive) value
func NewGreaterEqual(param interface{}) *Option {
	return &Option{Type: GreaterEqual, Param: param}
}

// NewLess makes an option that specifies a maximum (non-inclusive) value
func NewLess(param interface{}) *Option {
	return &Option{Type: Less, Param: param}
}

// NewLessEqual makes an option that specifies a maximum (inclusive) value
func NewLessEqual(param interface{}) *Option {
	return &Option{Type: LessEqual, Param: param}
}

// NewOneOf makes an option that specifies an enumerated value
func NewOneOf(param interface{}) *Option {
	return &Option{Type: OneOf, Param: param}
}
