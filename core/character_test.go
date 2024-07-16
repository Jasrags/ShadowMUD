package core_test

import (
	"testing"

	"shadowrunmud/core"
)

func NewStreetSamurai() *core.Character {
	// m, _ := metatype.LoadMetatype("ork")

	c := core.NewCharacter()
	// c.SetName("Street Samurai")
	// c.SetMetatype(m)
	// c.SetBody(7)
	// c.SetAgility(6)
	// c.SetReaction(5)
	// c.SetStrength(5)
	// c.SetWillpower(3)
	// c.SetLogic(2)
	// c.SetIntuition(3)
	// c.SetCharisma(2)
	// c.SetEssence(6)
	// c.SetEdge(1)

	return c
}

/*
METATYPE: ORK
B 7
A 6
R 5 (7)
S 5
W 3
L 2
I 3
C 2
ESS 0.88
EDG 1

Condition Monitor (P/S) 14 / 10
Armor 13
Limits: Physical 8 (9), Mental 4, Social 3
Physical Init 10 + 3D6
Active Skills:
    Automatics 5,
    Blades 5 (6),
    Longarms 3 (4),
    Pilot Ground 1,
    Pistols 4,
    Sneaking 2 (3),
    Unarmed Combat 2 (3)
Knowledge Skills:
    Great Restaurants 2,
    Law Enforcement 2,
    Poetry 1,
    Safe Houses 3
Languages:
    English N,
    Japanese 2
Qualities:
    Ambidextrous,
    Code of Honor (Bushido),
    Guts,
    Home Ground (Street Politics),
    Incompetent (Acting)
Augmentations:
    Cybereyes [Rating 3, flare compensation, low-light, smartlink, thermographic, vision enhancement 2, vision magnification],
    dermal plating 2 (alphaware),
    cyberarm [right, obvious, Strength 11, Agility 9, cyber sub-machine gun w/ external clip port],
    cyberarm [left, obvious, Strength 11, Agility 9, Armor 2, cyber spur, cyberarm slide],
    enhanced articulation,
    platelet factories,
    reflex recorder [Blades, Longarms, Sneaking, Unarmed Combat],
    synaptic booster 2,
    synthcardium 1
Vehicles
    Harley-Davidson Scorpion [Handling 4/3, Speed 4, Accel 2, Body 8, Armor 9, Pilot 1, Sensor 2]
Gear
    earbuds [audio enhancement (Rating 1),
    4 fake SINs [Rating 4, each with fake licenses (Rating 4, concealed carry, possession of firearms, possession of augmentations)], spatial recognizer],
    Hermes Ikon commlink (Rating 5),
    jammer (area, Rating 4),
    lined coat [9, chemical protection 3,
    fire resistance 3, non-conductivity 3],
    medkit (Rating 3),
    medkit (Rating 6),
    micro-transceiver,
    Middle Lifestyle (3 months),
    5 stim patches (Rating 6),
    2 trauma patches,
    white noise generator (Rating 6)
Weapons:
    Katana [Blade, Reach 1, Acc 7, DV 14P, AP –3]
    Sword [Blade, Reach 1, Acc 6, DV 14P, AP –2]
    Ares Light Fire 75 [Light Pistol, Acc 6 (8), DV 6P, AP —, SA, RC —,16(c),w/ 3 spare clips, 100 rounds regular ammo]
    Ares Predator V [Heavy Pistol, SA, Acc 5 (7), DV 8P, AP –5, RC —, Ammo 15 (c), APDS ammo (100 rounds) , 3 spare clips]
    HK-227 [SMG, Acc 5 (7), DV 8P, AP –1, SA/BF/FA, RC (1), 28 (c), w/ 3 spare clips, 100 rounds explosive ammo]
    FN HAR [Assault Rifle, Acc 5 (7), DV 10P, AP –6, SA/BF/FA, RC 2, 35 (c), w/intergral smartlink, 3 spare clips, 100 rounds APDS ammo]
    Enfield AS-7 [Shotgun, Acc 4 (6), DV 15P(f), AP +4, SA/BF, RC —,10(c) or 24 (d), w/ internal smartlink, 3 spare clips, 100 rounds flechette ammo]
    Ingram Valiant [LMG, Acc 5 (7), DV 10P, AP –3, BF/FA, RC 2 (3), 50 (c) or 100 (belt), w/ integral smartlink, 3 spare clips, 100 rounds explosive ammo]
    3 high explosive grenades [Grenade, non-aerodynamic, DV 16P, AP –2, Blast –2/m]
    Spurs (cyber) [Unarmed, Reach —, Acc 9, DV 14P, AP –2]
Contacts
    Fixer (Connection 4, Loyalty 2)
Starting ¥2,555 + (4D6 x 100)¥

*/

// var (
// streetSamurai = character.NewCharacter()
// )

// var (
// 	streetSamurai = character.Character{
// 		Name:      "Street Samurai",
// 		Metatype:  metatype.Metatypes["Human"],
// 		Body:      7,
// 		Agility:   6,
// 		Reaction:  5, // (7)
// 		Strength:  5,
// 		Willpower: 3,
// 		Logic:     2,
// 		Intuition: 3,
// 		Charisma:  2,
// 		Essence:   0.88,
// 		Edge:      1,
// 		// ActiveSkills: map[string]skill.ActiveSkill{},
// 		// ActiveSkills: map[string]skills.ActiveSkill{skill.ActiveSkillAutomatics: {Rating: 5},}
// 		// },
// 	}
// )

// func TestRollInitiative(t *testing.T) {
// 	c := character.NewCharacter()
// 	got := c.RollInitiative()
// 	want := 0

// 	if got != want {
// 		t.Errorf("RollInitiative() = %d, want %d", got, want)
// 	}
// }

// func TestGetInitiative(t *testing.T) {
// 	dt := []struct {
// 		Reaction  int
// 		Intuition int
// 		want      int
// 	}{
// 		{5, 7, 12},
// 		{5, 5, 10},
// 	}

// 	for _, tt := range dt {
// 		c := &character.Character{
// 			Reaction:  tt.Reaction,
// 			Intuition: tt.Intuition,
// 		}

// 		got := c.GetInitiative()

// 		if got != tt.want {
// 			t.Errorf("GetInitiative() = %d, want %d", got, tt.want)
// 		}
// 	}
// }

func TestGetPhysicalLimit(t *testing.T) {
	dt := []struct {
		Strength int
		Body     int
		Reaction int
		want     int
	}{
		{5, 7, 6, 8},
		{5, 7, 5, 8},
	}

	for _, tt := range dt {
		c := core.Character{
			Attributes: core.Attributes{
				Strength: core.AttributesInfo{Base: tt.Strength},
				Body:     core.AttributesInfo{Base: tt.Body},
				Reaction: core.AttributesInfo{Base: tt.Reaction},
			},
			// Strength: tt.Strength,
			// Body:     tt.Body,
			// Reaction: tt.Reaction,
		}

		got := c.GetPhysicalLimit()

		if got != tt.want {
			t.Errorf("GetPhysicalLimit() = %d, want %d", got, tt.want)
		}
	}
}

func TestGetMentalLimit(t *testing.T) {
	dt := []struct {
		Logic     int
		Intuition int
		Willpower int
		want      int
	}{
		{5, 7, 6, 8},
		{5, 7, 5, 8},
	}

	for _, tt := range dt {
		c := core.Character{
			Attributes: core.Attributes{
				Logic:     core.AttributesInfo{Base: tt.Logic},
				Intuition: core.AttributesInfo{Base: tt.Intuition},
				Willpower: core.AttributesInfo{Base: tt.Willpower},
			},
		}
		c.RecalculateAttributes()

		got := c.GetMentalLimit()

		if got != tt.want {
			t.Errorf("GetMentalLimit() = %d, want %d", got, tt.want)
		}
	}
}

func TestGetSocialLimit(t *testing.T) {
	dt := []struct {
		Charisma  int
		Willpower int
		Essence   float64
		want      int
	}{
		{5, 7, 6, 8},
		{5, 7, 3.4, 7},
	}

	for _, tt := range dt {
		c := core.Character{
			Attributes: core.Attributes{
				Charisma:  core.AttributesInfo{Base: tt.Charisma},
				Willpower: core.AttributesInfo{Base: tt.Willpower},
				Essence:   core.AttributesInfoF{Base: tt.Essence},
			},
		}
		c.RecalculateAttributes()

		got := c.GetSocialLimit()

		if got != tt.want {
			t.Errorf("GetSocialLimit() = %d, want %d", got, tt.want)
		}
	}
}
func TestRecalculate(t *testing.T) {
	tsk := core.TestSkill{
		Name:      "Automatics",
		Attribute: "Agility",
		Rank:      4,
	}

	tc := core.TestChar{
		Body:      core.TestAttribute{Base: 5},
		Agility:   core.TestAttribute{Base: 7},
		Reaction:  core.TestAttribute{Base: 6},
		Strength:  core.TestAttribute{Base: 8},
		Willpower: core.TestAttribute{Base: 5},
		Logic:     core.TestAttribute{Base: 7},
		Intuition: core.TestAttribute{Base: 6},
		Charisma:  core.TestAttribute{Base: 8},
		Essence:   6.0,
		Skills:    []core.TestSkill{tsk},
	}
	tc.Recalculate()

	// Check the recalculated values
	if tc.Body.Value != 5 {
		t.Errorf("Recalculate() failed for Body. Expected 5, got %d", tc.Body.Value)
	}
	if tc.Body.Base != 5 {
		t.Errorf("Recalculate() failed for Body. Expected base 5, got %d", tc.Body.Base)
	}
	if tc.Body.Mods != 0 {
		t.Errorf("Recalculate() failed for Body. Expected mod 0, got %d", tc.Body.Mods)
	}

	if tc.Agility.Value != 11 {
		t.Errorf("Recalculate() failed for Agility. Expected 11, got %d", tc.Agility.Value)
	}
	if tc.Agility.Base != 7 {
		t.Errorf("Recalculate() failed for Agility. Expected base 7, got %d", tc.Agility.Base)
	}
	if tc.Agility.Mods != 4 {
		t.Errorf("Recalculate() failed for Agility. Expected mod 4, got %d", tc.Agility.Mods)
	}

	if tc.Reaction.Value != 6 {
		t.Errorf("Recalculate() failed for Reaction. Expected 6, got %d", tc.Reaction.Value)
	}
	if tc.Reaction.Base != 6 {
		t.Errorf("Recalculate() failed for Reaction. Expected base 6, got %d", tc.Reaction.Base)
	}
	if tc.Reaction.Mods != 0 {
		t.Errorf("Recalculate() failed for Reaction. Expected mod 0, got %d", tc.Reaction.Mods)
	}

	if tc.Strength.Value != 8 {
		t.Errorf("Recalculate() failed for Strength. Expected 8, got %d", tc.Strength.Value)
	}
	if tc.Strength.Base != 8 {
		t.Errorf("Recalculate() failed for Strength. Expected base 8, got %d", tc.Strength.Base)
	}
	if tc.Strength.Mods != 0 {
		t.Errorf("Recalculate() failed for Strength. Expected mod 0, got %d", tc.Strength.Mods)
	}

	if tc.Willpower.Value != 5 {
		t.Errorf("Recalculate() failed for Willpower. Expected 5, got %d", tc.Willpower.Value)
	}
	if tc.Willpower.Base != 5 {
		t.Errorf("Recalculate() failed for Willpower. Expected base 5, got %d", tc.Willpower.Base)
	}
	if tc.Willpower.Mods != 0 {
		t.Errorf("Recalculate() failed for Willpower. Expected mod 0, got %d", tc.Willpower.Mods)
	}

	if tc.Logic.Value != 7 {
		t.Errorf("Recalculate() failed for Logic. Expected 7, got %d", tc.Logic.Value)
	}
	if tc.Logic.Base != 7 {
		t.Errorf("Recalculate() failed for Logic. Expected base 7, got %d", tc.Logic.Base)
	}
	if tc.Logic.Mods != 0 {
		t.Errorf("Recalculate() failed for Logic. Expected mod 0, got %d", tc.Logic.Mods)
	}

	if tc.Intuition.Value != 6 {
		t.Errorf("Recalculate() failed for Intuition. Expected 6, got %d", tc.Intuition.Value)
	}
	if tc.Intuition.Base != 6 {
		t.Errorf("Recalculate() failed for Intuition. Expected base 6, got %d", tc.Intuition.Base)
	}
	if tc.Intuition.Mods != 0 {
		t.Errorf("Recalculate() failed for Intuition. Expected mod 0, got %d", tc.Intuition.Mods)
	}

	if tc.Charisma.Value != 8 {
		t.Errorf("Recalculate() failed for Charisma. Expected 8, got %d", tc.Charisma.Value)
	}
	if tc.Charisma.Base != 8 {
		t.Errorf("Recalculate() failed for Charisma. Expected base 8, got %d", tc.Charisma.Base)
	}
	if tc.Charisma.Mods != 0 {
		t.Errorf("Recalculate() failed for Charisma. Expected mod 0, got %d", tc.Charisma.Mods)
	}
}
