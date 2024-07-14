package item

type CyberwareGrade struct {
	Name                  string
	EssenceCostMultiplier float64
	AvailMod              int
	CostMultiplier        float64
}

var CyberwareGrades = map[string]CyberwareGrade{
	"Standard": {
		Name:                  "Standard",
		EssenceCostMultiplier: 1.0,
		AvailMod:              0,
		CostMultiplier:        1.0,
	},
	"Alphaware": {
		Name:                  "Alphaware",
		EssenceCostMultiplier: 0.8,
		AvailMod:              2,
		CostMultiplier:        1.2,
	},
	"Betaware": {
		Name:                  "Betaware",
		EssenceCostMultiplier: 0.7,
		AvailMod:              4,
		CostMultiplier:        1.5,
	},
	"Deltaware": {
		Name:                  "Deltaware",
		EssenceCostMultiplier: 0.5,
		AvailMod:              8,
		CostMultiplier:        2.5,
	},
	"Used": {
		Name:                  "Used",
		EssenceCostMultiplier: 1.25,
		AvailMod:              -4,
		CostMultiplier:        0.75,
	},
}

type CyberwareType interface {
	GetName() string
	GetDescription() string
	GetEssenceCost() float64
	GetCapacity() int
	GetRating() int
	GetGrade() CyberwareGrade
	GetToggleAction() string
	IsActive() bool
	GetCost() int
	GetAvailability() string
	GetNotes() string
	GetRuleSource() string

	ToggleActivation()
	ModifyReaction() int
	ModifyInitiativeDice() int
}

type Cyberware struct {
	Name         string
	Description  string
	EssenceCost  float64
	Capacity     int
	Rating       int
	Grade        CyberwareGrade
	ToggleAction string
	IsActive     bool
	Cost         int
	Availability string
	Notes        string
	RuleSource   string
}

func (c *Cyberware) ModifyReaction() int {
	var n int
	if c.IsActive {
		n = c.Rating
	}

	return n
}

func (c *Cyberware) ModifyInitiativeDice() int {
	var n int
	if c.IsActive {
		n = c.Rating
	}

	return n
}

var CyberwareList = map[string]Cyberware{
	"Wired Reflexes R1": {
		Name:         "Wired Reflexes R1",
		Description:  "Wired Reflexes R1",
		EssenceCost:  2,
		Grade:        CyberwareGrades["Standard"],
		Cost:         39000,
		Availability: "8R",
		Notes:        "",
		RuleSource:   "Core",
	},
}

/*
De/Activate
    Complex Action
Enhancements
    +Rating Reaction
    +RatingD6 Initiative
Wireless
    The system is compatible with wireless reaction enhancers,
     and the total Reaction bonus from both systems can be above +4
     if both systems have wireless active.
*/
// Part 	Device 	            Essence 	Capacity 	Avail 	Cost 	Source
// Body 	Wired Reflexes R1 	2 	        - 	        8R 	    39,000¥ 	Core
// Body 	Wired Reflexes R2 	3 	        - 	        12R 	149,000¥ 	Core
// Body 	Wired Reflexes R3 	5 	        - 	        20R 	217,000¥ 	Core

// Part 	Device 						Essence 		Capacity 		Avail 			Cost 				Source
// Body 	Bone Lacing (Plastic) 		0.5 			- 				8R 				8,000¥ 				Core
// Body 	Bone Lacing (Aluminum) 		1 				- 				12R 			18,000¥ 			Core
// Body 	Bond Lacing (Titanium) 		1.5 			- 				16R 			30,000¥ 			Core
// Body 	Dermal Plating (R 1-6) 		Rating * 0.5 	- 				(Rating * 4)R 	Rating * 3,000¥ 	Core
// Body 	Fingertip Compartment 		0.1 			[1] 			4 				3,000¥ 				Core
// Body 	Grapple Gun 				0.5 			[4] 			8 				5,000¥ 				Core
// Body 	Internal Air Tank (R 1-3) 	0.25 			[3] 			Rating 			Rating * 4,500¥ 	Core
// Body 	Muscle Replacement (R 1-4) 	Rating 	- 		(Rating * 5)R 	Rating * 25,000¥ 	Core
// Body 	Reaction Enhancers (R 1-3) 	Rating * 0.3 	- 	(Rating * 5)R 	Rating * 13,000¥ 	Core
// Body 	Skillwires (R 1-6) 	Rating * 0.1 	- 	Rating * 4 	Rating * 20,000¥ 	Core
// Body 	Smuggling Compartment 	0.2 	[2] 	6 	7,500¥ 	Core
// Body 	Wired Reflexes R1 	2 	- 	8R 	39,000¥ 	Core
// Body 	Wired Reflexes R2 	3 	- 	12R 	149,000¥ 	Core
// Body 	Wired Reflexes R3 	5 	- 	20R 	217,000¥ 	Core
