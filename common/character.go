package common

import (
	"fmt"
	"io"
	"log"
	"math"
	"strings"

	"github.com/Jasrags/ShadowMUD/utils"
	"github.com/gliderlabs/ssh"
	"github.com/sirupsen/logrus"
	"golang.org/x/term"
)

const (
	CharacterDataPath       = "data/characters"
	CharacterFilename       = CharacterDataPath + "/%s.yaml"
	CharacterFileMinVersion = "0.0.1"
)

type AttributesInfo struct {
	Base  int `yaml:"base"`
	Mods  int `yaml:"-"`
	Value int `yaml:"-"`
}

func (ai *AttributesInfo) Reset() {
	ai.Mods = 0
	ai.Value = 0
}

func (ai *AttributesInfo) Recalculate() {
	ai.Value = ai.Base + ai.Mods
}

type AttributesInfoF struct {
	Base  float64 `yaml:"base"`
	Mods  float64 `yaml:"-"`
	Value float64 `yaml:"-"`
}

func (ai *AttributesInfoF) Reset() {
	ai.Mods = 0
	ai.Value = 0
}

func (ai *AttributesInfoF) Recalculate() {
	ai.Value = ai.Base + ai.Mods
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

type InitiativeDice struct {
	Physical AttributesInfo `yaml:"physical"`
	// Astral          AttributesInfo `yaml:"astral"`
	// MatrixAR        AttributesInfo `yaml:"matrix_ar"`
	// MatrixVRHotSim  AttributesInfo `yaml:"matrix_vr"`
	// MatrixVRColdSim AttributesInfo `yaml:"hot_sim"`
	// RiggerAR        AttributesInfo `yaml:"rigger_ar"`
}

type Equipment struct {
	Head          Armor  `yaml:"head,omitempty"`
	Body          Armor  `yaml:"body,omitempty"`
	Weapon        Weapon `yaml:"primary_weapon,omitempty"`
	OffHandWeapon Weapon `yaml:"off_hand_weapon,omitempty"`
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

func NewCharacter(s ssh.Session) *Character {
	pty, ptyWindow, isActive := s.Pty()
	if !isActive {
		logrus.Error("Session is not active")
	}

	return &Character{
		Session: s,
		Pty:     pty,
		Window:  ptyWindow,
		Term:    term.NewTerminal(s, ""),
	}
}

func (c *Character) Authenticate() bool {
	io.WriteString(c.Session, "Username: ")
	username, errReadLine := c.Term.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading username")
		return false
	}
	username = strings.TrimSpace(username)
	logrus.WithField("username", username).Info("Received username")

	passwordBytes, err := c.Term.ReadPassword("Password: ")
	if err != nil {
		log.Println("Error reading password:", err)
		return false
	}
	password := strings.TrimSpace(string(passwordBytes))
	logrus.WithField("password", password).Info("Received password")

	// Validate credentials
	if pass, ok := utils.Users[username]; ok && strings.EqualFold(pass, password) {
		logrus.Info("Authentication successful")

		return true
	}

	return false
}

func (c *Character) Load() {
	c.ID = "ce9f9d47-0a99-4ded-9e9a-30dc3b73f038"
	c.Name = "Test"
	roomSpec := &CoreRooms[0]
	c.Room = Room{
		ID:   roomSpec.ID,
		Spec: roomSpec,
	}
}

func (c *Character) GameLoop() error {
	for {
		io.WriteString(c.Session, ">")
		line, err := c.Term.ReadLine()
		if err != nil {
			return err
		}
		logrus.WithField("line", line).Info("Received line")
		io.WriteString(c.Session, "You typed: "+line+"\n")
	}
}

type Character struct {
	Session ssh.Session       `yaml:"-"`
	Pty     ssh.Pty           `yaml:"-"`
	Window  <-chan ssh.Window `yaml:"-"`
	Term    *term.Terminal    `yaml:"-"`

	// Personal Data
	ID       string   `yaml:"id"`
	Name     string   `yaml:"name"`
	Room     Room     `yaml:"room"`
	Metatype Metatype `yaml:"-"`
	// Ethnicity       string          `yaml:"ethnicity"`
	// Age             int             `yaml:"age"`
	// Sex             string          `yaml:"sex"`
	// Height          int             `yaml:"height"`
	// Weight          int             `yaml:"weight"`
	// StreetCred      int             `yaml:"street_cred"`
	// Notoriety       int             `yaml:"notoriety"`
	// PublicAwareness int             `yaml:"public_awareness"`
	// Karma           int             `yaml:"karma"`
	// TotalKarma      int             `yaml:"total_karma"`
	ConditionDamage ConditionDamage `yaml:"condition_damage"`
	// Attributes
	Attributes     Attributes     `yaml:"attributes"`
	InitiativeDice InitiativeDice `yaml:"initiative_dice"`
	Equipment      Equipment      `yaml:"equipment"`
	// Edge           int            `yaml:"edge"`
	// EdgePoints     int            `yaml:"edge_points"`
	// Derived Attributes
	// Magic         int `yaml:"-"`
	// Resonance     int `yaml:"-"`
	// PhysicalLimit int `yaml:"-"`
	// MentalLimit   int `yaml:"-"`
	// SocialLimit   int `yaml:"-"`
	// Initiative       int
	// MatrixInitiative int `yaml:"-"`
	// AstralInitiative int
	// Composure       int `yaml:"-"`
	// JudgeIntentions int `yaml:"-"`
	// Memory          int `yaml:"-"`
	// LiftCarry       int `yaml:"-"`
	// Movement        int `yaml:"-"`
	// Skills
	// ActiveSkills map[string]ActiveSkill `yaml:"active_skills"`
	// LanguageSkills  map[string]LanguageSkill  `yaml:"language_skills"`
	// KnowledgeSkills map[string]KnowledgeSkill `yaml:"knowledge_skills"`
	// Qualities       map[string]Quality        `yaml:"qualities"`
	// Contacts        map[string]Contact        `yaml:"contacts"`
	// Identities      map[string]string         `yaml:"identities"`
	// Lifestyles      map[string]string         `yaml:"lifestyles"`
	// Currancy        map[string]int            `yaml:"currancy"`
	// Weapons map[string]Weapon `yaml:"weapons"`
	// RangedWeapons map[string]WeaponRanged `yaml:"ranged_weapons"`
	// MeleeWeapons  map[string]WeaponMelee  `yaml:"melee_weapons"`
	// Armor []Armor `yaml:"armor"`
	// Cyberware     map[string]Cyberware    `yaml:"cyberware"`
	// Bioware       map[string]Bioware      `yaml:"bioware"`
	// Cyberdecks      map[string]string         `yaml:"cyberdecks"`
	// Augmentations   map[string]string         `yaml:"augmentations"`
	// Vehicals        map[string]string         `yaml:"vehicals"`
	// Gear map[string]string `yaml:"gear"`
	// AdeptPowers     map[string]string         `yaml:"adept_powers"`
}

// func (c *Character) MeleeAttack(target *Character) {
// 	attackPool := c.Attributes.Agility.Value + c.Skills.CloseCombat + c.Weapon.Accuracy
// 	defensePool := target.Attributes.Reaction.Value + target.Attributes.Intuition.Value

// 	attackHits, _ := c.RollDice(attackPool)
// 	defenseHits, _ := target.RollDice(defensePool)

// 	if attackHits > defenseHits {
// 		netHits := attackHits - defenseHits
// 		damage := c.Attributes.Strength.Value + c.Weapon.Damage + netHits

// 		// Apply additional damage for critical hits
// 		// if attackCriticalHits >= 2 { // Example: 2 or more 6s is a critical hit
// 		// 	damage += attackCriticalHits // Add critical hits to damage
// 		// }

// 		target.ConditionDamage.Physical -= damage
// 	}
// }

// func (c *Character) RollDice(pool int) (int, int) {
// 	rand.Seed(uint64(time.Now().UnixNano()))
// 	hits := 0
// 	criticalHits := 0
// 	for i := 0; i < pool; i++ {
// 		roll := rand.Intn(6) + 1
// 		if roll >= 5 {
// 			hits++
// 		}
// 		if roll == 6 {
// 			criticalHits++
// 		}
// 	}
// 	return hits, criticalHits
// }

// Will need to make this a function that can be called to recalculate
func (c *Character) GetConditionPhysical() int {
	return (c.Attributes.Body.Value / 2) + 8
}

// Will need to make this a function that can be called to recalculate
func (c *Character) GetConditionStun() int {
	return (c.Attributes.Willpower.Value / 2) + 8
}

// TODO: Indomitable quality can modify these limits
// Will need to make this a function that can be called to recalculate
func (c *Character) GetPhysicalLimit() int {
	s := float64(c.Attributes.Strength.Value)
	b := float64(c.Attributes.Body.Value)
	r := float64(c.Attributes.Reaction.Value)

	return int(math.Ceil((s*2 + b + r) / 3))
}

// Will need to make this a function that can be called to recalculate
func (c *Character) GetMentalLimit() int {
	l := float64(c.Attributes.Logic.Value)
	i := float64(c.Attributes.Intuition.Value)
	w := float64(c.Attributes.Willpower.Value)

	return int(math.Ceil((l*2 + i + w) / 3))
}

// Will need to make this a function that can be called to recalculate
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
	RiggerARInitiative        int // (Reaction + Intuition) + 1D6
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

/*
Use base initiative values to roll initiative

FINAL CALCULATIONS TABLE
Physical                Reaction + Intuition + 1D6
Astral                  Intuition x 2 + 2D6
Matrix AR               Reaction + Intuition + 1D6
Matrix VR (Hot Sim)     Data Processing + Intuition + 4D6
Matrix VR (Cold Sim)    Data Processing + Intuition + 3D6
Rigger AR               Reaction + Intuition + 1D6
*/
func (c *Character) RollInitiative() Initiatives {
	// TODO: Add DataProcessing
	total1, _ := utils.RollDice(c.InitiativeDice.Physical.Value)
	// total2, _ := utils.RollDice(c.InitiativeDice.Astral.Value)
	// total3, _ := utils.RollDice(c.InitiativeDice.MatrixAR.Value)
	// total4, _ := utils.RollDice(c.InitiativeDice.MatrixVRHotSim.Value)
	// total5, _ := utils.RollDice(c.InitiativeDice.MatrixVRColdSim.Value)
	// total6, _ := utils.RollDice(c.InitiativeDice.RiggerAR.Value)
	return Initiatives{
		Initiative: (c.Attributes.Reaction.Value + c.Attributes.Intuition.Value) + total1,
		// AstralInitiative:   (c.Attributes.Intuition.Value * 2) + total2,
		// MatrixARInitiative: (c.Attributes.Reaction.Value + c.Attributes.Intuition.Value) + total3,
		// MatrixVRHotSimInitiative:  (c.DataProcessing + c.Intuition)+total4,
		// MatrixVRColdSimInitiative: (c.DataProcessing + c.Intuition)+total5,
		// RiggerARInitiative: (c.Attributes.Reaction.Value + c.Attributes.Intuition.Value) + total6,
	}
}

/*
Composure (WIL + CHA)
There are many common occurrences in a shadowrunner’s life—vicious violence, death, metahuman misery, scary monsters and magic—that would make average citizens crumple into whimpering, traumatized rag-dolls. Whenever a character encounters a situation that she has not been hardened to, the gamemaster can call for a composure test to see whether she faces the situation with cool resolve, temporarily freezes with shock, or trembles and pisses herself.
Composure is a Willpower + Charisma Test, with a threshold based on the severity of the situation (keeping in mind how often the character has faced similar things in the past). Certain situations are bound to become routine to shadowrunners (getting shot at, attacked by a angry spirit, or seeing the remains of a ghoul’s meal); in these cases, gamemasters should no longer ask for composure tests.
*/

// func (c *Character) GetComposure() int {
// 	return (c.Attributes.Willpower.Value + c.Attributes.Charisma.Value)
// }

// /*
// Judge Intentions (INT + CHA)
// A character who wants to use her natural empathy to gauge another character’s emotional state, intentions, or honesty can make an Opposed Intuition + Charisma Test against the target’s Willpower + Charisma. Note that this sort of “psychological” evaluation is never a certainty—it’s just a way for a player to judge what her character “feels” about someone else. It should never serve as a lie detector or detailed psychological analysis. The gamemaster should simply use it as a way to convey gut feelings the character gets when dealing with another.
// */
// func (c *Character) GetJudgeIntentions() int {
// 	return (c.Attributes.Intuition.Value + c.Attributes.Charisma.Value)
// }

// /*
// Memory (LOG + WIL)
// If a character needs to remember an important fact or detail, the gamemaster can call for a Logic + Willpower Success Test. The threshold assigned should be based on how memorable or noticeable the item was: the face of the man that shot him seen close-up would be an easy task (threshold 1), whereas trying to remember the color of some random stranger’s tie glimpsed for only a moment would be more difficult (threshold 3, or even 4). Dice pool modifiers should be applied based on how far back the memory goes or anything that might have prevented a character from taking in all of the details (poor lighting, distractions, etc.).
// A character may also attempt to memorize something in advance. In this case, make a similar Logic + Willpower Test to determine how well the character retains the information. Memorizing long or drawn-out information may have a higher threshold. Each net hit from this test adds an extra die to any memory tests made to recall this information later on.
// A character who glitches on a memory test forgets some details or gets some parts of it wrong. A critical glitch means that the character has deluded himself into believing something entirely different.
// */
// func (c *Character) GetMemory() int {
// 	return (c.Attributes.Logic.Value + c.Attributes.Willpower.Value)
// }

// /*
// Lifting and Carrying (STR + BOD)
// A character can lift off the ground 15 kilograms per point Strength without making a test. If the character wishes to lift more than that, she may make a Strength + Body Test. Each net hit increases the weight she can lift by 15 kilograms more.
// A character can lift 5 kilograms per point Strength over her head without making a test. If the character wishes to lift more than that over her head, she may make a Strength + Body Test. Each net hit increases the weight she can lift by 5 kilograms more.
// Characters can lift and carry their Strength x 10 kilograms in weight without any sort of test. Lifting and carrying more than that calls for a Strength + Body Test. Each hit increases the weight she can lift by 10 kilograms more.
// */
// func (c *Character) GetLiftCarry() int {
// 	return (c.Attributes.Strength.Value + c.Attributes.Body.Value)
// }

// // TODO: Make movement work
// func (c *Character) GetMovement() int {
// 	return 0
// }

// // func (c *Character) Validate() error {
// // 	c.RecalculateAttributes()

// // 	return nil
// // }

// func (c *Character) AddCyberware(cyberware Cyberware) {
// 	c.Cyberware[cyberware.ID] = cyberware
// }

// func (c *Character) RemoveCyberware(id string) {
// 	delete(c.Cyberware, id)
// }

// func (c *Character) RecalculateCyberware() {
// 	// Apply essence modifiers
// 	for _, cw := range c.Cyberware {
// 		c.Attributes.Essence.Mods += cw.Spec.EssenceCost.Value
// 	}
// 	// Apply cyberware modifiers
// 	for _, cyberware := range c.Cyberware {
// 		for _, modifier := range cyberware.Modifiers {
// 			switch modifier.Effect {
// 			case "Increase":
// 				switch modifier.Type {
// 				case "Reaction":
// 					c.Attributes.Reaction.Mods += modifier.Value
// 				}
// 			}
// 		}
// 	}
// // }

// func (c *Character) RecalculateBioware() {
// 	for _, bw := range c.Bioware {
// 		c.Attributes.Essence.Mods += bw.Spec.EssenceCost
// }
// }

func (c *Character) Recalculate() {
	// c.RecalculateCyberware()
	// c.RecalculateBioware()
	c.RecalculateAttributes()
	// c.RecalculateInitiativeDice()
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
	c.Attributes.Essence.Recalculate()
}

func (c *Character) RecalculateInitiativeDice() {
	c.InitiativeDice.Physical.Recalculate()
	// c.InitiativeDice.Astral.Recalculate()
	// c.InitiativeDice.MatrixAR.Recalculate()
	// c.InitiativeDice.MatrixVRHotSim.Recalculate()
	// c.InitiativeDice.MatrixVRColdSim.Recalculate()
	// c.InitiativeDice.RiggerAR.Recalculate()
}

func (c *Character) Save() error {
	return utils.SaveStructToYAML(fmt.Sprintf(CharacterFilename, c.ID), c)
}

// LoadMetatypes loads metatypes from YAML files in a specified directory.
// It populates the global `Metatypes` map with the loaded metatypes.
// The function takes a `sync.WaitGroup` pointer as a parameter to indicate completion.
// It is expected to be called as a goroutine.
func LoadCharacter(id string) Character {
	logrus.WithFields(logrus.Fields{"id": id}).Info("Started loading character")

	var char Character
	if err := utils.LoadStructFromYAML(fmt.Sprintf(CharacterFilename, id), &char); err != nil {
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
