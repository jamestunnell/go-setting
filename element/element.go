package element

import (
	"github.com/jamestunnell/go-settings/option"
	"github.com/jamestunnell/go-settings/value"
)

// Element is a named setting group element, specifying type and zero or
// more options. If the default option is present then the element
// is considered to be 'required' and must be given a value.
type Element struct {
	fieldName string
	typ       value.Type
	options   []*option.Option
}

// New makes a new element.
func New(
	fieldName string,
	typ value.Type,
	options ...*option.Option,
) *Element {
	return &Element{fieldName: fieldName, typ: typ, options: options}
}

// Name returns the element name.
func (e *Element) Name() string {
	if nameOpt := e.Option(option.Name); nameOpt != nil {
		return nameOpt.Param.(string)
	}

	return e.fieldName
}

// Type returns the element type.
func (e *Element) Type() value.Type {
	return e.typ
}

// Required true if the element has the default option.
func (e *Element) Required() bool {
	return e.Option(option.Default) == nil
}

// DefaultVal returns the param value from the default option if found, nil otherwise.
func (e *Element) DefaultVal() interface{} {
	o := e.Option(option.Default)
	if o == nil {
		return nil
	}

	return o.Param
}

// Options returns the element options.
func (e *Element) Options() []*option.Option {
	return e.options
}

// Option returns the option with the given type if found, nil otherwise.
func (e *Element) Option(optType option.Type) *option.Option {
	for _, opt := range e.options {
		if opt.Type == optType {
			return opt
		}
	}

	return nil
}
