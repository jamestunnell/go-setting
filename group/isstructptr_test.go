package group_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-settings/group"
)

func TestIsStructPointerGivenNonStruct(t *testing.T) {
	typ, ok := group.IsStructPointer(reflect.TypeOf(5))

	assert.Nil(t, typ)
	assert.False(t, ok)
}

func TestFromStructPtrGivenStruct(t *testing.T) {
	x := struct{ X int }{X: 5}
	typ, ok := group.IsStructPointer(reflect.TypeOf(x))

	assert.Nil(t, typ)
	assert.False(t, ok)
}

func TestFromStructPtrGivenNonStructPtr(t *testing.T) {
	x := 5
	typ, ok := group.IsStructPointer(reflect.TypeOf(&x))

	assert.Nil(t, typ)
	assert.False(t, ok)
}

func TestFromStructPtrHappyPath(t *testing.T) {
	type MyStruct struct{ X int }
	x := &MyStruct{X: 5}
	typ, ok := group.IsStructPointer(reflect.TypeOf(x))

	assert.NotNil(t, typ)
	assert.True(t, ok)
	assert.Equal(t, "MyStruct", typ.Name())
}
