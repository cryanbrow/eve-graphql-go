package model

/*6type Alliance struct {
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

type Bloodline map[int]struct {
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

type CategoryID map[int]struct {
	CategoryGroups []*Group `yaml:"category_groups" json:"category_groups"`
	Name           struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"name" json:"name"`
	Published bool `yaml:"published" json:"published"`
	ID        int  `yaml:"id" json:"id"`
}

type Character struct {
	AllianceID     *int         `yaml:"alliance_id" json:"alliance_id"`
	Alliance       *Alliance    `yaml:"alliance" json:"alliance"`
	AncestryID     *int         `yaml:"ancestry_id" json:"ancestry_id"`
	Ancestry       *Ancestry    `yaml:"ancestry" json:"ancestry"`
	Birthday       *string      `yaml:"birthday" json:"birthday"`
	BloodlineID    *int         `yaml:"bloodline_id" json:"bloodline_id"`
	Bloodline      *Bloodline   `yaml:"bloodline" json:"bloodline"`
	CorporationID  *int         `yaml:"corporation_id" json:"corporation_id"`
	Corporation    *Corporation `yaml:"corporation" json:"corporation"`
	Description    *string      `yaml:"description" json:"description"`
	FactionID      *int         `yaml:"faction_id" json:"faction_id"`
	Faction        *Faction     `yaml:"faction" json:"faction"`
	Gender         *Gender      `yaml:"gender" json:"gender"`
	Name           *string      `yaml:"name" json:"name"`
	RaceID         *int         `yaml:"race_id" json:"race_id"`
	Race           *Race        `yaml:"race" json:"race"`
	SecurityStatus *float64     `yaml:"security_status" json:"security_status"`
	Title          *string      `yaml:"title" json:"title"`
}

type CharacterPortrait struct {
	Px128x128 *string `yaml:"px128x128" json:"px128x128"`
	Px256x256 *string `yaml:"px256x256" json:"px256x256"`
	Px512x512 *string `yaml:"px512x512" json:"px512x512"`
	Px64x64   *string `yaml:"px64x64" json:"px64x64"`
}

type Constellation struct {
	ConstellationID int       `yaml:"constellationID" json:"constellation_id"`
	Systems         []*int    `yaml:"systems" json:"systems"`
	SolarSystems    []*System `yaml:"solar_systems" json:"solar_systems"`
	Max             Position  `yaml:"max" json:"max"`
	Min             Position  `yaml:"min" json:"min"`
	Name            string    `yaml:"name" json:"name"`
	Region          *Region   `yaml:"region" json:"region"`
	RegionID        int       `yaml:"region_id" json:"region_id"`
	Position        Position  `yaml:"position" json:"position"`
}

type Corporation struct {
	Alliance      *Alliance  `yaml:"alliance" json:"alliance"`
	AllianceID    *int       `yaml:"alliance_id" json:"alliance_id"`
	Ceo           *Character `yaml:"ceo" json:"ceo"`
	CeoID         *int       `yaml:"ceo_id" json:"ceo_id"`
	Creator       *Character `yaml:"creator" json:"creator"`
	CreatorID     *int       `yaml:"creator_id" json:"creator_id"`
	DateFounded   *string    `yaml:"date_founded" json:"date_founded"`
	Description   *string    `yaml:"description" json:"description"`
	Faction       *Faction   `yaml:"faction" json:"faction"`
	FactionID     *int       `yaml:"faction_id" json:"faction_id"`
	HomeStation   *Station   `yaml:"home_station" json:"home_station"`
	HomeStationID *int       `yaml:"home_station_id" json:"home_station_id"`
	MemberCount   *int       `yaml:"member_count" json:"member_count"`
	Name          *string    `yaml:"name" json:"name"`
	Shares        *int       `yaml:"shares" json:"shares"`
	TaxRate       *float64   `yaml:"tax_rate" json:"tax_rate"`
	Ticker        *string    `yaml:"ticker" json:"ticker"`
	URL           *string    `yaml:"url" json:"url"`
	WarEligible   *bool      `yaml:"war_eligible" json:"war_eligible"`
}

type npcCorporation map[int]struct {
	AllowedMemberRaces []int           `yaml:"allowedMemberRaces" json:"allowedMemberRaces"`
	CeoID              int             `yaml:"ceoID" json:"ceoID"`
	CorporationTrades  map[int]float32 `yaml:"corporationTrades" json:"corporationTrades"`
	Deleted            bool            `yaml:"deleted" json:"deleted"`
	DescriptionID      struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"descriptionID" json:"descriptionID"`
	Divisions                 map[int]division `yaml:"divisions" json:"divisions"`
	EnemyID                   int              `yaml:"enemyID" json:"enemyID"`
	Extent                    string           `yaml:"extent" json:"extent"`
	FactionID                 int              `yaml:"factionID" json:"factionID"`
	FriendID                  int              `yaml:"friendID" json:"friendID"`
	HasPlayerPersonnelManager bool             `yaml:"hasPlayerPersonnelManager" json:"hasPlayerPersonnelManager"`
	IconID                    int              `yaml:"iconID" json:"iconID"`
	InitialPrice              int              `yaml:"initialPrice" json:"initialPrice"`
	Investors                 map[int]int      `yaml:"investors" json:"investors"`
	LPOfferTables             []int            `yaml:"lpOfferTables" json:"lpOfferTables"`
	MainActivityID            int              `yaml:"mainActivityID" json:"mainActivityID"`
	MemberLimit               int              `yaml:"memberLimit" json:"memberLimit"`
	MinSecurity               float32          `yaml:"minSecurity" json:"minSecurity"`
	MinimumJoinStanding       int              `yaml:"minimumJoinStanding" json:"minimumJoinStanding"`
	NameID                    struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"nameID" json:"nameID"`
	PublicShares               int     `yaml:"publicShares" json:"publicShares"`
	RaceID                     int     `yaml:"raceID" json:"raceID"`
	SendCharTerminationMessage bool    `yaml:"sendCharTerminationMessage" json:"sendCharTerminationMessage"`
	Shares                     int     `yaml:"shares" json:"shares"`
	Size                       string  `yaml:"size" json:"size"`
	SizeFactor                 float32 `yaml:"sizeFactor" json:"sizeFactor"`
	SolarSystemID              int     `yaml:"solarSystemID" json:"solarSystemID"`
	StationID                  int     `yaml:"stationID" json:"stationID"`
	TaxRate                    float32 `yaml:"taxRate" json:"taxRate"`
	TickerName                 string  `yaml:"tickerName" json:"tickerName"`
	UniqueName                 bool    `yaml:"uniqueName" json:"uniqueName"`
	ID                         int     `json:"id"`
}

type division struct {
	DivisionNumber int `yaml:"divisionNumber" json:"divisionNumber"`
	LeaderID       int `yaml:"leaderID" json:"leaderID"`
	Size           int `yaml:"size" json:"size"`
}

// A single corporation that a player has been part of.
type CorporationHistory struct {
	// unique id of the corporation
	CorporationID *int `yaml:"corporation_id" json:"corporation_id"`
	// unique id of the employment of the character
	RecordID *int `yaml:"record_id" json:"record_id"`
	// date the player started in RFC1123
	StartDate *string `yaml:"start_date" json:"start_date"`
	// corporation the player was employed by
	Employer *Corporation `yaml:"employer" json:"employer"`
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
*/
