package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/alliance"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/asset"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/character"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/corporation"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/dogma"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/market"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/universe"
	"github.com/cryanbrow/eve-graphql-go/graph/generated"
	"github.com/cryanbrow/eve-graphql-go/graph/generated/model"
	"github.com/cryanbrow/eve-graphql-go/graph/tracing"
)

func (r *allianceResolver) CreatorCorporation(ctx context.Context, obj *model.Alliance) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CreatorCorporation")
	defer span.End()
	return corporation.ByID(newCtx, obj.CreatorCorporationID)
}

func (r *allianceResolver) Creator(ctx context.Context, obj *model.Alliance) (*model.Character, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "Creator")
	defer span.End()
	return character.ByID(newCtx, obj.CreatorID)
}

func (r *allianceResolver) ExecutorCorporation(ctx context.Context, obj *model.Alliance) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ExecutorCorporation")
	defer span.End()
	return corporation.ByID(newCtx, obj.ExecutorCorporationID)
}

func (r *allianceResolver) Faction(ctx context.Context, obj *model.Alliance) (*model.Faction, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AllianceFaction")
	defer span.End()
	return universe.FactionByID(newCtx, obj.FactionID)
}

func (r *ancestryResolver) Bloodline(ctx context.Context, obj *model.Ancestry) (*model.Bloodline, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AncestryBloodline")
	defer span.End()
	return universe.BloodlineByID(newCtx, obj.BloodlineID)
}

func (r *assetResolver) Location(ctx context.Context, obj *model.Asset) (*model.Station, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AssetLocation")
	defer span.End()
	return universe.StationByID(newCtx, obj.LocationID)
}

func (r *assetResolver) ItemType(ctx context.Context, obj *model.Asset) (*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AssetItemTypes")
	defer span.End()
	return universe.ItemTypeByID(newCtx, obj.TypeID)
}

func (r *asteroid_beltResolver) System(ctx context.Context, obj *model.AsteroidBelt) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AstroidBeltSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SystemID)
}

func (r *characterResolver) Alliance(ctx context.Context, obj *model.Character) (*model.Alliance, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CharcterAlliance")
	defer span.End()
	return alliance.ByID(newCtx, obj.AllianceID)
}

func (r *characterResolver) Ancestry(ctx context.Context, obj *model.Character) (*model.Ancestry, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CharacterAncestry")
	defer span.End()
	return universe.AncestryByID(newCtx, obj.AncestryID)
}

func (r *characterResolver) Bloodline(ctx context.Context, obj *model.Character) (*model.Bloodline, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ChracterBloodline")
	defer span.End()
	return universe.BloodlineByID(newCtx, obj.BloodlineID)
}

func (r *characterResolver) Corporation(ctx context.Context, obj *model.Character) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CharacterCorporation")
	defer span.End()
	return corporation.ByID(newCtx, obj.CorporationID)
}

func (r *characterResolver) Faction(ctx context.Context, obj *model.Character) (*model.Faction, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CharacterFaction")
	defer span.End()
	return universe.FactionByID(newCtx, obj.FactionID)
}

func (r *characterResolver) Race(ctx context.Context, obj *model.Character) (*model.Race, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CharacterRace")
	defer span.End()
	return universe.RaceByID(newCtx, obj.RaceID)
}

func (r *constellationResolver) Region(ctx context.Context, obj *model.Constellation) (*model.Region, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ConstellationRegion")
	defer span.End()
	return universe.RegionByID(newCtx, obj.RegionID)
}

func (r *constellationResolver) SolarSystems(ctx context.Context, obj *model.Constellation) ([]*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ConstellationSolarSystem")
	defer span.End()
	return universe.SystemsByIDs(newCtx, obj.Systems)
}

func (r *corporationResolver) Alliance(ctx context.Context, obj *model.Corporation) (*model.Alliance, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CorporationAlliance")
	defer span.End()
	return alliance.ByID(newCtx, obj.AllianceID)
}

func (r *corporationResolver) Ceo(ctx context.Context, obj *model.Corporation) (*model.Character, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CorporationCeo")
	defer span.End()
	return character.ByID(newCtx, obj.CeoID)
}

func (r *corporationResolver) Creator(ctx context.Context, obj *model.Corporation) (*model.Character, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CorporationCreator")
	defer span.End()
	return character.ByID(newCtx, obj.CreatorID)
}

func (r *corporationResolver) Faction(ctx context.Context, obj *model.Corporation) (*model.Faction, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CorporationFaction")
	defer span.End()
	return universe.FactionByID(newCtx, obj.FactionID)
}

func (r *corporationResolver) HomeStation(ctx context.Context, obj *model.Corporation) (*model.Station, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CorporationHomeStation")
	defer span.End()
	return universe.StationByID(newCtx, obj.HomeStationID)
}

func (r *corporation_historyResolver) Employer(ctx context.Context, obj *model.CorporationHistory) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "Employer")
	defer span.End()
	return corporation.ByID(newCtx, obj.CorporationID)
}

func (r *dogma_attributeResolver) Attribute(ctx context.Context, obj *model.DogmaAttribute) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "DogmaAttributeAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.AttributeID)
}

func (r *dogma_effectResolver) Effect(ctx context.Context, obj *model.DogmaEffect) (*model.DogmaEffectDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "DogmaEffectEffect")
	defer span.End()
	return dogma.EffectByID(newCtx, obj.EffectID)
}

func (r *dogma_effect_detailResolver) DischargeAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "DischargeAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.DischargeAttributeID)
}

func (r *dogma_effect_detailResolver) DurationAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "DurationAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.DurationAttributeID)
}

func (r *dogma_effect_detailResolver) FalloffAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "FalloffAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.FalloffAttributeID)
}

func (r *dogma_effect_detailResolver) RangeAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "RangeAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.RangeAttributeID)
}

func (r *dogma_effect_detailResolver) TrackingSpeedAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "TrackingSpeedAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.TrackingSpeedAttributeID)
}

func (r *factionResolver) Corporation(ctx context.Context, obj *model.Faction) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "FactionCorporation")
	defer span.End()
	return corporation.ByID(newCtx, obj.CorporationID)
}

func (r *factionResolver) MilitiaCorporation(ctx context.Context, obj *model.Faction) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "MilitiaCorporation")
	defer span.End()
	return corporation.ByID(newCtx, obj.MilitiaCorporationID)
}

func (r *factionResolver) SolarSystem(ctx context.Context, obj *model.Faction) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "FactionSolarSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SolarSystemID)
}

func (r *groupResolver) Category(ctx context.Context, obj *model.Group) (*model.Category, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "GroupCategory")
	defer span.End()
	return universe.CategoryByID(newCtx, obj.CategoryID)
}

func (r *groupResolver) ItemTypes(ctx context.Context, obj *model.Group) ([]*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "GroupItemTypes")
	defer span.End()
	return universe.ItemTypesByIDs(newCtx, obj.Types)
}

func (r *item_typeResolver) Graphic(ctx context.Context, obj *model.ItemType) (*model.Graphic, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ItemTypeGraphic")
	defer span.End()
	return universe.GraphicByID(newCtx, obj.GraphicID)
}

func (r *item_typeResolver) Group(ctx context.Context, obj *model.ItemType) (*model.Group, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ItemTypeGroup")
	defer span.End()
	return universe.GroupByID(newCtx, obj.GroupID)
}

func (r *item_typeResolver) MarketGroup(ctx context.Context, obj *model.ItemType) (*model.MarketGroup, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ItemTypeMarketGroup")
	defer span.End()
	return market.GroupByID(newCtx, obj.MarketGroupID)
}

func (r *market_groupResolver) ParentGroup(ctx context.Context, obj *model.MarketGroup) (*model.Group, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "MarketGroupParentGroup")
	defer span.End()
	return universe.GroupByID(newCtx, obj.ParentGroupID)
}

func (r *market_groupResolver) TypesDetails(ctx context.Context, obj *model.MarketGroup) ([]*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "MarketGroupTypeDetails")
	defer span.End()
	return universe.ItemTypesByIDs(newCtx, obj.Types)
}

func (r *modifierResolver) ModifiedAttribute(ctx context.Context, obj *model.Modifier) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ModifierModifiedAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.ModifiedAttributeID)
}

func (r *modifierResolver) ModifyingAttribute(ctx context.Context, obj *model.Modifier) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ModifyingAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.ModifyingAttributeID)
}

func (r *moonResolver) System(ctx context.Context, obj *model.Moon) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "MoonSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SystemID)
}

func (r *orderResolver) Location(ctx context.Context, obj *model.Order) (*model.Station, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "OrderLocation")
	defer span.End()
	return universe.StationByID(newCtx, obj.LocationID)
}

func (r *orderResolver) System(ctx context.Context, obj *model.Order) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "OrderSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SystemID)
}

func (r *orderResolver) ItemType(ctx context.Context, obj *model.Order) (*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "OrderItemType")
	defer span.End()
	return universe.ItemTypeByID(newCtx, obj.TypeID)
}

func (r *planetResolver) System(ctx context.Context, obj *model.Planet) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "PlanetSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SystemID)
}

func (r *planetResolver) ItemType(ctx context.Context, obj *model.Planet) (*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "PlanetItemType")
	defer span.End()
	return universe.ItemTypeByID(newCtx, obj.TypeID)
}

func (r *queryResolver) AlliancesByID(ctx context.Context, id *int) (*model.Alliance, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AlliancesByID")
	defer span.End()
	return alliance.ByID(newCtx, id)
}

func (r *queryResolver) AlliancesCorporationsByID(ctx context.Context, id *int) ([]*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AlliancesCorporationsByID")
	defer span.End()
	return alliance.CorporationsByID(newCtx, id)
}

func (r *queryResolver) AlliancesIconByID(ctx context.Context, id *int) (*model.Icon, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AlliancesIconByID")
	defer span.End()
	return alliance.IconByID(newCtx, id)
}

func (r *queryResolver) AlliancesByName(ctx context.Context, name *string) (*model.Alliance, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AlliancesByName")
	defer span.End()
	return alliance.ByName(newCtx, name)
}

func (r *queryResolver) AlliancesCorporationsByName(ctx context.Context, name *string) ([]*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AlliancesCorporationsByName")
	defer span.End()
	return alliance.CorporationsByName(newCtx, name)
}

func (r *queryResolver) AlliancesIconByName(ctx context.Context, name *string) (*model.Icon, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AlliancesIconByName")
	defer span.End()
	return alliance.IconByName(newCtx, name)
}

func (r *queryResolver) AssetsByCharacterID(ctx context.Context, id *int) ([]*model.Asset, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AssetsByCharacterID")
	defer span.End()
	return asset.AssetsByCharacterID(newCtx, id)
}

func (r *queryResolver) AssetsByCorporationID(ctx context.Context, id *int) ([]*model.Asset, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AssetsByCorporationID")
	defer span.End()
	return asset.AssetsByCorporationID(newCtx, id)
}

func (r *queryResolver) AssetsByCharacterName(ctx context.Context, name *string) ([]*model.Asset, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AssetsByCharacterName")
	defer span.End()
	return asset.AssetsByCharacterName(newCtx, name)
}

func (r *queryResolver) AssetsByCorporationName(ctx context.Context, name *string) ([]*model.Asset, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AssetsByCorporationName")
	defer span.End()
	return asset.AssetsByCorporationName(newCtx, name)
}

func (r *queryResolver) BookmarksByCharacterID(ctx context.Context, id *int) ([]*model.Bookmark, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BookmarksFoldersByCharacterID(ctx context.Context, id *int) ([]*model.BookmarkFolder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BookmarksByCorporationID(ctx context.Context, id *int) ([]*model.Bookmark, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BookmarksFoldersByCorporationID(ctx context.Context, id *int) ([]*model.BookmarkFolder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BookmarksByCharacterName(ctx context.Context, name *string) ([]*model.Bookmark, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BookmarksFoldersByCharacterName(ctx context.Context, name *string) ([]*model.BookmarkFolder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BookmarksByCorporationName(ctx context.Context, name *string) ([]*model.Bookmark, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BookmarksFoldersByCorporationName(ctx context.Context, name *string) ([]*model.BookmarkFolder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CalendarEventsByCharacterID(ctx context.Context, id *int) ([]*model.EventSummary, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CalendarEventByID(ctx context.Context, characterID *int, eventID *int) (*model.EventDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CalendarEventAttendeesByID(ctx context.Context, characterID *int, eventID *int) ([]*model.EventAttendee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CalendarEventsByCharacterName(ctx context.Context, name *string) ([]*model.EventSummary, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CalendarEventByCharacterNameAndID(ctx context.Context, characterName *string, eventID *int) (*model.EventDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CalendarEventAttendeesByCharacterNameAndID(ctx context.Context, characterID *int, eventID *int) ([]*model.EventAttendee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterByID(ctx context.Context, id *int) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterResearchAgentsByID(ctx context.Context, id *int) ([]*model.ResearchAgent, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterBlueprintsByID(ctx context.Context, id *int) ([]*model.Blueprint, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterCorporationHistoryByID(ctx context.Context, id *int) ([]*model.CorporationHistory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterFatiqueByID(ctx context.Context, id *int) (*model.Fatique, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterMedalsByID(ctx context.Context, id *int) ([]*model.Medal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterNotificationsByID(ctx context.Context, id *int) ([]*model.Notification, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterPortraitByID(ctx context.Context, id *int) (*model.CharacterPortrait, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CharacterPortraitByID")
	defer span.End()
	return character.PortraitByID(newCtx, id)
}

func (r *queryResolver) CharacterRoleByID(ctx context.Context, id *int) (*model.Role, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterStandingsByID(ctx context.Context, id *int) ([]*model.Standing, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterTitlesByID(ctx context.Context, id *int) ([]*model.Title, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterByName(ctx context.Context, name *string) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterResearchAgentsByName(ctx context.Context, name *string) ([]*model.ResearchAgent, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterBlueprintsByName(ctx context.Context, name *string) ([]*model.Blueprint, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterCorporationHistoryByName(ctx context.Context, name *string) ([]*model.CorporationHistory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterFatiqueByName(ctx context.Context, name *string) (*model.Fatique, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterMedalsByName(ctx context.Context, name *string) ([]*model.Medal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterNotificationsByName(ctx context.Context, name *string) ([]*model.Notification, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterPortraitByName(ctx context.Context, name *string) (*model.CharacterPortrait, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterRoleByName(ctx context.Context, name *string) (*model.Role, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterStandingsByName(ctx context.Context, name *string) ([]*model.Standing, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CharacterTitlesByName(ctx context.Context, name *string) ([]*model.Title, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ClonesByID(ctx context.Context, id *int) ([]*model.Clone, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ClonesImplantsByID(ctx context.Context, id *int) ([]*model.Implant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OrdersForRegion(ctx context.Context, regionID int, orderType model.Ordertype, typeID *int, page int) (*model.OrderWrapper, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ResolverOrdersForRegion")
	defer span.End()
	return market.OrdersForRegion(newCtx, &regionID, &orderType, typeID, &page)
}

func (r *queryResolver) OrdersForRegionByName(ctx context.Context, region string, orderType model.Ordertype, typeName *string, page int) (*model.OrderWrapper, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ResolverOrdersForRegionByName")
	defer span.End()
	return market.OrdersForRegionByName(newCtx, &region, &orderType, typeName, &page)
}

func (r *queryResolver) SystemByID(ctx context.Context, id *int) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SystemByID")
	defer span.End()
	return universe.SystemByID(newCtx, id)
}

func (r *queryResolver) StationByID(ctx context.Context, id *int) (*model.Station, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StationByID")
	defer span.End()
	return universe.StationByID(newCtx, id)
}

func (r *queryResolver) PlanetByID(ctx context.Context, id *int) (*model.Planet, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "PlanetByID")
	defer span.End()
	return universe.PlanetByID(newCtx, id)
}

func (r *queryResolver) CorporationByID(ctx context.Context, id *int) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CorporationByID")
	defer span.End()
	return corporation.ByID(newCtx, id)
}

func (r *queryResolver) CorporationHistoryForCharacterID(ctx context.Context, id *int) ([]*model.CorporationHistory, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CorporationHistoryForCharacterID")
	defer span.End()
	return character.CorporationHistory(newCtx, id)
}

func (r *queryResolver) FactionByID(ctx context.Context, id *int) (*model.Faction, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "FactionByID")
	defer span.End()
	return universe.FactionByID(newCtx, id)
}

func (r *queryResolver) OrderHistory(ctx context.Context, regionID *int, typeID *int) ([]*model.OrderHistory, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "OrderHistory")
	defer span.End()
	return market.OrderHistory(newCtx, regionID, typeID)
}

func (r *regionResolver) ConstellationList(ctx context.Context, obj *model.Region) ([]*model.Constellation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "RegionConstellationList")
	defer span.End()
	return universe.ConstellationsByIDs(newCtx, obj.Constellations)
}

func (r *starResolver) SolarSystem(ctx context.Context, obj *model.Star) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StarSolarSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SolarSystemID)
}

func (r *starResolver) ItemType(ctx context.Context, obj *model.Star) (*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StarItemType")
	defer span.End()
	return universe.ItemTypeByID(newCtx, obj.TypeID)
}

func (r *stargateResolver) ItemType(ctx context.Context, obj *model.Stargate) (*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StragateItemType")
	defer span.End()
	return universe.ItemTypeByID(newCtx, obj.TypeID)
}

func (r *stargateDestinationResolver) Stargate(ctx context.Context, obj *model.StargateDestination) (*model.Stargate, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StargateDestinationStargate")
	defer span.End()
	return universe.StargateByID(newCtx, obj.StargateID)
}

func (r *stargateDestinationResolver) System(ctx context.Context, obj *model.StargateDestination) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StargateDestinationSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SystemID)
}

func (r *stationResolver) OwningCorporation(ctx context.Context, obj *model.Station) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "OwningCorporation")
	defer span.End()
	return corporation.ByID(newCtx, obj.Owner)
}

func (r *stationResolver) Race(ctx context.Context, obj *model.Station) (*model.Race, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StationRace")
	defer span.End()
	return universe.RaceByID(newCtx, obj.RaceID)
}

func (r *stationResolver) System(ctx context.Context, obj *model.Station) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StationSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SystemID)
}

func (r *stationResolver) StationType(ctx context.Context, obj *model.Station) (*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StationStationType")
	defer span.End()
	return universe.ItemTypeByID(newCtx, obj.TypeID)
}

func (r *systemResolver) Constellation(ctx context.Context, obj *model.System) (*model.Constellation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SystemConstellation")
	defer span.End()
	return universe.ConstellationByID(newCtx, obj.ConstellationID)
}

func (r *systemResolver) Star(ctx context.Context, obj *model.System) (*model.Star, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SystemStar")
	defer span.End()
	return universe.StarByID(newCtx, obj.StarID)
}

func (r *systemResolver) StargateList(ctx context.Context, obj *model.System) ([]*model.Stargate, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SystemStargateList")
	defer span.End()
	return universe.StargateDetails(newCtx, obj.Stargates)
}

func (r *systemResolver) StationList(ctx context.Context, obj *model.System) ([]*model.Station, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SystemStationList")
	defer span.End()
	return universe.StationsByIDs(newCtx, obj.Stations)
}

func (r *system_planetResolver) AsteroidBeltsProperties(ctx context.Context, obj *model.SystemPlanet) ([]*model.AsteroidBelt, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AsteroidBeltsProperties")
	defer span.End()
	return universe.AsteroidBeltDetails(newCtx, obj.AsteroidBelts)
}

func (r *system_planetResolver) MoonDetails(ctx context.Context, obj *model.SystemPlanet) ([]*model.Moon, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SystemPlanetMoonDetails")
	defer span.End()
	return universe.MoonDetails(newCtx, obj.Moons)
}

func (r *system_planetResolver) PlanetProperties(ctx context.Context, obj *model.SystemPlanet) (*model.Planet, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SystemPlanetPlanetProperties")
	defer span.End()
	return universe.PlanetByID(newCtx, obj.PlanetID)
}

// Alliance returns generated.AllianceResolver implementation.
func (r *Resolver) Alliance() generated.AllianceResolver { return &allianceResolver{r} }

// Ancestry returns generated.AncestryResolver implementation.
func (r *Resolver) Ancestry() generated.AncestryResolver { return &ancestryResolver{r} }

// Asset returns generated.AssetResolver implementation.
func (r *Resolver) Asset() generated.AssetResolver { return &assetResolver{r} }

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
type assetResolver struct{ *Resolver }
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
