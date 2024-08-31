package license

import "github.com/Jasrags/ShadowMUD/common/shared"

const (
	TypeAdeptLicense               Type = "Adept License"
	TypeAutomaticWeaponsLicense    Type = "Automatic Weapons License"
	TypeBluntWeaponsLicense        Type = "Blunt Weapons License"
	TypeBodyguardLicense           Type = "Bodyguard License"
	TypeBountyHuntersLicense       Type = "Bounty Hunter's License"
	TypeConcealedCarryPermit       Type = "Concealed Carry Permit"
	TypeCyberdeckLicense           Type = "Cyberdeck License"
	TypeDriversLicense             Type = "Driver's License"
	TypeDroneLicense               Type = "Drone License"
	TypeExoticWeaponsLicense       Type = "Exotic Weapons License"
	TypeExplosivesLicense          Type = "Explosives License"
	TypeFirearmsLicense            Type = "Firearms License"
	TypeHeavyWeaponsLicense        Type = "Heavy Weapons License"
	TypeHuntingLicense             Type = "Hunting License"
	TypeLargeBladesLicense         Type = "Large Blades License"
	TypeMageLicense                Type = "Mage License"
	TypeMarineLicense              Type = "Marine License"
	TypeMatrixSoftwareLicense      Type = "Matrix Software License"
	TypeMedicalLicense             Type = "Medical License"
	TypeMilitaryAmmunitionLicense  Type = "Military Ammunition License"
	TypeMilitaryArmorLicense       Type = "Military Armor License"
	TypeMilitaryWeaponsLicense     Type = "Military Weapons License"
	TypePetLicense                 Type = "Pet License"
	TypePilotLicense               Type = "Pilot License"
	TypePistolLicense              Type = "Pistol License"
	TypePrivateInvestigatorLicense Type = "Private Investigator License"
	TypeProjectileLicense          Type = "Projectile License"
	TypeRestrictedArmorLicense     Type = "Restricted Armor License"
	TypeRestrictedBiowareLicense   Type = "Restricted Bioware License"
	TypeRestrictedCyberwareLicense Type = "Restricted Cyberware License"
	TypeRifleLicense               Type = "Rifle License"
	TypeShotgunLicense             Type = "Shotgun License"
	TypeSmallBladesLicense         Type = "Small Blades License"
	TypeSummonerLicense            Type = "Summoner License"
	TypeTalismongerLicense         Type = "Talismonger License"
	TypeWeaponLicense              Type = "Weapon License"
)

type (
	Type string
	Spec struct {
		ID          string            `yaml:"id"`
		Name        string            `yaml:"name"`
		Description string            `yaml:"description"`
		Type        Type              `yaml:"type"`
		Cost        int               `yaml:"cost"`
		RuleSource  shared.RuleSource `yaml:"rule_source"`
	}
)
