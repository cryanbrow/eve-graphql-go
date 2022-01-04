package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/cryanbrow/eve-graphql-go/graph/generated"
	"github.com/cryanbrow/eve-graphql-go/graph/model"
)

func (r *queryResolver) OrdersForRegion(ctx context.Context, regionID *int, orderType *model.Ordertype, typeID *int) ([]*model.Order, error) {
	println(*regionID)
	url := fmt.Sprintf("https://esi.evetech.net/latest/markets/%s/orders/?datasource=tranquility&order_type=%s&page=%s&type_id=%s", strconv.Itoa(*regionID), *orderType, "1", strconv.Itoa(*typeID))
	request, err := http.NewRequest(http.MethodGet, url, nil)

	var orders []*model.Order
	if err != nil {
		log.Printf("Could not request orders by region. %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Cache-Control", "no-cache")

	response, err := Client.Do(request)
	if err != nil {
		log.Printf("Could not make request. %v", err)
		return orders, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response for body. %v", err)
		return orders, err
	}

	if err := json.Unmarshal(responseBytes, &orders); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		return orders, err
	}

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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) OrdersForregion(ctx context.Context, regionID *int, orderType *model.Ordertype, typeID *int) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}
