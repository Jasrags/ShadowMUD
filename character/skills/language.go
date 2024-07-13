package skills

type LanguageSkill struct {
	ID       int
	Name     string
	IsCommon bool
}

type (
	LanguageSkillIdx int
)

const (
	LanguageEnglish LanguageSkillIdx = iota
	LanguageSpanish
	LanguageLakota
	LanguageDakota
	LanguageDine
	LanguageRussian
	LanguageFrench
	LanguageItalian
	LanguageGerman
	LanguageAztlanerSpanish
	LanguageSperethiel
	LanguageOrzet
	LanguageJapanese
	LanguageMandarin
)

var languages = map[LanguageSkillIdx]LanguageSkill{
	LanguageEnglish:         {ID: 1, Name: "English", IsCommon: true},
	LanguageSpanish:         {ID: 2, Name: "Spanish", IsCommon: true},
	LanguageLakota:          {ID: 3, Name: "Lakota", IsCommon: true},
	LanguageDakota:          {ID: 4, Name: "Dakota", IsCommon: true},
	LanguageDine:            {ID: 5, Name: "Diné (Navajo)", IsCommon: true},
	LanguageRussian:         {ID: 6, Name: "Russian", IsCommon: true},
	LanguageFrench:          {ID: 7, Name: "French", IsCommon: true},
	LanguageItalian:         {ID: 8, Name: "Italian", IsCommon: true},
	LanguageGerman:          {ID: 9, Name: "German", IsCommon: true},
	LanguageAztlanerSpanish: {ID: 10, Name: "Aztlaner Spanish", IsCommon: true},
	LanguageSperethiel:      {ID: 11, Name: "Sperethiel (elven language)", IsCommon: false},
	LanguageOrzet:           {ID: 12, Name: "Or’zet (ork language)", IsCommon: false},
	LanguageJapanese:        {ID: 13, Name: "Japanese", IsCommon: false},
	LanguageMandarin:        {ID: 14, Name: "Mandarin", IsCommon: false},
}
