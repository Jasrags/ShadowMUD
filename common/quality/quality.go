package quality

import (
	"sync"

	"github.com/Jasrags/ShadowMUD/common/shared"
)

const (
	QualitiesFilepath = "_data/qualities"

	TypePositive Type = "Positive"
	TypeNegative Type = "Negative"
)

type (
	Type  string
	Specs map[string]Spec
	Spec  struct {
		ID            string            `yaml:"id"`
		Type          Type              `yaml:"type"`
		Name          string            `yaml:"name"`
		MaxRating     int               `yaml:"max_rating"`
		Description   string            `yaml:"description"`
		Prerequisites []string          `yaml:"prerequisites"`
		Modifiers     shared.Modifiers  `yaml:"modifiers"`
		Cost          int               `yaml:"cost"`
		RuleSource    shared.RuleSource `yaml:"rule_source"`
		Hidden        bool              `yaml:"hidden"`
	}
	Qualities map[string]*Quality
	Quality   struct {
		sync.Mutex `yaml:"-"`
		ID         string `yaml:"id"`
		Rating     int    `yaml:"rating"`
		Spec       *Spec  `yaml:"-"`
	}
)

func NewQuality(spec *Spec) *Quality {
	q := &Quality{
		ID:   spec.ID,
		Spec: spec,
	}

	return q
}
