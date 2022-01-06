// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AsteroidBelt struct {
	Name     *string   `json:"name"`
	Position *Position `json:"position"`
	System   *System   `json:"system"`
}

type DogmaAttribute struct {
	Attribute *DogmaAttributeDetail `json:"attribute"`
	Value     *float64              `json:"value"`
}

type Alliance struct {
	CreatorCorporation  *Corporation `json:"creator_corporation"`
	Creator             *Character   `json:"creator"`
	DateFounded         *string      `json:"date_founded"`
	ExecutorCorporation *Corporation `json:"executor_corporation"`
	Faction             *Faction     `json:"faction"`
	Name                *string      `json:"name"`
	Ticker              *string      `json:"ticker"`
}

type Ancestry struct {
	Bloodline        *Bloodline `json:"bloodline"`
	Description      *string    `json:"description"`
	Icon             *Icon      `json:"icon"`
	ID               *int       `json:"id"`
	Name             *string    `json:"name"`
	ShortDescription *string    `json:"short_description"`
}

type Bloodline struct {
	ID           *int         `json:"id"`
	Charisma     *int         `json:"charisma"`
	Corporation  *Corporation `json:"corporation"`
	Description  *string      `json:"description"`
	Intelligence *int         `json:"intelligence"`
	Memory       *int         `json:"memory"`
	Name         *string      `json:"name"`
	Perception   *int         `json:"perception"`
	Race         *Race        `json:"race"`
	ShipType     *ItemType    `json:"ship_type"`
	Willpower    *int         `json:"willpower"`
}

type Category struct {
	CategoryID     *int     `json:"category_id"`
	CategoryGroups []*Group `json:"category_groups"`
	Name           *string  `json:"name"`
	Published      *bool    `json:"published"`
}

type Character struct {
	Alliance       *Alliance    `json:"alliance"`
	Ancestry       *Ancestry    `json:"ancestry"`
	Birthday       *string      `json:"birthday"`
	Bloodline      *Bloodline   `json:"bloodline"`
	Corporation    *Corporation `json:"corporation"`
	Description    *string      `json:"description"`
	Faction        *Faction     `json:"faction"`
	Gender         *Gender      `json:"gender"`
	Name           *string      `json:"name"`
	Race           *Race        `json:"race"`
	SecurityStatus *float64     `json:"security_status"`
	Title          *string      `json:"title"`
}

type Constellation struct {
	ConstellationID *int      `json:"constellation_id"`
	Name            *string   `json:"name"`
	Position        *Position `json:"position"`
	Region          *Region   `json:"region"`
	SolarSystems    []*System `json:"solar_systems"`
}

type Corporation struct {
	Alliance    *Alliance  `json:"alliance"`
	Ceo         *Character `json:"ceo"`
	Creator     *Character `json:"creator"`
	DateFounded *string    `json:"date_founded"`
	Description *string    `json:"description"`
	Faction     *Faction   `json:"faction"`
	HomeStation *Station   `json:"home_station"`
	MemberCount *int       `json:"member_count"`
	Name        *string    `json:"name"`
	Shares      *int       `json:"shares"`
	TaxRate     *float64   `json:"tax_rate"`
	Ticker      *string    `json:"ticker"`
	URL         *string    `json:"url"`
	WarEligible *bool      `json:"war_eligible"`
}

type DogmaAttributeDetail struct {
	AttributeID  *int     `json:"attribute_id"`
	DefaultValue *float64 `json:"default_value"`
	Description  *string  `json:"description"`
	DisplayName  *string  `json:"display_name"`
	HighIsGood   *bool    `json:"high_is_good"`
	Icon         *Icon    `json:"icon"`
	Name         *string  `json:"name"`
	Published    *bool    `json:"published"`
	Stackable    *bool    `json:"stackable"`
	Unit         *Unit    `json:"unit"`
}

type DogmaEffect struct {
	Effect    *DogmaEffectDetail `json:"effect"`
	IsDefault *bool              `json:"is_default"`
}

type DogmaEffectDetail struct {
	Description              *string         `json:"description"`
	DisallowAutoRepeat       *bool           `json:"disallow_auto_repeat"`
	DischargeAttribute       *DogmaAttribute `json:"discharge_attribute"`
	DisplayName              *string         `json:"display_name"`
	DurationAttribute        *DogmaAttribute `json:"duration_attribute"`
	EffectCategory           *int            `json:"effect_category"`
	EffectID                 *int            `json:"effect_id"`
	ElectronicChance         *bool           `json:"electronic_chance"`
	FalloffAttributeID       *int            `json:"falloff_attribute_id"`
	Icon                     *Icon           `json:"icon"`
	IsAssistance             *bool           `json:"is_assistance"`
	IsOffensive              *bool           `json:"is_offensive"`
	IsWarpSafe               *bool           `json:"is_warp_safe"`
	Modifiers                []*Modifier     `json:"modifiers"`
	Name                     *string         `json:"name"`
	PostExpression           *int            `json:"post_expression"`
	PreExpression            *int            `json:"pre_expression"`
	Published                *bool           `json:"published"`
	RangeAttributeID         *int            `json:"range_attribute_id"`
	RangeChange              *bool           `json:"range_change"`
	TrackingSpeedAttributeID *int            `json:"tracking_speed_attribute_id"`
}

type Faction struct {
	Corporation        *Corporation `json:"corporation"`
	Description        *string      `json:"description"`
	FactionID          *int         `json:"faction_id"`
	IsUnique           *bool        `json:"is_unique"`
	MilitiaCorporation *Corporation `json:"militia_corporation"`
	Name               *string      `json:"name"`
	SizeFactor         *float64     `json:"size_factor"`
	SolarSystem        *System      `json:"solar_system"`
	StationCount       *int         `json:"station_count"`
	StationSystemCount *int         `json:"station_system_count"`
}

type Graphic struct {
	CollisionFile  *string `json:"collision_file"`
	GraphicFile    *string `json:"graphic_file"`
	GraphicID      *int    `json:"graphic_id"`
	IconFolder     *string `json:"icon_folder"`
	SofDna         *string `json:"sof_dna"`
	SofFactionName *string `json:"sof_faction_name"`
	SofHullName    *string `json:"sof_hull_name"`
	SofRaceName    *string `json:"sof_race_name"`
}

type Group struct {
	Category  *Category   `json:"category"`
	GroupID   *int        `json:"group_id"`
	Name      *string     `json:"name"`
	Published *bool       `json:"published"`
	ItemTypes []*ItemType `json:"item_types"`
}

type Icon struct {
	ID *int `json:"id"`
}

type ItemType struct {
	TypeID          *int              `json:"type_id"`
	Capacity        *float64          `json:"capacity"`
	Description     *string           `json:"description"`
	DogmaAttributes []*DogmaAttribute `json:"dogma_attributes"`
	DogmaEffects    []*DogmaEffect    `json:"dogma_effects"`
	Graphic         *Graphic          `json:"graphic"`
	Group           *Group            `json:"group"`
	Icon            *Icon             `json:"icon"`
	MarketGroup     *MarketGroup      `json:"market_group"`
	Mass            *float64          `json:"mass"`
	Name            *string           `json:"name"`
	PackagedVolume  *float64          `json:"packaged_volume"`
	PortionSize     *int              `json:"portion_size"`
	Published       *bool             `json:"published"`
	Radius          *float64          `json:"radius"`
	Volume          *float64          `json:"volume"`
}

type MarketGroup struct {
	Description *string     `json:"description"`
	ID          *int        `json:"id"`
	Name        *string     `json:"name"`
	ParentGroup *Group      `json:"parent_group"`
	Types       []*ItemType `json:"types"`
}

type Modifier struct {
	Domain               *string `json:"domain"`
	EffectID             *int    `json:"effect_id"`
	Func                 *string `json:"func"`
	ModifiedAttributeID  *int    `json:"modified_attribute_id"`
	ModifyingAttributeID *int    `json:"modifying_attribute_id"`
	Operator             *int    `json:"operator"`
}

type Moon struct {
	MoonID   *int      `json:"moon_id"`
	Name     *string   `json:"name"`
	Position *Position `json:"position"`
	System   *System   `json:"system"`
}

type Order struct {
	Duration     *int      `json:"duration"`
	Isbuyorder   *bool     `json:"isbuyorder"`
	Issued       *string   `json:"issued"`
	Location     *Station  `json:"location"`
	Locationid   *int      `json:"locationid"`
	Minvolume    *int      `json:"minvolume"`
	Orderid      int       `json:"orderid"`
	Price        *float64  `json:"price"`
	Range        *Range    `json:"range"`
	System       *System   `json:"system"`
	Systemid     *int      `json:"systemid"`
	Itemtype     *ItemType `json:"itemtype"`
	Typeid       *int      `json:"typeid"`
	Volumeremain *int      `json:"volumeremain"`
	Volumetotal  *int      `json:"volumetotal"`
}

type Planet struct {
	Name     *string   `json:"name"`
	PlanetID *int      `json:"planet_id"`
	Position *Position `json:"position"`
	System   *System   `json:"system"`
	ItemType *ItemType `json:"item_type"`
}

type Position struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
	Z *float64 `json:"z"`
}

type Race struct {
	Alliance    *Alliance `json:"alliance"`
	Description *string   `json:"description"`
	Name        *string   `json:"name"`
	RaceID      *int      `json:"race_id"`
}

type Region struct {
	ConstellationList []*Constellation `json:"constellation_list"`
	Description       *string          `json:"description"`
	Name              *string          `json:"name"`
	RegionID          *int             `json:"region_id"`
}

type Star struct {
	Age           *int           `json:"age"`
	Luminosity    *float64       `json:"luminosity"`
	Name          *string        `json:"name"`
	Radius        *int           `json:"radius"`
	SolarSystem   *System        `json:"solar_system"`
	SpectralClass *SpectralClass `json:"spectral_class"`
	StarID        *int           `json:"star_id"`
	Temperature   *int           `json:"temperature"`
	ItemType      *ItemType      `json:"item_type"`
}

type Stargate struct {
	Destination *StargateDestination `json:"destination"`
	Name        *string              `json:"name"`
	Position    *Position            `json:"position"`
	StargateID  *int                 `json:"stargate_id"`
	System      *System              `json:"system"`
	ItemType    *ItemType            `json:"item_type"`
}

type StargateDestination struct {
	Stargate *Stargate `json:"stargate"`
	System   *System   `json:"system"`
}

type Station struct {
	MaxDockableShipVolume    *float64     `json:"max_dockable_ship_volume"`
	Name                     *string      `json:"name"`
	OfficeRentalCost         *float64     `json:"office_rental_cost"`
	OwningCorporation        *Corporation `json:"owning_corporation"`
	Position                 *Position    `json:"position"`
	Race                     *Race        `json:"race"`
	ReprocessingEfficiency   *float64     `json:"reprocessing_efficiency"`
	ReprocessingStationsTake *float64     `json:"reprocessing_stations_take"`
	Services                 []*Services  `json:"services"`
	StationID                *int         `json:"station_id"`
	System                   *System      `json:"system"`
	StationType              *ItemType    `json:"station_type"`
}

type System struct {
	Constellation *Constellation  `json:"constellation"`
	Name          *string         `json:"name"`
	Planets       []*SystemPlanet `json:"planets"`
	Position      *Position       `json:"position"`
	SecurityClass *string         `json:"security_class"`
	Star          *Star           `json:"star"`
	StargateList  []*Stargate     `json:"stargate_list"`
	StationList   []*Station      `json:"station_list"`
	StationIds    []*int          `json:"station_ids"`
	SystemID      *int            `json:"system_id"`
}

type SystemPlanet struct {
	AsteroidBelts    []*AsteroidBelt `json:"asteroid_belts"`
	MoonList         []*Moon         `json:"moon_list"`
	PlanetProperties *Planet         `json:"planet_properties"`
}

type Unit struct {
	ID *int `json:"id"`
}

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

var AllGender = []Gender{
	GenderMale,
	GenderFemale,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Ordertype string

const (
	OrdertypeBuy  Ordertype = "buy"
	OrdertypeSell Ordertype = "sell"
	OrdertypeAll  Ordertype = "all"
)

var AllOrdertype = []Ordertype{
	OrdertypeBuy,
	OrdertypeSell,
	OrdertypeAll,
}

func (e Ordertype) IsValid() bool {
	switch e {
	case OrdertypeBuy, OrdertypeSell, OrdertypeAll:
		return true
	}
	return false
}

func (e Ordertype) String() string {
	return string(e)
}

func (e *Ordertype) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Ordertype(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ordertype", str)
	}
	return nil
}

func (e Ordertype) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Range string

const (
	RangeStation     Range = "station"
	RangeRegion      Range = "region"
	RangeSolarsystem Range = "solarsystem"
	RangeRange1      Range = "range_1"
	RangeRange2      Range = "range_2"
	RangeRange3      Range = "range_3"
	RangeRange4      Range = "range_4"
	RangeRange5      Range = "range_5"
	RangeRange10     Range = "range_10"
	RangeRange20     Range = "range_20"
	RangeRange30     Range = "range_30"
	RangeRange40     Range = "range_40"
)

var AllRange = []Range{
	RangeStation,
	RangeRegion,
	RangeSolarsystem,
	RangeRange1,
	RangeRange2,
	RangeRange3,
	RangeRange4,
	RangeRange5,
	RangeRange10,
	RangeRange20,
	RangeRange30,
	RangeRange40,
}

func (e Range) IsValid() bool {
	switch e {
	case RangeStation, RangeRegion, RangeSolarsystem, RangeRange1, RangeRange2, RangeRange3, RangeRange4, RangeRange5, RangeRange10, RangeRange20, RangeRange30, RangeRange40:
		return true
	}
	return false
}

func (e Range) String() string {
	return string(e)
}

func (e *Range) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Range(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid range", str)
	}
	return nil
}

func (e Range) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Services string

const (
	ServicesBountyMissions       Services = "bounty_missions"
	ServicesAssasinationMissions Services = "assasination_missions"
	ServicesCourierMissions      Services = "courier_missions"
	ServicesInterbus             Services = "interbus"
	ServicesReprocessingPlant    Services = "reprocessing_plant"
	ServicesRefinery             Services = "refinery"
	ServicesMarket               Services = "market"
	ServicesBlackMarket          Services = "black_market"
	ServicesStockExchange        Services = "stock_exchange"
	ServicesCloning              Services = "cloning"
	ServicesSurgery              Services = "surgery"
	ServicesDnaTherapy           Services = "dna_therapy"
	ServicesRepairFacilities     Services = "repair_facilities"
	ServicesFactory              Services = "factory"
	ServicesLabratory            Services = "labratory"
	ServicesGambling             Services = "gambling"
	ServicesFitting              Services = "fitting"
	ServicesPaintshop            Services = "paintshop"
	ServicesNews                 Services = "news"
	ServicesStorage              Services = "storage"
	ServicesInsurance            Services = "insurance"
	ServicesDocking              Services = "docking"
	ServicesOfficeRental         Services = "office_rental"
	ServicesJumpCloneFacility    Services = "jump_clone_facility"
	ServicesLoyaltyPointStore    Services = "loyalty_point_store"
	ServicesNavyOffices          Services = "navy_offices"
	ServicesSecurityOffices      Services = "security_offices"
)

var AllServices = []Services{
	ServicesBountyMissions,
	ServicesAssasinationMissions,
	ServicesCourierMissions,
	ServicesInterbus,
	ServicesReprocessingPlant,
	ServicesRefinery,
	ServicesMarket,
	ServicesBlackMarket,
	ServicesStockExchange,
	ServicesCloning,
	ServicesSurgery,
	ServicesDnaTherapy,
	ServicesRepairFacilities,
	ServicesFactory,
	ServicesLabratory,
	ServicesGambling,
	ServicesFitting,
	ServicesPaintshop,
	ServicesNews,
	ServicesStorage,
	ServicesInsurance,
	ServicesDocking,
	ServicesOfficeRental,
	ServicesJumpCloneFacility,
	ServicesLoyaltyPointStore,
	ServicesNavyOffices,
	ServicesSecurityOffices,
}

func (e Services) IsValid() bool {
	switch e {
	case ServicesBountyMissions, ServicesAssasinationMissions, ServicesCourierMissions, ServicesInterbus, ServicesReprocessingPlant, ServicesRefinery, ServicesMarket, ServicesBlackMarket, ServicesStockExchange, ServicesCloning, ServicesSurgery, ServicesDnaTherapy, ServicesRepairFacilities, ServicesFactory, ServicesLabratory, ServicesGambling, ServicesFitting, ServicesPaintshop, ServicesNews, ServicesStorage, ServicesInsurance, ServicesDocking, ServicesOfficeRental, ServicesJumpCloneFacility, ServicesLoyaltyPointStore, ServicesNavyOffices, ServicesSecurityOffices:
		return true
	}
	return false
}

func (e Services) String() string {
	return string(e)
}

func (e *Services) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Services(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid services", str)
	}
	return nil
}

func (e Services) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SpectralClass string

const (
	SpectralClassK2V   SpectralClass = "K2_V"
	SpectralClassK4V   SpectralClass = "K4_V"
	SpectralClassG2V   SpectralClass = "G2_V"
	SpectralClassG8V   SpectralClass = "G8_V"
	SpectralClassM7V   SpectralClass = "M7_V"
	SpectralClassK7V   SpectralClass = "K7_V"
	SpectralClassM2V   SpectralClass = "M2_V"
	SpectralClassK5V   SpectralClass = "K5_V"
	SpectralClassM3V   SpectralClass = "M3_V"
	SpectralClassG0V   SpectralClass = "G0_V"
	SpectralClassG7V   SpectralClass = "G7_V"
	SpectralClassG3V   SpectralClass = "G3_V"
	SpectralClassF9V   SpectralClass = "F9_V"
	SpectralClassG5V   SpectralClass = "G5_V"
	SpectralClassF6V   SpectralClass = "F6_V"
	SpectralClassK8V   SpectralClass = "K8_V"
	SpectralClassK9V   SpectralClass = "K9_V"
	SpectralClassK6V   SpectralClass = "K6_V"
	SpectralClassG9V   SpectralClass = "G9_V"
	SpectralClassG6V   SpectralClass = "G6_V"
	SpectralClassG4Vi  SpectralClass = "G4_VI"
	SpectralClassG4V   SpectralClass = "G4_V"
	SpectralClassF8V   SpectralClass = "F8_V"
	SpectralClassF2V   SpectralClass = "F2_V"
	SpectralClassF1V   SpectralClass = "F1_V"
	SpectralClassK3V   SpectralClass = "K3_V"
	SpectralClassF0Vi  SpectralClass = "F0_VI"
	SpectralClassG1Vi  SpectralClass = "G1_VI"
	SpectralClassG0Vi  SpectralClass = "G0_VI"
	SpectralClassK1V   SpectralClass = "K1_V"
	SpectralClassM4V   SpectralClass = "M4_V"
	SpectralClassM1V   SpectralClass = "M1_V"
	SpectralClassM6V   SpectralClass = "M6_V"
	SpectralClassM0V   SpectralClass = "M0_V"
	SpectralClassK2Iv  SpectralClass = "K2_IV"
	SpectralClassG2Vi  SpectralClass = "G2_VI"
	SpectralClassK0V   SpectralClass = "K0_V"
	SpectralClassK5Iv  SpectralClass = "K5_IV"
	SpectralClassF5Vi  SpectralClass = "F5_VI"
	SpectralClassG6Vi  SpectralClass = "G6_VI"
	SpectralClassF6Vi  SpectralClass = "F6_VI"
	SpectralClassF2Iv  SpectralClass = "F2_IV"
	SpectralClassG3Vi  SpectralClass = "G3_VI"
	SpectralClassM8V   SpectralClass = "M8_V"
	SpectralClassF1Vi  SpectralClass = "F1_VI"
	SpectralClassK1Iv  SpectralClass = "K1_IV"
	SpectralClassF7V   SpectralClass = "F7_V"
	SpectralClassG5Vi  SpectralClass = "G5_VI"
	SpectralClassM5V   SpectralClass = "M5_V"
	SpectralClassG7Vi  SpectralClass = "G7_VI"
	SpectralClassF5V   SpectralClass = "F5_V"
	SpectralClassF4Vi  SpectralClass = "F4_VI"
	SpectralClassF8Vi  SpectralClass = "F8_VI"
	SpectralClassK3Iv  SpectralClass = "K3_IV"
	SpectralClassF4Iv  SpectralClass = "F4_IV"
	SpectralClassF0V   SpectralClass = "F0_V"
	SpectralClassG7Iv  SpectralClass = "G7_IV"
	SpectralClassG8Vi  SpectralClass = "G8_VI"
	SpectralClassF2Vi  SpectralClass = "F2_VI"
	SpectralClassF4V   SpectralClass = "F4_V"
	SpectralClassF7Vi  SpectralClass = "F7_VI"
	SpectralClassF3V   SpectralClass = "F3_V"
	SpectralClassG1V   SpectralClass = "G1_V"
	SpectralClassG9Vi  SpectralClass = "G9_VI"
	SpectralClassF3Iv  SpectralClass = "F3_IV"
	SpectralClassF9Vi  SpectralClass = "F9_VI"
	SpectralClassM9V   SpectralClass = "M9_V"
	SpectralClassK0Iv  SpectralClass = "K0_IV"
	SpectralClassF1Iv  SpectralClass = "F1_IV"
	SpectralClassG4Iv  SpectralClass = "G4_IV"
	SpectralClassF3Vi  SpectralClass = "F3_VI"
	SpectralClassK4Iv  SpectralClass = "K4_IV"
	SpectralClassG5Iv  SpectralClass = "G5_IV"
	SpectralClassG3Iv  SpectralClass = "G3_IV"
	SpectralClassG1Iv  SpectralClass = "G1_IV"
	SpectralClassK7Iv  SpectralClass = "K7_IV"
	SpectralClassG0Iv  SpectralClass = "G0_IV"
	SpectralClassK6Iv  SpectralClass = "K6_IV"
	SpectralClassK9Iv  SpectralClass = "K9_IV"
	SpectralClassG2Iv  SpectralClass = "G2_IV"
	SpectralClassF9Iv  SpectralClass = "F9_IV"
	SpectralClassF0Iv  SpectralClass = "F0_IV"
	SpectralClassK8Iv  SpectralClass = "K8_IV"
	SpectralClassG8Iv  SpectralClass = "G8_IV"
	SpectralClassF6Iv  SpectralClass = "F6_IV"
	SpectralClassF5Iv  SpectralClass = "F5_IV"
	SpectralClassA0    SpectralClass = "A0"
	SpectralClassA0iv  SpectralClass = "A0IV"
	SpectralClassA0iv2 SpectralClass = "A0IV2"
)

var AllSpectralClass = []SpectralClass{
	SpectralClassK2V,
	SpectralClassK4V,
	SpectralClassG2V,
	SpectralClassG8V,
	SpectralClassM7V,
	SpectralClassK7V,
	SpectralClassM2V,
	SpectralClassK5V,
	SpectralClassM3V,
	SpectralClassG0V,
	SpectralClassG7V,
	SpectralClassG3V,
	SpectralClassF9V,
	SpectralClassG5V,
	SpectralClassF6V,
	SpectralClassK8V,
	SpectralClassK9V,
	SpectralClassK6V,
	SpectralClassG9V,
	SpectralClassG6V,
	SpectralClassG4Vi,
	SpectralClassG4V,
	SpectralClassF8V,
	SpectralClassF2V,
	SpectralClassF1V,
	SpectralClassK3V,
	SpectralClassF0Vi,
	SpectralClassG1Vi,
	SpectralClassG0Vi,
	SpectralClassK1V,
	SpectralClassM4V,
	SpectralClassM1V,
	SpectralClassM6V,
	SpectralClassM0V,
	SpectralClassK2Iv,
	SpectralClassG2Vi,
	SpectralClassK0V,
	SpectralClassK5Iv,
	SpectralClassF5Vi,
	SpectralClassG6Vi,
	SpectralClassF6Vi,
	SpectralClassF2Iv,
	SpectralClassG3Vi,
	SpectralClassM8V,
	SpectralClassF1Vi,
	SpectralClassK1Iv,
	SpectralClassF7V,
	SpectralClassG5Vi,
	SpectralClassM5V,
	SpectralClassG7Vi,
	SpectralClassF5V,
	SpectralClassF4Vi,
	SpectralClassF8Vi,
	SpectralClassK3Iv,
	SpectralClassF4Iv,
	SpectralClassF0V,
	SpectralClassG7Iv,
	SpectralClassG8Vi,
	SpectralClassF2Vi,
	SpectralClassF4V,
	SpectralClassF7Vi,
	SpectralClassF3V,
	SpectralClassG1V,
	SpectralClassG9Vi,
	SpectralClassF3Iv,
	SpectralClassF9Vi,
	SpectralClassM9V,
	SpectralClassK0Iv,
	SpectralClassF1Iv,
	SpectralClassG4Iv,
	SpectralClassF3Vi,
	SpectralClassK4Iv,
	SpectralClassG5Iv,
	SpectralClassG3Iv,
	SpectralClassG1Iv,
	SpectralClassK7Iv,
	SpectralClassG0Iv,
	SpectralClassK6Iv,
	SpectralClassK9Iv,
	SpectralClassG2Iv,
	SpectralClassF9Iv,
	SpectralClassF0Iv,
	SpectralClassK8Iv,
	SpectralClassG8Iv,
	SpectralClassF6Iv,
	SpectralClassF5Iv,
	SpectralClassA0,
	SpectralClassA0iv,
	SpectralClassA0iv2,
}

func (e SpectralClass) IsValid() bool {
	switch e {
	case SpectralClassK2V, SpectralClassK4V, SpectralClassG2V, SpectralClassG8V, SpectralClassM7V, SpectralClassK7V, SpectralClassM2V, SpectralClassK5V, SpectralClassM3V, SpectralClassG0V, SpectralClassG7V, SpectralClassG3V, SpectralClassF9V, SpectralClassG5V, SpectralClassF6V, SpectralClassK8V, SpectralClassK9V, SpectralClassK6V, SpectralClassG9V, SpectralClassG6V, SpectralClassG4Vi, SpectralClassG4V, SpectralClassF8V, SpectralClassF2V, SpectralClassF1V, SpectralClassK3V, SpectralClassF0Vi, SpectralClassG1Vi, SpectralClassG0Vi, SpectralClassK1V, SpectralClassM4V, SpectralClassM1V, SpectralClassM6V, SpectralClassM0V, SpectralClassK2Iv, SpectralClassG2Vi, SpectralClassK0V, SpectralClassK5Iv, SpectralClassF5Vi, SpectralClassG6Vi, SpectralClassF6Vi, SpectralClassF2Iv, SpectralClassG3Vi, SpectralClassM8V, SpectralClassF1Vi, SpectralClassK1Iv, SpectralClassF7V, SpectralClassG5Vi, SpectralClassM5V, SpectralClassG7Vi, SpectralClassF5V, SpectralClassF4Vi, SpectralClassF8Vi, SpectralClassK3Iv, SpectralClassF4Iv, SpectralClassF0V, SpectralClassG7Iv, SpectralClassG8Vi, SpectralClassF2Vi, SpectralClassF4V, SpectralClassF7Vi, SpectralClassF3V, SpectralClassG1V, SpectralClassG9Vi, SpectralClassF3Iv, SpectralClassF9Vi, SpectralClassM9V, SpectralClassK0Iv, SpectralClassF1Iv, SpectralClassG4Iv, SpectralClassF3Vi, SpectralClassK4Iv, SpectralClassG5Iv, SpectralClassG3Iv, SpectralClassG1Iv, SpectralClassK7Iv, SpectralClassG0Iv, SpectralClassK6Iv, SpectralClassK9Iv, SpectralClassG2Iv, SpectralClassF9Iv, SpectralClassF0Iv, SpectralClassK8Iv, SpectralClassG8Iv, SpectralClassF6Iv, SpectralClassF5Iv, SpectralClassA0, SpectralClassA0iv, SpectralClassA0iv2:
		return true
	}
	return false
}

func (e SpectralClass) String() string {
	return string(e)
}

func (e *SpectralClass) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SpectralClass(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid spectral_class", str)
	}
	return nil
}

func (e SpectralClass) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
