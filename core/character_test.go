package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Jasrags/ShadowMUD/core"
)

func TestCharacterStreetSamurai(t *testing.T) {

	//  Armor: 13
	//  TODO: Physical: 14, Stun: 10},
	//  Limits: Physical: 8 (9), Mental: 4, Social: 3
	//  Physical Init: 10 + 3D6
	//  Active Skills:
	//      Automatics 5
	//      Blades 5 (6)
	//      Longarms 3 (4)
	//      Pilot Ground 1
	//      Pistols 4
	//      Sneaking 2 (3)
	//      Unarmed Combat 2 (3)
	//  Knowledge Skills:
	//      Great Restaurants 2
	//      Law Enforcement 2
	//      Poetry 1
	//      Safe Houses 3
	//  Languages:
	//      English 10
	//      Japanese 2
	//  Qualities:
	//      Ambidextrous
	//      Code of Honor (Bushido)
	//      Guts
	//      Home Ground (Street Politics)
	//      Incompetent (Acting)
	//  Augmentations:
	//      Cybereyes [Rating 3, flare compensation, low-light, smartlink, thermographic, vision enhancement 2, vision magnification],
	//      dermal plating 2 (alphaware),
	//      cyberarm [right, obvious, Strength 11, Agility 9, cyber sub-machine gun w/ external clip port],
	//      cyberarm [left, obvious, Strength 11, Agility 9, Armor 2, cyber spur, cyberarm slide],
	//      enhanced articulation,
	//      platelet factories,
	//      reflex recorder [Blades, Longarms, Sneaking, Unarmed Combat],
	//      synaptic booster 2,
	//      synthcardium 1
	//  Vehicles:
	//      Harley-Davidson Scorpion [Handling 4/3, Speed 4, Accel 2, Body 8, Armor 9, Pilot 1, Sensor 2]
	//  Gear:
	//      earbuds [audio enhancement (Rating 1)],
	//      4 fake SINs [Rating 4, each with fake licenses (Rating 4, concealed carry, possession of firearms, possession of augmentations)],
	//      spatial recognizer,
	//      Hermes Ikon commlink (Rating 5),
	//      jammer (area, Rating 4),
	//      lined coat [9, chemical protection 3, fire resistance 3, non-conductivity 3],
	//      medkit (Rating 3),
	//      medkit (Rating 6),
	//      micro-transceiver,
	//      Middle Lifestyle (3 months),
	//      5 stim patches (Rating 6),
	//      2 trauma patches,
	//      white noise generator (Rating 6)
	//  Weapons:
	//      Katana [Blade, Reach 1, Acc 7, DV 14P, AP –3]
	//      Sword [Blade, Reach 1, Acc 6, DV 14P, AP –2]
	//      Ares Light Fire 75 [Light Pistol, Acc 6 (8), DV 6P, AP —, SA, RC —,16 (c),w/ 3 spare clips, 100 rounds regular ammo]
	//      Ares Predator V [Heavy Pistol, SA, Acc 5 (7), DV 8P, AP –5, RC —, Ammo 15 (c), APDS ammo (100 rounds) , 3 spare clips]
	//      HK-227 [SMG, Acc 5 (7), DV 8P, AP –1, SA/BF/FA, RC (1), 28 (c), w/ 3 spare clips, 100 rounds explosive ammo]
	//      FN HAR [Assault Rifle, Acc 5 (7), DV 10P, AP –6, SA/BF/FA, RC 2, 35 (c), w/ intergral smartlink, 3 spare clips, 100 rounds APDS ammo]
	//      Enfield AS-7 [Shotgun, Acc 4 (6), DV 15P(f), AP +4, SA/BF, RC —,10(c) or 24 (d), w/ internal smartlink, 3 spare clips, 100 rounds flechette ammo]
	//      Ingram Valiant [LMG, Acc 5 (7), DV 10P, AP –3, BF/FA, RC 2 (3), 50 (c) or 100 (belt), w/ integral smartlink, 3 spare clips, 100 rounds explosive ammo]
	//      3 high explosive grenades [Grenade, non-aerodynamic, DV 16P, AP –2, Blast –2/m]
	//      Spurs (cyber) [Unarmed, Reach —, Acc 9, DV 14P, AP –2]
	//  Contacts:
	//      Fixer (Connection 4, Loyalty 2)
	//  Starting ¥2,555 + (4D6 x 100)¥

	c := core.Character{
		ID:              "street-samurai",
		Name:            "Street Samurai",
		MetatypeID:      "ork",
		Age:             30,
		Ethnicity:       "Japanese",
		Sex:             "Male",
		Height:          180,
		Weight:          90,
		StreetCred:      1,
		Notoriety:       1,
		PublicAwareness: 1,
		Karma:           1,
		TotalKarma:      1,
		ConditionDamage: core.ConditionDamage{
			Physical: 1,
			Stun:     1,
		},
		Attributes: core.Attributes{
			Body:      core.AttributesInfo{Base: 7},
			Agility:   core.AttributesInfo{Base: 6},
			Reaction:  core.AttributesInfo{Base: 5}, // (7)
			Strength:  core.AttributesInfo{Base: 5},
			Willpower: core.AttributesInfo{Base: 3},
			Logic:     core.AttributesInfo{Base: 2},
			Intuition: core.AttributesInfo{Base: 3},
			Charisma:  core.AttributesInfo{Base: 2},
			Essence:   core.AttributesInfoF{Base: 6}, // 0.88
		},
		InitiativeDice: core.InitiativeDice{
			Physical:        core.AttributesInfo{Base: 1}, // 10+3d6
			Astral:          core.AttributesInfo{Base: 2},
			MatrixAR:        core.AttributesInfo{Base: 1},
			MatrixVRHotSim:  core.AttributesInfo{Base: 4},
			MatrixVRColdSim: core.AttributesInfo{Base: 3},
			RiggerAR:        core.AttributesInfo{Base: 1},
		},
		Edge:       1,
		EdgePoints: 1,
		Magic:      1,
		Resonance:  1,
		ActiveSkills: map[string]core.ActiveSkill{
			"Automatics":         {Name: "Automatics", Rating: 5},
			"Blades":             {Name: "Blades", Rating: 5},   // (6)
			"Longarms":           {Name: "Longarms", Rating: 3}, // (4)
			"Pilot Ground Craft": {Name: "Pilot Ground Craft", Rating: 1},
			"Pistols":            {Name: "Pistols", Rating: 4},
			"Sneaking":           {Name: "Sneaking", Rating: 2},      // (3)
			"Unarmed Combat":     {Name: "Unarmed Comba", Rating: 2}, // (3)
		},
		LanguageSkills: map[string]core.LanguageSkill{
			"English":  {Name: "English", Rating: 10},
			"Japanese": {Name: "Japanese", Rating: 2},
		},
		KnowledgeSkills: map[string]core.KnowledgeSkill{
			"Great Restaurants": {Name: "Great Restaurants", Rating: 2},
			"Law Enforcement":   {Name: "Law Enforcement", Rating: 2},
			"Poetry":            {Name: "Poetry", Rating: 1},
			"Safe Houses":       {Name: "Safe Houses", Rating: 3},
		},
		Qualities: map[string]core.Quality{
			"Ambidextrous":                  {Name: "Ambidextrous", Type: core.QualityTypePositive},
			"Code of Honor (Bushido)":       {Name: "Code of Honor (Bushido)", Type: core.QualityTypePositive},
			"Guts":                          {Name: "Guts", Type: core.QualityTypePositive},
			"Home Ground (Street Politics)": {Name: "Home Ground (Street Politics)", Type: core.QualityTypePositive},
			"Incompetent (Acting)":          {Name: "Incompetent (Acting)", Type: core.QualityTypeNegative},
		},
		Contacts: map[string]core.Contact{
			"Fixer": {Type: core.ContactTypeFixer, Connection: 4, Loyalty: 2},
		},
		// Identities: map[string]core.Identity{},
		// Lifestyles: map[string]core.Lifestyle{},
		// Currancy: map[string]core.Currency{},
		// RangedWeapons: map[string]core.RangedWeapon{},
		// MeleeWeapons: map[string]core.MeleeWeapon{},
		// Armor: map[string]string{},
		// Cyberware: map[string]core.Cyberware{},
		// Cyberdecks: map[string]core.Cyberdeck{},
		// Vehicles: map[string]core.Vehicle{},
		// Programs: map[string]core.Program{},
		// Gear: map[string]core.Gear{},
		// AdeptPowers: map[string]core.AdeptPower{},
	}

	c.Recalculate()

	if ok := assert.Equal(t, "street-samurai", c.ID); !ok {
		t.Errorf("ID Expected %s, got %s", "street-samurai", c.ID)
	}
	if ok := assert.Equal(t, "Street Samurai", c.Name); !ok {
		t.Errorf("Name Expected %s, got %s", "Street Samurai", c.Name)
	}
	// TODO: What are we doing with references?
	// if ok := assert.Equal(t, "ork", c.MetatypeID); !ok {
	//     t.Errorf("Metatype Expected %s, got %s", "ork", c.MetatypeID)
	// }
	if ok := assert.Equal(t, "Japanese", c.Ethnicity); !ok {
		t.Errorf("Sex Expected %s, got %s", "Japanese", c.Ethnicity)
	}
	if ok := assert.Equal(t, "Male", c.Sex); !ok {
		t.Errorf("Sex Expected %s, got %s", "Male", c.Sex)
	}
	if ok := assert.Equal(t, 30, c.Age); !ok {
		t.Errorf("Age Expected %d, got %d", 30, c.Age)
	}
	if ok := assert.Equal(t, 180, c.Height); !ok {
		t.Errorf("Height Expected %d, got %d", 180, c.Height)
	}
	if ok := assert.Equal(t, 90, c.Weight); !ok {
		t.Errorf("Weight Expected %d, got %d", 90, c.Weight)
	}
	if ok := assert.Equal(t, 1, c.StreetCred); !ok {
		t.Errorf("StreetCred Expected %d, got %d", 1, c.StreetCred)
	}
	if ok := assert.Equal(t, 1, c.Notoriety); !ok {
		t.Errorf("Notoriety Expected %d, got %d", 1, c.Notoriety)
	}
	if ok := assert.Equal(t, 1, c.PublicAwareness); !ok {
		t.Errorf("PublicAwareness Expected %d, got %d", 1, c.PublicAwareness)
	}
	if ok := assert.Equal(t, 1, c.Karma); !ok {
		t.Errorf("Karma Expected %d, got %d", 1, c.Karma)
	}
	if ok := assert.Equal(t, 1, c.TotalKarma); !ok {
		t.Errorf("TotalKarma Expected %d, got %d", 1, c.TotalKarma)
	}
	// TODO: Uncomment when recalculated
	// if ok := assert.Equal(t, 14, c.GetConditionPhysical()); !ok {
	// 	t.Errorf("ConditionDamage.Physical Expected %d, got %d", 14, c.ConditionDamage.Physical)
	// }
	// TODO: Uncomment when recalculated
	// if ok := assert.Equal(t, 10, c.GetConditionStun()); !ok {
	// 	t.Errorf("ConditionDamage.Stun Expected %d, got %d", 10, c.ConditionDamage.Stun)
	// }

	td1 := []struct {
		name      string
		attribute core.AttributesInfo
		base      int
		mods      int
		value     int
	}{
		// Attributes
		{"Body", c.Attributes.Body, 7, 0, 7},
		{"Agility", c.Attributes.Agility, 6, 0, 6},
		// TODO: Uncomment when recalculated
		// {"Reaction", c.Attributes.Reaction, 5, 2, 7},
		{"Strength", c.Attributes.Strength, 5, 0, 5},
		{"Willpower", c.Attributes.Willpower, 3, 0, 3},
		{"Logic", c.Attributes.Logic, 2, 0, 2},
		{"Intuition", c.Attributes.Intuition, 3, 0, 3},
		{"Charisma", c.Attributes.Charisma, 2, 0, 2},
		// Initiative dice
		{"Physical", c.InitiativeDice.Physical, 1, 0, 1},
		{"Astral", c.InitiativeDice.Astral, 2, 0, 2},
		{"MatrixAR", c.InitiativeDice.MatrixAR, 1, 0, 1},
		{"MatrixVRHotSim", c.InitiativeDice.MatrixVRHotSim, 4, 0, 4},
		{"MatrixVRColdSim", c.InitiativeDice.MatrixVRColdSim, 3, 0, 3},
		{"RiggerAR", c.InitiativeDice.RiggerAR, 1, 0, 1},
	}

	for _, tt := range td1 {
		t.Run(tt.name, func(t *testing.T) {
			if ok := assert.Equal(t, tt.base, tt.attribute.Base); !ok {
				t.Errorf("%s.Base Expected %d, got %d", tt.name, tt.base, tt.attribute.Base)
			}
			if ok := assert.Equal(t, tt.mods, tt.attribute.Mods); !ok {
				t.Errorf("%s.Mods Expected %d, got %d", tt.name, tt.mods, tt.attribute.Mods)
			}
			if ok := assert.Equal(t, tt.value, tt.attribute.Value); !ok {
				t.Errorf("%s.Value Expected %d, got %d", tt.name, tt.value, tt.attribute.Value)
			}
		})
	}

	td2 := []struct {
		name      string
		attribute core.AttributesInfoF
		base      float64
		mods      float64
		value     float64
	}{
		// TODO: Uncomment when recalculated
		// {"Essence", c.Attributes.Essence, 6, 5.12, 0.88},
	}

	for _, tt := range td2 {
		t.Run(tt.name, func(t *testing.T) {
			if ok := assert.Equal(t, tt.base, tt.attribute.Base); !ok {
				t.Errorf("%s.Base Expected %f, got %f", tt.name, tt.base, tt.attribute.Base)
			}
			if ok := assert.Equal(t, tt.mods, tt.attribute.Mods); !ok {
				t.Errorf("%s.Mods Expected %f, got %f", tt.name, tt.mods, tt.attribute.Mods)
			}
			if ok := assert.Equal(t, tt.value, tt.attribute.Value); !ok {
				t.Errorf("%s.Value Expected %f, got %f", tt.name, tt.value, tt.attribute.Value)
			}
		})
	}

	// filename := "/Users/jrags/Code/Jasrags/ShadowMUD/test/data/characters/street_samurai.yaml"

	// util.SaveStructToYAML(filename, c)
}

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
		}
		c.RecalculateAttributes()

		got := c.GetPhysicalLimit()

		if ok := assert.Equal(t, tt.want, got); !ok {
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

		if ok := assert.Equal(t, tt.want, got); !ok {
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

func TestRecalculateCyberware(t *testing.T) {
	c := core.Character{
		Attributes: core.Attributes{
			Reaction: core.AttributesInfo{Base: 5},
		},
		Cyberware: map[string]core.Cyberware{
			"wired-reflexes-r1": {
				EssenceCost: 2.0,
				Rating:      1,
				Modifiers: []core.CyberwareModifier{
					{Type: "Reaction", Effect: "Increase", Value: 1},
				},
			},
		},
	}

	c.Attributes.Reaction.Recalculate()
	got1 := c.Attributes.Reaction.Base
	want1 := 5

	if !assert.Equal(t, want1, got1) {
		t.Errorf("RecalculateCyberware() = %d, want %d", got1, want1)
	}

	c.RecalculateCyberware()
	c.Attributes.Reaction.Recalculate()

	got2 := c.Attributes.Reaction.Value
	want2 := 6

	if !assert.Equal(t, want2, got2) {
		t.Errorf("RecalculateCyberware() = %d, want %d", got2, want2)
	}
}
