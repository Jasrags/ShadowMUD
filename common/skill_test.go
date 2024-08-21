package common_test

// const ActiveSkillsFile = "/Users/jrags/Code/Jasrags/ShadowMUD/_data/skills/active.yaml"
// const KnowledgeSkillsFile = "/Users/jrags/Code/Jasrags/ShadowMUD/_data/skills/knowledge.yaml"

// const SkillXmlFile = "/Users/jrags/Code/Jasrags/ShadowMUD/chummer_data/skills.xml"

// type Skills []common.Skill

// func TestReadFiles(t *testing.T) {
// 	list := []common.Skill{}
// 	if err := utils.LoadStructFromYAML(fmt.Sprintf("../%s/%s.yaml", common.SkillsFilepath, "active"), &list); err != nil {
// 		logrus.WithError(err).Error("Error loading active skills")
// 	}
// }

// func loadAndParseXMLFile(filename string) (*YourXMLStruct, error) {
// 	xmlData, err := os.Open(SkillXmlFile)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer xmlData.Close()

// 	var result YourXMLStruct
// 	err = xml.Unmarshal(xmlData, &result)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &result, nil
// }

// // Define your XML structure here
// type YourXMLStruct struct {
// 	ID   string `yaml:"id"`
// 	Name string `yaml:"name"`
// }
