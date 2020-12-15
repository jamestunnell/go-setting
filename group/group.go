package group

import "github.com/jamestunnell/go-settings/element"

// Group and its elements represent a struct and all its exported fields.
type Group struct {
	structptr interface{}
	name      string
	elements  []*element.Element
	subgroups []*Group
}

// StructPtr returns the struct pointer used to create the group.
func (g *Group) StructPtr() interface{} {
	return g.structptr
}

// Name returns the name of the group.
func (g *Group) Name() string {
	return g.name
}

// Elements returns the setting elements.
func (g *Group) Elements() []*element.Element {
	return g.elements
}

// Element looks for a setting element by name.
// Returns nil if not found.
func (g *Group) Element(name string) *element.Element {
	return findElement(g.elements, name)
}

// Groups returns all of the groups that are nested under
// the current group.
func (g *Group) Groups() []*Group {
	return g.subgroups
}

// Group looks for a sub-group by name.
// Returns nil if not found.
func (g *Group) Group(name string) *Group {
	return findGroup(g.subgroups, name)
}

func findElement(elems []*element.Element, name string) *element.Element {
	for _, elem := range elems {
		if elem.Name() == name {
			return elem
		}
	}

	return nil
}

func findGroup(subgroups []*Group, name string) *Group {
	for _, subsetting := range subgroups {
		if subsetting.Name() == name {
			return subsetting
		}
	}

	return nil
}
