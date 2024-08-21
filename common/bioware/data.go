package bioware

import "github.com/Jasrags/ShadowMUD/common/shared"

var (
	CoreBioware = []Spec{
		{
			ID:           "adrenaline_pump",
			Name:         "Adrenaline Pump",
			Description:  "Adrenaline Pump Description",
			BiowareType:  BiowareTypeBasic,
			EssenceCost:  0.75, // Rating * 0.75
			Capacity:     Capacity{},
			ToggleAction: shared.ActionFree,
			IsActive:     false,
			Modifiers:    shared.Modifiers{},
			Cost:         55000, // Rating * 55,000¥
			Availability: 0,     // (Rating * 6)F
			Legality:     shared.LegalityTypeLegal,
			Notes:        "",
			RuleSource:   shared.RuleSourceSR5Core,
		},
	}
)
