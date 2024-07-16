package core

type TestAttribute struct {
	Base  int
	Mods  int
	Value int
}

func (ta *TestAttribute) Reset() {
	ta.Value = 0
	ta.Mods = 0
}
func (ta *TestAttribute) Recalculate() {
	ta.Value = ta.Base + ta.Mods
}

type TestChar struct {
	Body      TestAttribute
	Agility   TestAttribute
	Reaction  TestAttribute
	Strength  TestAttribute
	Willpower TestAttribute
	Logic     TestAttribute
	Intuition TestAttribute
	Charisma  TestAttribute
	Essence   float32
	Skills    []TestSkill
}

func (tc *TestChar) Recalculate() {
	tc.Body.Reset()
	tc.Agility.Reset()
	tc.Reaction.Reset()
	tc.Strength.Reset()
	tc.Willpower.Reset()
	tc.Logic.Reset()
	tc.Intuition.Reset()
	tc.Charisma.Reset()

	for _, skill := range tc.Skills {
		switch skill.Attribute {
		case "Agility":
			tc.Agility.Mods += skill.Rank
		}
	}

	tc.Body.Recalculate()
	tc.Agility.Recalculate()
	tc.Reaction.Recalculate()
	tc.Strength.Recalculate()
	tc.Willpower.Recalculate()
	tc.Logic.Recalculate()
	tc.Intuition.Recalculate()
	tc.Charisma.Recalculate()

}

type TestSkill struct {
	Name      string
	Attribute string
	Rank      int
}

// var (
// 	tsk = TestSkill{
// 		Name:      "Automatics",
// 		Attribute: "Agility",
// 		Rank:      4,
// 	}

// 	tc = TestChar{
// 		Body:      TestAttribute{Base: 5},
// 		Agility:   TestAttribute{Base: 7},
// 		Reaction:  TestAttribute{Base: 6},
// 		Strength:  TestAttribute{Base: 8},
// 		Willpower: TestAttribute{Base: 5},
// 		Logic:     TestAttribute{Base: 7},
// 		Intuition: TestAttribute{Base: 6},
// 		Charisma:  TestAttribute{Base: 8},
// 		Essence:   6.0,
// 		Skills:    []TestSkill{tsk},
// 	}
// )
