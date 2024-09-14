package jason

import (
	"strings"

	"github.com/Jasrags/ShadowMUD/common/metatype"
	"github.com/Jasrags/ShadowMUD/common/shared"
	"github.com/Jasrags/ShadowMUD/utils"
	"github.com/gliderlabs/ssh"
	"github.com/sirupsen/logrus"

	"github.com/i582/cfmt/cmd/cfmt"
)

type Builder struct {
	cfg *Config

	Karma int

	Name       string
	Metatype   *metatype.Metatype
	MagicType  *MagicType
	Attributes shared.Attributes
	Skills     []string

	// Character *Character
}

func NewBuilder(cfg *Config) *Builder {
	return &Builder{
		cfg: cfg,

		Karma: cfg.BuildPointKarma,
		// Character: NewCharacter(cfg),
	}
}

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
		b.Karma += b.Metatype.PointCost
		// TODO: refund any spent karma on attributes
		b.Attributes.Reset()
	}

	// Check if we have enough karma to set the new metatype
	if b.Karma < m.PointCost {
		return cfmt.Sprintf("Not enough karma to set metatype: %s (Need %d, Have %d)", id, m.PointCost, b.Karma)
	}

	// Pay the cost in karma and set the metatype
	b.Karma -= m.PointCost
	b.Metatype = &m

	// Set the attributes to the minimum starting values for the metatype
	b.Attributes.Body.Base = m.Attributes.Body.Min
	b.Attributes.Agility.Base = m.Attributes.Agility.Min
	b.Attributes.Reaction.Base = m.Attributes.Reaction.Min
	b.Attributes.Strength.Base = m.Attributes.Strength.Min
	b.Attributes.Willpower.Base = m.Attributes.Willpower.Min
	b.Attributes.Logic.Base = m.Attributes.Logic.Min
	b.Attributes.Intuition.Base = m.Attributes.Intuition.Min
	b.Attributes.Charisma.Base = m.Attributes.Charisma.Min
	b.Attributes.Essence.Base = m.Attributes.Essence.Max

	b.Attributes.Recalculate()

	return cfmt.Sprintf("Metatype set to '%s' for (%d) karma", m.Name, m.PointCost)
}

func (b *Builder) SetMagicType(s ssh.Session, id string) string {
	// Check if the metatype is set
	if b.Metatype == nil {
		return cfmt.Sprintf("Metatype must be set before setting magic type")
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
		b.Attributes.Magic.Reset()
		b.Attributes.Resonance.Reset()
	}

	// Check if the magic type exists
	m, ok := CoreMagicTypes[id]
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
		b.Attributes.Magic.Base = b.Metatype.Attributes.Magic.Min
	case "technomancer":
		b.Attributes.Resonance.Base = b.Metatype.Attributes.Resonance.Min
	}

	return cfmt.Sprintf("Magic type set to %s for (%d) karma", id, m.PointCost)
}

func (b *Builder) SetAttribute(id string, value int) {

}

// func (b *Builder) GetMagicType() character.MagicType {
// 	return b.Character.GetMagicType()
// }
