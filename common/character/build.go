package character

import (
	"fmt"

	"github.com/Jasrags/ShadowMUD/common/metatype"
	"github.com/Jasrags/ShadowMUD/common/shared"
)

/*
POINT BUY
The Point Buy method has greater flexibility than any other system; the tradeoff, of course, is that the complete range of options available can make character creation a somewhat time-consuming process. For many, though, the time investment is worth it, as they have the chance to design a character precisely the way they want it to be.

In the Point Buy system, you start with 800 Karma. The first thing you have to do is purchase a metatype, as per the Metatype Cost Table.

Once you have purchased your metatype, set your attributes at the minimum levels using the Metatype Attribute Table on p. 66, SR5, or p. 106 of this book. From this point on, the Point Buy system generally works similar to Character Advancement (p 103, SR5), only you are advancing a character with the minimum attributes for their metatype and no skills. The player uses Karma to buy attributes, skills, qualities, contacts (per the rules on p. 98, SR5), gear (at the cost of 1 Karma for every 2,000 nuyen; a maximum of 200 Karma can be spent this way), and anything else needed to flesh the character out. The things that must be taken in consideration are the following: First, as with the Priority System, characters at creation may only have 1 Mental or Physical attribute at their natural maximum (the special attributes—Edge, Magic, and Resonance—do not fall under this limit). Second, if characters want to use Magic or the Resonance, they must buy one of the additional qualities below:

Adept (20 Karma): This makes a character an adept, able to channel mana into physical abilities. They get a Magic Rating of 1 and can buy more ranks with Karma. As with customary character creation, the character gets free power points equal to their Magic Rating. For more information on adepts, see p. 69, SR5.

Aspected Magician (15 Karma): Selecting this quality allows a character to be an aspected magician, meaning they are skilled in one particular area of magic—Sorcery, Conjuring, or Enchanting. They get a Magic Rating of 1 and can buy more ranks with Karma. For more information on the abilities and limitations of aspected magicians, see p. 69, SR5.

Magician (30 Karma): This makes the character a magic-user, able to cast spells, conjure spirits, and use other magical abilities. They get a Magic Rating of 1 and can buy more ranks with Karma. For more information on magicians and what they can do, see p. 69, SR5.

Mystic Adept (35 Karma): This makes the character a mystic adept, a hybrid of magician and adept who can cast spells while also gaining some of the physical abilities of an adept. They get a Magic Rating of 1 and can buy more ranks with Karma. They do not gain free power points; instead, they need to buy power points at a cost of 5 Karma per power point (to a maximum number equal to their Magic Rating).

Technomancer (15 Karma): With this quality, a character becomes a technomancer. They gain the Resonance attribute at a level of 1 and can buy more ranks with Karma. For more information on technomancers, see the Life as a Technomancer in 2075 sidebar on p. 69, SR5, as well as game rules starting on p. 249, SR5.

Note that leftover Karma from the point-buy process cannot be carried over from character creation—it’s use it or lose it! As with the priority system, no more than 5,000 nuyen can be carried over from character creation. Characters roll for starting nuyen per their purchased lifestyle, using the Starting Nuyen Table, p. 95, SR5.

ATTRIBUTE AND SKILL TABLES
The cost to improve an attribute is new Rating x 5 Karma. The calculations for these improvements have already been made in the Karma Advancement Table. To use the table, first find your current Rating in the Starting Rating column on the left, then move to the right until you are in the column whose header matches your desired new rating. For example, if you are raising an Attribute from 4 to 5, find 4 in the Starting Rating column, and move to the right along the row until you find the desired Rating (column 5, in this case). In this case the entry is 25, which means you need to pay 25 Karma for the attribute rating increase (which is equal to Rating 5 x 5 Karma). If you wanted to go from 4 to 6, you’d move one column further to the right and see that you needed a total of 55 Karma to make this increase.

The maximum number of ratings you can increase a single Attribute by in any given period of downtime during a campaign is 2. If you wish to raise the Attribute any further, you will have to wait for more downtime.

The skill table works on a similar principle, though Active Skill ratings costs are computed at new Rating x 2. If you are purchasing a brand new skill, find the desired rating on the table and pay that cumulative amount. For example, if you are purchasing the running skill for the first time, and are buying it up to Rating 3, you will pay12 Karma. To go from 7 to 8 in a skill, you will pay 16 Karma (rating 8 x 2 Karma). To calculate the cost of jumping more than one level, subtract the number in the column with your current Rating from the number in the column with your desired higher Rating. Knowledge and Language skills work in a similar manner, though their cost is only new Rating x 1. A character may raise the Rating of their Active, Knowledge, or Language skills up by a maximum of 3 rating points per any one downtime. To raise the skill(s) any further, they have to wait for another period of downtime.

Active Skill Groups cost new Rating x 5 to raise and can only be raised by 1 rating per downtime.

LEARNING COMPLEX FORMS
To gain access to a new complex form, a technomancer must spend 4 Karma. Details on learning complex forms can be found in the Resonance Library section (p. 252).

LEARNING MAGIC
Aspected magicians, magicians, and mystic adepts may purchase new spells, rituals, or preparations to use (see Magic, p. 276). The magic user must spend 5 Karma to learn the spell, ritual, or preparation. For details on how to learn magic, see p. 299.

QUALITIES
There are two ways for a character to pick up new qualities. First, they can be assigned by the gamemaster as a result of events or actions in the course of a campaign. Positive qualities may be assigned as reward for good roleplaying, while Negative qualities may be assigned if something traumatic or significant happens or the character does something for which the Negative quality is a reasonable consequence (“reasonable” is defined by the gamemaster). A player may also purchase Positive qualities for his character at any time during game play. The cost for purchasing a Positive quality during game play is the listed Karma cost x 2. Similarly, if a character wishes to get rid of a Negative quality, has met any stipulated requirements, and the gamemaster has given the player permission, the player may do so at a rate of listed Karma x 2.
*/

const (
	TotalBuildPoints          = 800
	KarmaNuyenConversionRate  = 2000
	KarmaNuyenConversionLimit = 200
	AttributeCost             = 5 // New rating * 5
	ActiveSkillCost           = 2 // New rating * 2
	KnowledgeSkillCost        = 1 // New rating * 2
	ActiveSkillGroupCost      = 5 // New rating * 5
	SpellCost                 = 5 // New spell
	ComplexFormCost           = 4 // New complex form
	PositiveQualityCostFactor = 2 // 2 * Karma
	NegativeQualityCostFactor = 2 // Remove 2 * Karma

	ChangeActiveSkill    = "active"
	ChangeKnowledgeSkill = "knowledge"
	ChangeSkillGroup     = "group"
	ChangeAttribute      = "attribute"

	// KarmaSpendMetatype   KarmaSpend = "metatype"
	// KarmaSpendAttributes KarmaSpend = "attributes"
	// KarmaSpendSkills     KarmaSpend = "skills"
	// KarmaSpendQualities  KarmaSpend = "qualities"
	// KarmaSpendMagic      KarmaSpend = "magic"

	// NewSkillSpecializationCost = 7
	// NewKnowledgeSkillCost      = 1
	// NewPositiveQualityCost     = 2 // * Karma
	// RemoveNegativeQualityCost  = 2 // * Bonus Karma Value
	// NewComplexFormCost         = 4
	// NewInitiateLevelCost       = 10 // + Grade * 3
	// NewSpell                   = 5
)

type (
	// KarmaSpend     string
	MagicTypeCosts map[MagicType]int

// MetatypeCosts  map[metatype.MetatypeName]int
)

func CalculateChangeCost(changeType string, currentValue, desiredValue int) int {
	if desiredValue == currentValue {
		return 0
	}

	var cost int
	switch changeType {
	case ChangeAttribute:
		cost = AttributeCost
	case ChangeActiveSkill:
		cost = ActiveSkillCost
	case ChangeKnowledgeSkill:
		cost = KnowledgeSkillCost
	case ChangeSkillGroup:
		cost = ActiveSkillGroupCost
	default:
		cost = 0
	}

	totalCost := 0
	if desiredValue < currentValue {
		for i := currentValue; i > desiredValue; i-- {
			totalCost -= i * cost
		}
	} else {
		for i := currentValue + 1; i <= desiredValue; i++ {
			totalCost += i * cost
		}
	}

	return totalCost
}

func GetMagicTypeCost(m MagicType) int {
	v, ok := magicTypeCosts[m]
	if !ok {
		return 0
	}
	return v
}

var (
	// MagicTypeCosts is a map of magic type to the cost of that magic type
	magicTypeCosts = map[MagicType]int{
		MagicTypeNone:             0,
		MagicTypeAdept:            20,
		MagicTypeMagician:         15,
		MagicTypeAspectedMagician: 30,
		MagicTypeMysticAdept:      35,
		MagicTypeTechnomancer:     15,
	}
)

type PointBuilder struct {
	Name           string
	Metatype       *metatype.Metatype
	MagicType      MagicType
	MaxedAttribute shared.AttributeType
	Attributes     map[shared.AttributeType]int
	Skills         map[string]int
	Qualities      map[string]int
	BuildPoints    int
}

// Initialize a new PointBuilder
func NewPointBuilder() *PointBuilder {
	return &PointBuilder{
		Attributes: map[shared.AttributeType]int{
			shared.AttributeBody:      0,
			shared.AttributeAgility:   0,
			shared.AttributeReaction:  0,
			shared.AttributeStrength:  0,
			shared.AttributeWillpower: 0,
			shared.AttributeLogic:     0,
			shared.AttributeIntuition: 0,
			shared.AttributeCharisma:  0,
			shared.AttributeEdge:      0,
			shared.AttributeEssence:   6,
			shared.AttributeMagic:     0,
			shared.AttributeResonance: 0,
		},
		Skills:      make(map[string]int),
		Qualities:   make(map[string]int),
		BuildPoints: TotalBuildPoints,
	}
}

func (pb *PointBuilder) SetName(name string) {
	pb.Name = name
}

// Set the metatype and adjust build points
func (pb *PointBuilder) SetMetatype(m *metatype.Metatype) error {
	if pb.BuildPoints < m.PointCost {
		return fmt.Errorf("not enough build points")
	}
	pb.Metatype = m
	pb.BuildPoints -= m.PointCost

	// Set attributes to metatype minimums
	// Magic and Resonance are set in magic type
	pb.Attributes[shared.AttributeBody] = m.Attributes.Body.Min
	pb.Attributes[shared.AttributeAgility] = m.Attributes.Agility.Min
	pb.Attributes[shared.AttributeReaction] = m.Attributes.Reaction.Min
	pb.Attributes[shared.AttributeStrength] = m.Attributes.Strength.Min
	pb.Attributes[shared.AttributeWillpower] = m.Attributes.Willpower.Min
	pb.Attributes[shared.AttributeLogic] = m.Attributes.Logic.Min
	pb.Attributes[shared.AttributeIntuition] = m.Attributes.Intuition.Min
	pb.Attributes[shared.AttributeCharisma] = m.Attributes.Charisma.Min
	pb.Attributes[shared.AttributeEdge] = m.Attributes.Edge.Min
	pb.Attributes[shared.AttributeEssence] = 6

	return nil
}

func (pb *PointBuilder) RemoveMetatype() {
	if pb.Metatype == nil {
		return
	}

	// Set attributes to 0
	// Magic and Resonance are set with SetMagicType
	pb.Attributes[shared.AttributeBody] = 0
	pb.Attributes[shared.AttributeAgility] = 0
	pb.Attributes[shared.AttributeReaction] = 0
	pb.Attributes[shared.AttributeStrength] = 0
	pb.Attributes[shared.AttributeWillpower] = 0
	pb.Attributes[shared.AttributeLogic] = 0
	pb.Attributes[shared.AttributeIntuition] = 0
	pb.Attributes[shared.AttributeCharisma] = 0
	pb.Attributes[shared.AttributeEdge] = 0
	pb.Attributes[shared.AttributeEssence] = 0

	pb.BuildPoints += pb.Metatype.PointCost
	pb.Metatype = nil
}

func (pb *PointBuilder) SetMagicType(magicType MagicType) error {
	if pb.Metatype == nil {
		return fmt.Errorf("metatype not set")
	}

	pb.MagicType = magicType

	// If we are not magic users just set the name
	if magicType == MagicTypeNone {
		return nil
	}

	pb.BuildPoints -= magicTypeCosts[magicType]

	switch magicType {
	case MagicTypeTechnomancer:
		pb.Attributes[shared.AttributeResonance] = pb.Metatype.Attributes.Resonance.Min
	default:
		pb.Attributes[shared.AttributeMagic] = pb.Metatype.Attributes.Magic.Min
	}

	return nil
}

// Remove the magic type and adjust build points
func (pb *PointBuilder) RemoveMagicType() {
	pb.Attributes[shared.AttributeResonance] = 0
	pb.Attributes[shared.AttributeMagic] = 0

	pb.BuildPoints += magicTypeCosts[pb.MagicType]
	pb.MagicType = MagicTypeNone
}

// TODO: Add more detailed info into error messages (i.e. what is the current value, what is the max, etc.)
func (pb *PointBuilder) AdjustAttribute(attribute shared.AttributeType, newValue int) error {
	// Check if metatype is set
	if pb.Metatype == nil {
		return fmt.Errorf("metatype not set")
	}
	// Check if magic type is set
	if pb.MagicType == "" {
		return fmt.Errorf("magic type not set")
	}
	// Check if attribute is valid
	if _, ok := pb.Attributes[attribute]; !ok {
		return fmt.Errorf("'%s' is not a valid attribute", attribute)
	}
	// Check if magic type is set for magic
	if attribute == shared.AttributeMagic && pb.MagicType == MagicTypeNone {
		return fmt.Errorf("can not adjust magic without being a magic user")
	}
	// Check if magic type is set for resonance
	if attribute == shared.AttributeResonance && pb.MagicType != MagicTypeTechnomancer {
		return fmt.Errorf("can not adjust resonance without being a technomancer")
	}

	min, max := getMetatypeMinMax(attribute, pb)

	// Check if the attribute is at the metatype minimum
	if newValue < min {
		return fmt.Errorf("'%s' (%d) can not be lowered below metatype minimum (%d)", attribute, newValue, min)
	}
	// Check if the attribute is at the metatype maximum
	if newValue > max {
		return fmt.Errorf("'%s' (%d) can not be raised above metatype maximum (%d)", attribute, newValue, max)
	}
	// Check if there is already an attribute at the metatype maximum
	if newValue == max && pb.MaxedAttribute != "" {
		return fmt.Errorf("you may only have one attribute at the metatype maximum")
	}

	currentValue := pb.Attributes[attribute]
	cost := CalculateChangeCost(ChangeAttribute, currentValue, newValue)

	// Check if there are enough build points remaining
	if pb.BuildPoints < cost {
		return fmt.Errorf("not enough remaining build points (%d) for this change (%d)", pb.BuildPoints, cost)
	}

	pb.Attributes[attribute] = newValue
	pb.BuildPoints -= cost
	pb.MaxedAttribute = attribute

	return nil
}

func getMetatypeMinMax(attribute shared.AttributeType, pb *PointBuilder) (int, int) {
	var min, max int
	switch attribute {
	case shared.AttributeBody:
		min = pb.Metatype.Attributes.Body.Min
		max = pb.Metatype.Attributes.Body.Max
	case shared.AttributeAgility:
		min = pb.Metatype.Attributes.Agility.Min
		max = pb.Metatype.Attributes.Agility.Max
	case shared.AttributeReaction:
		min = pb.Metatype.Attributes.Reaction.Min
		max = pb.Metatype.Attributes.Reaction.Max
	case shared.AttributeStrength:
		min = pb.Metatype.Attributes.Strength.Min
		max = pb.Metatype.Attributes.Strength.Max
	case shared.AttributeWillpower:
		min = pb.Metatype.Attributes.Willpower.Min
		max = pb.Metatype.Attributes.Willpower.Max
	case shared.AttributeLogic:
		min = pb.Metatype.Attributes.Logic.Min
		max = pb.Metatype.Attributes.Logic.Max
	case shared.AttributeIntuition:
		min = pb.Metatype.Attributes.Intuition.Min
		max = pb.Metatype.Attributes.Intuition.Max
	case shared.AttributeCharisma:
		min = pb.Metatype.Attributes.Charisma.Min
		max = pb.Metatype.Attributes.Charisma.Max
	case shared.AttributeEdge:
		min = pb.Metatype.Attributes.Edge.Min
		max = pb.Metatype.Attributes.Edge.Max
	case shared.AttributeMagic:
		min = pb.Metatype.Attributes.Magic.Min
		max = pb.Metatype.Attributes.Magic.Max
	case shared.AttributeResonance:
		min = pb.Metatype.Attributes.Resonance.Min
		max = pb.Metatype.Attributes.Resonance.Max
	default:
		min = 0
		max = 0
	}

	return min, max
}

// // TODO: Add more detailed info into error messages (i.e. what is the current value, what is the max, etc.)
// func (pb *PointBuilder) DecreaseAttribute(attribute shared.AttributeType) error {
// 	if pb.Metatype == nil {
// 		return fmt.Errorf("metatype not set")
// 	}
// 	if pb.MagicType == "" {
// 		return fmt.Errorf("magic type not set")
// 	}
// 	if attribute == shared.AttributeMagic && pb.MagicType == MagicTypeNone {
// 		return fmt.Errorf("can not adjust magic without being a magic user")
// 	}
// 	if attribute == shared.AttributeResonance && pb.MagicType != MagicTypeTechnomancer {
// 		return fmt.Errorf("can not adjust resonance without being a technomancer")
// 	}

// 	current := pb.Attributes[attribute]
// 	next := current - 1

// 	min, _ := getMetatypeMinMax(attribute, pb)

// 	if next < min {
// 		return fmt.Errorf("attribute can not be lowered below metatype minimum")
// 	}

// 	// If this is the attribute at the max, clear it
// 	if attribute == pb.MaxedAttribute {
// 		pb.MaxedAttribute = ""
// 	}

// 	cost := current * AttributeCost
// 	pb.Attributes[attribute] -= next
// 	pb.BuildPoints += cost

// 	return nil
// }

// Allocate build points to skills
// TODO: Restrict advancement of magic and resonance skills to the required magic type
func (pb *PointBuilder) AdjustSkill(changeType, skill string, value int) error {
	if pb.Metatype == nil {
		return fmt.Errorf("metatype not set")
	}
	if pb.MagicType == "" {
		return fmt.Errorf("magic type not set")
	}

	current, ok := pb.Skills[skill]
	if !ok {
		current = 0
	}

	if value < 0 {
		return fmt.Errorf("skill value can not be negative")
	}

	if value == current {
		return nil
	}

	if value > 13 {
		return fmt.Errorf("skill value can not be greater than 13")
	}

	cost := CalculateChangeCost(changeType, current, value)

	if pb.BuildPoints < cost {
		return fmt.Errorf("not enough build points")
	}

	pb.Skills[skill] = value
	pb.BuildPoints -= cost

	return nil
}

// Allocate build points to qualities
func (pb *PointBuilder) AllocateQuality(quality string, cost int, positive bool) error {
	var totalCost int
	if positive {
		totalCost = cost * PositiveQualityCostFactor
	} else {
		totalCost = cost * NegativeQualityCostFactor
	}
	if pb.BuildPoints < totalCost {
		return fmt.Errorf("not enough build points")
	}
	pb.Qualities[quality] = cost
	pb.BuildPoints -= totalCost
	return nil
}

// Build the final Character
// func (pb *PointBuilder) Build() *Character {
// 	return &Character{
// 		Name:        pb.name,
// 		Metatype:    pb.metatype,
// 		Attributes:  pb.attributes,
// 		Skills:      pb.skills,
// 		Qualities:   pb.qualities,
// 		BuildPoints: pb.buildPoints,
// 	}
// }

// func (b *PointBuilder) AddSkill(skill string, rating int) {
// 	b.Skills[skill] = rating
// }

// func (b *PointBuilder) SetMetatype(m metatype.MetatypeName) {
// 	// b.Character.Metatype = m
// 	b.Metatype = metatypeCosts[m]
// }

// func (b *PointBuilder) AddMagicType(m MagicType) {
// 	// b.Character.Magic = m
// 	b.Magic = magicTypeCosts[m]
// }

// func (b *PointBuilder) Validate() error {
// 	// Check if metatype is set
// 	if b.Metatype.ID == "" {
// 		return errors.New("Metatype is not set")
// 	}
// 	// Check if magic type is set
// 	// Check if attributes are set
// 	// Check if skills are set
// 	// Check if qualities are set
// 	return nil
// }

// func (b *PointBuilder) Build() {
// 	// Select metatype
// 	// Set attributes to metatype minimums
// 	// Select optional magic type
// 	// Spend on attributes
// 	// Spend on skills
// 	// Spend on qualities
// 	// b.buildMetatype()
// 	// b.buildMagic()
// 	// // TODO: Set metatype minimums and advance with points
// 	// b.buildAttributes()
// 	// b.buildSkills()
// 	// b.buildQualities()
// }
