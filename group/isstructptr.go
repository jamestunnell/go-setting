package group

import "reflect"

// IsStructPointer returns true only if the given type is for a
// struct pointer.
// Also returns the struct (non-pointer) type.
func IsStructPointer(t reflect.Type) (reflect.Type, bool) {
	if t.Kind() != reflect.Ptr {
		return nil, false
	}

	t2 := t.Elem()
	if t2.Kind() != reflect.Struct {
		return nil, false
	}

	return t2, true
}
