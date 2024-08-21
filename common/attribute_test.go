package common_test

// func TestNewAttribute(t *testing.T) {
// 	name := common.AttributeType("TestAttribute")
// 	base := 10

// 	attr := common.NewAttribute(name, base)

// 	assert.Equal(t, base, attr.Base)
// }

// func TestAttribute_SetBase(t *testing.T) {
// 	name := common.AttributeType("TestAttribute")
// 	base := 10
// 	newBase := 15

// 	attr := common.NewAttribute(name, base)
// 	attr.SetBase(newBase)

// 	assert.Equal(t, newBase, attr.Base)
// }

// func TestAttribute_AddDelta(t *testing.T) {
// 	name := common.AttributeType("TestAttribute")
// 	base := 10
// 	delta := 5
// 	expectedTotalValue := base + delta

// 	attr := common.NewAttribute(name, base)
// 	attr.AddDelta(delta)

// 	assert.Equal(t, expectedTotalValue, attr.TotalValue)
// }

// func TestAttribute_SubDelta(t *testing.T) {
// 	name := common.AttributeType("TestAttribute")
// 	base := 10
// 	delta := 5
// 	expectedTotalValue := base - delta

// 	attr := common.NewAttribute(name, base)
// 	attr.SubDelta(delta)

// 	assert.Equal(t, expectedTotalValue, attr.TotalValue)
// }

// func TestAttribute_Recalculate(t *testing.T) {
// 	name := common.AttributeType("TestAttribute")
// 	base := 10
// 	delta := 5
// 	expectedTotalValue := base + delta

// 	attr := common.NewAttribute(name, base)
// 	attr.AddDelta(delta)

// 	assert.Equal(t, expectedTotalValue, attr.TotalValue)
// }

// func TestAttributes_Recalculate(t *testing.T) {
// 	attrs := &common.Attributes{
// 		Body:      *common.NewAttribute("Body", 10),
// 		Agility:   *common.NewAttribute("Agility", 10),
// 		Reaction:  *common.NewAttribute("Reaction", 10),
// 		Strength:  *common.NewAttribute("Strength", 10),
// 		Willpower: *common.NewAttribute("Willpower", 10),
// 		Logic:     *common.NewAttribute("Logic", 10),
// 		Intuition: *common.NewAttribute("Intuition", 10),
// 		Charisma:  *common.NewAttribute("Charisma", 10),
// 		Edge:      *common.NewAttribute("Edge", 10),
// 		Essence:   *common.NewAttribute("Essence", 10.0),
// 		Magic:     *common.NewAttribute("Magic", 10),
// 		Resonance: *common.NewAttribute("Resonance", 10),
// 	}

// 	attrs.Body.AddDelta(5)
// 	attrs.Essence.SubDelta(5)
// 	attrs.Recalculate()

// 	assert.Equal(t, 15, attrs.Body.TotalValue)
// 	assert.Equal(t, 5.0, attrs.Essence.TotalValue)
// }
