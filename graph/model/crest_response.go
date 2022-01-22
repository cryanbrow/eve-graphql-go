package model

type Names struct {
	Agents         []*NameTuple `json:"agents"`
	Alliances      []*NameTuple `json:"alliances"`
	Characters     []*NameTuple `json:"characters"`
	Constellations []*NameTuple `json:"constellations"`
	Corporations   []*NameTuple `json:"corporations"`
	Factions       []*NameTuple `json:"factions"`
	InventoryTypes []*NameTuple `json:"inventory_types"`
	Regions        []*NameTuple `json:"regions"`
	Systems        []*NameTuple `json:"systems"`
}

type NameTuple struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

const AGENTS = "agents"
const ALLIANCES = "alliances"
const CHARACTERS = "characters"
const CONSTELLATIONS = "constellations"
const CORPORATIONS = "corporations"
const FACTIONS = "factions"
const INVENTORY_TYPES = "inventory_types"
const REGIONS = "regions"
const SYSTEMS = "systems"
