package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cryanbrow/eve-graphql-go/graph/generated"
	"github.com/cryanbrow/eve-graphql-go/graph/model"
)

func (r *queryResolver) OrdersForRegion(ctx context.Context, regionID *int, orderType *model.Ordertype, typeID *int) ([]*model.Order, error) {
	var duration int = 1
	var isBuyOrder bool = true
	var float float64 = 0.0
	order := &model.Order{
		Duration:     &duration,
		IsBuyOrder:   &isBuyOrder,
		Issued:       nil,
		Location:     nil,
		MinVolume:    &duration,
		OrderID:      nil,
		Price:        &float,
		Range:        nil,
		System:       nil,
		ItemType:     nil,
		VolumeRemain: &duration,
		VolumeTotal:  &duration,
	}

	var orders []*model.Order

	orders = append(orders, order)
	return orders, nil
}

func (r *queryResolver) SystemByID(ctx context.Context, id *int) (*model.System, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) StationByID(ctx context.Context, id *int) (*model.Station, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
