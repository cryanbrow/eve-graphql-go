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

func (r *orderResolver) Location(ctx context.Context, obj *model.Order) (*model.Station, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *orderResolver) System(ctx context.Context, obj *model.Order) (*model.System, error) {
	return dao.SystemByID(obj.SystemID)
}

func (r *queryResolver) OrdersForRegion(ctx context.Context, regionID *int, orderType *model.Ordertype, typeID *int) ([]*model.Order, error) {
	return dao.OrdersForRegion(regionID, orderType, typeID)
}

func (r *queryResolver) SystemByID(ctx context.Context, id *int) (*model.System, error) {
	return dao.SystemByID(id)
}

func (r *queryResolver) StationByID(ctx context.Context, id *int) (*model.Station, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *system_planetResolver) MoonDetails(ctx context.Context, obj *model.SystemPlanet) ([]*model.Moon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *system_planetResolver) PlanetProperties(ctx context.Context, obj *model.SystemPlanet) (*model.Planet, error) {
	return dao.PlanetByID(obj.PlanetID)
}

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// System_planet returns generated.System_planetResolver implementation.
func (r *Resolver) System_planet() generated.System_planetResolver { return &system_planetResolver{r} }

type orderResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type system_planetResolver struct{ *Resolver }
