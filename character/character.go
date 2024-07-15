package character

import (
	"fmt"
	"math"

	"shadowrunmud/character/item"
	"shadowrunmud/character/metatype"
	"shadowrunmud/character/skill"
	"shadowrunmud/util"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	CharacterDataPath       = "data/characters"
	CharacterFilename       = CharacterDataPath + "/%s.yaml"
	CharacterFileMinVersion = "0.0.1"
)

// type Character interface {
// 	GetID() string
// 	SetID(string)
// 	GetName() string
// 	SetName(string)
// 	GetMetatypeName() string
// 	SetMetatypeName(string)
// 	GetMetatype() metatype.Metatype
// 	SetMetatype(metatype.Metatype)
// 	GetEthnicity() string
// 	SetEthnicity(string)
// 	GetAge() int
// 	SetAge(int)
// 	GetSex() string
// 	SetSex(string)
// 	GetHeight() int
// 	SetHeight(int)
// 	GetWeight() int
// 	SetWeight(int)
// 	GetStreetCred() int
// 	SetStreetCred(int)
// 	GetNotoriety() int
// 	SetNotoriety(int)
// 	GetPublicAwareness() int
// 	SetPublicAwareness(int)
// 	GetKarma() int
// 	SetKarma(int)
// 	GetTotalKarma() int
// 	SetTotalKarma(int)
// 	GetBody() int
// 	SetBody(int)
// 	GetAgility() int
// 	SetAgility(int)
// 	GetReaction() int
// 	SetReaction(int)
// 	GetStrength() int
// 	SetStrength(int)
// 	GetWillpower() int
// 	SetWillpower(int)
// 	GetLogic() int
// 	SetLogic(int)
// 	GetIntuition() int
// 	SetIntuition(int)
// 	GetCharisma() int
// 	SetCharisma(int)
// 	GetEdge() int
// 	SetEdge(int)
// 	GetEdgePoints() int
// 	SetEdgePoints(int)
// 	GetEssence() float64
// 	SetEssence(float64)
// 	GetMagic() int
// 	SetMagic(int)
// 	GetResonance() int
// 	SetResonance(int)
// 	GetPhysicalLimit() int
// 	GetMentalLimit() int
// 	GetSocialLimit() int
// 	GetInitiative() Initiatives
// 	RollInitiative() Initiatives
// 	GetComposure() int
// 	GetJudgeIntentions() int
// 	GetMemory() int
// 	GetLiftCarry() int
// 	GetMovement() int
// 	GetActiveSkills() map[string]skill.ActiveSkill
// 	SetActiveSkills(map[string]skill.ActiveSkill)
// 	GetLanguageSkills() map[string]skill.LanguageSkill
// 	SetLanguageSkills(map[string]skill.LanguageSkill)
// 	GetKnowledgeSkills() map[string]skill.KnowledgeSkill
// 	SetKnowledgeSkills(map[string]skill.KnowledgeSkill)
// 	GetQualities() map[string]string
// 	SetQualities(map[string]string)
// 	GetContacts() map[string]string
// 	SetContacts(map[string]string)
// 	GetIdentities() map[string]string
// 	SetIdentities(map[string]string)
// 	GetLifestyles() map[string]string
// 	SetLifestyles(map[string]string)
// 	GetCurrancy() map[string]int
// 	SetCurrancy(map[string]int)
// 	GetRangeWeapons() map[string]item.WeaponRanged
// 	SetRangeWeapons(map[string]item.WeaponRanged)
// 	GetMeleeWeapons() map[string]item.WeaponMelee
// 	SetMeleeWeapons(map[string]item.WeaponMelee)
// 	GetArmor() map[string]string
// 	SetArmor(map[string]string)
// 	GetCyberdecks() map[string]string
// 	SetCyberdecks(map[string]string)
// 	GetAugmentations() map[string]string
// 	SetAugmentations(map[string]string)
// 	GetVehicals() map[string]string
// 	SetVehicals(map[string]string)
// 	GetGear() map[string]string
// 	SetGear(map[string]string)
// 	GetAdeptPowers() map[string]string
// 	SetAdeptPowers(map[string]string)

// 	Save() error
// }

func NewCharacter() *Character {
	uuid := uuid.New().String()
	return &Character{
		ID: uuid,
	}
}

type Character struct {
	// Personal Data
	ID              string            `yaml:"id"`
	Name            string            `yaml:"name"`
	MetatypeName    string            `yaml:"metatype_name"`
	Metatype        metatype.Metatype `yaml:"-"`
	Ethnicity       string            `yaml:"ethnicity"`
	Age             int               `yaml:"age"`
	Sex             string            `yaml:"sex"`
	Height          int               `yaml:"height"`
	Weight          int               `yaml:"weight"`
	StreetCred      int               `yaml:"street_cred"`
	Notoriety       int               `yaml:"notoriety"`
	PublicAwareness int               `yaml:"public_awareness"`
	Karma           int               `yaml:"karma"`
	TotalKarma      int               `yaml:"total_karma"`
	// Attributes
	Body       int `yaml:"body"`
	Agility    int `yaml:"agility"`
	Reaction   int `yaml:"reaction"`
	Strength   int `yaml:"strength"`
	Willpower  int `yaml:"willpower"`
	Logic      int `yaml:"logic"`
	Intuition  int `yaml:"intuition"`
	Charisma   int `yaml:"charisma"`
	Edge       int `yaml:"edge"`
	EdgePoints int `yaml:"edge_points"`
	// Derived Attributes
	Essence       float64 `yaml:"-"`
	Magic         int     `yaml:"-"`
	Resonance     int     `yaml:"-"`
	PhysicalLimit int     `yaml:"-"`
	MentalLimit   int     `yaml:"-"`
	SocialLimit   int     `yaml:"-"`
	// Initiative       int
	MatrixInitiative int `yaml:"-"`
	// AstralInitiative int
	Composure       int `yaml:"-"`
	JudgeIntentions int `yaml:"-"`
	Memory          int `yaml:"-"`
	// LiftCarry       int `yaml:"-"`
	// Movement        int `yaml:"-"`
	// Skills
	ActiveSkills    map[string]skill.ActiveSkill    `yaml:"active_skills"`
	LanguageSkills  map[string]skill.LanguageSkill  `yaml:"language_skills"`
	KnowledgeSkills map[string]skill.KnowledgeSkill `yaml:"knowledge_skills"`
	Qualities       map[string]string               `yaml:"qualities"`
	Contacts        map[string]string               `yaml:"contacts"`
	Identities      map[string]string               `yaml:"identities"`
	Lifestyles      map[string]string               `yaml:"lifestyles"`
	Currancy        map[string]int                  `yaml:"currancy"`
	RangedWeapons   map[string]item.WeaponRanged    `yaml:"ranged_weapons"`
	MeleeWeapons    map[string]item.WeaponMelee     `yaml:"melee_weapons"`
	Armor           map[string]string               `yaml:"armor"`
	Cyberdecks      map[string]string               `yaml:"cyberdecks"`
	Augmentations   map[string]string               `yaml:"augmentations"`
	Vehicals        map[string]string               `yaml:"vehicals"`
	Gear            map[string]string               `yaml:"gear"`
	AdeptPowers     map[string]string               `yaml:"adept_powers"`
}

// Personal Data
func (c *char) GetID() string {
	return c.ID
}

func (c *char) SetID(id string) {
	c.ID = id
}

func (c *char) GetName() string {
	return c.Name
}

func (c *char) SetName(name string) {
	c.Name = name
}

func (c *char) GetMetatypeName() string {
	return c.MetatypeName
}

func (c *char) SetMetatypeName(name string) {
	c.MetatypeName = name
}

func (c *char) GetMetatype() metatype.Metatype {
	return c.Metatype
}

func (c *char) SetMetatype(m metatype.Metatype) {
	c.Metatype = m
}

func (c *char) GetEthnicity() string {
	return c.Ethnicity
}

func (c *char) SetEthnicity(e string) {
	c.Ethnicity = e
}

func (c *char) GetAge() int {
	return c.Age
}

func (c *char) SetAge(a int) {
	c.Age = a
}

func (c *char) GetSex() string {
	return c.Sex
}

func (c *char) SetSex(s string) {
	c.Sex = s
}

func (c *char) GetHeight() int {
	return c.Height
}

func (c *char) SetHeight(h int) {
	c.Height = h
}

func (c *char) GetWeight() int {
	return c.Weight
}

func (c *char) SetWeight(w int) {
	c.Weight = w
}

func (c *char) GetStreetCred() int {
	return c.StreetCred
}

func (c *char) SetStreetCred(sc int) {
	c.StreetCred = sc
}

func (c *char) GetNotoriety() int {
	return c.Notoriety
}

func (c *char) SetNotoriety(n int) {
	c.Notoriety = n
}

func (c *char) GetPublicAwareness() int {
	return c.PublicAwareness
}

func (c *char) SetPublicAwareness(pa int) {
	c.PublicAwareness = pa
}

func (c *char) GetKarma() int {
	return c.Karma
}

func (c *char) SetKarma(k int) {
	c.Karma = k
}

func (c *char) GetTotalKarma() int {
	return c.TotalKarma
}

func (c *char) SetTotalKarma(tk int) {
	c.TotalKarma = tk
}

// Attributes
func (c *char) GetBody() int {
	return c.Body
}

func (c *char) SetBody(b int) {
	c.Body = b
}

func (c *char) GetAgility() int {
	return c.Agility
}

func (c *char) SetAgility(a int) {
	c.Agility = a
}

func (c *char) GetReaction() int {
	return c.Reaction
}

func (c *char) SetReaction(r int) {
	c.Reaction = r
}

func (c *char) GetStrength() int {
	return c.Strength
}

func (c *char) SetStrength(s int) {
	c.Strength = s
}

func (c *char) GetWillpower() int {
	return c.Willpower
}

func (c *char) SetWillpower(w int) {
	c.Willpower = w
}

func (c *char) GetLogic() int {
	return c.Logic
}

func (c *char) SetLogic(l int) {
	c.Logic = l
}

func (c *char) GetIntuition() int {
	return c.Intuition
}

func (c *char) SetIntuition(i int) {
	c.Intuition = i
}

func (c *char) GetCharisma() int {
	return c.Charisma
}

func (c *char) SetCharisma(ch int) {
	c.Charisma = ch
}

func (c *char) GetEdge() int {
	return c.Edge
}

func (c *char) SetEdge(e int) {
	c.Edge = e
}

func (c *char) GetEdgePoints() int {
	return c.EdgePoints
}

func (c *char) SetEdgePoints(ep int) {
	c.EdgePoints = ep
}

// Derived Attributes

func (c *char) GetEssence() float64 {
	return c.Essence
}

func (c *char) SetEssence(e float64) {
	c.Essence = e
}

func (c *char) GetMagic() int {
	return c.Magic
}

func (c *char) SetMagic(m int) {
	c.Magic = m
}

func (c *char) GetResonance() int {
	return c.Resonance
}

func (c *char) SetResonance(r int) {
	c.Resonance = r
}

func (c *char) GetPhysicalLimit() int {
	s := float64(c.Strength)
	b := float64(c.Body)
	r := float64(c.Reaction)

	return int(math.Ceil((s*2 + b + r) / 3))
}

func (c *char) GetMentalLimit() int {
	l := float64(c.Logic)
	i := float64(c.Intuition)
	w := float64(c.Willpower)

	return int(math.Ceil((l*2 + i + w) / 3))
}

func (c *char) GetSocialLimit() int {
	ch := float64(c.Charisma)
	w := float64(c.Willpower)
	e := c.Essence

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
func (c *char) GetInitiative() Initiatives {
	// TODO: Add DataProcessing
	return Initiatives{
		Initiative:         (c.Reaction + c.Intuition),
		AstralInitiative:   (c.Intuition * 2),
		MatrixARInitiative: (c.Reaction + c.Intuition),
		// MatrixVRHotSimInitiative:  (c.DataProcessing + c.Intuition),
		// MatrixVRColdSimInitiative: (c.DataProcessing + c.Intuition),
	}
}

// Use base initiative values to roll initiative
func (c *char) RollInitiative() Initiatives {
	// TODO: Add DataProcessing
	total1, _ := util.RollDice(1)
	total2, _ := util.RollDice(2)
	total3, _ := util.RollDice(1)
	// total4, _ := util.RollDice(4)
	// total5, _ := util.RollDice(3)
	return Initiatives{
		Initiative:         (c.Reaction + c.Intuition) + total1,
		AstralInitiative:   (c.Intuition * 2) + total2,
		MatrixARInitiative: (c.Reaction + c.Intuition) + total3,
		// MatrixVRHotSimInitiative:  (c.DataProcessing + c.Intuition)+total4,
		// MatrixVRColdSimInitiative: (c.DataProcessing + c.Intuition)+total5,
	}
}

/*
Composure (WIL + CHA)
There are many common occurrences in a shadowrunner’s life—vicious violence, death, metahuman misery, scary monsters and magic—that would make average citizens crumple into whimpering, traumatized rag-dolls. Whenever a character encounters a situation that she has not been hardened to, the gamemaster can call for a composure test to see whether she faces the situation with cool resolve, temporarily freezes with shock, or trembles and pisses herself.
Composure is a Willpower + Charisma Test, with a threshold based on the severity of the situation (keeping in mind how often the character has faced similar things in the past). Certain situations are bound to become routine to shadowrunners (getting shot at, attacked by a angry spirit, or seeing the remains of a ghoul’s meal); in these cases, gamemasters should no longer ask for composure tests.
*/
func (c *char) GetComposure() int {
	return (c.Willpower + c.Charisma)
}

/*
Judge Intentions (INT + CHA)
A character who wants to use her natural empathy to gauge another character’s emotional state, intentions, or honesty can make an Opposed Intuition + Charisma Test against the target’s Willpower + Charisma. Note that this sort of “psychological” evaluation is never a certainty—it’s just a way for a player to judge what her character “feels” about someone else. It should never serve as a lie detector or detailed psychological analysis. The gamemaster should simply use it as a way to convey gut feelings the character gets when dealing with another.
*/
func (c *char) GetJudgeIntentions() int {
	return (c.Intuition + c.Charisma)
}

/*
Memory (LOG + WIL)
If a character needs to remember an important fact or detail, the gamemaster can call for a Logic + Willpower Success Test. The threshold assigned should be based on how memorable or noticeable the item was: the face of the man that shot him seen close-up would be an easy task (threshold 1), whereas trying to remember the color of some random stranger’s tie glimpsed for only a moment would be more difficult (threshold 3, or even 4). Dice pool modifiers should be applied based on how far back the memory goes or anything that might have prevented a character from taking in all of the details (poor lighting, distractions, etc.).
A character may also attempt to memorize something in advance. In this case, make a similar Logic + Willpower Test to determine how well the character retains the information. Memorizing long or drawn-out information may have a higher threshold. Each net hit from this test adds an extra die to any memory tests made to recall this information later on.
A character who glitches on a memory test forgets some details or gets some parts of it wrong. A critical glitch means that the character has deluded himself into believing something entirely different.
*/
func (c *char) GetMemory() int {
	return (c.Logic + c.Willpower)
}

/*
Lifting and Carrying (STR + BOD)
A character can lift off the ground 15 kilograms per point Strength without making a test. If the character wishes to lift more than that, she may make a Strength + Body Test. Each net hit increases the weight she can lift by 15 kilograms more.
A character can lift 5 kilograms per point Strength over her head without making a test. If the character wishes to lift more than that over her head, she may make a Strength + Body Test. Each net hit increases the weight she can lift by 5 kilograms more.
Characters can lift and carry their Strength x 10 kilograms in weight without any sort of test. Lifting and carrying more than that calls for a Strength + Body Test. Each hit increases the weight she can lift by 10 kilograms more.
*/
func (c *char) GetLiftCarry() int {
	return (c.Strength + c.Body)
}

// TODO: Make movement work
func (c *char) GetMovement() int {
	return 0
}

// Skills
func (c *char) GetActiveSkills() map[string]skill.ActiveSkill {
	return c.ActiveSkills
}

func (c *char) SetActiveSkills(skills map[string]skill.ActiveSkill) {
	c.ActiveSkills = skills
}

func (c *char) GetLanguageSkills() map[string]skill.LanguageSkill {
	return c.LanguageSkills
}

func (c *char) SetLanguageSkills(skills map[string]skill.LanguageSkill) {
	c.LanguageSkills = skills
}

func (c *char) GetKnowledgeSkills() map[string]skill.KnowledgeSkill {
	return c.KnowledgeSkills
}

func (c *char) SetKnowledgeSkills(skills map[string]skill.KnowledgeSkill) {
	c.KnowledgeSkills = skills
}

func (c *char) GetQualities() map[string]string {
	return c.Qualities
}

func (c *char) SetQualities(qualities map[string]string) {
	c.Qualities = qualities
}

func (c *char) GetContacts() map[string]string {
	return c.Contacts
}

func (c *char) SetContacts(contacts map[string]string) {
	c.Contacts = contacts
}

func (c *char) GetIdentities() map[string]string {
	return c.Identities
}

func (c *char) SetIdentities(identities map[string]string) {
	c.Identities = identities
}

func (c *char) GetLifestyles() map[string]string {
	return c.Lifestyles
}

func (c *char) SetLifestyles(lifestyles map[string]string) {
	c.Lifestyles = lifestyles
}

func (c *char) GetCurrancy() map[string]int {
	return c.Currancy
}

func (c *char) SetCurrancy(currancy map[string]int) {
	c.Currancy = currancy
}

func (c *char) GetRangeWeapons() map[string]item.WeaponRanged {
	return c.RangedWeapons
}

func (c *char) SetRangeWeapons(weapons map[string]item.WeaponRanged) {
	c.RangedWeapons = weapons
}

func (c *char) GetMeleeWeapons() map[string]item.WeaponMelee {
	return c.MeleeWeapons
}

func (c *char) SetMeleeWeapons(weapons map[string]item.WeaponMelee) {
	c.MeleeWeapons = weapons
}

func (c *char) GetArmor() map[string]string {
	return c.Armor
}

func (c *char) SetArmor(armor map[string]string) {
	c.Armor = armor
}

func (c *char) GetCyberdecks() map[string]string {
	return c.Cyberdecks
}

func (c *char) SetCyberdecks(cyberdecks map[string]string) {
	c.Cyberdecks = cyberdecks
}

func (c *char) GetAugmentations() map[string]string {
	return c.Augmentations
}

func (c *char) SetAugmentations(augmentations map[string]string) {
	c.Augmentations = augmentations
}

func (c *char) GetVehicals() map[string]string {
	return c.Vehicals
}

func (c *char) SetVehicals(vehicals map[string]string) {
	c.Vehicals = vehicals
}

func (c *char) GetGear() map[string]string {
	return c.Gear
}

func (c *char) SetGear(gear map[string]string) {
	c.Gear = gear
}

func (c *char) GetAdeptPowers() map[string]string {
	return c.AdeptPowers
}

func (c *char) SetAdeptPowers(adeptPowers map[string]string) {
	c.AdeptPowers = adeptPowers
}

func (c *char) Save() error {
	return util.SaveStructToYAML(fmt.Sprintf(CharacterFilename, c.ID), c)
}

// LoadMetatypes loads metatypes from YAML files in a specified directory.
// It populates the global `Metatypes` map with the loaded metatypes.
// The function takes a `sync.WaitGroup` pointer as a parameter to indicate completion.
// It is expected to be called as a goroutine.
func LoadCharacter(id string) Character {
	logrus.WithFields(logrus.Fields{"id": id}).Debug("Started loading character")

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
