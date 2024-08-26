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

func TestIncreaseAttribute(t *testing.T) {
	data := []struct {
		attribute               shared.AttributeType
		magicType               character.MagicType
		expectedPointsRemaining int
		expectedError           error
	}{
		{shared.AttributeBody, character.MagicTypeNone, 790, nil},
		// {shared.AttributeMagic, character.MagicTypeNone, 800, fmt.Errorf("can not adjust magic without being a magic user")},
		// {shared.AttributeResonance, fmt.Errorf("can not adjust resonance without being a technomancer")},
	}

	// pb := character.NewPointBuilder()
	// pb.MagicType = character.MagicTypeMagician
	// pb.Metatype = &metatype.Metatype{
	// 	Name: "Human",
	// 	Attributes: metatype.Attributes{
	// 		Body:      metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
	// 		Magic:     metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
	// 		Resonance: metatype.Attribute[int]{Min: 0, Max: 6, AugMax: 10},
	// 	}}
	// pb.MagicType = character.MagicTypeMagician
	// pb.Attributes[shared.AttributeBody] = 1
	// pb.Attributes[shared.AttributeMagic] = 1
	// pb.Attributes[shared.AttributeResonance] = 1

	for _, d := range data {
		pb := character.NewPointBuilder()
		pb.SetMagicType(d.magicType)
		pb.Metatype = &metatype.Metatype{
			Name: "Human",
			Attributes: metatype.Attributes{
				Body:      metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
				Magic:     metatype.Attribute[int]{Min: 1, Max: 6, AugMax: 10},
				Resonance: metatype.Attribute[int]{Min: 0, Max: 6, AugMax: 10},
			}}

		err := pb.IncreaseAttribute(d.attribute)

		if d.expectedError != nil {
			assert.Error(t, err)
			assert.Equal(t, d.expectedError, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, pb.Attributes[d.attribute], 2)
			assert.Equal(t, pb.BuildPoints, d.expectedPointsRemaining)
		}
	}
}
