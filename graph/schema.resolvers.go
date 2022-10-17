package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/alliance"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/asset"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/corporation"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/dogma"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/eve_character"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/market"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/universe"
	"github.com/cryanbrow/eve-graphql-go/graph/generated"
	"github.com/cryanbrow/eve-graphql-go/graph/generated/model"
	"github.com/cryanbrow/eve-graphql-go/graph/tracing"
)

// CreatorCorporation is the resolver for the creator_corporation field.
func (r *allianceResolver) CreatorCorporation(ctx context.Context, obj *model.Alliance) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CreatorCorporation")
	defer span.End()
	return corporation.ByID(newCtx, obj.CreatorCorporationID)
}

// Creator is the resolver for the creator field.
func (r *allianceResolver) Creator(ctx context.Context, obj *model.Alliance) (*model.Character, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "Creator")
	defer span.End()
	return eve_character.ByID(newCtx, obj.CreatorID)
}

// ExecutorCorporation is the resolver for the executor_corporation field.
func (r *allianceResolver) ExecutorCorporation(ctx context.Context, obj *model.Alliance) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ExecutorCorporation")
	defer span.End()
	return corporation.ByID(newCtx, obj.ExecutorCorporationID)
}

// Faction is the resolver for the faction field.
func (r *allianceResolver) Faction(ctx context.Context, obj *model.Alliance) (*model.Faction, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AllianceFaction")
	defer span.End()
	return universe.FactionByID(newCtx, obj.FactionID)
}

// Bloodline is the resolver for the bloodline field.
func (r *ancestryResolver) Bloodline(ctx context.Context, obj *model.Ancestry) (*model.Bloodline, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AncestryBloodline")
	defer span.End()
	return universe.BloodlineByID(newCtx, obj.BloodlineID)
}

// Location is the resolver for the location field.
func (r *assetResolver) Location(ctx context.Context, obj *model.Asset) (*model.Station, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AssetLocation")
	defer span.End()
	return universe.StationByID(newCtx, obj.LocationID)
}

// ItemType is the resolver for the item_type field.
func (r *assetResolver) ItemType(ctx context.Context, obj *model.Asset) (*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AssetItemTypes")
	defer span.End()
	return universe.ItemTypeByID(newCtx, obj.TypeID)
}

// System is the resolver for the system field.
func (r *asteroid_beltResolver) System(ctx context.Context, obj *model.AsteroidBelt) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AstroidBeltSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SystemID)
}

// Alliance is the resolver for the alliance field.
func (r *characterResolver) Alliance(ctx context.Context, obj *model.Character) (*model.Alliance, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CharcterAlliance")
	defer span.End()
	return alliance.ByID(newCtx, obj.AllianceID)
}

// Ancestry is the resolver for the ancestry field.
func (r *characterResolver) Ancestry(ctx context.Context, obj *model.Character) (*model.Ancestry, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CharacterAncestry")
	defer span.End()
	return universe.AncestryByID(newCtx, obj.AncestryID)
}

// Bloodline is the resolver for the bloodline field.
func (r *characterResolver) Bloodline(ctx context.Context, obj *model.Character) (*model.Bloodline, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CharacterBloodline")
	defer span.End()
	return universe.BloodlineByID(newCtx, obj.BloodlineID)
}

// Corporation is the resolver for the corporation field.
func (r *characterResolver) Corporation(ctx context.Context, obj *model.Character) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CharacterCorporation")
	defer span.End()
	return corporation.ByID(newCtx, obj.CorporationID)
}

// Faction is the resolver for the faction field.
func (r *characterResolver) Faction(ctx context.Context, obj *model.Character) (*model.Faction, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CharacterFaction")
	defer span.End()
	return universe.FactionByID(newCtx, obj.FactionID)
}

// Race is the resolver for the race field.
func (r *characterResolver) Race(ctx context.Context, obj *model.Character) (*model.Race, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CharacterRace")
	defer span.End()
	return universe.RaceByID(newCtx, obj.RaceID)
}

// Region is the resolver for the region field.
func (r *constellationResolver) Region(ctx context.Context, obj *model.Constellation) (*model.Region, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ConstellationRegion")
	defer span.End()
	return universe.RegionByID(newCtx, obj.RegionID)
}

// SolarSystems is the resolver for the solar_systems field.
func (r *constellationResolver) SolarSystems(ctx context.Context, obj *model.Constellation) ([]*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ConstellationSolarSystem")
	defer span.End()
	return universe.SystemsByIDs(newCtx, obj.Systems)
}

// Alliance is the resolver for the alliance field.
func (r *corporationResolver) Alliance(ctx context.Context, obj *model.Corporation) (*model.Alliance, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CorporationAlliance")
	defer span.End()
	return alliance.ByID(newCtx, obj.AllianceID)
}

// Ceo is the resolver for the ceo field.
func (r *corporationResolver) Ceo(ctx context.Context, obj *model.Corporation) (*model.Character, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CorporationCeo")
	defer span.End()
	return eve_character.ByID(newCtx, obj.CeoID)
}

// Creator is the resolver for the creator field.
func (r *corporationResolver) Creator(ctx context.Context, obj *model.Corporation) (*model.Character, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CorporationCreator")
	defer span.End()
	return eve_character.ByID(newCtx, obj.CreatorID)
}

// Faction is the resolver for the faction field.
func (r *corporationResolver) Faction(ctx context.Context, obj *model.Corporation) (*model.Faction, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CorporationFaction")
	defer span.End()
	return universe.FactionByID(newCtx, obj.FactionID)
}

// HomeStation is the resolver for the home_station field.
func (r *corporationResolver) HomeStation(ctx context.Context, obj *model.Corporation) (*model.Station, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CorporationHomeStation")
	defer span.End()
	return universe.StationByID(newCtx, obj.HomeStationID)
}

// Employer is the resolver for the employer field.
func (r *corporation_historyResolver) Employer(ctx context.Context, obj *model.CorporationHistory) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "Employer")
	defer span.End()
	return corporation.ByID(newCtx, obj.CorporationID)
}

// Attribute is the resolver for the attribute field.
func (r *dogma_attributeResolver) Attribute(ctx context.Context, obj *model.DogmaAttribute) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "DogmaAttributeAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.AttributeID)
}

// Effect is the resolver for the effect field.
func (r *dogma_effectResolver) Effect(ctx context.Context, obj *model.DogmaEffect) (*model.DogmaEffectDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "DogmaEffectEffect")
	defer span.End()
	return dogma.EffectByID(newCtx, obj.EffectID)
}

// DischargeAttribute is the resolver for the discharge_attribute field.
func (r *dogma_effect_detailResolver) DischargeAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "DischargeAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.DischargeAttributeID)
}

// DurationAttribute is the resolver for the duration_attribute field.
func (r *dogma_effect_detailResolver) DurationAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "DurationAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.DurationAttributeID)
}

// FalloffAttribute is the resolver for the falloff_attribute field.
func (r *dogma_effect_detailResolver) FalloffAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "FalloffAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.FalloffAttributeID)
}

// RangeAttribute is the resolver for the range_attribute field.
func (r *dogma_effect_detailResolver) RangeAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "RangeAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.RangeAttributeID)
}

// TrackingSpeedAttribute is the resolver for the tracking_speed_attribute field.
func (r *dogma_effect_detailResolver) TrackingSpeedAttribute(ctx context.Context, obj *model.DogmaEffectDetail) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "TrackingSpeedAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.TrackingSpeedAttributeID)
}

// Corporation is the resolver for the corporation field.
func (r *factionResolver) Corporation(ctx context.Context, obj *model.Faction) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "FactionCorporation")
	defer span.End()
	return corporation.ByID(newCtx, obj.CorporationID)
}

// MilitiaCorporation is the resolver for the militia_corporation field.
func (r *factionResolver) MilitiaCorporation(ctx context.Context, obj *model.Faction) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "MilitiaCorporation")
	defer span.End()
	return corporation.ByID(newCtx, obj.MilitiaCorporationID)
}

// SolarSystem is the resolver for the solar_system field.
func (r *factionResolver) SolarSystem(ctx context.Context, obj *model.Faction) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "FactionSolarSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SolarSystemID)
}

// Category is the resolver for the category field.
func (r *groupResolver) Category(ctx context.Context, obj *model.Group) (*model.Category, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "GroupCategory")
	defer span.End()
	return universe.CategoryByID(newCtx, obj.CategoryID)
}

// ItemTypes is the resolver for the item_types field.
func (r *groupResolver) ItemTypes(ctx context.Context, obj *model.Group) ([]*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "GroupItemTypes")
	defer span.End()
	return universe.ItemTypesByIDs(newCtx, obj.Types)
}

// Graphic is the resolver for the graphic field.
func (r *item_typeResolver) Graphic(ctx context.Context, obj *model.ItemType) (*model.Graphic, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ItemTypeGraphic")
	defer span.End()
	return universe.GraphicByID(newCtx, obj.GraphicID)
}

// Group is the resolver for the group field.
func (r *item_typeResolver) Group(ctx context.Context, obj *model.ItemType) (*model.Group, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ItemTypeGroup")
	defer span.End()
	return universe.GroupByID(newCtx, obj.GroupID)
}

// MarketGroup is the resolver for the market_group field.
func (r *item_typeResolver) MarketGroup(ctx context.Context, obj *model.ItemType) (*model.MarketGroup, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ItemTypeMarketGroup")
	defer span.End()
	return market.GroupByID(newCtx, obj.MarketGroupID)
}

// ParentGroup is the resolver for the parent_group field.
func (r *market_groupResolver) ParentGroup(ctx context.Context, obj *model.MarketGroup) (*model.Group, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "MarketGroupParentGroup")
	defer span.End()
	return universe.GroupByID(newCtx, obj.ParentGroupID)
}

// TypesDetails is the resolver for the types_details field.
func (r *market_groupResolver) TypesDetails(ctx context.Context, obj *model.MarketGroup) ([]*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "MarketGroupTypeDetails")
	defer span.End()
	return universe.ItemTypesByIDs(newCtx, obj.Types)
}

// ModifiedAttribute is the resolver for the modified_attribute field.
func (r *modifierResolver) ModifiedAttribute(ctx context.Context, obj *model.Modifier) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ModifierModifiedAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.ModifiedAttributeID)
}

// ModifyingAttribute is the resolver for the modifying_attribute field.
func (r *modifierResolver) ModifyingAttribute(ctx context.Context, obj *model.Modifier) (*model.DogmaAttributeDetail, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ModifyingAttribute")
	defer span.End()
	return dogma.AttributeByID(newCtx, obj.ModifyingAttributeID)
}

// System is the resolver for the system field.
func (r *moonResolver) System(ctx context.Context, obj *model.Moon) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "MoonSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SystemID)
}

// Location is the resolver for the location field.
func (r *orderResolver) Location(ctx context.Context, obj *model.Order) (*model.Station, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "OrderLocation")
	defer span.End()
	return universe.StationByID(newCtx, obj.LocationID)
}

// System is the resolver for the system field.
func (r *orderResolver) System(ctx context.Context, obj *model.Order) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "OrderSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SystemID)
}

// ItemType is the resolver for the item_type field.
func (r *orderResolver) ItemType(ctx context.Context, obj *model.Order) (*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "OrderItemType")
	defer span.End()
	return universe.ItemTypeByID(newCtx, obj.TypeID)
}

// System is the resolver for the system field.
func (r *planetResolver) System(ctx context.Context, obj *model.Planet) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "PlanetSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SystemID)
}

// ItemType is the resolver for the item_type field.
func (r *planetResolver) ItemType(ctx context.Context, obj *model.Planet) (*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "PlanetItemType")
	defer span.End()
	return universe.ItemTypeByID(newCtx, obj.TypeID)
}

// AlliancesByID is the resolver for the alliances_byID field.
func (r *queryResolver) AlliancesByID(ctx context.Context, id *int) (*model.Alliance, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AlliancesByID")
	defer span.End()
	return alliance.ByID(newCtx, id)
}

// AlliancesCorporationsByID is the resolver for the alliances_corporationsByID field.
func (r *queryResolver) AlliancesCorporationsByID(ctx context.Context, id *int) ([]*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AlliancesCorporationsByID")
	defer span.End()
	return alliance.CorporationsByID(newCtx, id)
}

// AlliancesIconByID is the resolver for the alliances_iconByID field.
func (r *queryResolver) AlliancesIconByID(ctx context.Context, id *int) (*model.Icon, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AlliancesIconByID")
	defer span.End()
	return alliance.IconByID(newCtx, id)
}

// AlliancesByName is the resolver for the alliances_byName field.
func (r *queryResolver) AlliancesByName(ctx context.Context, name *string) (*model.Alliance, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AlliancesByName")
	defer span.End()
	return alliance.ByName(newCtx, name)
}

// AlliancesCorporationsByName is the resolver for the alliances_corporationsByName field.
func (r *queryResolver) AlliancesCorporationsByName(ctx context.Context, name *string) ([]*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AlliancesCorporationsByName")
	defer span.End()
	return alliance.CorporationsByName(newCtx, name)
}

// AlliancesIconByName is the resolver for the alliances_iconByName field.
func (r *queryResolver) AlliancesIconByName(ctx context.Context, name *string) (*model.Icon, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AlliancesIconByName")
	defer span.End()
	return alliance.IconByName(newCtx, name)
}

// AssetsByCharacterID is the resolver for the assets_byCharacterID field.
func (r *queryResolver) AssetsByCharacterID(ctx context.Context, id *int) ([]*model.Asset, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AssetsByCharacterID")
	defer span.End()
	return asset.AssetsByCharacterID(newCtx, id)
}

// AssetsByCorporationID is the resolver for the assets_byCorporationID field.
func (r *queryResolver) AssetsByCorporationID(ctx context.Context, id *int) ([]*model.Asset, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AssetsByCorporationID")
	defer span.End()
	return asset.AssetsByCorporationID(newCtx, id)
}

// AssetsByCharacterName is the resolver for the assets_byCharacterName field.
func (r *queryResolver) AssetsByCharacterName(ctx context.Context, name *string) ([]*model.Asset, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AssetsByCharacterName")
	defer span.End()
	return asset.AssetsByCharacterName(newCtx, name)
}

// AssetsByCorporationName is the resolver for the assets_byCorporationName field.
func (r *queryResolver) AssetsByCorporationName(ctx context.Context, name *string) ([]*model.Asset, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AssetsByCorporationName")
	defer span.End()
	return asset.AssetsByCorporationName(newCtx, name)
}

// BookmarksByCharacterID is the resolver for the bookmarks_byCharacterID field.
func (r *queryResolver) BookmarksByCharacterID(ctx context.Context, id *int) ([]*model.Bookmark, error) {
	panic(fmt.Errorf("not implemented"))
}

// BookmarksFoldersByCharacterID is the resolver for the bookmarks_foldersByCharacterID field.
func (r *queryResolver) BookmarksFoldersByCharacterID(ctx context.Context, id *int) ([]*model.BookmarkFolder, error) {
	panic(fmt.Errorf("not implemented"))
}

// BookmarksByCorporationID is the resolver for the bookmarks_byCorporationID field.
func (r *queryResolver) BookmarksByCorporationID(ctx context.Context, id *int) ([]*model.Bookmark, error) {
	panic(fmt.Errorf("not implemented"))
}

// BookmarksFoldersByCorporationID is the resolver for the bookmarks_foldersByCorporationID field.
func (r *queryResolver) BookmarksFoldersByCorporationID(ctx context.Context, id *int) ([]*model.BookmarkFolder, error) {
	panic(fmt.Errorf("not implemented"))
}

// BookmarksByCharacterName is the resolver for the bookmarks_byCharacterName field.
func (r *queryResolver) BookmarksByCharacterName(ctx context.Context, name *string) ([]*model.Bookmark, error) {
	panic(fmt.Errorf("not implemented"))
}

// BookmarksFoldersByCharacterName is the resolver for the bookmarks_foldersByCharacterName field.
func (r *queryResolver) BookmarksFoldersByCharacterName(ctx context.Context, name *string) ([]*model.BookmarkFolder, error) {
	panic(fmt.Errorf("not implemented"))
}

// BookmarksByCorporationName is the resolver for the bookmarks_byCorporationName field.
func (r *queryResolver) BookmarksByCorporationName(ctx context.Context, name *string) ([]*model.Bookmark, error) {
	panic(fmt.Errorf("not implemented"))
}

// BookmarksFoldersByCorporationName is the resolver for the bookmarks_foldersByCorporationName field.
func (r *queryResolver) BookmarksFoldersByCorporationName(ctx context.Context, name *string) ([]*model.BookmarkFolder, error) {
	panic(fmt.Errorf("not implemented"))
}

// CalendarEventsByCharacterID is the resolver for the calendar_eventsByCharacterID field.
func (r *queryResolver) CalendarEventsByCharacterID(ctx context.Context, id *int) ([]*model.EventSummary, error) {
	panic(fmt.Errorf("not implemented"))
}

// CalendarEventByID is the resolver for the calendar_eventByID field.
func (r *queryResolver) CalendarEventByID(ctx context.Context, characterID *int, eventID *int) (*model.EventDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

// CalendarEventAttendeesByID is the resolver for the calendar_eventAttendeesByID field.
func (r *queryResolver) CalendarEventAttendeesByID(ctx context.Context, characterID *int, eventID *int) ([]*model.EventAttendee, error) {
	panic(fmt.Errorf("not implemented"))
}

// CalendarEventsByCharacterName is the resolver for the calendar_eventsByCharacterName field.
func (r *queryResolver) CalendarEventsByCharacterName(ctx context.Context, name *string) ([]*model.EventSummary, error) {
	panic(fmt.Errorf("not implemented"))
}

// CalendarEventByCharacterNameAndID is the resolver for the calendar_eventByCharacterNameAndID field.
func (r *queryResolver) CalendarEventByCharacterNameAndID(ctx context.Context, characterName *string, eventID *int) (*model.EventDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

// CalendarEventAttendeesByCharacterNameAndID is the resolver for the calendar_eventAttendeesByCharacterNameAndID field.
func (r *queryResolver) CalendarEventAttendeesByCharacterNameAndID(ctx context.Context, characterID *int, eventID *int) ([]*model.EventAttendee, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterByID is the resolver for the character_byID field.
func (r *queryResolver) CharacterByID(ctx context.Context, id *int) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterResearchAgentsByID is the resolver for the character_researchAgentsByID field.
func (r *queryResolver) CharacterResearchAgentsByID(ctx context.Context, id *int) ([]*model.ResearchAgent, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterBlueprintsByID is the resolver for the character_blueprintsByID field.
func (r *queryResolver) CharacterBlueprintsByID(ctx context.Context, id *int) ([]*model.Blueprint, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterCorporationHistoryByID is the resolver for the character_corporationHistoryById field.
func (r *queryResolver) CharacterCorporationHistoryByID(ctx context.Context, id *int) ([]*model.CorporationHistory, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterFatiqueByID is the resolver for the character_fatiqueByID field.
func (r *queryResolver) CharacterFatiqueByID(ctx context.Context, id *int) (*model.Fatique, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterMedalsByID is the resolver for the character_medalsByID field.
func (r *queryResolver) CharacterMedalsByID(ctx context.Context, id *int) ([]*model.Medal, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterNotificationsByID is the resolver for the character_notificationsByID field.
func (r *queryResolver) CharacterNotificationsByID(ctx context.Context, id *int) ([]*model.Notification, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterPortraitByID is the resolver for the character_portraitByID field.
func (r *queryResolver) CharacterPortraitByID(ctx context.Context, id *int) (*model.CharacterPortrait, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CharacterPortraitByID")
	defer span.End()
	return eve_character.PortraitByID(newCtx, id)
}

// CharacterRoleByID is the resolver for the character_roleByID field.
func (r *queryResolver) CharacterRoleByID(ctx context.Context, id *int) (*model.Role, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterStandingsByID is the resolver for the character_standingsByID field.
func (r *queryResolver) CharacterStandingsByID(ctx context.Context, id *int) ([]*model.Standing, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterTitlesByID is the resolver for the character_titlesByID field.
func (r *queryResolver) CharacterTitlesByID(ctx context.Context, id *int) ([]*model.Title, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterByName is the resolver for the character_byName field.
func (r *queryResolver) CharacterByName(ctx context.Context, name *string) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterResearchAgentsByName is the resolver for the character_researchAgentsByName field.
func (r *queryResolver) CharacterResearchAgentsByName(ctx context.Context, name *string) ([]*model.ResearchAgent, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterBlueprintsByName is the resolver for the character_blueprintsByName field.
func (r *queryResolver) CharacterBlueprintsByName(ctx context.Context, name *string) ([]*model.Blueprint, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterCorporationHistoryByName is the resolver for the character_corporationHistoryByName field.
func (r *queryResolver) CharacterCorporationHistoryByName(ctx context.Context, name *string) ([]*model.CorporationHistory, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterFatiqueByName is the resolver for the character_fatiqueByName field.
func (r *queryResolver) CharacterFatiqueByName(ctx context.Context, name *string) (*model.Fatique, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterMedalsByName is the resolver for the character_medalsByName field.
func (r *queryResolver) CharacterMedalsByName(ctx context.Context, name *string) ([]*model.Medal, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterNotificationsByName is the resolver for the character_notificationsByName field.
func (r *queryResolver) CharacterNotificationsByName(ctx context.Context, name *string) ([]*model.Notification, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterPortraitByName is the resolver for the character_portraitByName field.
func (r *queryResolver) CharacterPortraitByName(ctx context.Context, name *string) (*model.CharacterPortrait, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterRoleByName is the resolver for the character_roleByName field.
func (r *queryResolver) CharacterRoleByName(ctx context.Context, name *string) (*model.Role, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterStandingsByName is the resolver for the character_standingsByName field.
func (r *queryResolver) CharacterStandingsByName(ctx context.Context, name *string) ([]*model.Standing, error) {
	panic(fmt.Errorf("not implemented"))
}

// CharacterTitlesByName is the resolver for the character_titlesByName field.
func (r *queryResolver) CharacterTitlesByName(ctx context.Context, name *string) ([]*model.Title, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClonesByID is the resolver for the clones_byID field.
func (r *queryResolver) ClonesByID(ctx context.Context, id *int) ([]*model.Clone, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClonesImplantsByID is the resolver for the clones_implantsByID field.
func (r *queryResolver) ClonesImplantsByID(ctx context.Context, id *int) ([]*model.Implant, error) {
	panic(fmt.Errorf("not implemented"))
}

// OrdersForRegion is the resolver for the ordersForRegion field.
func (r *queryResolver) OrdersForRegion(ctx context.Context, regionID int, orderType model.Ordertype, typeID *int, page int) (*model.OrderWrapper, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ResolverOrdersForRegion")
	defer span.End()
	return market.OrdersForRegion(newCtx, &regionID, &orderType, typeID, &page)
}

// OrdersForRegionByName is the resolver for the ordersForRegionByName field.
func (r *queryResolver) OrdersForRegionByName(ctx context.Context, region string, orderType model.Ordertype, typeName *string, page int) (*model.OrderWrapper, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "ResolverOrdersForRegionByName")
	defer span.End()
	return market.OrdersForRegionByName(newCtx, &region, &orderType, typeName, &page)
}

// SkillsCharacterAttributesByID is the resolver for the skills_characterAttributesByID field.
func (r *queryResolver) SkillsCharacterAttributesByID(ctx context.Context, characterID int) (*model.Attributes, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SkillsCharacterAttributesByID")
	defer span.End()
	return eve_character.AttributesByID(newCtx, &characterID)
}

// SkillsCharacterAttributesByName is the resolver for the skills_characterAttributesByName field.
func (r *queryResolver) SkillsCharacterAttributesByName(ctx context.Context, character string) (*model.Attributes, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SkillsCharacterAttributesByID")
	defer span.End()
	return eve_character.AttributesByName(newCtx, &character)
}

// SkillsQueueByCharacterID is the resolver for the skills_QueueByCharacterID field.
func (r *queryResolver) SkillsQueueByCharacterID(ctx context.Context, characterID int) ([]*model.SkillQueueItem, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SkillsQueueByCharacterID")
	defer span.End()
	return eve_character.SkillQueueByID(newCtx, &characterID)
}

// SkillsQueueByCharacterName is the resolver for the skills_QueueByCharacterName field.
func (r *queryResolver) SkillsQueueByCharacterName(ctx context.Context, character string) ([]*model.SkillQueueItem, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SkillsQueueByCharacterName")
	defer span.End()
	return eve_character.SkillQueueByName(newCtx, &character)
}

// SkillsByCharacterID is the resolver for the skills_ByCharacterID field.
func (r *queryResolver) SkillsByCharacterID(ctx context.Context, characterID int) (*model.SkillWrapper, error) {
	panic(fmt.Errorf("not implemented"))
}

// SkillsByCharacterName is the resolver for the skills_ByCharacterName field.
func (r *queryResolver) SkillsByCharacterName(ctx context.Context, character string) (*model.SkillWrapper, error) {
	panic(fmt.Errorf("not implemented"))
}

// SystemByID is the resolver for the systemById field.
func (r *queryResolver) SystemByID(ctx context.Context, id *int) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SystemByID")
	defer span.End()
	return universe.SystemByID(newCtx, id)
}

// StationByID is the resolver for the stationById field.
func (r *queryResolver) StationByID(ctx context.Context, id *int) (*model.Station, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StationByID")
	defer span.End()
	return universe.StationByID(newCtx, id)
}

// PlanetByID is the resolver for the planetById field.
func (r *queryResolver) PlanetByID(ctx context.Context, id *int) (*model.Planet, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "PlanetByID")
	defer span.End()
	return universe.PlanetByID(newCtx, id)
}

// CorporationByID is the resolver for the corporationById field.
func (r *queryResolver) CorporationByID(ctx context.Context, id *int) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CorporationByID")
	defer span.End()
	return corporation.ByID(newCtx, id)
}

// CorporationHistoryForCharacterID is the resolver for the corporationHistoryForCharacterId field.
func (r *queryResolver) CorporationHistoryForCharacterID(ctx context.Context, id *int) ([]*model.CorporationHistory, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "CorporationHistoryForCharacterID")
	defer span.End()
	return eve_character.CorporationHistory(newCtx, id)
}

// FactionByID is the resolver for the factionByID field.
func (r *queryResolver) FactionByID(ctx context.Context, id *int) (*model.Faction, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "FactionByID")
	defer span.End()
	return universe.FactionByID(newCtx, id)
}

// OrderHistory is the resolver for the orderHistory field.
func (r *queryResolver) OrderHistory(ctx context.Context, regionID *int, typeID *int) ([]*model.OrderHistory, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "OrderHistory")
	defer span.End()
	return market.OrderHistory(newCtx, regionID, typeID)
}

// ConstellationList is the resolver for the constellation_list field.
func (r *regionResolver) ConstellationList(ctx context.Context, obj *model.Region) ([]*model.Constellation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "RegionConstellationList")
	defer span.End()
	return universe.ConstellationsByIDs(newCtx, obj.Constellations)
}

// ItemType is the resolver for the item_type field.
func (r *skillQueueItemResolver) ItemType(ctx context.Context, obj *model.SkillQueueItem) (*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SkillQueueItemType")
	defer span.End()
	return universe.ItemTypeByID(newCtx, obj.SkillID)
}

// SolarSystem is the resolver for the solar_system field.
func (r *starResolver) SolarSystem(ctx context.Context, obj *model.Star) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StarSolarSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SolarSystemID)
}

// ItemType is the resolver for the item_type field.
func (r *starResolver) ItemType(ctx context.Context, obj *model.Star) (*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StarItemType")
	defer span.End()
	return universe.ItemTypeByID(newCtx, obj.TypeID)
}

// ItemType is the resolver for the item_type field.
func (r *stargateResolver) ItemType(ctx context.Context, obj *model.Stargate) (*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StragateItemType")
	defer span.End()
	return universe.ItemTypeByID(newCtx, obj.TypeID)
}

// Stargate is the resolver for the stargate field.
func (r *stargateDestinationResolver) Stargate(ctx context.Context, obj *model.StargateDestination) (*model.Stargate, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StargateDestinationStargate")
	defer span.End()
	return universe.StargateByID(newCtx, obj.StargateID)
}

// System is the resolver for the system field.
func (r *stargateDestinationResolver) System(ctx context.Context, obj *model.StargateDestination) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StargateDestinationSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SystemID)
}

// OwningCorporation is the resolver for the owning_corporation field.
func (r *stationResolver) OwningCorporation(ctx context.Context, obj *model.Station) (*model.Corporation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "OwningCorporation")
	defer span.End()
	return corporation.ByID(newCtx, obj.Owner)
}

// Race is the resolver for the race field.
func (r *stationResolver) Race(ctx context.Context, obj *model.Station) (*model.Race, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StationRace")
	defer span.End()
	return universe.RaceByID(newCtx, obj.RaceID)
}

// System is the resolver for the system field.
func (r *stationResolver) System(ctx context.Context, obj *model.Station) (*model.System, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StationSystem")
	defer span.End()
	return universe.SystemByID(newCtx, obj.SystemID)
}

// StationType is the resolver for the station_type field.
func (r *stationResolver) StationType(ctx context.Context, obj *model.Station) (*model.ItemType, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "StationStationType")
	defer span.End()
	return universe.ItemTypeByID(newCtx, obj.TypeID)
}

// Constellation is the resolver for the constellation field.
func (r *systemResolver) Constellation(ctx context.Context, obj *model.System) (*model.Constellation, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SystemConstellation")
	defer span.End()
	return universe.ConstellationByID(newCtx, obj.ConstellationID)
}

// Star is the resolver for the star field.
func (r *systemResolver) Star(ctx context.Context, obj *model.System) (*model.Star, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SystemStar")
	defer span.End()
	return universe.StarByID(newCtx, obj.StarID)
}

// StargateList is the resolver for the stargate_list field.
func (r *systemResolver) StargateList(ctx context.Context, obj *model.System) ([]*model.Stargate, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SystemStargateList")
	defer span.End()
	return universe.StargateDetails(newCtx, obj.Stargates)
}

// StationList is the resolver for the station_list field.
func (r *systemResolver) StationList(ctx context.Context, obj *model.System) ([]*model.Station, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SystemStationList")
	defer span.End()
	return universe.StationsByIDs(newCtx, obj.Stations)
}

// AsteroidBeltsProperties is the resolver for the asteroid_belts_properties field.
func (r *system_planetResolver) AsteroidBeltsProperties(ctx context.Context, obj *model.SystemPlanet) ([]*model.AsteroidBelt, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "AsteroidBeltsProperties")
	defer span.End()
	return universe.AsteroidBeltDetails(newCtx, obj.AsteroidBelts)
}

// MoonDetails is the resolver for the moon_details field.
func (r *system_planetResolver) MoonDetails(ctx context.Context, obj *model.SystemPlanet) ([]*model.Moon, error) {
	newCtx, span := tracing.TraceProvider.Tracer(tracerName).Start(ctx, "SystemPlanetMoonDetails")
	defer span.End()
	return universe.MoonDetails(newCtx, obj.Moons)
}

// PlanetProperties is the resolver for the planet_properties field.
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

// SkillQueueItem returns generated.SkillQueueItemResolver implementation.
func (r *Resolver) SkillQueueItem() generated.SkillQueueItemResolver {
	return &skillQueueItemResolver{r}
}

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
type skillQueueItemResolver struct{ *Resolver }
type starResolver struct{ *Resolver }
type stargateResolver struct{ *Resolver }
type stargateDestinationResolver struct{ *Resolver }
type stationResolver struct{ *Resolver }
type systemResolver struct{ *Resolver }
type system_planetResolver struct{ *Resolver }
