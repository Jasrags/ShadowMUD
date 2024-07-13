package skills

type KnowledgeSkill struct {
	ID       int
	Name     string
	IsCommon bool
}

type (
	KnowledgeSkillIdx int
)

const (
	KnowledgeSkillSeattleStreetGangs KnowledgeSkillIdx = iota
)
