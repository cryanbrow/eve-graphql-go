package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/alliance"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/character"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/corporation"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/dogma"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/market"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/universe"
	"github.com/cryanbrow/eve-graphql-go/graph/generated"
	"github.com/cryanbrow/eve-graphql-go/graph/generated/model"
)

func (r *allianceResolver) CreatorCorporation(ctx context.Context, obj *model.Alliance) (*model.Corporation, error) {
	newCtx, span := traceProvider.Tracer(tracer_name).Start(ctx, "CreatorCorporation")
	defer span.End()
	return corporation.CorporationByID(obj.CreatorCorporationID, newCtx)
}

func (r *allianceResolver) Creator(ctx context.Context, obj *model.Alliance) (*model.Character, error) {
	return character.CharacterByID(obj.CreatorID)
}

func (r *allianceResolver) ExecutorCorporation(ctx context.Context, obj *model.Alliance) (*model.Corporation, error) {
	newCtx, span := traceProvider.Tracer(tracer_name).Start(ctx, "ExecutorCorporation")
	defer span.End()
	return corporation.CorporationByID(obj.ExecutorCorporationID, newCtx)
}

func (r *allianceResolver) Faction(ctx context.Context, obj *model.Alliance) (*model.Faction, error) {
	return universe.FactionByID(obj.FactionID)
}

func (r *ancestryResolver) Bloodline(ctx context.Context, obj *model.Ancestry) (*model.Bloodline, error) {
	return universe.BloodlineByID(obj.BloodlineID)
}

func (r *asteroid_beltResolver) System(ctx context.Context, obj *model.AsteroidBelt) (*model.System, error) {
	return universe.SystemByID(obj.SystemID)
}

func (r *characterResolver) Alliance(ctx context.Context, obj *model.Character) (*model.Alliance, error) {
	return alliance.AllianceByID(obj.AllianceID)
}

func (r *characterResolver) Ancestry(ctx context.Context, obj *model.Character) (*model.Ancestry, error) {
	return universe.AncestryByID(obj.AncestryID)
}

func (r *characterResolver) Bloodline(ctx context.Context, obj *model.Character) (*model.Bloodline, error) {
	return universe.BloodlineByID(obj.BloodlineID)
}

func (r *characterResolver) Corporation(ctx context.Context, obj *model.Character) (*model.Corporation, error) {
	newCtx, span := traceProvider.Tracer(tracer_name).Start(ctx, "CharacterCorporation")
	defer span.End()
	return corporation.CorporationByID(obj.CorporationID, newCtx)
}

func (r *characterResolver) Faction(ctx context.Context, obj *model.Character) (*model.Faction, error) {
	return universe.FactionByID(obj.FactionID)
}

func (r *characterResolver) Race(ctx context.Context, obj *model.Character) (*model.Race, error) {
	return universe.RaceByID(obj.RaceID)
}

func (r *constellationResolver) Region(ctx context.Context, obj *model.Constellation) (*model.Region, error) {
	return universe.RegionByID(obj.RegionID)
}

func (r *constellationResolver) SolarSystems(ctx context.Context, obj *model.Constellation) ([]*model.System, error) {
	return universe.SystemsByIDs(obj.Systems)
}

func (r *corporationResolver) Alliance(ctx context.Context, obj *model.Corporation) (*model.Alliance, error) {
	return alliance.AllianceByID(obj.AllianceID)
}

func (r *corporationResolver) Ceo(ctx context.Context, obj *model.Corporation) (*model.Character, error) {
	return character.CharacterByID(obj.CeoID)
}

func (r *corporationResolver) Creator(ctx context.Context, obj *model.Corporation) (*model.Character, error) {
	return character.CharacterByID(obj.CreatorID)
}

func (r *corporationResolver) Faction(ctx context.Context, obj *model.Corporation) (*model.Faction, error) {
	return universe.FactionByID(obj.FactionID)
}

func (r *corporationResolver) HomeStation(ctx context.Context, obj *model.Corporation) (*model.Station, error) {
	return universe.StationByID(obj.HomeStationID)
}

func (r *corporation_historyResolver) Employer(ctx context.Context, obj *model.CorporationHistory) (*model.Corporation, error) {
	newCtx, span := traceProvider.Tracer(tracer_name).Start(ctx, "Employer")
	defer span.End()
	return corporation.CorporationByID(obj.CorporationID, newCtx)
}

func (r *dogma_attributeResolver) Attribute(ctx context.Context, obj *model.DogmaAttribute) (*model.DogmaAttributeDetail, error) {
	return dogma.DogmaAttributeByID(obj.AttributeID)
}

func (r *dogma_effectResolver) Effect(ctx context.Context, obj *model.DogmaEffect) (*model.DogmaEffectDetail, error) {
	return dogma.DogmaEffectByID(obj.EffectID)
}

func (r *dogma_effect_detailResolver) DischargeAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	return dogma.DogmaAttributeByID(obj.DischargeAttributeID)
}

func (r *dogma_effect_detailResolver) DurationAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	return dogma.DogmaAttributeByID(obj.DurationAttributeID)
}

func (r *dogma_effect_detailResolver) FalloffAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	return dogma.DogmaAttributeByID(obj.FalloffAttributeID)
}

func (r *dogma_effect_detailResolver) RangeAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	return dogma.DogmaAttributeByID(obj.RangeAttributeID)
}

func (r *dogma_effect_detailResolver) TrackingSpeedAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	return dogma.DogmaAttributeByID(obj.TrackingSpeedAttributeID)
}

func (r *factionResolver) Corporation(ctx context.Context, obj *model.Faction) (*model.Corporation, error) {
	newCtx, span := traceProvider.Tracer(tracer_name).Start(ctx, "FactionCorporation")
	defer span.End()
	return corporation.CorporationByID(obj.CorporationID, newCtx)
}

func (r *factionResolver) MilitiaCorporation(ctx context.Context, obj *model.Faction) (*model.Corporation, error) {
	newCtx, span := traceProvider.Tracer(tracer_name).Start(ctx, "MilitiaCorporation")
	defer span.End()
	return corporation.CorporationByID(obj.MilitiaCorporationID, newCtx)
}

func (r *factionResolver) SolarSystem(ctx context.Context, obj *model.Faction) (*model.System, error) {
	return universe.SystemByID(obj.SolarSystemID)
}

func (r *groupResolver) Category(ctx context.Context, obj *model.Group) (*model.Category, error) {
	return universe.CategoryByID(obj.CategoryID)
}

func (r *groupResolver) ItemTypes(ctx context.Context, obj *model.Group) ([]*model.ItemType, error) {
	return universe.ItemTypesByIDs(obj.Types)
}

func (r *item_typeResolver) Graphic(ctx context.Context, obj *model.ItemType) (*model.Graphic, error) {
	return universe.GraphicByID(obj.GraphicID)
}

func (r *item_typeResolver) Group(ctx context.Context, obj *model.ItemType) (*model.Group, error) {
	return universe.GroupByID(obj.GroupID)
}

func (r *item_typeResolver) MarketGroup(ctx context.Context, obj *model.ItemType) (*model.MarketGroup, error) {
	return market.MarketGroupByID(obj.MarketGroupID)
}

func (r *market_groupResolver) ParentGroup(ctx context.Context, obj *model.MarketGroup) (*model.Group, error) {
	return universe.GroupByID(obj.ParentGroupID)
}

func (r *market_groupResolver) TypesDetails(ctx context.Context, obj *model.MarketGroup) ([]*model.ItemType, error) {
	return universe.ItemTypesByIDs(obj.Types)
}

func (r *modifierResolver) ModifiedAttribute(ctx context.Context, obj *model.Modifier) (*model.DogmaAttributeDetail, error) {
	return dogma.DogmaAttributeByID(obj.ModifiedAttributeID)
}

func (r *modifierResolver) ModifyingAttribute(ctx context.Context, obj *model.Modifier) (*model.DogmaAttributeDetail, error) {
	return dogma.DogmaAttributeByID(obj.ModifyingAttributeID)
}

func (r *moonResolver) System(ctx context.Context, obj *model.Moon) (*model.System, error) {
	return universe.SystemByID(obj.SystemID)
}

func (r *orderResolver) Location(ctx context.Context, obj *model.Order) (*model.Station, error) {
	return universe.StationByID(obj.LocationID)
}

func (r *orderResolver) System(ctx context.Context, obj *model.Order) (*model.System, error) {
	return universe.SystemByID(obj.SystemID)
}

func (r *orderResolver) ItemType(ctx context.Context, obj *model.Order) (*model.ItemType, error) {
	return universe.ItemTypeByID(obj.TypeID)
}

func (r *planetResolver) System(ctx context.Context, obj *model.Planet) (*model.System, error) {
	return universe.SystemByID(obj.SystemID)
}

func (r *planetResolver) ItemType(ctx context.Context, obj *model.Planet) (*model.ItemType, error) {
	return universe.ItemTypeByID(obj.TypeID)
}

func (r *queryResolver) OrdersForRegion(ctx context.Context, regionID int, orderType model.Ordertype, typeID *int, page int) (*model.OrderWrapper, error) {
	return market.OrdersForRegion(&regionID, &orderType, typeID, &page)
}

func (r *queryResolver) OrdersForRegionByName(ctx context.Context, region string, orderType model.Ordertype, typeName *string, page int) (*model.OrderWrapper, error) {
	return market.OrdersForRegionByName(&region, &orderType, typeName, &page)
}

func (r *queryResolver) SystemByID(ctx context.Context, id *int) (*model.System, error) {
	return universe.SystemByID(id)
}

func (r *queryResolver) StationByID(ctx context.Context, id *int) (*model.Station, error) {
	return universe.StationByID(id)
}

func (r *queryResolver) PlanetByID(ctx context.Context, id *int) (*model.Planet, error) {
	return universe.PlanetByID(id)
}

func (r *queryResolver) CorporationByID(ctx context.Context, id *int) (*model.Corporation, error) {
	newCtx, span := traceProvider.Tracer(tracer_name).Start(ctx, "CorporationByID")
	defer span.End()
	return corporation.CorporationByID(id, newCtx)
}

func (r *queryResolver) CorporationHistoryForCharacterID(ctx context.Context, id *int) ([]*model.CorporationHistory, error) {
	return character.CorporationHistory(id)
}

func (r *queryResolver) CharacterPortraitByID(ctx context.Context, id *int) (*model.CharacterPortrait, error) {
	return character.CharacterPortraitByID(id)
}

func (r *queryResolver) FactionByID(ctx context.Context, id *int) (*model.Faction, error) {
	return universe.FactionByID(id)
}

func (r *queryResolver) OrderHistory(ctx context.Context, regionID *int, typeID *int) ([]*model.OrderHistory, error) {
	return market.OrderHistory(regionID, typeID)
}

func (r *regionResolver) ConstellationList(ctx context.Context, obj *model.Region) ([]*model.Constellation, error) {
	return universe.ConstellationsByIDs(obj.Constellations)
}

func (r *starResolver) SolarSystem(ctx context.Context, obj *model.Star) (*model.System, error) {
	return universe.SystemByID(obj.SolarSystemID)
}

func (r *starResolver) ItemType(ctx context.Context, obj *model.Star) (*model.ItemType, error) {
	return universe.ItemTypeByID(obj.TypeID)
}

func (r *stargateResolver) ItemType(ctx context.Context, obj *model.Stargate) (*model.ItemType, error) {
	return universe.ItemTypeByID(obj.TypeID)
}

func (r *stargateDestinationResolver) Stargate(ctx context.Context, obj *model.StargateDestination) (*model.Stargate, error) {
	return universe.StargateByID(obj.StargateID)
}

func (r *stargateDestinationResolver) System(ctx context.Context, obj *model.StargateDestination) (*model.System, error) {
	return universe.SystemByID(obj.SystemID)
}

func (r *stationResolver) OwningCorporation(ctx context.Context, obj *model.Station) (*model.Corporation, error) {
	newCtx, span := traceProvider.Tracer(tracer_name).Start(ctx, "OwningCorporation")
	defer span.End()
	return corporation.CorporationByID(obj.Owner, newCtx)
}

func (r *stationResolver) Race(ctx context.Context, obj *model.Station) (*model.Race, error) {
	return universe.RaceByID(obj.RaceID)
}

func (r *stationResolver) System(ctx context.Context, obj *model.Station) (*model.System, error) {
	return universe.SystemByID(obj.SystemID)
}

func (r *stationResolver) StationType(ctx context.Context, obj *model.Station) (*model.ItemType, error) {
	return universe.ItemTypeByID(obj.TypeID)
}

func (r *systemResolver) Constellation(ctx context.Context, obj *model.System) (*model.Constellation, error) {
	return universe.ConstellationByID(obj.ConstellationID)
}

func (r *systemResolver) Star(ctx context.Context, obj *model.System) (*model.Star, error) {
	return universe.StarByID(obj.StarID)
}

func (r *systemResolver) StargateList(ctx context.Context, obj *model.System) ([]*model.Stargate, error) {
	return universe.StargateDetails(obj.Stargates)
}

func (r *systemResolver) StationList(ctx context.Context, obj *model.System) ([]*model.Station, error) {
	return universe.StationsByIDs(obj.Stations)
}

func (r *system_planetResolver) AsteroidBeltsProperties(ctx context.Context, obj *model.SystemPlanet) ([]*model.AsteroidBelt, error) {
	return universe.AsteroidBeltDetails(obj.AsteroidBelts)
}

func (r *system_planetResolver) MoonDetails(ctx context.Context, obj *model.SystemPlanet) ([]*model.Moon, error) {
	return universe.MoonDetails(obj.Moons)
}

func (r *system_planetResolver) PlanetProperties(ctx context.Context, obj *model.SystemPlanet) (*model.Planet, error) {
	return universe.PlanetByID(obj.PlanetID)
}

// Alliance returns generated.AllianceResolver implementation.
func (r *Resolver) Alliance() generated.AllianceResolver { return &allianceResolver{r} }

// Ancestry returns generated.AncestryResolver implementation.
func (r *Resolver) Ancestry() generated.AncestryResolver { return &ancestryResolver{r} }

// Asteroid_belt returns generated.Asteroid_beltResolver implementation.
func (r *Resolver) Asteroid_belt() generated.Asteroid_beltResolver { return &asteroid_beltResolver{r} }

// Character returns generated.CharacterResolver implementation.
func (r *Resolver) Character() generated.CharacterResolver { return &characterResolver{r} }

// Constellation returns generated.ConstellationResolver implementation.
func (r *Resolver) Constellation() generated.ConstellationResolver { return &constellationResolver{r} }

// Corporation returns generated.CorporationResolver implementation.
func (r *Resolver) Corporation() generated.CorporationResolver { return &corporationResolver{r} }

// Corporation_history returns generated.Corporation_historyResolver implementation.
func (r *Resolver) Corporation_history() generated.Corporation_historyResolver {
	return &corporation_historyResolver{r}
}

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

// Moon returns generated.MoonResolver implementation.
func (r *Resolver) Moon() generated.MoonResolver { return &moonResolver{r} }

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

// Planet returns generated.PlanetResolver implementation.
func (r *Resolver) Planet() generated.PlanetResolver { return &planetResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Region returns generated.RegionResolver implementation.
func (r *Resolver) Region() generated.RegionResolver { return &regionResolver{r} }

// Star returns generated.StarResolver implementation.
func (r *Resolver) Star() generated.StarResolver { return &starResolver{r} }

// Stargate returns generated.StargateResolver implementation.
func (r *Resolver) Stargate() generated.StargateResolver { return &stargateResolver{r} }

// StargateDestination returns generated.StargateDestinationResolver implementation.
func (r *Resolver) StargateDestination() generated.StargateDestinationResolver {
	return &stargateDestinationResolver{r}
}

// Station returns generated.StationResolver implementation.
func (r *Resolver) Station() generated.StationResolver { return &stationResolver{r} }

// System returns generated.SystemResolver implementation.
func (r *Resolver) System() generated.SystemResolver { return &systemResolver{r} }

// System_planet returns generated.System_planetResolver implementation.
func (r *Resolver) System_planet() generated.System_planetResolver { return &system_planetResolver{r} }

type allianceResolver struct{ *Resolver }
type ancestryResolver struct{ *Resolver }
type asteroid_beltResolver struct{ *Resolver }
type characterResolver struct{ *Resolver }
type constellationResolver struct{ *Resolver }
type corporationResolver struct{ *Resolver }
type corporation_historyResolver struct{ *Resolver }
type dogma_attributeResolver struct{ *Resolver }
type dogma_effectResolver struct{ *Resolver }
type dogma_effect_detailResolver struct{ *Resolver }
type factionResolver struct{ *Resolver }
type groupResolver struct{ *Resolver }
type item_typeResolver struct{ *Resolver }
type market_groupResolver struct{ *Resolver }
type modifierResolver struct{ *Resolver }
type moonResolver struct{ *Resolver }
type orderResolver struct{ *Resolver }
type planetResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type regionResolver struct{ *Resolver }
type starResolver struct{ *Resolver }
type stargateResolver struct{ *Resolver }
type stargateDestinationResolver struct{ *Resolver }
type stationResolver struct{ *Resolver }
type systemResolver struct{ *Resolver }
type system_planetResolver struct{ *Resolver }
