package setting_test

import (
	"testing"

	"github.com/jamestunnell/go-setting"
	"github.com/jamestunnell/go-setting/constraint"
	"github.com/jamestunnell/go-setting/value"
	"github.com/stretchr/testify/assert"
)

func TestFindElementJustElements(t *testing.T) {
	g := newTestGroup()

	assert.NotNil(t, g.FindElement("A"))
	assert.NotNil(t, g.FindElement("B"))

	assert.Nil(t, g.FindElement("C"))
}

func TestFindElementWithSubgroup(t *testing.T) {
	g := &setting.Group{
		Elements: map[string]*setting.Element{},
		Subgroups: map[string]*setting.Group{
			"X": newTestGroup(),
			"Y": newTestGroup(),
		},
	}

	assert.NotNil(t, g.FindElement("X", "A"))
	assert.NotNil(t, g.FindElement("X", "B"))
	assert.NotNil(t, g.FindElement("Y", "B"))
	assert.NotNil(t, g.FindElement("Y", "B"))

	assert.Nil(t, g.FindElement("X", "C"))
}

func TestFindElementNoPathElems(t *testing.T) {
	g := newTestGroup()

	assert.Nil(t, g.FindElement())
}

func newTestGroup() *setting.Group {
	a := setting.NewElement(
		value.NewFloat(7.2), constraint.NewLess(value.NewFloat(10.0)))
	b := setting.NewElement(
		value.NewInt(25),
		constraint.NewLessEqual(value.NewInt(25)),
		constraint.NewGreater(value.NewInt(20)))
	g := &setting.Group{
		Elements:  map[string]*setting.Element{"A": a, "B": b},
		Subgroups: map[string]*setting.Group{},
	}

	return g
}
