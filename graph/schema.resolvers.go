package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	dao "github.com/cryanbrow/eve-graphql-go/graph/data_access"
	"github.com/cryanbrow/eve-graphql-go/graph/generated"
	"github.com/cryanbrow/eve-graphql-go/graph/model"
)

func (r *asteroid_beltResolver) System(ctx context.Context, obj *model.AsteroidBelt) (*model.System, error) {
	return dao.SystemByID(obj.SystemID)
}

func (r *corporationResolver) Alliance(ctx context.Context, obj *model.Corporation) (*model.Alliance, error) {
	return dao.AllianceByID(obj.AllianceID)
}

func (r *corporationResolver) Ceo(ctx context.Context, obj *model.Corporation) (*model.Character, error) {
	return dao.CharacterByID(obj.CeoID)
}

func (r *corporationResolver) Creator(ctx context.Context, obj *model.Corporation) (*model.Character, error) {
	return dao.CharacterByID(obj.CreatorID)
}

func (r *corporationResolver) Faction(ctx context.Context, obj *model.Corporation) (*model.Faction, error) {
	return dao.FactionByID(obj.FactionID)
}

func (r *corporationResolver) HomeStation(ctx context.Context, obj *model.Corporation) (*model.Station, error) {
	return dao.StationByID(obj.HomeStationID)
}

func (r *dogma_attributeResolver) Attribute(ctx context.Context, obj *model.DogmaAttribute) (*model.DogmaAttributeDetail, error) {
	return dao.DogmaAttributeByID(obj.AttributeID)
}

func (r *dogma_effectResolver) Effect(ctx context.Context, obj *model.DogmaEffect) (*model.DogmaEffectDetail, error) {
	return dao.DogmaEffectByID(obj.EffectID)
}

func (r *dogma_effect_detailResolver) DischargeAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	return dao.DogmaAttributeByID(obj.DischargeAttributeID)
}

func (r *dogma_effect_detailResolver) DurationAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	return dao.DogmaAttributeByID(obj.DurationAttributeID)
}

func (r *dogma_effect_detailResolver) FalloffAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	return dao.DogmaAttributeByID(obj.FalloffAttributeID)
}

func (r *dogma_effect_detailResolver) RangeAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	return dao.DogmaAttributeByID(obj.RangeAttributeID)
}

func (r *dogma_effect_detailResolver) TrackingSpeedAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	return dao.DogmaAttributeByID(obj.TrackingSpeedAttributeID)
}

func (r *factionResolver) Corporation(ctx context.Context, obj *model.Faction) (*model.Corporation, error) {
	return dao.CorporationByID(obj.CorporationID)
}

func (r *factionResolver) MilitiaCorporation(ctx context.Context, obj *model.Faction) (*model.Corporation, error) {
	return dao.CorporationByID(obj.MilitiaCorporationID)
}

func (r *factionResolver) SolarSystem(ctx context.Context, obj *model.Faction) (*model.System, error) {
	return dao.SystemByID(obj.SolarSystemID)
}

func (r *groupResolver) Category(ctx context.Context, obj *model.Group) (*model.Category, error) {
	return dao.CategoryByID(obj.CategoryID)
}

func (r *groupResolver) ItemTypes(ctx context.Context, obj *model.Group) ([]*model.ItemType, error) {
	return dao.ItemTypesByIDs(obj.Types)
}

func (r *item_typeResolver) Graphic(ctx context.Context, obj *model.ItemType) (*model.Graphic, error) {
	return dao.GraphicByID(obj.GraphicID)
}

func (r *item_typeResolver) Group(ctx context.Context, obj *model.ItemType) (*model.Group, error) {
	return dao.GroupByID(obj.GroupID)
}

func (r *item_typeResolver) MarketGroup(ctx context.Context, obj *model.ItemType) (*model.MarketGroup, error) {
	return dao.MarketGroupByID(obj.MarketGroupID)
}

func (r *market_groupResolver) ParentGroup(ctx context.Context, obj *model.MarketGroup) (*model.Group, error) {
	return dao.GroupByID(obj.ParentGroupID)
}

func (r *market_groupResolver) TypesDetails(ctx context.Context, obj *model.MarketGroup) ([]*model.ItemType, error) {
	return dao.ItemTypesByIDs(obj.Types)
}

func (r *modifierResolver) ModifiedAttribute(ctx context.Context, obj *model.Modifier) (*model.DogmaAttributeDetail, error) {
	return dao.DogmaAttributeByID(obj.ModifiedAttributeID)
}

func (r *modifierResolver) ModifyingAttribute(ctx context.Context, obj *model.Modifier) (*model.DogmaAttributeDetail, error) {
	return dao.DogmaAttributeByID(obj.ModifyingAttributeID)
}

func (r *orderResolver) Location(ctx context.Context, obj *model.Order) (*model.Station, error) {
	return dao.StationByID(obj.LocationID)
}

func (r *orderResolver) System(ctx context.Context, obj *model.Order) (*model.System, error) {
	return dao.SystemByID(obj.SystemID)
}

func (r *orderResolver) ItemType(ctx context.Context, obj *model.Order) (*model.ItemType, error) {
	return dao.ItemTypeByID(obj.TypeID)
}

func (r *planetResolver) System(ctx context.Context, obj *model.Planet) (*model.System, error) {
	return dao.SystemByID(obj.SystemID)
}

func (r *planetResolver) ItemType(ctx context.Context, obj *model.Planet) (*model.ItemType, error) {
	return dao.ItemTypeByID(obj.TypeID)
}

func (r *queryResolver) OrdersForRegion(ctx context.Context, regionID *int, orderType *model.Ordertype, typeID *int) ([]*model.Order, error) {
	return dao.OrdersForRegion(regionID, orderType, typeID)
}

func (r *queryResolver) SystemByID(ctx context.Context, id *int) (*model.System, error) {
	return dao.SystemByID(id)
}

func (r *queryResolver) StationByID(ctx context.Context, id *int) (*model.Station, error) {
	return dao.StationByID(id)
}

func (r *queryResolver) PlanetByID(ctx context.Context, id *int) (*model.Planet, error) {
	return dao.PlanetByID(id)
}

func (r *queryResolver) CorporationByID(ctx context.Context, id *int) (*model.Corporation, error) {
	return dao.CorporationByID(id)
}

func (r *queryResolver) FactionByID(ctx context.Context, id *int) (*model.Faction, error) {
	return dao.FactionByID(id)
}

func (r *system_planetResolver) AsteroidBeltsProperties(ctx context.Context, obj *model.SystemPlanet) ([]*model.AsteroidBelt, error) {
	return dao.AsteroidBeltDetails(obj.AsteroidBelts)
}

func (r *system_planetResolver) MoonDetails(ctx context.Context, obj *model.SystemPlanet) ([]*model.Moon, error) {
	return dao.MoonDetails(obj.Moons)
}

func (r *system_planetResolver) PlanetProperties(ctx context.Context, obj *model.SystemPlanet) (*model.Planet, error) {
	return dao.PlanetByID(obj.PlanetID)
}

// Asteroid_belt returns generated.Asteroid_beltResolver implementation.
func (r *Resolver) Asteroid_belt() generated.Asteroid_beltResolver { return &asteroid_beltResolver{r} }

// Corporation returns generated.CorporationResolver implementation.
func (r *Resolver) Corporation() generated.CorporationResolver { return &corporationResolver{r} }

// Dogma_attribute returns generated.Dogma_attributeResolver implementation.
func (r *Resolver) Dogma_attribute() generated.Dogma_attributeResolver {
	return &dogma_attributeResolver{r}
}

// Dogma_effect returns generated.Dogma_effectResolver implementation.
func (r *Resolver) Dogma_effect() generated.Dogma_effectResolver { return &dogma_effectResolver{r} }

// Dogma_effect_detail returns generated.Dogma_effect_detailResolver implementation.
func (r *Resolver) Dogma_effect_detail() generated.Dogma_effect_detailResolver {
	return &dogma_effect_detailResolver{r}
}

// Faction returns generated.FactionResolver implementation.
func (r *Resolver) Faction() generated.FactionResolver { return &factionResolver{r} }

// Group returns generated.GroupResolver implementation.
func (r *Resolver) Group() generated.GroupResolver { return &groupResolver{r} }

// Item_type returns generated.Item_typeResolver implementation.
func (r *Resolver) Item_type() generated.Item_typeResolver { return &item_typeResolver{r} }

// Market_group returns generated.Market_groupResolver implementation.
func (r *Resolver) Market_group() generated.Market_groupResolver { return &market_groupResolver{r} }

// Modifier returns generated.ModifierResolver implementation.
func (r *Resolver) Modifier() generated.ModifierResolver { return &modifierResolver{r} }

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

// Planet returns generated.PlanetResolver implementation.
func (r *Resolver) Planet() generated.PlanetResolver { return &planetResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// System_planet returns generated.System_planetResolver implementation.
func (r *Resolver) System_planet() generated.System_planetResolver { return &system_planetResolver{r} }

type asteroid_beltResolver struct{ *Resolver }
type corporationResolver struct{ *Resolver }
type dogma_attributeResolver struct{ *Resolver }
type dogma_effectResolver struct{ *Resolver }
type dogma_effect_detailResolver struct{ *Resolver }
type factionResolver struct{ *Resolver }
type groupResolver struct{ *Resolver }
type item_typeResolver struct{ *Resolver }
type market_groupResolver struct{ *Resolver }
type modifierResolver struct{ *Resolver }
type orderResolver struct{ *Resolver }
type planetResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type system_planetResolver struct{ *Resolver }
