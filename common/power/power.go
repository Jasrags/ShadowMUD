package power

import "github.com/Jasrags/ShadowMUD/common/shared"

const ()

type (
	Spec struct {
		ID          string            `yaml:"id"`
		Name        string            `yaml:"name"`
		Description string            `yaml:"description"`
		Points      float64           `yaml:"points"`
		Levels      bool              `yaml:"levels"`
		Limit       int               `yaml:"limit"`
		Action      string            `yaml:"action"`
		AdeptWay    float64           `yaml:"adept_way"`
		RuleSource  shared.RuleSource `yaml:"rule_source"`
		// "adeptwayrequires": {
		//     "required": {
		//       "oneof": {
		//         "quality": [
		//           "The Artist's Way",
		//           "The Beast's Way",
		//           "The Spiritual Way"
		//         ]
		//       }
		//     }
		//   },
		//   "bonus": {
		//     "unlockskills": {
		//       "+content": "Name",
		//       "+@name": "Assensing"
		//     }
		//   }
		// },
	}
	EnhancementSpec struct {
		ID          string            `yaml:"id"`
		Name        string            `yaml:"name"`
		Description string            `yaml:"description"`
		Cost        int               `yaml:"cost"`
		RuleSource  shared.RuleSource `yaml:"rule_source"`
		// "bonus": {
		//   "unarmeddv": "1"
		// },
		// "required": {
		//     "allof": {
		//       "quality": "The Artist's Way"
		//     }
		//   }
	}
)
