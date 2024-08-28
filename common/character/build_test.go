package character_test

import (
	"fmt"
	"testing"

	"github.com/Jasrags/ShadowMUD/common/character"
	"github.com/Jasrags/ShadowMUD/common/metatype"
	"github.com/Jasrags/ShadowMUD/common/shared"

	"github.com/stretchr/testify/assert"
)

func TestSetName(t *testing.T) {
	name := "John Doe"
	pb := character.NewPointBuilder()
	pb.SetName(name)
	assert.Equal(t, name, pb.Name)
}

func TestSetMetatype(t *testing.T) {
	data := []struct {
		metatype            *metatype.Metatype
		expectedBuildPoints int
		expectedError       error
	}{
		{&metatype.Metatype{Name: "Human", PointCost: 0}, character.TotalBuildPoints, nil},
		{&metatype.Metatype{Name: "Elf", PointCost: 40}, character.TotalBuildPoints - 40, nil},
		{&metatype.Metatype{Name: "Dwarf", PointCost: 50}, character.TotalBuildPoints - 50, nil},
		{&metatype.Metatype{Name: "Ork", PointCost: 50}, character.TotalBuildPoints - 50, nil},
		{&metatype.Metatype{Name: "Troll", PointCost: 90}, character.TotalBuildPoints - 90, nil},
		{&metatype.Metatype{Name: "Invalid", PointCost: character.TotalBuildPoints + 10}, character.TotalBuildPoints, fmt.Errorf("not enough build points")},
	}

	for _, d := range data {
		pb := character.NewPointBuilder()
		err := pb.SetMetatype(d.metatype)

		if d.expectedError != nil {
			assert.Error(t, err)
			assert.Equal(t, d.expectedError, err)
			continue
		} else {
			assert.NoError(t, err)
		}

		assert.Equal(t, d.metatype, pb.Metatype)
		assert.Equal(t, d.expectedBuildPoints, pb.BuildPoints)
	}
}

func TestRemoveMetatype(t *testing.T) {
	data := []struct {
		metatype *metatype.Metatype
	}{
		{&metatype.Metatype{Name: "Human", PointCost: 0}},
		{&metatype.Metatype{Name: "Elf", PointCost: 40}},
		{&metatype.Metatype{Name: "Dwarf", PointCost: 50}},
		{&metatype.Metatype{Name: "Ork", PointCost: 50}},
		{&metatype.Metatype{Name: "Troll", PointCost: 90}},
	}

	for _, d := range data {
		pb := character.NewPointBuilder()
		pb.SetMetatype(d.metatype)

		pb.RemoveMetatype()
		assert.Empty(t, pb.Metatype)
		assert.Equal(t, character.TotalBuildPoints, pb.BuildPoints)
	}
}

var metatypeHuman = &metatype.Metatype{
	Name: "Human",
	Attributes: metatype.Attributes{
		Body:      metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
		Magic:     metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
		Resonance: metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
	}}

func TestSetMagicType(t *testing.T) {
	data := []struct {
		magicType              character.MagicType
		expectedBuildPoints    int
		expectedMagicValue     int
		expectedResonanceValue int
	}{
		{character.MagicTypeNone, 800, 0, 0},             // 0
		{character.MagicTypeAdept, 780, 1, 0},            // 20
		{character.MagicTypeMagician, 785, 1, 0},         // 15
		{character.MagicTypeAspectedMagician, 770, 1, 0}, // 30
		{character.MagicTypeMysticAdept, 765, 1, 0},      // 35
		{character.MagicTypeTechnomancer, 785, 0, 1},     // 15
	}

	for _, d := range data {
		pb := character.NewPointBuilder()
		pb.SetMetatype(metatypeHuman)
		err := pb.SetMagicType(d.magicType)

		assert.NoError(t, err)
		assert.Equal(t, d.magicType, pb.MagicType)
		assert.Equal(t, d.expectedBuildPoints, pb.BuildPoints)
		assert.Equal(t, d.expectedMagicValue, pb.Attributes[shared.AttributeMagic])
		assert.Equal(t, d.expectedResonanceValue, pb.Attributes[shared.AttributeResonance])
	}
}

func TestRemoveMagicType(t *testing.T) {
	pb := character.NewPointBuilder()
	errMetatype := pb.SetMetatype(metatypeHuman)
	errMagic := pb.SetMagicType(character.MagicTypeAdept)
	pb.RemoveMagicType()

	assert.NoError(t, errMetatype)
	assert.NoError(t, errMagic)
	assert.Equal(t, character.MagicTypeNone, pb.MagicType)
	assert.Equal(t, 800, pb.BuildPoints)
	assert.Equal(t, 0, pb.Attributes[shared.AttributeMagic])
	assert.Equal(t, 0, pb.Attributes[shared.AttributeResonance])
}

// TODO: Test for missing metatype
// TODO: Test for missing magic type
// TODO: Test for exceeding the max attribute value
// TODO: Test for having more than one maxed attribute
func TestAdjustAttribute(t *testing.T) {
	data := []struct {
		attribute           shared.AttributeType
		magicType           character.MagicType
		newValue            int
		expectedNewValue    int
		expectedBuildPoints int
		expectedError       error
	}{
		{shared.AttributeBody, character.MagicTypeNone, 2, 2, 790, nil},
		{shared.AttributeMagic, character.MagicTypeAdept, 2, 2, 770, nil},
		{shared.AttributeResonance, character.MagicTypeTechnomancer, 2, 2, 775, nil},
		{shared.AttributeMagic, character.MagicTypeNone, 0, 1, 800, fmt.Errorf("can not adjust magic without being a magic user")},
		{shared.AttributeResonance, character.MagicTypeNone, 0, 1, 800, fmt.Errorf("can not adjust resonance without being a technomancer")},
		{shared.AttributeBody, character.MagicTypeNone, 0, 1, 800, fmt.Errorf("'Body' (0) can not be lowered below metatype minimum (1)")},
		{shared.AttributeBody, character.MagicTypeNone, 14, 1, 800, fmt.Errorf("'Body' (14) can not be raised above metatype maximum (6)")},
	}

	for _, d := range data {
		pb := character.NewPointBuilder()
		pb.SetMetatype(metatypeHuman)
		pb.SetMagicType(d.magicType)

		err := pb.AdjustAttribute(d.attribute, d.newValue)
		if d.expectedError != nil {
			assert.Error(t, err)
			assert.Equal(t, d.expectedError, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, d.expectedNewValue, pb.Attributes[d.attribute])
			assert.Equal(t, d.expectedBuildPoints, pb.BuildPoints)
		}
	}
}

func TestAdjustSkill(t *testing.T) {
	data := []struct {
		changeType          string
		skill               string
		newValue            int
		expectedNewValue    int
		expectedBuildPoints int
		expectedError       error
	}{
		// TODO: Test for not enough build points
		// TODO: Test for metatype not set
		// TODO: Test for magictype not set
		// TODO: Test for skills not available to the magic type
		{character.ChangeActiveSkill, "Running", 0, 0, 800, nil},                                                      // Active skill mo change
		{character.ChangeActiveSkill, "Running", 1, 1, 798, nil},                                                      // Active skill increate from 0 to 1
		{character.ChangeActiveSkill, "Sneaking", 2, 2, 796, nil},                                                     // Active skill increase from 1 to 2
		{character.ChangeActiveSkill, "Sneaking", 1, 1, 800, nil},                                                     // Active skill decrease from 2 to 1
		{character.ChangeKnowledgeSkill, "English", 1, 1, 799, nil},                                                   // Knowledge skill 0 to 1
		{character.ChangeSkillGroup, "Automatics", 1, 1, 795, nil},                                                    // Skill group 0 to 1
		{character.ChangeSkillGroup, "Running", -1, 0, 800, fmt.Errorf("skill value can not be negative")},            // Skill group decrease below 0
		{character.ChangeKnowledgeSkill, "English", 14, 0, 800, fmt.Errorf("skill value can not be greater than 13")}, // Knowledge skill increase above 13
		{character.ChangeSkillGroup, "Automatics", 13, 0, 800, nil},
		// {"Hacking", 11, 690, fmt.Errorf("not enough build points")},
	}

	for _, d := range data {
		pb := character.NewPointBuilder()
		pb.SetMetatype(&metatype.Metatype{Name: "Human"})
		pb.SetMagicType(character.MagicTypeNone)
		pb.Skills = map[string]int{
			"Sneaking": 1,
		}
		err := pb.AdjustSkill(d.changeType, d.skill, d.newValue)
		if d.expectedError != nil {
			assert.Error(t, err)
			assert.Equal(t, d.expectedError, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, d.expectedNewValue, pb.Skills[d.skill])
			assert.Equal(t, d.expectedBuildPoints, pb.BuildPoints)
		}
	}
}

// func TestCalculateAttributeChangeCost(t *testing.T) {
// 	data := []struct {
// 		currentValue int
// 		desiredValue int
// 		expectedCost int
// 	}{
// 		{1, 1, 0},     // No change
// 		{1, 2, 10},    // 1 -> 2
// 		{2, 1, -10},   // 2 -> 1
// 		{12, 13, 65},  // 12 -> 13
// 		{13, 12, -65}, // 13 -> 12
// 	}

// 	for _, d := range data {
// 		cost := character.CalculateAttributeChangeCost(d.currentValue, d.desiredValue)
// 		assert.Equal(t, d.expectedCost, cost)
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
