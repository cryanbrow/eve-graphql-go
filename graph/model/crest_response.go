package model

//Names is a struct of possible types of Names that can be returned for IDs
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

//NameTuple is a representation of a name as it relates to an ID.
type NameTuple struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

const Agents = "agents"
const Alliances = "alliances"
const Characters = "characters"
const Constellations = "constellations"
const Corporations = "corporations"
const Factions = "factions"
const InventoryTypes = "inventory_types"
const Regions = "regions"
const Systems = "systems"
