package settings

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"

	"github.com/jamestunnell/go-settings/element"
	"github.com/jamestunnell/go-settings/value"
)

var matchElemName = regexp.MustCompile(`^[_A-Za-z_][0-9A-Za-z_]*$`)

// FromStructPtr makes a new Settings using the given struct pointer.
// Returns non-nil error in case of failure.
// Failure can be due to: given value is not a pointer to a struct,
// unsupported field type, missing 'settings' tag, invalid element name,
// duplicate element name, invalid option format, invalid option param format,
// an option that is not applicable to the field type, or an option that is not
// compatible with another option for the same element.
func FromStructPtr(structptr interface{}) (*Settings, error) {
	t, ok := IsStructPointer(structptr)
	if !ok {
		return nil, errors.New("not a pointer to struct")
	}

	name := t.Name()
	elems := []*element.Element{}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		valType := getFieldValType(f.Type)
		if valType == -1 {
			err := fmt.Errorf("field %s has unsupported type %s", f.Name, f.Type.Name())
			return nil, err
		}

		elem, err := element.FromTag(f.Name, f.Tag, valType)
		if err != nil {
			err = fmt.Errorf("failed to make element for field %s: %w", f.Name, err)
			return nil, err
		}

		elemName := elem.Name()
		if !matchElemName.MatchString(elemName) {
			return nil, fmt.Errorf("element name '%s' has invalid format", elemName)
		}

		if findElement(elems, elemName) != nil {
			return nil, fmt.Errorf("element name '%s' is duplicated", elemName)
		}

		elems = append(elems, elem)
	}

	s := &Settings{name: name, elements: elems, structptr: structptr}

	return s, nil
}

func getFieldValType(t reflect.Type) value.Type {
	if t.Kind() == reflect.Slice {
		switch t.Elem().Name() {
		case "float64":
			return value.Float64s
		case "uint64":
			return value.UInt64s
		case "int64":
			return value.Int64s
		case "bool":
			return value.Bools
		case "string":
			return value.Strings
		}
	} else {
		switch t.Name() {
		case "float64":
			return value.Float64
		case "uint64":
			return value.UInt64
		case "int64":
			return value.Int64
		case "bool":
			return value.Bool
		case "string":
			return value.String
		}
	}

	return value.Type(-1)
}
