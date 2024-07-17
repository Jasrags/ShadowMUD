package core

const (
	WeaponRangedDataPath    = "data/items/weapons/ranged"
	WeaponRangedFilename    = WeaponRangedDataPath + "/%s.yaml"
	WeaponRangedeMinVersion = "0.0.1"
)

type WeaponRanged struct {
	ID               string             `yaml:"id,omitempty"`
	Name             string             `yaml:"name,omitempty"`
	Description      string             `yaml:"description,omitempty"`
	Accuracy         int                `yaml:"accuracy,omitempty"`
	DamageValue      int                `yaml:"damage_value,omitempty"`
	DamageType       DamageType         `yaml:"damage_type,omitempty"`
	ArmorPenatration int                `yaml:"armor_penatration,omitempty"`
	Modes            []WeaponFiringMode `yaml:"modes,omitempty"`
	Recoil           int                `yaml:"recoil,omitempty"`
	AmmoType         string             `yaml:"ammo_type,omitempty"`
	AmmoCapacity     int                `yaml:"ammo_capacity,omitempty"`
	Availability     string             `yaml:"availability,omitempty"`
	ItemTags         []ItemTag          `yaml:"tags"`
	Modifiers        []Modifier         `yaml:"modifiers"`
	Cost             int                `yaml:"cost,omitempty"`
	RuleSource       string             `yaml:"rule_source,omitempty"`
}

var CoreWeaponRanged = []WeaponRanged{}

// Throwing Weapons
// ================
// | Weapon | Acc | DV | AP | Modes | RC | Ammo | Avail | Cost | Source |
// |--------|-----|----|----|-------|----|------|-------|------|--------|
// | Shuriken | Physical | (STR+1)P | -1 | – | – | 4R | 25¥ | Core |
// | Throwing Knife | Physical | (STR+1)P | -1 | – | – | 4R | 25¥ | Core |
// Ballistic Projectiles
// =====================
// | Weapon | Acc | DV | AP | Modes | RC | Ammo | Avail | Cost | Source |
// |--------|-----|----|----|-------|----|------|-------|------|--------|
// | Bow | 6 | (Rating+2)P | -(Rating/4) | – | – | – | Rating | Rating×100¥ | Core |
// | Light Crossbow | 7 | 5P | -1 | – | – | 4 (m) | 2 | 300¥ | Core |
// | Medium Crossbow | 6 | 7P | -2 | – | – | 4 (m) | 4R | 500¥ | Core |
// | Heavy Crossbow | 5 | 10P | -3 | – | – | 4 (m) | 8R | 1,000¥ | Core |
// Exotic Ranged Weapons
// =====================
// | Weapon | Acc | DV | AP | Modes | RC | Ammo | Avail | Cost | Source |
// |--------|-----|----|----|-------|----|------|-------|------|--------|
// | Grapple gun | 3 | 7S | -2 | SS | – | 1 (ml) | 8R | 500¥ | Core |
// Tasers
// ======
// | Weapon | Acc | DV  | AP | Modes | RC | Ammo | Avail | Cost | Source    |
// |--------|-----|-----|----|-------|----|------|-------|------|-----------|
// | Defiance EX Shocker | 4 | 9S(e) | -5 | SS | – | 4 (m) | – | 250¥ | Core |
// | Yamaha Pulsar | 5 | 7S(e) | -5 | SA | – | 4 (m) | – | 180¥ | Core |
// Pistols
// -------
// Hold-Out Pistols
// ----------------
// | Weapon | Acc | DV  | AP | Modes | RC | Ammo | Avail | Cost | Source    |
// |--------|-----|-----|----|-------|----|------|-------|------|-----------|
// | Fichetti Tiffani Needler | 5 | 8P(f) | +5 | SA | – | 4 (c) | 6R | 1,000¥ | Core |
// | Streetline Special | 4 | 6P | – | SA | – | 6 (c) | 4R | 120¥ | Core |
// | Walther Palm Pistol | 4 | 7P | – | SS/BF | – | 2 (b) | 4R | 180¥ | Core |
// Light Pistols
// -------------
// | Weapon | Acc | DV | AP | Modes | RC | Ammo | Avail | Cost | Source |
// |--------|-----|----|----|-------|----|------|-------|------|--------|
// | Ares Light Fire 70 | 7 | 8P | – | SA | – | 16 (c) | 3R | 200¥ | Core |
// | Ares Light Fire 75 | 6 (8) | 6P | – | SA | – | 16 (c) | 6F | 1,250¥ | Core |
// | Beretta 201T | 6 | 6P | – | SA/BF | – (1) | 21 (c) | 7R | 210¥ | Core |
// | Colt America L36 | 7 | 7P | – | SA | – | 11 (c) | 4R | 320¥ | Core |
// | Fichetti Security 600 | 6 (7) | 7P | – | SA | – (1) | 30 (c) | 6R | 350¥ | Core |
// | Taurus Omni-6, light pistol rounds | 5 (6) | 6P | – | SA | – | 6 (cy) | 3R | 300¥ | Core |
// | Taurus Omni-6, heavy pistol rounds | 5 (6) | 7P | -1 | SS | – | 6 (cy) | 3R | 300¥ | Core |
// Heavy Pistols
// -------------
// | Weapon | Acc | DV  | AP | Modes | RC | Ammo | Avail | Cost | Source    |
// |--------|-----|-----|----|-------|----|------|-------|------|-----------|
// | Ares Predator V | 5 (7) | 8P  | -1 | SA | – | 15 (c) | 5R | 725¥ | Core |
// | Ares Viper Silvergun | 4 | 9P(f) | +4 | SA/BF | – | 30 (c) | 8F | 380¥ | Core |
// | Browning Ultra-Power | 5 (6) | 8P | -1 | SA | – | 10 (c) | 4R | 640¥ | Core |
// | Colt Government 2066 | 6 | 7P | -1 | SA | – | 14 (c) | 7R | 425¥ | Core |
// | Remington Roomsweeper | 4 | 7P | -1 | SA | – | 8 (m) | 6R | 250¥ | Core |
// | Ruger Super Warhawk | 5 | 9P | -2 | SS | – | 6 (cy) | 4R | 400¥ | Core |
// Machine Pistols
// ---------------
// | Weapon | Acc | DV | AP | Modes | RC | Ammo | Avail | Cost | Source |
// |--------|-----|----|----|-------|----|------|-------|------|--------|
// | Ares Crusader II | 5 (7) | 7P | – | SA/BF | 2 | 40 (c) | 9R | 830¥ | Core |
// | Ceska Black Scorpion | 5 | 6P | – | SA/BF | – (1) | 35 (c) | 6R | 270¥ | Core |
// | Steyr TMP | 4 | 7P | – | SA/BF/FA | – | 30 (c) | 8R | 350¥ | Core |
// Submachine Guns
// ===============
// | Weapon | Acc | DV | AP | Modes | RC | Ammo | Avail | Cost | Source |
// |--------|-----|----|----|-------|----|------|-------|------|--------|
// | Colt Cobra TZ-120 | 4 (5) | 7P | – | SA/BF/FA | 2 (3) | 32 (c) | 5R | 660¥ | Core |
// | FN P93 Praetor | 6 | 8P | – | SA/BF/FA | 1 (2) | 50 (c) | 11F | 900¥ | Core |
// | HK-227 | 5 (7) | 7P | – | SA/BF/FA | – (1) | 28 (c) | 8R | 730¥ | Core |
// | Ingram Smartgun X | 4 (6) | 8P | – | BF/FA | 2 | 32 (c) | 6R | 800¥ | Core |
// | SCK Model 100 | 5 (7) | 8P | – | SA/BF | – (1) | 30 (c) | 6R | 875¥ | Core |
// | Uzi IV | 4 (5) | 7P | – | BF | – (1) | 24 (c) | 4R | 450¥ | Core |
// Assault Rifles
// --------------
// | Weapon | Acc | DV  | AP | Modes | RC | Ammo | Avail | Cost | Source    |
// |--------|-----|-----|----|-------|----|------|-------|------|-----------|
// | AK-97  | 4   | 10P | -2 | SA/BF/FA | –  | 38 (c) | 4R    | 950¥ | Core      |
// | Ares Alpha  | 5 (7) | 11P | -2 | SA/BF/FA | 2  | 42 (c) | 11F   | 2,650¥ | Core |
// | Colt M23  | 4   | 9P  | -2 | SA/BF/FA | –  | 40 (c) | 4R    | 550¥ | Core      |
// | FN HAR  | 5 (6) | 10P | -2 | SA/BF/FA | 2  | 35 (c) | 8R    | 1,500¥ | Core |
// | Yamaha Raiden  | 6 (8) | 11P | -2 | BF/FA | 2  | 60 (c) | 14F   | 2,600¥ | Core |
// Sniper Rifles
// -------------
// Weapon | Acc | DV | AP | Modes | RC | Ammo | Avail | Cost | Source
// -------|-----|----|----|-------|----|------|-------|------|--------
// Ares Desert Strike | 7 | 13P | -4 | SA | – (1) | 14 (c) | 10F | 17,500¥ | Core
// Cavalier Arms Crockett EBR | 6 | 12P | -3 | SA/BF | – (1) | 20 (c) | 12F | 10,300¥ | Core
// Ranger Arms SM-5 | 8 | 14P | -5 | SA | – (1) | 15 (c) | 16F | 28,000¥ | Core
// Remington 950 | 7 | 12P | -4 | SS | – | 5 (m) | 4R | 2,100¥ | Core
// Ruger 100 | 6 | 11P | -3 | SS | – (1) | 8 (m) | 4R | 1,300¥ | Core
// Shotguns
// ========
// | Weapon | Acc | DV  | AP | Modes    | RC | Ammo | Avail | Cost  | Source    |
// |--------|-----|-----|----|----------|----|------|-------|-------|-----------|
// | Defiance T-250 | 4   | 10P | -1 | SS/SA    | –  | 5 (m)            | 4R  | 450¥   | Core     |
// | Defiance T-250, short-barreled | 4 | 9P | -1 | SS/SA    | –  | 5 (m)            | 4R  | 450¥   | Core     |
// | Enfield AS-7 | 4 (5) | 13P | -1 | SA/BF    | –  | 10 (c) or 24 (d) | 12F | 1,100¥ | Core     |
// | PJSS Model 55 | 6   | 11P | -1 | SS       | – (1) | 2 (b)            | 9R  | 1,000¥ | Core     |
// Light Machine Guns
// ------------------
// | Weapon | Acc | DV  | AP | Modes | RC | Ammo | Avail | Cost  | Source    |
// |--------|-----|-----|----|-------|----|------|-------|-------|-----------|
// | Ingram Valiant         | 5 (6) | 9P  | -2 | BF/FA | 2 (3) | 50 (c) or 100 (belt) | 12F   | 5,800¥ | Core    |
// Medium Machine Guns
// -------------------
// | Weapon              | Acc | DV  | AP | Modes | RC    | Ammo              | Avail | Cost   | Source    |
// |---------------------|-----|-----|----|-------|-------|-------------------|-------|--------|-----------|
// | Stoner-Ares M202    | 5   | 10P | -3 | FA    | –     | 50 (c) or 100 (belt) | 12F   | 7,000¥ | Core      |
// Heavy Machine Guns
// ------------------
// | Weapon         | Acc | DV  | AP | Modes | RC   | Ammo              | Avail | Cost    | Source    |
// |----------------|-----|-----|----|-------|------|-------------------|-------|---------|-----------|
// | RPK HMG        | 5   | 12P | -4 | FA    | – (6) | 50 (c) or 100 (belt) | 16F   | 16,300¥ | Core      |
// Exotic Firearms
// ===============
// | Weapon                        | Acc | DV   | AP  | Modes | RC | Ammo | Avail | Cost    | Source    |
// |-------------------------------|-----|------|-----|-------|----|------|-------|---------|-----------|
// | Ares S-III Super Squirt       | 3   | Chem | SA  | –     | –  | 20   | 7R    | 950¥    | Core      |
// | Fichetti Pain Inducer         | 3   | Spec | –   | SS    | –  | Spec | 11R   | 5,000¥  | Core      |
// | Parashield Dart Pistol        | 5   | Drug | SA  | –     | –  | 5    | 4R    | 600¥    | Core      |
// | Parashield Dart Rifle         | 6   | Drug | SA  | –     | –  | 6    | 6R    | 1,200¥  | Core      |
// Assault Cannons
// ----------------
// | Weapon                           | Acc | DV  | AP | Modes | RC   | Ammo              | Avail | Cost    | Source    |
// |----------------------------------|-----|-----|----|-------|------|-------------------|-------|---------|-----------|
// | Krime Cannon                     | 4   | 16P | -6 | SA    | – (1) | 6 (m)             | 20F   | 21,000¥ | Core      |
// | Panther XXL                      | 5 (7) | 17P | -6 | SA    | –    | 15 (c)            | 20F   | 43,000¥ | Core      |
// Grenade Launchers
// -----------------
// | Weapon                      | Acc | DV      | AP | Modes | RC | Ammo | Avail | Cost    | Source    |
// |-----------------------------|-----|---------|----|-------|----|------|-------|---------|-----------|
// | Ares Alpha, Grenade Launcher| 4 (6) | Grenade | SS | –     | 6  | –    | –     | Core    |           |
// | Ares Antioch-2              | 4 (6) | Grenade | SS | –     | 8  | 8F   | 3,200¥ | Core    |           |
// | ArmTech MGL-12              | 4   | Grenade | SA | –     | 12 | 10F  | 5,000¥ | Core    |           |
// Missile Launchers
// -----------------
// | Weapon                      | Acc | DV      | AP | Modes | RC   | Ammo | Avail | Cost    | Source    |
// |-----------------------------|-----|---------|----|-------|------|------|-------|---------|-----------|
// | Aztechnology Striker        | 5   | Missile | SS | –     | 2 (ml) | 10F  | 1,200¥ | Core    |
// | Onotari Interceptor         | 4 (6) | Missile | SS | –     | 2 (ml) | 18F  | 14,000¥ | Core    |
