package lifestyle

import "github.com/Jasrags/ShadowMUD/common/shared"

const (
	CategoryEntertainmentAsset   Category = "Entertainment - Asset"
	CategoryEntertainmentService Category = "Entertainment - Service"
	CategoryEntertainmentOuting  Category = "Entertainment - Outing"
	CategoryPositive             Category = "Positive"
	CategoryNegative             Category = "Negative"
	CategoryContracts            Category = "Contracts"

	GridSubscriptionGlobal GridSubscription = "Global Grid"
	GridSubscriptionLocal  GridSubscription = "Local Grid"
	GridSubscriptionPublic GridSubscription = "Public Grid"
)

type (
	Category         string
	GridSubscription string
	Comfort          string
	Security         string
	Neighborhood     string
	Quality          string
	City             string
	SecurityRating   string
	Spec             struct {
		ID              string             `yaml:"id"`
		Name            string             `yaml:"name"`
		Description     string             `yaml:"description"`
		Cost            int                `yaml:"cost"`
		FreeGrids       []GridSubscription `yaml:"free_grids"`
		Dice            string             `yaml:"dice"`
		LP              string             `yaml:"lp"`
		CostForArea     string             `yaml:"cost_for_area"`
		CostForComforts string             `yaml:"cost_for_comforts"`
		CostForSecurity string             `yaml:"cost_for_security"`
		Multiplier      string             `yaml:"multiplier"`
		RuleSource      shared.RuleSource  `yaml:"rule_source"`
	}
)
