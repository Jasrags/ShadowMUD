package common

import "sync"

const (
	ZonesDataPath = "data/zones"
	ZoneFilename  = ZonesDataPath + "/%s.yaml"
)

type Zones map[string]*Zone

type ZoneSpec struct {
	ID          string     `yaml:"id"`
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	RuleSource  RuleSource `yaml:"rule_source"`
}

func NewZone(spec *ZoneSpec) *Zone {
	return &Zone{
		ID:    spec.ID,
		Spec:  spec,
		Rooms: make(Rooms),
	}
}

type Zone struct {
	sync.Mutex
	ID    string    `yaml:"id"`
	Spec  *ZoneSpec `yaml:"-"`
	Rooms Rooms     `yaml:"-"`
}

var CoreZones = []ZoneSpec{
	{
		ID:          "the_void",
		Name:        "The Void",
		Description: "The Void is a place of nothingness. It is a place where the laws of physics do not apply. It is a place where time and space are meaningless. It is a place where the mind can wander and the soul can rest. It is a place where the universe is born and where it dies. It is a place where the gods dwell and where the demons lurk. It is a place where the past, present, and future are one. It is a place where the light and the dark are one. It is a place where the living and the dead are one. It is a place where the mind and the body are one. It is a place where the self and the other are one. It is a place where the dream and the reality are one. It is a place where the truth and the lie are one. It is a place where the beginning and the end are one. It is a place where the alpha and the omega are one. It is a place where the light and the dark are one. It is a place where the living and the dead are one. It is a place where the mind and the body are one. It is a place where the self and the other are one. It is a place where the dream and the reality are one. It is a place where the truth and the lie are one. It is a place where the beginning and the end are one. It is a place where the alpha and the omega are one. It is a place where the light and the dark are one. It is a place where the living and the dead are one. It is a place where the mind and the body are one. It is a place where the self and the other are one. It is a place where the dream and the reality are one. It is a place where the truth and the lie are one. It is a place where the beginning and the end are one. It is a place where the alpha and the omega are one. It is a place where the light and the dark are one. It is a place where the living and the dead are one. It is a place where the mind and the body are one. It is a place where the self and the other are one. It is a place where the dream and the reality are one. It is a place where the truth and the lie are one. It is a place where the beginning and the end are one.",
	},
	{
		ID:          "seattle",
		Name:        "Seattle",
		Description: "Seattle is a city in the UCAS, located in the Pacific Northwest. It is the largest city in the UCAS and the largest metroplex in the world. Seattle is a major hub for trade, commerce, and cult ure, and is home to a number of megacorporations, including Ares Macrotechnology, Aztechnology, and NeoNET.",
		RuleSource:  RuleSourceSR5Core,
	},
}
