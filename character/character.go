package character

import (
	"math"
	"shadowrunmud/character/metatype"
	"time"

	"golang.org/x/exp/rand"
)

type Character struct {
	// Personal Data
	Name            string
	Metatype        metatype.Metatype
	Ethnicity       string
	Age             int
	Sex             string
	Height          int
	Weight          int
	StreetCred      int
	Notoriety       int
	PublicAwareness int
	Karma           int
	TotalKarma      int
	// Attributes
	Body       int
	Agility    int
	Reaction   int
	Strength   int
	Willpower  int
	Logic      int
	Intuition  int
	Charisma   int
	Edge       int
	EdgePoints int
	// Derived Attributes
	Essence       float64
	Magic         int
	Resonance     int
	PhysicalLimit int
	MentalLimit   int
	SocialLimit   int
	// Initiative       int
	MatrixInitiative int
	// AstralInitiative int
	Composure       int
	JudgeIntentions int
	Memory          int
	LiftCarry       int
	Movement        int
	// Skills
	// ActiveSkills map[string]skills.ActiveSkill
	// LanguageSkills  map[string]skills.LanguageSkill
	// KnowledgeSkills map[string]skills.KnowledgeSkill
	// Qualities       map[string]string
	// Contacts        map[string]string
	// Identities      map[string]string
	// Lifestyles      map[string]string
	// Currancy        map[string]int
	// RangedWeapons map[string]item.WeaponRanged
	// MeleeWeapons  map[string]item.WeaponMelee
	// Armor           map[string]string
	// Cyberdecks    map[string]string
	// Augmentations map[string]string
	// Vehicals      map[string]string
	// Gear          map[string]string
	// AdeptPowers   map[string]string
}

func (c *Character) GetBody() int {
	return c.Body
}

func (c *Character) GetAgility() int {
	return c.Agility
}

func (c *Character) GetReaction() int {
	return c.Reaction
}

func (c *Character) GetStrength() int {
	return c.Strength
}

func (c *Character) GetWillpower() int {
	return c.Willpower
}

func (c *Character) GetLogic() int {
	return c.Logic
}

func (c *Character) GetIntuition() int {
	return c.Intuition
}

func (c *Character) GetCharisma() int {
	return c.Charisma
}

func (c *Character) GetEdge() int {
	return c.Edge
}

func (c *Character) GetEdgePoints() int {
	return c.EdgePoints
}

func (c *Character) GetEssence() float64 {
	return c.Essence
}

func (c *Character) GetMagic() int {
	return c.Magic
}

func (c *Character) GetResonance() int {
	return c.Resonance
}

func (c *Character) GetPhysicalLimit() int {
	s := float64(c.Strength)
	b := float64(c.Body)
	r := float64(c.Reaction)

	return int(math.Ceil((s*2 + b + r) / 3))
}

func (c *Character) GetMentalLimit() int {
	l := float64(c.Logic)
	i := float64(c.Intuition)
	w := float64(c.Willpower)

	return int(math.Ceil((l*2 + i + w) / 3))
}

func (c *Character) GetSocialLimit() int {
	ch := float64(c.Charisma)
	w := float64(c.Willpower)
	e := c.Essence

	return int(math.Ceil((ch*2 + w + e) / 3))
}

func (c *Character) GetInitiative() int {
	// TODO: Add appropriate attribute and Initiative Dice bonuses
	return c.Reaction + c.Intuition
}

func (c *Character) GetMatrixInitiative() int {
	return c.MatrixInitiative
}

func (c *Character) GetAstralInitiative() int {
	return c.Intuition * 2
}

func (c *Character) GetComposure() int {
	return c.Composure
}

func (c *Character) GetJudgeIntentions() int {
	return c.JudgeIntentions
}

func (c *Character) GetMemory() int {
	return c.Memory
}

func (c *Character) GetLiftCarry() int {
	return c.LiftCarry
}

func (c *Character) GetMovement() int {
	return c.Movement
}

func rollDice(numDice int) (int, []int) {
	rand.Seed(uint64(time.Now().UnixNano()))

	diceRolled := make([]int, numDice)
	total := 0
	for i := 0; i < numDice; i++ {
		roll := rand.Intn(6) + 1
		total += roll
		diceRolled[i] = roll
	}

	return total, diceRolled
}

func (c *Character) RollInitiative() int {
	total, _ := rollDice(1)
	return ((c.Reaction + c.Intuition) * total)
}

var streetSamurai = Character{
	Name:      "Street Samurai",
	Metatype:  metatype.Metatypes["Human"],
	Body:      7,
	Agility:   6,
	Reaction:  5, // (7)
	Strength:  5,
	Willpower: 3,
	Logic:     2,
	Intuition: 3,
	Charisma:  2,
	Essence:   0.88,
	Edge:      1,
	// ActiveSkills: map[string]skills.ActiveSkill{skill.ActiveSkillAutomatics: {Rating: 5},}
	// },
}

/*
FINAL CALCULATIONS TABLE
MECHANIC							FORMULA														AUGMENTATION BONUSES
Initiative							(Reaction + Intuition) + 1D6								Add appropriate attribute and Initiative Dice bonuses
Astral Initiative					(Intuition x 2) + 2D6										—
Matrix AR Initiative				(Reaction + Intuition) + 1D6								—
Matrix VR Initiative (Hot Sim)		(Data Processing + Intuition) + 4D6							—
Matrix VR Initiative (Cold Sim)		(Data Processing + Intuition) + 3D6							—

Inherent Limits						Add appropriate attribute(s); calculate as listed below		—
Mental 								[(Logic x 2) + Intuition + Willpower] / 3 (round up)		—
Social								[(Charisma x 2) + Willpower + Essence] / 3 (round up)		—
Physical							[(Strength x 2) + Body + Reaction] / 3 (round up)			—

Condition Monitor Boxes
Physical 							[Body x 2] + 8												Add bonuses to Body before calculating; round up final results
Stun								[Willpower x 2] + 8											Add bonuses to Willpower before calculating; round up final results
Overflow							Body + Augmentation bonuses									-

Living Persona
Attack								Charisma													—
Data processing						Logic														—
Device Rating						Intuition													—
Firewall							Willpower													—
Sleaze								Resonance													—
Reputation
Notoriety							Public Awareness 											Street Cred

*/
// Overflow Attack Device Rating Sleaze Notoriety

// (Intuition + Reaction) + 1D6
// Add appropriate attribute and Initiative Dice bonuses
// FORMULA
// AUGMENTATION BONUSES
// Astral Initiative
// (Intuition x 2) + 2D6
// —
// (Intuition + Reaction) + 1D6 — (Data Processing + Intuition) + 4D6 — [(Logic x 2) + Intuition + Willpower] / 3 (round up) — [(Charisma x 2) + Willpower + Essence] / 3 (round up) —
// Matrix VR Initiative (Cold Sim)
// (Data Processing + Intuition) + 3D6
// —
// Inherent Limits
// Add appropriate attribute(s); calculate as listed below
// —
// Physical
// [(Strength x 2) + Body + Reaction] / 3 (round up)
// —
// Condition Monitor Boxes
// Calculate as listed below
// —
// Reputation
// [Body / 2] + 8
// Add bonuses to Body before calculating; round up final results
// Stun
// [Willpower / 2] + 8
// Add bonuses to Willpower before calculating; round up final results
// Body + Augmentation bonuses — Charisma — Resonance — Intuition —
// Public Awareness
// Street Cred

// Condition Monitor (P/S)
// 14 / 10
// Armor
// 13
// Limits
// Physical 8 (9), Mental 4, Social 3
// Physical Init
// 10 + 3D6
// Active Skills
// Automatics 5, Blades 5 (6), Longarms 3 (4), Pilot Ground 1, Pistols 4, Sneaking 2 (3), Unarmed Combat 2 (3)
// Knowledge Skills
// Great Restaurants 2, Law Enforcement 2, Poetry 1, Safe Houses 3
// Languages
// English N, Japanese 2
// Qualities
// Ambidextrous, Code of Honor (Bushido), Guts, Home Ground (Street Politics), Incompetent (Acting)
// Augmentations
// Cybereyes [Rating 3, flare compensation, low-light, smartlink, thermographic, vision enhancement 2, vision magnification], dermal plating 2 (alphaware), cyberarm [right, obvious, Strength 11, Agility 9, cyber sub-machine gun w/ external clip port], cyberarm [left, obvious, Strength 11, Agility 9, Armor 2, cyber spur, cyberarm slide], enhanced articulation, platelet factories, reflex recorder [Blades, Longarms, Sneaking, Unarmed Combat], synaptic booster 2, synthcardium 1
// Vehicles
// Harley-Davidson Scorpion [Handling 4/3, Speed 4, Accel 2, Body 8, Armor 9, Pilot 1, Sensor 2]
// Gear
// earbuds [audio enhancement (Rating 1), 4 fake SINs [Rating 4, each with fake licenses (Rating 4, concealed carry, possession of firearms, possession of augmentations)], spatial recognizer], Hermes Ikon commlink (Rating 5), jammer (area, Rating 4), lined coat [9, chemical protection 3, fire resistance 3, non-conductivity 3], medkit (Rating 3), medkit (Rating 6), micro-transceiver, Middle Lifestyle (3 months), 5 stim patches (Rating 6), 2 trauma patches, white noise generator (Rating 6)
// Weapons
// Katana [Blade, Reach 1, Acc 7, DV 14P, AP –3]
// Sword [Blade, Reach 1, Acc 6, DV 14P, AP –2]
// Ares Light Fire 75 [Light Pistol, Acc 6 (8), DV 6P, AP —, SA, RC —,16
// (c),w/ 3 spare clips, 100 rounds regular ammo]
// Ares Predator V [Heavy Pistol, SA, Acc 5 (7), DV 8P, AP –5, RC —, Ammo
// 15 (c), APDS ammo (100 rounds) , 3 spare clips]
// HK-227 [SMG, Acc 5 (7), DV 8P, AP –1, SA/BF/FA, RC (1), 28 (c), w/ 3 spare
// clips, 100 rounds explosive ammo]
// FN HAR [Assault Rifle, Acc 5 (7), DV 10P, AP –6, SA/BF/FA, RC 2, 35 (c), w/
// intergral smartlink, 3 spare clips, 100 rounds APDS ammo]
// Enfield AS-7 [Shotgun, Acc 4 (6), DV 15P(f), AP +4, SA/BF, RC —,10(c) or
// 24 (d), w/ internal smartlink, 3 spare clips, 100 rounds flechette ammo] Ingram Valiant [LMG, Acc 5 (7), DV 10P, AP –3, BF/FA, RC 2 (3), 50 (c) or 100 (belt), w/ integral smartlink, 3 spare clips, 100 rounds explosive
// ammo]
// 3 high explosive grenades [Grenade, non-aerodynamic, DV 16P, AP –2,
// Blast –2/m]
// Spurs (cyber) [Unarmed, Reach —, Acc 9, DV 14P, AP –2]
// Contacts
// Fixer (Connection 4, Loyalty 2)
// Starting ¥
// 2,555 + (4D6 x 100)¥
