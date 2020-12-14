package settings

import "github.com/jamestunnell/go-settings/element"

// Settings and its elements represent a struct and all its fields that
// are marked with the 'settings' tag.
type Settings struct {
	structptr interface{}
	name      string
	elements  []*element.Element
}

// StructPtr returns the struct pointer used to create the settings.
func (s *Settings) StructPtr() interface{} {
	return s.structptr
}

// Name returns the name of the settings, which is the name of the
// struct type from the struct pointer used to create the settings.
func (s *Settings) Name() string {
	return s.name
}

// Elements returns the setting elements.
func (s *Settings) Elements() []*element.Element {
	return s.elements
}

// Element looks for a setting element by name.
// Returns nil if the element is not found.
func (s *Settings) Element(name string) *element.Element {
	return findElement(s.elements, name)
}

func findElement(elems []*element.Element, name string) *element.Element {
	for _, elem := range elems {
		if elem.Name() == name {
			return elem
		}
	}

	return nil
}
