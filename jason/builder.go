package jason

import (
	"strconv"
	"strings"

	"github.com/Jasrags/ShadowMUD/common/character"
	"github.com/Jasrags/ShadowMUD/common/magic"
	"github.com/Jasrags/ShadowMUD/common/metatype"
	"github.com/Jasrags/ShadowMUD/utils"

	"github.com/gliderlabs/ssh"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sirupsen/logrus"
)

func NewAttributes() Attributes {
	return Attributes{
		"agility":   {Name: "Agility"},
		"body":      {Name: "Body"},
		"charisma":  {Name: "Charisma"},
		"intuition": {Name: "Intuition"},
		"logic":     {Name: "Logic"},
		"magic":     {Name: "Magic"},
		"reaction":  {Name: "Reaction"},
		"resonance": {Name: "Resonance"},
		"strength":  {Name: "Strength"},
		"willpower": {Name: "Willpower"},
		"edge":      {Name: "Edge"},
	}
}

type Attributes map[string]*Attribute[int]

func (a *Attributes) Reset() {
	for _, attr := range *a {
		attr.Reset()
	}
}

func (a *Attributes) Recalculate() {
	for _, attr := range *a {
		attr.Recalculate()
	}
}

type Attribute[T int | float64] struct {
	Name       string `yaml:"name"`
	Base       T      `yaml:"base"`
	Delta      T      `yaml:"delta"`
	TotalValue T      `yaml:"total_value"`
	PointCost  int    `yaml:"point_cost"`
}

func (a *Attribute[T]) SetBase(value T) {
	a.Base = value
	a.Recalculate()
}

func (a *Attribute[T]) Recalculate() {
	a.TotalValue = a.Base + a.Delta
}

func (a *Attribute[T]) Reset() {
	a.Base = 0
	a.Delta = 0
	a.TotalValue = 0
}

type (
	Qualities map[string]*Quality
	Quality   struct {
		ID     string `yaml:"id"`
		Rating int    `yaml:"rating"`
		// Spec   *quality.Spec `yaml:"-"`
	}
	Skills map[string]*Skill
	Skill  struct {
		ID             string `yaml:"id"`
		Specialization string `yaml:"specialization"`
		Rating         int    `yaml:"rating"`
		// Spec           *skill.Spec `yaml:"-"`
	}
)

type Builder struct {
	cfg *Config

	Karma int

	Name       string
	Metatype   *metatype.Metatype
	MagicType  *magic.MagicType
	Attributes Attributes
	Skills     Skills
	Qualties   Qualities
	Essence    Attribute[float64]

	// Character *Character
}

func NewBuilder(cfg *Config) *Builder {
	return &Builder{
		cfg: cfg,

		Karma:      cfg.BuildPointKarma,
		Attributes: NewAttributes(),
	}
}

// func (b *Builder) getMetatypeMinMax(id string) (int, int) {
// 	if b.Metatype == nil {
// 		return 0, 0
// 	}

// 	if _, ok := b.Metatype.Attributes[id]; !ok {
// 		return 0, 0
// 	}

// 	min := b.Metatype.Attributes[id].Min
// 	max := b.Metatype.Attributes[id].Max

// 	return min, max
// }

func (b *Builder) SetName(name string) string {
	if len(name) < b.cfg.CharacterNameMinLength || len(name) > b.cfg.CharacterNameMaxLength {
		return cfmt.Sprintf("{{Name '%s' must be between %d and %d characters}}::#ff8700", name, b.cfg.CharacterNameMinLength, b.cfg.CharacterNameMaxLength)
	}

	if !b.cfg.CharacterNameRegex.MatchString(name) {
		return cfmt.Sprintf("{{Name '%s' must contain only letters, numbers, and underscores}}::#ff8700", name)
	}

	for _, n := range b.cfg.BannedNames {
		if strings.EqualFold(name, n) {
			return cfmt.Sprintf("{{Name '%s' is not allowed}}::#ff8700", name)
		}
	}

	for _, n := range b.cfg.CharacterNames {
		if strings.EqualFold(name, n) {
			return cfmt.Sprintf("{{Name '%s' is not allowed}}::#ff8700", name)
		}
	}

	b.Name = name

	return cfmt.Sprintf("{{Name set to: %s}}::#0000ff", name)
}

func (b *Builder) SetMetatype(s ssh.Session, id string) string {
	id = strings.ToLower(id)

	// Check if the metatype exists
	m, ok := metatype.CoreMetatypes[id]
	if !ok {
		return cfmt.Sprintf("Invalid metatype ID: %s", id)
	}

	// Reset and refund karma attributes if we have a metatype already
	if b.Metatype != nil {
		ok, err := utils.PromptConfirmInput(s, "Changing your metatype will reset your all your attributes and refund any spent karma.\n{{Are you sure you want to continue? (y/n)}}::yellow")
		if err != nil {
			logrus.WithError(err).Error("Error prompting for confirmation")
		}
		if !ok {
			return cfmt.Sprintf("{{Metatype change cancelled.}}::cyan\n")
		}
		// Refund karma for the old metatype
		b.Karma += b.Metatype.PointCost

		// Refund karma for any spent attributes and reset them
		for _, a := range b.Attributes {
			b.Karma += a.PointCost
			a.Reset()
		}
	}

	// Check if we have enough karma to set the new metatype
	if b.Karma < m.PointCost {
		return cfmt.Sprintf("Not enough karma to set metatype: %s (Need %d, Have %d)", id, m.PointCost, b.Karma)
	}

	// Pay the cost in karma and set the metatype
	b.Karma -= m.PointCost
	b.Metatype = &m

	// Set the attributes to the minimum starting values for the metatype
	b.Attributes["agility"].Base = m.Attributes["agility"].Min
	b.Attributes["body"].Base = m.Attributes["body"].Min
	b.Attributes["charisma"].Base = m.Attributes["charisma"].Min
	b.Attributes["edge"].Base = m.Attributes["edge"].Min
	b.Attributes["intuition"].Base = m.Attributes["intuition"].Min
	b.Attributes["logic"].Base = m.Attributes["logic"].Min
	b.Attributes["reaction"].Base = m.Attributes["reaction"].Min
	b.Attributes["strength"].Base = m.Attributes["strength"].Min
	b.Attributes["willpower"].Base = m.Attributes["willpower"].Min
	b.Attributes["magic"].Base = m.Attributes["magic"].Min
	b.Attributes["resonance"].Base = m.Attributes["resonance"].Min
	b.Attributes.Recalculate()

	b.Essence.Base = m.Essence.Max
	b.Essence.Recalculate()

	// Add metatype qualities
	if len(m.Qualities) < 1 {
		for _, quality := range m.Qualities {
			b.Qualties[quality] = &Quality{ID: quality}
		}
	}

	return cfmt.Sprintf("Metatype set to '%s' for (%d) karma", m.Name, m.PointCost)
}

func (b *Builder) SetMagicType(s ssh.Session, id string) string {
	// Check if the metatype is set
	if b.Metatype == nil {
		return cfmt.Sprintf("{{Metatype must be set before setting magic type}}::yellow")
	}

	// Reset and refund karma attributes if we have a magic type already
	if b.MagicType != nil {
		ok, err := utils.PromptConfirmInput(s, "Changing your magic type will reset your magic/resonance attributes and refund any spent karma.\n{{Are you sure you want to continue? (y/n)}}::yellow")
		if err != nil {
			logrus.WithError(err).Error("Error prompting for confirmation")
		}
		if !ok {
			return cfmt.Sprintf("{{Magic type change cancelled.}}::cyan")
		}

		b.Karma += b.MagicType.PointCost
		b.MagicType = nil
		// TODO: refund any spent karma on magic and resonance
		b.Attributes["magic"].Reset()
		b.Attributes["resonance"].Reset()
	}

	// Check if the magic type exists
	m, ok := magic.CoreMagicTypes[id]
	if !ok {
		return cfmt.Sprintf("Invalid magic type ID: %s", id)
	}

	// Check if we have enough karma to set the new magic type
	if b.Karma < m.PointCost {
		return cfmt.Sprintf("Not enough karma to set magic type: %s (Need %d, Have %d)", m.Name, m.PointCost, b.Karma)
	}

	// Pay the cost in karma and set the magic type
	b.Karma -= m.PointCost
	b.MagicType = &m

	// Set the magic or resonance attribute based on the magic type
	switch id {
	case "none":
	case "adept":
		fallthrough
	case "magician":
		fallthrough
	case "aspected_magician":
		fallthrough
	case "mystic_adept":
		b.Attributes["magic"].Base = b.Metatype.Attributes["magic"].Min
	case "technomancer":
		b.Attributes["resonance"].Base = b.Metatype.Attributes["resonance"].Min
	}

	return cfmt.Sprintf("Magic type set to %s for (%d) karma", id, m.PointCost)
}

func (b *Builder) SetAttribute(id, value string) string {
	// Check if metatype and magic type are set
	if b.Metatype == nil {
		return cfmt.Sprintf("{{Metatype must be set before setting attributes}}::yellow")
	}

	if b.MagicType == nil {
		return cfmt.Sprintf("{{Magic type must be set before setting attributes}}::yellow")
	}

	id = strings.ToLower(id)

	v, errAtio := strconv.Atoi(value)
	if errAtio != nil {
		return cfmt.Sprintf("Invalid value for attribute '%s': %s", id, value)
	}

	if id == "magic" && b.MagicType.ID == "none" {
		return cfmt.Sprintf("{{Cannot set magic attribute when magic type is none}}::yellow")
	}
	if id == "magic" && b.MagicType.ID == "technomancer" {
		return cfmt.Sprintf("{{Cannot set magic attribute when magic type is technomancer}}::yellow")
	}
	if id == "resonance" && b.MagicType.ID != "technomancer" {
		return cfmt.Sprintf("{{Cannot set resonance attribute when magic type is not technomancer}}::yellow")
	}

	// Check if the id is a valid attribute
	attr, ok := b.Attributes[id]
	if !ok {
		return cfmt.Sprintf("Invalid attribute ID: %s", id)
	}

	// Check if the id is a valid metatdata attribute
	mattr, ok := b.Metatype.Attributes[id]
	if !ok {
		return cfmt.Sprintf("Invalid metatype attribute ID: %s", id)
	}

	cost := character.CalculateChangeCost(character.ChangeAttribute, b.Attributes[id].Base, v)

	logrus.WithFields(logrus.Fields{"id": id, "base": b.Attributes[id].Base, "value": v, "cost": cost}).Debug("Calculating attribute change cost")
	// Check if the value is within the allowed range
	if v < mattr.Min || v > mattr.Max {
		return cfmt.Sprintf("{{'%s' must be between %d and %d}}::yellow", attr.Name, mattr.Min, mattr.Max)
	}

	attr.Base = v
	attr.PointCost = cost
	b.Karma -= cost

	return cfmt.Sprintf("Attribute '%s' set to %d for (%d) karma", id, v, cost)
}

// func (b *Builder) AddQuality(id string) string {
// 	// Check if metatype is set
// 	if b.Metatype == nil {
// 		return cfmt.Sprintf("{{Metatype must be set before adding qualities}}::yellow")
// 	}

// 	// Check if the quality exists
// 	// q, ok := character.CoreQualities[id]
// 	// if !ok {
// 	// return cfmt.Sprintf("Invalid quality ID: %s", id)
// 	// }

// 	// Check if the quality is already added
// 	for _, quality := range b.Qualties {
// 		if quality == id {
// 			return cfmt.Sprintf("Quality '%s' is already added", id)
// 		}
// 	}

// 	// Check if we have enough karma to add the quality
// 	if b.Karma < q.PointCost {
// 		return cfmt.Sprintf("Not enough karma to add quality: %s (Need %d, Have %d)", id, q.PointCost, b.Karma)
// 	}

// 	// Pay the cost in karma and add the quality
// 	b.Karma -= q.PointCost
// 	b.Qualties = append(b.Qualties, id)

// 	return cfmt.Sprintf("Quality '%s' added for (%d) karma", id, q.PointCost)
// }
