package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	dao "github.com/cryanbrow/eve-graphql-go/graph/data_access"
	"github.com/cryanbrow/eve-graphql-go/graph/generated"
	"github.com/cryanbrow/eve-graphql-go/graph/model"
)

func (r *asteroid_beltResolver) System(ctx context.Context, obj *model.AsteroidBelt) (*model.System, error) {
	return dao.SystemByID(obj.SystemID)
}

func (r *item_typeResolver) MarketGroup(ctx context.Context, obj *model.ItemType) (*model.MarketGroup, error) {
	return dao.MarketGroupByID(obj.MarketGroupID)
}

func (r *market_groupResolver) ParentGroup(ctx context.Context, obj *model.MarketGroup) (*model.Group, error) {

	panic(fmt.Errorf("not implemented"))
}

func (r *market_groupResolver) TypesDetails(ctx context.Context, obj *model.MarketGroup) ([]*model.ItemType, error) {
	return dao.ItemTypesByIDs(obj.Types)
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

func (r *queryResolver) OrdersForRegion(ctx context.Context, regionID *int, orderType *model.Ordertype, typeID *int) ([]*model.Order, error) {
	return dao.OrdersForRegion(regionID, orderType, typeID)
}

func (r *queryResolver) SystemByID(ctx context.Context, id *int) (*model.System, error) {
	return dao.SystemByID(id)
}

func (r *queryResolver) StationByID(ctx context.Context, id *int) (*model.Station, error) {
	return dao.StationByID(id)
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

// Item_type returns generated.Item_typeResolver implementation.
func (r *Resolver) Item_type() generated.Item_typeResolver { return &item_typeResolver{r} }

// Market_group returns generated.Market_groupResolver implementation.
func (r *Resolver) Market_group() generated.Market_groupResolver { return &market_groupResolver{r} }

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// System_planet returns generated.System_planetResolver implementation.
func (r *Resolver) System_planet() generated.System_planetResolver { return &system_planetResolver{r} }

type asteroid_beltResolver struct{ *Resolver }
type item_typeResolver struct{ *Resolver }
type market_groupResolver struct{ *Resolver }
type orderResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type system_planetResolver struct{ *Resolver }
