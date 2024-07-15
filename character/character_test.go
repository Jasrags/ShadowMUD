package character_test

import (
	"testing"

	"shadowrunmud/character"
)

func NewStreetSamurai() *character.Character {
	// m, _ := metatype.LoadMetatype("ork")

	c := character.NewCharacter()
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
		c := character.Character{
			Attributes: character.Attributes{
				Strength: character.AttributesInfo{Base: tt.Strength},
				Body:     character.AttributesInfo{Base: tt.Body},
				Reaction: character.AttributesInfo{Base: tt.Reaction},
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
		c := character.Character{
			Attributes: character.Attributes{
				Logic:     character.AttributesInfo{Base: tt.Logic},
				Intuition: character.AttributesInfo{Base: tt.Intuition},
				Willpower: character.AttributesInfo{Base: tt.Willpower},
			},
			// Logic:     tt.Logic,
			// Intuition: tt.Intuition,
			// Willpower: tt.Willpower,
		}

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
		c := character.Character{
			Attributes: character.Attributes{
				Charisma:  character.AttributesInfo{Base: tt.Charisma},
				Willpower: character.AttributesInfo{Base: tt.Willpower},
				Essence:   character.AttributesInfoF{Base: tt.Essence},
			},
			// Charisma:  tt.Charisma,
			// Willpower: tt.Willpower,
			// Essence: tt.Essence,
		}

		got := c.GetSocialLimit()

		if got != tt.want {
			t.Errorf("GetSocialLimit() = %d, want %d", got, tt.want)
		}
	}
}
