package jason

import (
	"github.com/Jasrags/ShadowMUD/common/metatype"
	"github.com/Jasrags/ShadowMUD/common/shared"
)

type Character struct {
	cfg *Config

	Name string

	MetatypeID string
	Metatype   *metatype.Metatype
	MagicType  string
	Attributes shared.Attributes

	TotalKarma int
	Karma      int

	// Skills []string
	// Qualities []string
}

func NewCharacter(cfg *Config) *Character {
	return &Character{
		cfg: cfg,

		// MagicType: character.MagicTypeNone,
	}
}

// func (c *Character) SetName(name string) string {
// 	if len(name) < c.cfg.CharacterNameMinLength || len(name) > c.cfg.CharacterNameMaxLength {
// 		return cfmt.Sprintf("Name '%s' must be between %d and %d characters", name, c.cfg.CharacterNameMinLength, c.cfg.CharacterNameMaxLength)
// 	}

// 	if !c.cfg.CharacterNameRegex.MatchString(name) {
// 		return cfmt.Sprintf("Name '%s' must contain only letters, numbers, and underscores", name)
// 	}

// 	for _, n := range c.cfg.BannedNames {
// 		if strings.EqualFold(name, n) {
// 			return cfmt.Sprintf("Name '%s' is not allowed", name)
// 		}
// 	}

// 	for _, n := range c.cfg.CharacterNames {
// 		if strings.EqualFold(name, n) {
// 			return cfmt.Sprintf("Name '%s' is not allowed", name)
// 		}
// 	}

// 	c.Name = name

// 	return ""
// }

// func (c *Character) GetName() string {
// 	return c.Name
// }

// func (c *Character) SetMetatype(m metatype.Metatype) string {
// 	c.MetatypeID = m.ID
// 	c.Metatype = &m

// 	// Apply metatype Min as starting attribute base value (Except for essence which is metatype max)
// 	c.Attributes.Body.Base = m.Attributes.Body.Min
// 	c.Attributes.Agility.Base = m.Attributes.Agility.Min
// 	c.Attributes.Reaction.Base = m.Attributes.Reaction.Min
// 	c.Attributes.Strength.Base = m.Attributes.Strength.Min
// 	c.Attributes.Willpower.Base = m.Attributes.Willpower.Min
// 	c.Attributes.Logic.Base = m.Attributes.Logic.Min
// 	c.Attributes.Intuition.Base = m.Attributes.Intuition.Min
// 	c.Attributes.Charisma.Base = m.Attributes.Charisma.Min
// 	c.Attributes.Edge.Base = m.Attributes.Edge.Min
// 	c.Attributes.Essence.Base = m.Attributes.Essence.Max
// 	// TODO: Apply any racial qualities or restrictions

// 	return ""
// }

// func (c *Character) GetMetatype() metatype.Metatype {
// 	return c.Metatype
// }

// func (c *Character) SetMagicType(magicType character.MagicType) string {
// 	c.MagicType = magicType

// 	return ""
// }

// func (c *Character) GetMagicType() character.MagicType {
// 	return c.MagicType
// }

// func (c *Character) AddKarma(karma int) {
// 	c.Karma += karma
// 	c.TotalKarma += karma
// }

// func (c *Character) RemoveKarma(karma int) string {
// 	if c.Karma < karma {
// 		return cfmt.Sprintf("Not enough karma to remove: %d (Have %d)", karma, c.Karma)
// 	}

// 	c.Karma -= karma

// 	return ""
// }
