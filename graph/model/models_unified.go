package model

type Alliance struct {
	CreatorCorporationID  *int         `json:"creator_corporation_id"`
	CreatorCorporation    *Corporation `json:"creator_corporation"`
	CreatorID             *int         `json:"creator_id"`
	Creator               *Character   `json:"creator"`
	DateFounded           *string      `json:"date_founded"`
	ExecutorCorporationID *int         `json:"executor_corporation_id"`
	ExecutorCorporation   *Corporation `json:"executor_corporation"`
	FactionID             *int         `json:"faction_id"`
	Faction               *Faction     `json:"faction"`
	Name                  *string      `json:"name"`
	Ticker                *string      `json:"ticker"`
}

type Ancestry struct {
	BloodlineID   int        `yaml:"bloodlineID" json:"bloodlineID"`
	Bloodline     *Bloodline `json:"bloodline"`
	Charisma      int        `yaml:"charisma" json:"charisma"`
	DescriptionID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"descriptionID" json:"descriptionID"`
	IconID       int `yaml:"iconID" json:"iconID"`
	Intelligence int `yaml:"intelligence" json:"intelligence"`
	Memory       int `yaml:"memory" json:"memory"`
	NameID       struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"nameID" json:"nameID"`
	Perception       int    `yaml:"perception" json:"perception"`
	ShortDescription string `yaml:"shortDescription" json:"shortDescription"`
	Willpower        int    `yaml:"willpower" json:"willpower"`
	ID               int    `json:"id"`
}

type AsteroidBelt struct {
	ID         int            `yaml:"id" json:"id"`
	Name       string         `yaml:"name" json:"name"`
	Position   Position       `yaml:"position" json:"position"`
	Statistics StatisticsType `yaml:"statistics" json:"statistics"`
	System     *System        `json:"system"`
	SystemID   int            `yaml:"system_id" json:"system_id"`
	TypeID     int            `yaml:"typeID" json:"type_id"`
}

type bloodline map[int]struct {
	Charisma      int          `yaml:"charisma" json:"charisma"`
	Corporation   *Corporation `json:"corporation"`
	CorporationID int          `yaml:"corporationID" json:"corporationID"`
	DescriptionID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"descriptionID" json:"descriptionID"`
	IconID       int `yaml:"iconID" json:"iconID"`
	Intelligence int `yaml:"intelligence" json:"intelligence"`
	Memory       int `yaml:"memory" json:"memory"`
	NameID       struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"nameID" json:"nameID"`
	Perception int       `yaml:"perception" json:"perception"`
	Race       *Race     `json:"race"`
	RaceID     int       `yaml:"raceID" json:"raceID"`
	ShipTypeID *int      `json:"ship_type_id"`
	ShipType   *ItemType `json:"ship_type"`
	Willpower  int       `yaml:"willpower" json:"willpower"`
	ID         int       `yaml:"id" json:"id"`
}

type StatisticsType struct {
	Density        float32 `yaml:"density" json:"density"`
	Eccentricity   float64 `yaml:"eccentricity" json:"eccentricity"`
	EscapeVelocity float64 `yaml:"escapeVelocity" json:"escape_velocity"`
	Fragmented     bool    `yaml:"fragmented" json:"fragmented"`
	Life           float64 `yaml:"life" json:"life"`
	Locked         bool    `yaml:"locked" json:"locked"`
	MassDust       float64 `yaml:"massDust" json:"mass_dust"`
	MassGas        float64 `yaml:"massGas" json:"mass_gas"`
	OrbitPeriod    float64 `yaml:"orbitPeriod" json:"orbit_period"`
	OrbitRadius    float64 `yaml:"orbitRadius" json:"orbit_radius"`
	Pressure       float64 `yaml:"pressure" json:"pressure"`
	Radius         float64 `yaml:"radius" json:"radius"`
	RotationRate   float64 `yaml:"rotationRate" json:"rotation_rate"`
	SpectralClass  string  `yaml:"spectralClass" json:"spectral_class"`
	SurfaceGravity float64 `yaml:"surfaceGravity" json:"surface_gravity"`
	Temperature    float64 `yaml:"temperature" json:"temperature"`
}

type Position struct {
	X float64 `yaml:"x" json:"x"`
	Y float64 `yaml:"y" json:"y"`
	Z float64 `yaml:"z" json:"z"`
}
