package setting

// MapByName is an alias
type MapByName = map[string]*Group

// Group and its elements represent a struct and its fields.
type Group struct {
	Elements  map[string]*Element
	Subgroups map[string]*Group
}

// FindElement searches for an element according to the given path.
// Returns nil if the element is not found or the path is empty.
func (g *Group) FindElement(path ...string) *Element {
	switch len(path) {
	case 0:
		return nil
	case 1:
		wantName := path[0]
		for name, elem := range g.Elements {
			if name == wantName {
				return elem
			}
		}
	default:
		wantName := path[0]
		for name, subgroup := range g.Subgroups {
			if name == wantName {
				return subgroup.FindElement(path[1:]...)
			}
		}
	}

	return nil
}
