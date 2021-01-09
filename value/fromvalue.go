package value

import "reflect"

// FromValue makes a new value from the given reflect.Value.
// Returns nil if the given value type is not supported.
func FromValue(v reflect.Value) Value {
	t := v.Type()

	if t.Kind() == reflect.Ptr {
		t = t.Elem()

		if t.Kind() == reflect.Slice {
			switch t.Elem().Name() {
			case "float64":
				return NewFloatSliceFromPtr(v.Interface().(*[]float64))
			case "uint64":
				return NewUIntSliceFromPtr(v.Interface().(*[]uint64))
			case "int64":
				return NewIntSliceFromPtr(v.Interface().(*[]int64))
			case "bool":
				return NewBoolSliceFromPtr(v.Interface().(*[]bool))
			case "string":
				return NewStringSliceFromPtr(v.Interface().(*[]string))
			}
		} else {
			switch t.Name() {
			case "float64":
				return NewFloatFromPtr(v.Interface().(*float64))
			case "uint64":
				return NewUIntFromPtr(v.Interface().(*uint64))
			case "int64":
				return NewIntFromPtr(v.Interface().(*int64))
			case "bool":
				return NewBoolFromPtr(v.Interface().(*bool))
			case "string":
				return NewStringFromPtr(v.Interface().(*string))
			}
		}
	} else if t.Kind() == reflect.Slice {
		switch t.Elem().Name() {
		case "float64":
			return NewFloatSlice(v.Interface().([]float64)...)
		case "uint64":
			return NewUIntSlice(v.Interface().([]uint64)...)
		case "int64":
			return NewIntSlice(v.Interface().([]int64)...)
		case "bool":
			return NewBoolSlice(v.Interface().([]bool)...)
		case "string":
			return NewStringSlice(v.Interface().([]string)...)
		}
	} else {
		switch t.Name() {
		case "float64":
			return NewFloat(v.Interface().(float64))
		case "uint64":
			return NewUInt(v.Interface().(uint64))
		case "int64":
			return NewInt(v.Interface().(int64))
		case "bool":
			return NewBool(v.Interface().(bool))
		case "string":
			return NewString(v.Interface().(string))
		}
	}

	return nil
}
