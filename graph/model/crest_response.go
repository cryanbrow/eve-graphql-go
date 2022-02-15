package model

// Names is a struct of possible types of Names that can be returned for IDs
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

// NameTuple is a representation of a name as it relates to an ID.
type NameTuple struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

// Agents is a constant for NameByID
const Agents = "agents"

// Alliances is a constant for NameByID
const Alliances = "alliances"

// Characters is a constant for NameByID
const Characters = "characters"

// Constellations is a constant for NameByID
const Constellations = "constellations"

// Corporations is a constant for NameByID
const Corporations = "corporations"

// Factions is a constant for NameByID
const Factions = "factions"

// InventoryTypes is a constant for NameByID
const InventoryTypes = "inventory_types"

// Regions is a constant for NameByID
const Regions = "regions"

// Systems is a constant for NameByID
const Systems = "systems"
