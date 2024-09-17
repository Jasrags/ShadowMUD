package character_test

import (
	"testing"

	"github.com/Jasrags/ShadowMUD/common/character"

	"github.com/stretchr/testify/assert"
)

// var (
// 	cfg           = &config.Server{}
// 	metatypeHuman = &metatype.Metatype{
// 		Name: "Human",
// 		Attributes: metatype.Attributes{
// 			Body:      metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
// 			Agility:   metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
// 			Reaction:  metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
// 			Strength:  metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
// 			Willpower: metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
// 			Logic:     metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
// 			Intuition: metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
// 			Charisma:  metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
// 			Magic:     metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
// 			Resonance: metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
// 		},
// 	}
// )

// func TestSetName(t *testing.T) {
// 	name := "John Doe"
// 	pb := character.NewPointBuilder(cfg, nil)
// 	pb.SetName(name)
// 	assert.Equal(t, name, pb.Name)
// }

// func TestSetMetatype(t *testing.T) {
// 	data := []struct {
// 		metatype            *metatype.Metatype
// 		expectedBuildPoints int
// 		expectedError       error
// 	}{
// 		{&metatype.Metatype{Name: "Human", PointCost: 0}, character.TotalBuildPoints, nil},
// 		{&metatype.Metatype{Name: "Elf", PointCost: 40}, character.TotalBuildPoints - 40, nil},
// 		{&metatype.Metatype{Name: "Dwarf", PointCost: 50}, character.TotalBuildPoints - 50, nil},
// 		{&metatype.Metatype{Name: "Ork", PointCost: 50}, character.TotalBuildPoints - 50, nil},
// 		{&metatype.Metatype{Name: "Troll", PointCost: 90}, character.TotalBuildPoints - 90, nil},
// 		{&metatype.Metatype{Name: "Invalid", PointCost: character.TotalBuildPoints + 10}, character.TotalBuildPoints, fmt.Errorf("not enough build points")},
// 	}

// 	for _, d := range data {
// 		pb := character.NewPointBuilder(cfg, nil)
// 		err := pb.SetMetatype(d.metatype)

// 		if d.expectedError != nil {
// 			assert.Error(t, err)
// 			assert.Equal(t, d.expectedError, err)
// 			continue
// 		} else {
// 			assert.NoError(t, err)
// 		}

// 		assert.Equal(t, d.metatype, pb.Metatype)
// 		assert.Equal(t, d.expectedBuildPoints, pb.BuildPoints)
// 	}
// }

// func TestRemoveMetatype(t *testing.T) {
// 	data := []struct {
// 		metatype            *metatype.Metatype
// 		expectedBuildPoints int
// 		expectedError       error
// 	}{
// 		{&metatype.Metatype{Name: "Human", PointCost: 0}, 800, nil},
// 		{&metatype.Metatype{Name: "Elf", PointCost: 40}, 760, nil},
// 		{&metatype.Metatype{Name: "Dwarf", PointCost: 50}, 750, nil},
// 		{&metatype.Metatype{Name: "Ork", PointCost: 50}, 750, nil},
// 		{&metatype.Metatype{Name: "Troll", PointCost: 90}, 710, nil},
// 	}

// 	for _, d := range data {
// 		pb := character.NewPointBuilder(cfg, nil)
// 		pb.SetMetatype(d.metatype)

// 		pb.RemoveMetatype()
// 		assert.Empty(t, pb.Metatype)
// 		assert.Equal(t, character.TotalBuildPoints, pb.BuildPoints)
// 	}
// }

// func TestSetMagicType(t *testing.T) {
// 	data := []struct {
// 		magicType              character.MagicType
// 		expectedBuildPoints    int
// 		expectedMagicValue     int
// 		expectedResonanceValue int
// 	}{
// 		{character.MagicTypeNone, 800, 0, 0},             // 0
// 		{character.MagicTypeAdept, 780, 1, 0},            // 20
// 		{character.MagicTypeMagician, 785, 1, 0},         // 15
// 		{character.MagicTypeAspectedMagician, 770, 1, 0}, // 30
// 		{character.MagicTypeMysticAdept, 765, 1, 0},      // 35
// 		{character.MagicTypeTechnomancer, 785, 0, 1},     // 15
// 	}

// 	for _, d := range data {
// 		pb := character.NewPointBuilder(cfg, nil)
// 		pb.SetMetatype(metatypeHuman)
// 		err := pb.SetMagicType(d.magicType)

// 		assert.NoError(t, err)
// 		assert.Equal(t, d.magicType, pb.MagicType)
// 		assert.Equal(t, d.expectedBuildPoints, pb.BuildPoints)
// 		assert.Equal(t, d.expectedMagicValue, pb.Attributes[shared.AttributeMagic])
// 		assert.Equal(t, d.expectedResonanceValue, pb.Attributes[shared.AttributeResonance])
// 	}
// }

// func TestRemoveMagicType(t *testing.T) {
// 	pb := character.NewPointBuilder(cfg, nil)
// 	pb.SetMetatype(metatypeHuman)
// 	pb.MagicType = character.MagicTypeAdept
// 	pb.BuildPoints = 780

// 	// errMagic := pb.SetMagicType(character.MagicTypeAdept)
// 	pb.RemoveMagicType()
// 	assert.Equal(t, character.MagicTypeNone, pb.MagicType)
// 	assert.Equal(t, 800, pb.BuildPoints)
// 	assert.Equal(t, 0, pb.Attributes[shared.AttributeMagic])
// 	assert.Equal(t, 0, pb.Attributes[shared.AttributeResonance])
// }

// // TODO: Test for already having an attribute at metatype max
// // TODO: Test for not enough build points
// func TestAdjustAttribute(t *testing.T) {
// 	data := []struct {
// 		magicType           character.MagicType
// 		attribute           shared.AttributeType
// 		newValue            int
// 		expectedBuildPoints int
// 		expected            error
// 	}{
// 		{character.MagicTypeNone, shared.AttributeBody, 1, 800, nil}, // Minimum value
// 		{character.MagicTypeNone, shared.AttributeBody, 6, 800, nil}, // Maximum value
// 		{character.MagicTypeNone, shared.AttributeBody, 0, 800, fmt.Errorf("'Body' (0) can not be lowered below metatype minimum (1)")},
// 		{character.MagicTypeNone, shared.AttributeBody, 7, 800, fmt.Errorf("'Body' (7) can not be raised above metatype maximum (6)")},
// 		{character.MagicTypeNone, shared.AttributeAgility, 1, 800, nil}, // Minimum value
// 		{character.MagicTypeNone, shared.AttributeAgility, 6, 800, nil}, // Maximum value
// 		{character.MagicTypeNone, shared.AttributeAgility, 0, 800, fmt.Errorf("'Agility' (0) can not be lowered below metatype minimum (1)")},
// 		{character.MagicTypeNone, shared.AttributeAgility, 7, 800, fmt.Errorf("'Agility' (7) can not be raised above metatype maximum (6)")},
// 		{character.MagicTypeNone, shared.AttributeReaction, 1, 800, nil}, // Minimum value
// 		{character.MagicTypeNone, shared.AttributeReaction, 6, 800, nil}, // Maximum value
// 		{character.MagicTypeNone, shared.AttributeReaction, 0, 800, fmt.Errorf("'Reaction' (0) can not be lowered below metatype minimum (1)")},
// 		{character.MagicTypeNone, shared.AttributeReaction, 7, 800, fmt.Errorf("'Reaction' (7) can not be raised above metatype maximum (6)")},
// 		{character.MagicTypeNone, shared.AttributeStrength, 1, 800, nil}, // Minimum value
// 		{character.MagicTypeNone, shared.AttributeStrength, 6, 800, nil}, // Maximum value
// 		{character.MagicTypeNone, shared.AttributeStrength, 0, 800, fmt.Errorf("'Strength' (0) can not be lowered below metatype minimum (1)")},
// 		{character.MagicTypeNone, shared.AttributeStrength, 7, 800, fmt.Errorf("'Strength' (7) can not be raised above metatype maximum (6)")},
// 		{character.MagicTypeNone, shared.AttributeWillpower, 1, 800, nil}, // Minimum value
// 		{character.MagicTypeNone, shared.AttributeWillpower, 6, 800, nil}, // Maximum value
// 		{character.MagicTypeNone, shared.AttributeWillpower, 0, 800, fmt.Errorf("'Willpower' (0) can not be lowered below metatype minimum (1)")},
// 		{character.MagicTypeNone, shared.AttributeWillpower, 7, 800, fmt.Errorf("'Willpower' (7) can not be raised above metatype maximum (6)")},
// 		{character.MagicTypeNone, shared.AttributeLogic, 1, 800, nil}, // Minimum value
// 		{character.MagicTypeNone, shared.AttributeLogic, 6, 800, nil}, // Maximum value
// 		{character.MagicTypeNone, shared.AttributeLogic, 0, 800, fmt.Errorf("'Logic' (0) can not be lowered below metatype minimum (1)")},
// 		{character.MagicTypeNone, shared.AttributeLogic, 7, 800, fmt.Errorf("'Logic' (7) can not be raised above metatype maximum (6)")},
// 		{character.MagicTypeNone, shared.AttributeIntuition, 1, 800, nil}, // Minimum value
// 		{character.MagicTypeNone, shared.AttributeIntuition, 6, 800, nil}, // Maximum value
// 		{character.MagicTypeNone, shared.AttributeIntuition, 0, 800, fmt.Errorf("'Intuition' (0) can not be lowered below metatype minimum (1)")},
// 		{character.MagicTypeNone, shared.AttributeIntuition, 7, 800, fmt.Errorf("'Intuition' (7) can not be raised above metatype maximum (6)")},
// 		{character.MagicTypeNone, shared.AttributeCharisma, 1, 800, nil}, // Minimum value
// 		{character.MagicTypeNone, shared.AttributeCharisma, 6, 800, nil}, // Maximum value
// 		{character.MagicTypeNone, shared.AttributeCharisma, 0, 800, fmt.Errorf("'Charisma' (0) can not be lowered below metatype minimum (1)")},
// 		{character.MagicTypeNone, shared.AttributeCharisma, 7, 800, fmt.Errorf("'Charisma' (7) can not be raised above metatype maximum (6)")},
// 		{character.MagicTypeNone, shared.AttributeMagic, 1, 800, fmt.Errorf("can not adjust magic without being a magic user")},
// 		{character.MagicTypeNone, shared.AttributeResonance, 1, 800, fmt.Errorf("can not adjust resonance without being a technomancer")},
// 	}

// 	for _, d := range data {
// 		pb := character.NewPointBuilder(cfg, nil)
// 		pb.SetMetatype(metatypeHuman)
// 		pb.SetMagicType(d.magicType)

// 		err := pb.AdjustAttribute(d.attribute, d.newValue)
// 		if d.expected != nil {
// 			assert.Error(t, err)
// 			assert.Equal(t, d.expected, err)
// 		} else {
// 			assert.NoError(t, err)
// 			assert.Contains(t, pb.Attributes, d.attribute)
// 			assert.Equal(t, d.newValue, pb.Attributes[d.attribute])
// 		}
// 	}
// }

// func TestAdjustSkill(t *testing.T) {
// 	data := []struct {
// 		changeType          string
// 		skill               string
// 		newValue            int
// 		expectedNewValue    int
// 		expectedBuildPoints int
// 		expectedError       error
// 	}{
// 		// TODO: Test for not enough build points
// 		// TODO: Test for metatype not set
// 		// TODO: Test for magictype not set
// 		// TODO: Test for skills not available to the magic type
// 		{character.ChangeActiveSkill, "Running", 0, 0, 800, nil},                                                      // Active skill mo change
// 		{character.ChangeActiveSkill, "Running", 1, 1, 798, nil},                                                      // Active skill increate from 0 to 1
// 		{character.ChangeActiveSkill, "Sneaking", 2, 2, 796, nil},                                                     // Active skill increase from 1 to 2
// 		{character.ChangeActiveSkill, "Sneaking", 1, 1, 800, nil},                                                     // Active skill decrease from 2 to 1
// 		{character.ChangeKnowledgeSkill, "English", 1, 1, 799, nil},                                                   // Knowledge skill 0 to 1
// 		{character.ChangeSkillGroup, "Automatics", 1, 1, 795, nil},                                                    // Skill group 0 to 1
// 		{character.ChangeSkillGroup, "Running", -1, 0, 800, fmt.Errorf("skill value can not be negative")},            // Skill group decrease below 0
// 		{character.ChangeKnowledgeSkill, "English", 14, 0, 800, fmt.Errorf("skill value can not be greater than 13")}, // Knowledge skill increase above 13
// 		// {character.ChangeActiveSkill, "Hacking", 11, 0, 690, fmt.Errorf("not enough build points")},
// 	}

// 	for _, d := range data {
// 		pb := character.NewPointBuilder(cfg, nil)
// 		pb.SetMetatype(&metatype.Metatype{Name: "Human"})
// 		pb.SetMagicType(character.MagicTypeNone)
// 		pb.Skills = map[string]int{
// 			"Sneaking": 1,
// 		}
// 		err := pb.AdjustSkill(d.changeType, d.skill, d.newValue)
// 		if d.expectedError != nil {
// 			assert.Error(t, err)
// 			assert.Equal(t, d.expectedError, err)
// 		} else {
// 			assert.NoError(t, err)
// 			assert.Equal(t, d.expectedNewValue, pb.Skills[d.skill])
// 			assert.Equal(t, d.expectedBuildPoints, pb.BuildPoints)
// 		}
// 	}
// }

func TestCalculateChangeCost(t *testing.T) {
	data := []struct {
		changeType   string
		currentValue int
		desiredValue int
		expectedCost int
	}{
		// Active Skills
		{character.ChangeActiveSkill, 1, 1, 0},     // No change
		{character.ChangeActiveSkill, 1, 2, 4},     // 1 -> 2
		{character.ChangeActiveSkill, 2, 1, -4},    // 2 -> 1
		{character.ChangeActiveSkill, 12, 13, 26},  // 12 -> 13
		{character.ChangeActiveSkill, 13, 12, -26}, // 13 -> 12
		// Knowledge Skills
		{character.ChangeKnowledgeSkill, 1, 1, 0},     // No change
		{character.ChangeKnowledgeSkill, 1, 2, 2},     // 1 -> 2
		{character.ChangeKnowledgeSkill, 2, 1, -2},    // 2 -> 1
		{character.ChangeKnowledgeSkill, 12, 13, 13},  // 12 -> 13
		{character.ChangeKnowledgeSkill, 13, 12, -13}, // 13 -> 12
		// Skill groups
		{character.ChangeSkillGroup, 1, 1, 0},     // No change
		{character.ChangeSkillGroup, 1, 2, 10},    // 1 -> 2
		{character.ChangeSkillGroup, 2, 1, -10},   // 2 -> 1
		{character.ChangeSkillGroup, 12, 13, 65},  // 12 -> 13
		{character.ChangeSkillGroup, 13, 12, -65}, // 13 -> 12
		// Attributes
		{character.ChangeAttribute, 1, 1, 0},     // No change
		{character.ChangeAttribute, 1, 2, 10},    // 1 -> 2
		{character.ChangeAttribute, 2, 1, -10},   // 2 -> 1
		{character.ChangeAttribute, 12, 13, 65},  // 12 -> 13
		{character.ChangeAttribute, 13, 12, -65}, // 13 -> 12
	}

	for _, d := range data {
		cost := character.CalculateChangeCost(d.changeType, d.currentValue, d.desiredValue)
		assert.Equal(t, d.expectedCost, cost)
	}
}

// func TestAllocateQuality(t *testing.T) {
// 	data := []struct {
// 		buildPoints         int
// 		quality             string
// 		value               int
// 		positive            bool
// 		expectedBuildPoints int
// 		expectedError       error
// 	}{
// 		// {character.TotalBuildPoints, "Ambidextrous", 10, true, 790, nil},      // Add positive quality
// 		// {character.TotalBuildPoints, "Addiction (Mild)", 10, false, 810, nil}, // Add negative quality
// 		// {character.TotalBuildPoints, "Test", 10, true, 0, fmt.Errorf("quality already added")},
// 		// {1, "Not Enough", 10, true, 800, fmt.Errorf("not enough build points")},
// 	}

// 	for _, d := range data {
// 		pb := character.NewPointBuilder(cfg, nil)
// 		pb.SetMetatype(&metatype.Metatype{Name: "Human"})
// 		pb.SetMagicType(character.MagicTypeNone)
// 		pb.Qualities = map[string]int{"Test": 10}

// 		err := pb.AllocateQuality(d.quality, d.value, d.positive)
// 		if d.expectedError != nil {
// 			assert.Error(t, err)
// 			assert.Equal(t, d.expectedError, err)
// 		} else {
// 			assert.NoError(t, err)
// 			assert.Contains(t, pb.Qualities, d.quality)
// 			assert.Equal(t, d.expectedBuildPoints, pb.BuildPoints)
// 			assert.Equal(t, d.value, pb.Qualities[d.quality])
// 		}
// 	}

// }

// func TestRemoveQuality(t *testing.T) {
// 	data := []struct {
// 		quality             string
// 		expectedBuildPoints int
// 		expectedError       error
// 	}{
// 		{"Test", character.TotalBuildPoints, nil},
// 		{"Not Found", character.TotalBuildPoints, fmt.Errorf("quality not found")},
// 	}

// 	for _, d := range data {
// 		pb := character.NewPointBuilder(cfg, nil)
// 		pb.Qualities = map[string]int{"Test": 10}
// 		pb.BuildPoints = character.TotalBuildPoints - 10

// 		err := pb.RemoveQuality(d.quality)
// 		if d.expectedError != nil {
// 			assert.Error(t, err)
// 			assert.Equal(t, d.expectedError, err)
// 		} else {
// 			assert.NoError(t, err)
// 			assert.NotContains(t, pb.Qualities, d.quality)
// 			assert.Equal(t, d.expectedBuildPoints, pb.BuildPoints)
// 		}
// 	}
// }

// func TestPurchaseNuyen(t *testing.T) {
// 	data := []struct {
// 		buildPoints         int
// 		cost                int
// 		expectedNuyen       int
// 		expectedKarma       int
// 		expectedBuildPoints int
// 		expectedError       error
// 	}{
// 		{800, 1, 2000, 1, 799, nil},
// 		// {800, 0, 0, 1000, 0, 0, 800, fmt.Errorf("not enough build points")},
// 		// {800, 0, 0, character.KarmaNuyenConversionLimit, character.KarmaNuyenConversionLimit * character.KarmaNuyenConversionRate, character.KarmaNuyenConversionLimit, 0, nil},
// 		// {800, 0, character.KarmaNuyenConversionLimit - 1, 100, 100 * character.KarmaNuyenConversionRate, character.KarmaNuyenConversionLimit, 700, nil},
// 		{800, 201, 0, 0, 800, fmt.Errorf("can not convert more than %d karma to nuyen", character.KarmaNuyenConversionLimit)},
// 		{50, 100, 0, 0, 50, fmt.Errorf("not enough build points")},
// 	}

// 	for _, d := range data {
// 		pb := character.NewPointBuilder(cfg, nil)
// 		pb.BuildPoints = d.buildPoints

// 		err := pb.PurchaseNuyen(d.cost)
// 		if d.expectedError != nil {
// 			assert.Error(t, err)
// 			assert.Equal(t, d.expectedError, err)
// 		} else {
// 			assert.NoError(t, err)
// 			assert.Equal(t, d.expectedNuyen, pb.Nuyen)
// 			assert.Equal(t, d.expectedKarma, pb.KarmaForNuyen)
// 			assert.Equal(t, d.expectedBuildPoints, pb.BuildPoints)
// 			// assert.Equal(t, d.expectedBuildPoints, pb.BuildPoints)
// 			// assert.Equal(t, d.value, pb.Qualities[d.quality])
// 		}
// 	}
// }

// func TestAddSpell(t *testing.T) {
// 	data := []struct {
// 		buildPoints         int
// 		spell               string
// 		expectedBuildPoints int
// 		expectedError       error
// 	}{
// 		{character.TotalBuildPoints, "Acid Stream", character.TotalBuildPoints - character.SpellCost, nil},
// 		{0, "Acid Stream", character.TotalBuildPoints, fmt.Errorf("not enough build points")},
// 		{character.TotalBuildPoints - character.SpellCost, "Acid Test", character.TotalBuildPoints - character.SpellCost, fmt.Errorf("spell already added")},
// 	}

// 	for _, d := range data {
// 		pb := character.NewPointBuilder(cfg, nil)
// 		pb.Spells = map[string]int{"Acid Test": character.SpellCost}
// 		pb.BuildPoints = d.buildPoints

// 		err := pb.AddSpell(d.spell)
// 		if d.expectedError != nil {
// 			assert.Error(t, err)
// 			assert.Equal(t, d.expectedError, err)
// 		} else {
// 			assert.NoError(t, err)
// 			assert.Contains(t, pb.Spells, d.spell)
// 			assert.Equal(t, d.expectedBuildPoints, pb.BuildPoints)
// 		}
// 	}
// }

// func TestRemoveSpell(t *testing.T) {
// 	data := []struct {
// 		spell               string
// 		expectedBuildPoints int
// 		expectedError       error
// 	}{
// 		{"Acid Stream", character.TotalBuildPoints, nil},
// 		{"Acid Test", character.TotalBuildPoints, fmt.Errorf("spell not found")},
// 	}

// 	for _, d := range data {
// 		pb := character.NewPointBuilder(cfg, nil)
// 		pb.BuildPoints = character.TotalBuildPoints - character.SpellCost
// 		pb.Spells = map[string]int{"Acid Stream": character.SpellCost}

// 		err := pb.RemoveSpell(d.spell)
// 		if d.expectedError != nil {
// 			assert.Error(t, err)
// 			assert.Equal(t, d.expectedError, err)
// 		} else {
// 			assert.NoError(t, err)
// 			assert.NotContains(t, pb.Spells, d.spell)
// 			assert.Equal(t, d.expectedBuildPoints, pb.BuildPoints)
// 		}
// 	}
// }

// func TestAddComplexForm(t *testing.T) {
// 	data := []struct {
// 		buildPoints         int
// 		complexForm         string
// 		expectedBuildPoints int
// 		expectedError       error
// 	}{
// 		{character.TotalBuildPoints, "Resonance Spike", character.TotalBuildPoints - character.ComplexFormCost, nil},
// 		{0, "Resonance Spike", character.TotalBuildPoints, fmt.Errorf("not enough build points")},
// 		{character.TotalBuildPoints - character.ComplexFormCost, "Resonance Test", character.TotalBuildPoints - character.ComplexFormCost, fmt.Errorf("complex form already added")},
// 	}

// 	for _, d := range data {
// 		pb := character.NewPointBuilder(cfg, nil)
// 		pb.ComplexForms = map[string]int{"Resonance Test": character.ComplexFormCost}
// 		pb.BuildPoints = d.buildPoints

// 		err := pb.AddComplexForm(d.complexForm)
// 		if d.expectedError != nil {
// 			assert.Error(t, err)
// 			assert.Equal(t, d.expectedError, err)
// 		} else {
// 			assert.NoError(t, err)
// 			assert.Contains(t, pb.ComplexForms, d.complexForm)
// 			assert.Equal(t, d.expectedBuildPoints, pb.BuildPoints)
// 		}
// 	}
// }

// func TestRemoveComplexForm(t *testing.T) {
// 	data := []struct {
// 		complexForm         string
// 		expectedBuildPoints int
// 		expectedError       error
// 	}{
// 		{"Resonance Spike", character.TotalBuildPoints, nil},
// 		{"Resonance Test", character.TotalBuildPoints, fmt.Errorf("complex form not found")},
// 	}

// 	for _, d := range data {
// 		pb := character.NewPointBuilder(cfg, nil)
// 		pb.BuildPoints = character.TotalBuildPoints - character.ComplexFormCost
// 		pb.ComplexForms = map[string]int{"Resonance Spike": character.ComplexFormCost}

// 		err := pb.RemoveComplexForm(d.complexForm)
// 		if d.expectedError != nil {
// 			assert.Error(t, err)
// 			assert.Equal(t, d.expectedError, err)
// 		} else {
// 			assert.NoError(t, err)
// 			assert.NotContains(t, pb.ComplexForms, d.complexForm)
// 			assert.Equal(t, d.expectedBuildPoints, pb.BuildPoints)
// 		}
// 	}
// }
