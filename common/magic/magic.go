package magic

import "github.com/Jasrags/ShadowMUD/common/shared"

type (
	MagicTypes map[string]MagicType
	MagicType  struct {
		ID          string
		Name        string
		Description string
		PointCost   int
		Hidden      bool
		RuleSource  shared.RuleSource
	}
)
