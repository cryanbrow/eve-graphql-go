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