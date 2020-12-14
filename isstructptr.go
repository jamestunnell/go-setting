package settings

import "reflect"

// IsStructPointer returns true only if the given value is a struct pointer.
func IsStructPointer(s interface{}) (reflect.Type, bool) {
	tPtr := reflect.TypeOf(s)

	if tPtr.Kind() != reflect.Ptr {
		return nil, false
	}

	t := tPtr.Elem()
	if t.Kind() != reflect.Struct {
		return nil, false
	}

	return t, true
}
