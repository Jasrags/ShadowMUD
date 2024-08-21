package shared_test

import (
	"testing"

	"github.com/Jasrags/ShadowMUD/common/shared"

	"github.com/stretchr/testify/assert"
)

func TestNewAttribute(t *testing.T) {
	name := shared.AttributeType("TestAttribute")
	base := 10

	attr := shared.NewAttribute(name, base)

	assert.Equal(t, base, attr.Base)
}

func TestAttribute_SetBase(t *testing.T) {
	name := shared.AttributeType("TestAttribute")
	base := 10
	newBase := 15

	attr := shared.NewAttribute(name, base)
	attr.SetBase(newBase)

	assert.Equal(t, newBase, attr.Base)
}

func TestAttribute_AddDelta(t *testing.T) {
	name := shared.AttributeType("TestAttribute")
	base := 10
	delta := 5
	expectedTotalValue := base + delta

	attr := shared.NewAttribute(name, base)
	attr.AddDelta(delta)

	assert.Equal(t, expectedTotalValue, attr.TotalValue)
}

func TestAttribute_SubDelta(t *testing.T) {
	name := shared.AttributeType("TestAttribute")
	base := 10
	delta := 5
	expectedTotalValue := base - delta

	attr := shared.NewAttribute(name, base)
	attr.SubDelta(delta)

	assert.Equal(t, expectedTotalValue, attr.TotalValue)
}

func TestAttribute_Recalculate(t *testing.T) {
	name := shared.AttributeType("TestAttribute")
	base := 10
	delta := 5
	expectedTotalValue := base + delta

	attr := shared.NewAttribute(name, base)
	attr.AddDelta(delta)

	assert.Equal(t, expectedTotalValue, attr.TotalValue)
}

func TestAttributes_Recalculate(t *testing.T) {
	attrs := &shared.Attributes{
		Body:      *shared.NewAttribute("Body", 10),
		Agility:   *shared.NewAttribute("Agility", 10),
		Reaction:  *shared.NewAttribute("Reaction", 10),
		Strength:  *shared.NewAttribute("Strength", 10),
		Willpower: *shared.NewAttribute("Willpower", 10),
		Logic:     *shared.NewAttribute("Logic", 10),
		Intuition: *shared.NewAttribute("Intuition", 10),
		Charisma:  *shared.NewAttribute("Charisma", 10),
		Edge:      *shared.NewAttribute("Edge", 10),
		Essence:   *shared.NewAttribute("Essence", 10.0),
		Magic:     *shared.NewAttribute("Magic", 10),
		Resonance: *shared.NewAttribute("Resonance", 10),
	}

	attrs.Body.AddDelta(5)
	attrs.Essence.SubDelta(5)
	attrs.Recalculate()

	assert.Equal(t, 15, attrs.Body.TotalValue)
	assert.Equal(t, 5.0, attrs.Essence.TotalValue)
}
