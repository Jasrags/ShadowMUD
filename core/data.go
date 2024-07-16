package core

type CoreDataSkills struct {
	Language  map[string]LanguageSkill
	Active    map[string]ActiveSkill
	Knowledge map[string]KnowledgeSkill
}

func NewCoreDataSkills() *CoreDataSkills {
	return &CoreDataSkills{
		Language:  make(map[string]LanguageSkill),
		Active:    make(map[string]ActiveSkill),
		Knowledge: make(map[string]KnowledgeSkill),
	}
}

type CoreData struct {
	Metatypes     map[string]Metatype
	Qualities     map[string]Quality
	CyberwareList map[string]Cyberware
	Skills        *CoreDataSkills
}

func NewCoreData() *CoreData {
	return &CoreData{
		Metatypes:     make(map[string]Metatype),
		Qualities:     make(map[string]Quality),
		CyberwareList: make(map[string]Cyberware),
		Skills:        NewCoreDataSkills(),
	}
}

var LoadedCoreData = NewCoreData()

func LoadCoreData() {
	LoadedCoreData.Metatypes = LoadMetatypes()
	LoadedCoreData.Qualities = LoadQualities()
	LoadedCoreData.CyberwareList = LoadCyberware()
	LoadedCoreData.Skills.Language = LoadLanguageSkills()
	LoadedCoreData.Skills.Active = LoadActiveSkills()
	LoadedCoreData.Skills.Knowledge = LoadKnowledgeSkills()
}
