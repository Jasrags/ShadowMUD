package core

import (
	"fmt"
	"math"

	"shadowrunmud/core/util"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	CharacterDataPath       = "data/characters"
	CharacterFilename       = CharacterDataPath + "/%s.yaml"
	CharacterFileMinVersion = "0.0.1"
)

func NewCharacter() *Character {
	uuid := uuid.New().String()
	return &Character{
		ID: uuid,
	}
}

type Attributes struct {
	Body      AttributesInfo  `yaml:"body"`
	Agility   AttributesInfo  `yaml:"agility"`
	Reaction  AttributesInfo  `yaml:"reaction"`
	Strength  AttributesInfo  `yaml:"strength"`
	Willpower AttributesInfo  `yaml:"willpower"`
	Logic     AttributesInfo  `yaml:"logic"`
	Intuition AttributesInfo  `yaml:"intuition"`
	Charisma  AttributesInfo  `yaml:"charisma"`
	Essence   AttributesInfoF `yaml:"essence"`
}

type AttributesInfo struct {
	Value int `yaml:"-"`
	Base  int `yaml:"base"`
	Mods  int `yaml:"-"`
}

func (ai *AttributesInfo) Recalculate() {
	ai.Value = ai.Base + ai.Mods
}

type AttributesInfoF struct {
	Value float64 `yaml:"-"`
	Base  float64 `yaml:"base"`
	Mods  float64 `yaml:"-"`
}

type Equipment struct {
	// Weapons   map[string]Weapon    `yaml:"weapons"`
	// Armor     map[string]Armor     `yaml:"armor"`
	// Cyberware map[string]Cyberware `yaml:"cyberware"`
	// Gear      map[string]Gear      `yaml:"gear"`
}

/*
If the damage is Stun, it carries over into the Physical damage track.
For every two full boxes of excess Stun damage, carry over 1 box to
the Physical damage track
• If a character takes more Physical damage than he has boxes in the
Physical damage track, the character is in trouble. Overflowing the
Physical damage track means he’s near death. Instant death occurs only
if damage overflows the Physical damage track by more than the character’s
Body attribute. One point over that limit and his memory will be toasted
at their favorite shadowrunner bar.
*/
type ConditionDamage struct {
	Physical int `yaml:"physical"`
	Stun     int `yaml:"stun"`
}

type Character struct {
	// Personal Data
	ID              string          `yaml:"id"`
	Name            string          `yaml:"name"`
	MetatypeName    string          `yaml:"metatype_name"`
	MetatypeID      string          `yaml:"metatype_id"`
	Metatype        Metatype        `yaml:"-"`
	Ethnicity       string          `yaml:"ethnicity"`
	Age             int             `yaml:"age"`
	Sex             string          `yaml:"sex"`
	Height          int             `yaml:"height"`
	Weight          int             `yaml:"weight"`
	StreetCred      int             `yaml:"street_cred"`
	Notoriety       int             `yaml:"notoriety"`
	PublicAwareness int             `yaml:"public_awareness"`
	Karma           int             `yaml:"karma"`
	TotalKarma      int             `yaml:"total_karma"`
	ConditionDamage ConditionDamage `yaml:"condition_damage"`
	// Attributes
	Attributes Attributes `yaml:"attributes"`
	Edge       int        `yaml:"edge"`
	EdgePoints int        `yaml:"edge_points"`
	// Derived Attributes
	// Essence       float64 `yaml:"-"`
	Magic         int `yaml:"-"`
	Resonance     int `yaml:"-"`
	PhysicalLimit int `yaml:"-"`
	MentalLimit   int `yaml:"-"`
	SocialLimit   int `yaml:"-"`
	// Initiative       int
	MatrixInitiative int `yaml:"-"`
	// AstralInitiative int
	Composure       int `yaml:"-"`
	JudgeIntentions int `yaml:"-"`
	Memory          int `yaml:"-"`
	// LiftCarry       int `yaml:"-"`
	// Movement        int `yaml:"-"`
	// Skills
	ActiveSkills    map[string]ActiveSkill    `yaml:"active_skills"`
	LanguageSkills  map[string]LanguageSkill  `yaml:"language_skills"`
	KnowledgeSkills map[string]KnowledgeSkill `yaml:"knowledge_skills"`
	Qualities       map[string]string         `yaml:"qualities"`
	Contacts        map[string]string         `yaml:"contacts"`
	Identities      map[string]string         `yaml:"identities"`
	Lifestyles      map[string]string         `yaml:"lifestyles"`
	Currancy        map[string]int            `yaml:"currancy"`
	RangedWeapons   map[string]WeaponRanged   `yaml:"ranged_weapons"`
	MeleeWeapons    map[string]WeaponMelee    `yaml:"melee_weapons"`
	Armor           map[string]string         `yaml:"armor"`
	Cyberdecks      map[string]string         `yaml:"cyberdecks"`
	Augmentations   map[string]string         `yaml:"augmentations"`
	Vehicals        map[string]string         `yaml:"vehicals"`
	Gear            map[string]string         `yaml:"gear"`
	AdeptPowers     map[string]string         `yaml:"adept_powers"`
}

func (c *Character) GetConditionPhysical() int {
	return (c.Attributes.Body.Value / 2) + 8
}

func (c *Character) GetConditionStun() int {
	return (c.Attributes.Willpower.Value / 2) + 8
}

// TODO: Indomitable quality can modify these limits
func (c *Character) GetPhysicalLimit() int {
	s := float64(c.Attributes.Strength.Value)
	b := float64(c.Attributes.Body.Value)
	r := float64(c.Attributes.Reaction.Value)

	return int(math.Ceil((s*2 + b + r) / 3))
}

func (c *Character) GetMentalLimit() int {
	l := float64(c.Attributes.Logic.Value)
	i := float64(c.Attributes.Intuition.Value)
	w := float64(c.Attributes.Willpower.Value)

	return int(math.Ceil((l*2 + i + w) / 3))
}

func (c *Character) GetSocialLimit() int {
	ch := float64(c.Attributes.Charisma.Value)
	w := float64(c.Attributes.Willpower.Value)
	e := c.Attributes.Essence.Value

	return int(math.Ceil((ch*2 + w + e) / 3))
}

type Initiatives struct {
	Initiative                int // (Reaction + Intuition) + 1D6
	AstralInitiative          int // (Intuition x 2) + 2D6
	MatrixARInitiative        int // (Reaction + Intuition) + 1D6
	MatrixVRHotSimInitiative  int // (Data Processing + Intuition) + 4D6
	MatrixVRColdSimInitiative int // (Data Processing + Intuition) + 3D6
}

// Return base initiative values
func (c *Character) GetInitiative() Initiatives {
	// TODO: Add DataProcessing
	return Initiatives{
		Initiative:         (c.Attributes.Reaction.Value + c.Attributes.Intuition.Value),
		AstralInitiative:   (c.Attributes.Intuition.Value * 2),
		MatrixARInitiative: (c.Attributes.Reaction.Value + c.Attributes.Intuition.Value),
		// MatrixVRHotSimInitiative:  (c.DataProcessing + c.Intuition),
		// MatrixVRColdSimInitiative: (c.DataProcessing + c.Intuition),
	}
}

// Use base initiative values to roll initiative
func (c *Character) RollInitiative() Initiatives {
	// TODO: Add DataProcessing
	total1, _ := util.RollDice(1)
	total2, _ := util.RollDice(2)
	total3, _ := util.RollDice(1)
	// total4, _ := util.RollDice(4)
	// total5, _ := util.RollDice(3)
	return Initiatives{
		Initiative:         (c.Attributes.Reaction.Value + c.Attributes.Intuition.Value) + total1,
		AstralInitiative:   (c.Attributes.Intuition.Value * 2) + total2,
		MatrixARInitiative: (c.Attributes.Reaction.Value + c.Attributes.Intuition.Value) + total3,
		// MatrixVRHotSimInitiative:  (c.DataProcessing + c.Intuition)+total4,
		// MatrixVRColdSimInitiative: (c.DataProcessing + c.Intuition)+total5,
	}
}

/*
Composure (WIL + CHA)
There are many common occurrences in a shadowrunner’s life—vicious violence, death, metahuman misery, scary monsters and magic—that would make average citizens crumple into whimpering, traumatized rag-dolls. Whenever a character encounters a situation that she has not been hardened to, the gamemaster can call for a composure test to see whether she faces the situation with cool resolve, temporarily freezes with shock, or trembles and pisses herself.
Composure is a Willpower + Charisma Test, with a threshold based on the severity of the situation (keeping in mind how often the character has faced similar things in the past). Certain situations are bound to become routine to shadowrunners (getting shot at, attacked by a angry spirit, or seeing the remains of a ghoul’s meal); in these cases, gamemasters should no longer ask for composure tests.
*/
func (c *Character) GetComposure() int {
	return (c.Attributes.Willpower.Value + c.Attributes.Charisma.Value)
}

/*
Judge Intentions (INT + CHA)
A character who wants to use her natural empathy to gauge another character’s emotional state, intentions, or honesty can make an Opposed Intuition + Charisma Test against the target’s Willpower + Charisma. Note that this sort of “psychological” evaluation is never a certainty—it’s just a way for a player to judge what her character “feels” about someone else. It should never serve as a lie detector or detailed psychological analysis. The gamemaster should simply use it as a way to convey gut feelings the character gets when dealing with another.
*/
func (c *Character) GetJudgeIntentions() int {
	return (c.Attributes.Intuition.Value + c.Attributes.Charisma.Value)
}

/*
Memory (LOG + WIL)
If a character needs to remember an important fact or detail, the gamemaster can call for a Logic + Willpower Success Test. The threshold assigned should be based on how memorable or noticeable the item was: the face of the man that shot him seen close-up would be an easy task (threshold 1), whereas trying to remember the color of some random stranger’s tie glimpsed for only a moment would be more difficult (threshold 3, or even 4). Dice pool modifiers should be applied based on how far back the memory goes or anything that might have prevented a character from taking in all of the details (poor lighting, distractions, etc.).
A character may also attempt to memorize something in advance. In this case, make a similar Logic + Willpower Test to determine how well the character retains the information. Memorizing long or drawn-out information may have a higher threshold. Each net hit from this test adds an extra die to any memory tests made to recall this information later on.
A character who glitches on a memory test forgets some details or gets some parts of it wrong. A critical glitch means that the character has deluded himself into believing something entirely different.
*/
func (c *Character) GetMemory() int {
	return (c.Attributes.Logic.Value + c.Attributes.Willpower.Value)
}

/*
Lifting and Carrying (STR + BOD)
A character can lift off the ground 15 kilograms per point Strength without making a test. If the character wishes to lift more than that, she may make a Strength + Body Test. Each net hit increases the weight she can lift by 15 kilograms more.
A character can lift 5 kilograms per point Strength over her head without making a test. If the character wishes to lift more than that over her head, she may make a Strength + Body Test. Each net hit increases the weight she can lift by 5 kilograms more.
Characters can lift and carry their Strength x 10 kilograms in weight without any sort of test. Lifting and carrying more than that calls for a Strength + Body Test. Each hit increases the weight she can lift by 10 kilograms more.
*/
func (c *Character) GetLiftCarry() int {
	return (c.Attributes.Strength.Value + c.Attributes.Body.Value)
}

// TODO: Make movement work
func (c *Character) GetMovement() int {
	return 0
}

func (c *Character) Validate() error {
	c.RecalculateAttributes()

	return nil
}

func (c *Character) RecalculateAttributes() {
	c.Attributes.Body.Recalculate()
	c.Attributes.Agility.Recalculate()
	c.Attributes.Reaction.Recalculate()
	c.Attributes.Strength.Recalculate()
	c.Attributes.Willpower.Recalculate()
	c.Attributes.Logic.Recalculate()
	c.Attributes.Intuition.Recalculate()
	c.Attributes.Charisma.Recalculate()
}

func (c *Character) Save() error {
	return util.SaveStructToYAML(fmt.Sprintf(CharacterFilename, c.ID), c)
}

// LoadMetatypes loads metatypes from YAML files in a specified directory.
// It populates the global `Metatypes` map with the loaded metatypes.
// The function takes a `sync.WaitGroup` pointer as a parameter to indicate completion.
// It is expected to be called as a goroutine.
func LoadCharacter(id string) Character {
	logrus.WithFields(logrus.Fields{"id": id}).Info("Started loading character")

	var char Character
	if err := util.LoadStructFromYAML(fmt.Sprintf(CharacterFilename, id), &char); err != nil {
		logrus.WithFields(logrus.Fields{"id": id}).WithError(err).Fatal("Could not load character")
	}

	// if char.GetMetatypeName() != "" {
	// 	m, _ := metatype.LoadMetatype(char.GetMetatypeName())
	// 	char.SetMetatype(m)
	// }

	logrus.WithFields(logrus.Fields{"id": id}).Info("Loaded character file")

	return char
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
