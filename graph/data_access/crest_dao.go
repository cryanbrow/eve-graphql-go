package data_access

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	cache "github.com/cryanbrow/eve-graphql-go/graph/caching"
	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	log "github.com/sirupsen/logrus"

	"github.com/cryanbrow/eve-graphql-go/graph/model"
)

func OrdersForRegion(regionID *int, orderType *model.Ordertype, typeID *int, page *int) ([]*model.Order, error) {
	log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID, "orderType": orderType}).Info("OrdersForRegion Called")
	orders := make([]*model.Order, 0)
	base_url := fmt.Sprintf("%s/markets/%s/orders/", baseUriESI, strconv.Itoa(*regionID))

	redis_key := "OrdersForRegion:" + strconv.Itoa(*regionID) + ":" + orderType.String()

	query_params := make([]configuration.Key_value, 2)
	kv := new(configuration.Key_value)
	kv.Key = "page"
	kv.Value = strconv.Itoa(*page)
	query_params = append(query_params, *kv)

	if typeID != nil {
		redis_key = redis_key + ":" + strconv.Itoa(*typeID)
		kv.Key = "type_id"
		kv.Value = strconv.Itoa(*typeID)
		query_params = append(query_params, *kv)
	}

	redis_key = redis_key + ":" + strconv.Itoa(*page)

	orderResult, _, err := ordersForRegionREST(base_url, query_params, redis_key)

	if err == nil {
		orders = append(orders, orderResult...)
	} else {
		log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID, "orderType": orderType}).Errorf("First page query for Orders has error : %v", err)
		println(err)
	}

	return orders, nil
}

func OrdersForRegionByName(region *string, orderType *model.Ordertype, typeName *string, page *int) ([]*model.Order, error) {
	regionID, err := idForName(region, model.REGIONS)
	if err != nil {
		return nil, errors.New("unknown name for region")
	}
	typeID, err := idForName(typeName, model.INVENTORY_TYPES)
	if err != nil {
		return nil, errors.New("unknown name for typeName")
	}
	orders, err := OrdersForRegion(&regionID, orderType, &typeID, page)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func OrderHistory(regionID *int, typeID *int) ([]*model.OrderHistory, error) {
	if regionID == nil || typeID == nil {
		return nil, errors.New("nil id")
	}
	var orderHistory []*model.OrderHistory = make([]*model.OrderHistory, 0)
	base_url := fmt.Sprintf("%s/markets/%s/history", baseUriESI, strconv.Itoa(*regionID))
	redis_key := "OrderHistoryByID:" + strconv.Itoa(*regionID) + ":" + strconv.Itoa(*typeID)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return orderHistory, err
	}

	if err := json.Unmarshal(responseBytes, &orderHistory); err != nil {
		log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return orderHistory, err
	}

	return orderHistory, nil
}

func ordersForRegionREST(url string, additional_query_params []configuration.Key_value, redis_key string) ([]*model.Order, int, error) {
	var orders []*model.Order
	var pages = 0
	var buffer bytes.Buffer
	responseBytes, header, err := makeRESTCall(url, http.MethodGet, buffer, additional_query_params, redis_key)
	if err != nil {
		return orders, 0, err
	}

	pages, _ = strconv.Atoi(header.Get("x-pages"))

	if err := json.Unmarshal(responseBytes, &orders); err != nil {
		log.WithFields(log.Fields{"url": url}).Errorf("Could not unmarshal reponseBytes. : %v", err.Error())
		return orders, 0, err
	}

	return orders, pages, nil
}

func SystemByArray(ids []*int) ([]*model.System, error) {
	systemDetails := make([]*model.System, 0)
	for _, element := range ids {
		system, err := SystemByID(element)
		if err == nil {
			systemDetails = append(systemDetails, system)
		} else {
			return nil, err
		}
	}
	return systemDetails, nil
}

func SystemByID(id *int) (*model.System, error) {
	if id == nil {
		return nil, errors.New("nil id")
	}
	var system *model.System = new(model.System)
	base_url := fmt.Sprintf("%s/universe/systems/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "SystemByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return system, err
	}

	if err := json.Unmarshal(responseBytes, &system); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return system, err
	}

	return system, nil
}

func StationByArray(ids []*int) ([]*model.Station, error) {
	stationDetails := make([]*model.Station, 0)
	for _, element := range ids {
		station, err := StationByID(element)
		if err == nil {
			stationDetails = append(stationDetails, station)
		} else {
			return nil, err
		}
	}
	return stationDetails, nil
}

func StationByID(id *int) (*model.Station, error) {
	if id == nil {
		return nil, errors.New("nil id")
	}
	if *id > 2147483647 {
		return nil, nil
	}
	var station *model.Station = new(model.Station)
	base_url := fmt.Sprintf("%s/universe/stations/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "StationByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return station, err
	}

	if err := json.Unmarshal(responseBytes, &station); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return station, err
	}

	return station, nil
}

func CorporationByID(id *int) (*model.Corporation, error) {
	var corporation *model.Corporation = new(model.Corporation)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/corporations/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "CorporationByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return corporation, err
	}

	if err := json.Unmarshal(responseBytes, &corporation); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return corporation, err
	}

	return corporation, nil
}

func AllianceByID(id *int) (*model.Alliance, error) {
	var alliance *model.Alliance = new(model.Alliance)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/alliances/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "AllianceByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return alliance, err
	}

	if err := json.Unmarshal(responseBytes, &alliance); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return alliance, err
	}

	return alliance, nil
}

func CharacterByID(id *int) (*model.Character, error) {
	var character *model.Character = new(model.Character)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/characters/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "CharacterByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return character, err
	}

	if err := json.Unmarshal(responseBytes, &character); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return character, err
	}

	return character, nil
}

func PlanetByID(id *int) (*model.Planet, error) {
	var planet *model.Planet = new(model.Planet)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/planets/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "PlanetByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return planet, err
	}

	if err := json.Unmarshal(responseBytes, &planet); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return planet, err
	}

	return planet, nil
}

func StargateDetails(stargates []*int) ([]*model.Stargate, error) {
	stargateDetails := make([]*model.Stargate, 0)
	for _, element := range stargates {
		stargate, err := StargateByID(element)
		if err == nil {
			stargateDetails = append(stargateDetails, stargate)
		} else {
			return nil, err
		}
	}
	return stargateDetails, nil
}

func StargateByID(id *int) (*model.Stargate, error) {
	var stargate *model.Stargate = new(model.Stargate)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/stargates/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "StargateByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return stargate, err
	}

	if err := json.Unmarshal(responseBytes, &stargate); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return stargate, err
	}

	return stargate, nil
}

func MoonDetails(moons []*int) ([]*model.Moon, error) {
	moonDetails := make([]*model.Moon, 0)
	for _, element := range moons {
		moon, err := MoonByID(element)
		if err == nil {
			moonDetails = append(moonDetails, moon)
		} else {
			return nil, err
		}
	}
	return moonDetails, nil
}

func MoonByID(id *int) (*model.Moon, error) {
	var moon *model.Moon = new(model.Moon)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/moons/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "MoonByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return moon, err
	}

	if err := json.Unmarshal(responseBytes, &moon); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return moon, err
	}

	return moon, nil
}

func ItemTypesByIDs(itemTypes []*int) ([]*model.ItemType, error) {
	itemTypeDetails := make([]*model.ItemType, 0)
	for _, element := range itemTypes {
		itemType, err := ItemTypeByID(element)
		if err == nil {
			itemTypeDetails = append(itemTypeDetails, itemType)
		} else {
			return nil, err
		}
	}
	return itemTypeDetails, nil
}

func ItemTypeByID(id *int) (*model.ItemType, error) {
	var itemType *model.ItemType = new(model.ItemType)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/types/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "ItemTypeByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return itemType, err
	}

	if err := json.Unmarshal(responseBytes, &itemType); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return itemType, err
	}

	return itemType, nil
}

func AsteroidBeltDetails(asteroidBelts []*int) ([]*model.AsteroidBelt, error) {
	asteroidBeltDetails := make([]*model.AsteroidBelt, 0)
	for _, element := range asteroidBelts {
		asteroidBelt, err := AsteroidBeltByID(element)
		if err == nil {
			asteroidBeltDetails = append(asteroidBeltDetails, asteroidBelt)
		} else {
			return nil, err
		}
	}
	return asteroidBeltDetails, nil
}

func AsteroidBeltByID(id *int) (*model.AsteroidBelt, error) {
	var asteroidBelt *model.AsteroidBelt = new(model.AsteroidBelt)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/asteroid_belts/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "AsteroidBeltByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return asteroidBelt, err
	}

	if err := json.Unmarshal(responseBytes, &asteroidBelt); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return asteroidBelt, err
	}

	return asteroidBelt, nil
}

func MarketGroupByID(id *int) (*model.MarketGroup, error) {
	var marketGroup *model.MarketGroup = new(model.MarketGroup)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/markets/groups/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "MarketGroupByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return marketGroup, err
	}

	if err := json.Unmarshal(responseBytes, &marketGroup); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return marketGroup, err
	}

	return marketGroup, nil
}

func GroupByID(id *int) (*model.Group, error) {
	var group *model.Group = new(model.Group)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/groups/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "GroupByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return group, err
	}

	if err := json.Unmarshal(responseBytes, &group); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return group, err
	}

	return group, nil
}

func ConstellationsByIDs(ids []*int) ([]*model.Constellation, error) {
	constellationDetails := make([]*model.Constellation, 0)
	for _, element := range ids {
		constellation, err := ConstellationByID(element)
		if err == nil {
			constellationDetails = append(constellationDetails, constellation)
		} else {
			return nil, err
		}
	}
	return constellationDetails, nil
}

func ConstellationByID(id *int) (*model.Constellation, error) {
	var constellation *model.Constellation = new(model.Constellation)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/constellations/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "ConstellationByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return constellation, err
	}

	if err := json.Unmarshal(responseBytes, &constellation); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return constellation, err
	}

	return constellation, nil
}

func StarByID(id *int) (*model.Star, error) {
	var star *model.Star = new(model.Star)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/stars/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "StarByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return star, err
	}

	if err := json.Unmarshal(responseBytes, &star); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return star, err
	}

	return star, nil
}

func GraphicByID(id *int) (*model.Graphic, error) {
	var graphic *model.Graphic = new(model.Graphic)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/graphics/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "GraphicByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return graphic, err
	}

	if err := json.Unmarshal(responseBytes, &graphic); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return graphic, err
	}

	return graphic, nil
}

func DogmaAttributeByID(id *int) (*model.DogmaAttributeDetail, error) {
	var dogmaAttribute *model.DogmaAttributeDetail = new(model.DogmaAttributeDetail)
	if id == nil {
		return nil, nil
	}
	base_url := fmt.Sprintf("%s/dogma/attributes/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "DogmaAttributeByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return dogmaAttribute, err
	}

	if err := json.Unmarshal(responseBytes, &dogmaAttribute); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return dogmaAttribute, err
	}

	return dogmaAttribute, nil
}

func DogmaEffectByID(id *int) (*model.DogmaEffectDetail, error) {
	var dogmaEffect *model.DogmaEffectDetail = new(model.DogmaEffectDetail)
	if id == nil {
		return nil, nil
	}
	base_url := fmt.Sprintf("%s/dogma/effects/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "DogmaEffectByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return dogmaEffect, err
	}

	if err := json.Unmarshal(responseBytes, &dogmaEffect); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return dogmaEffect, err
	}

	return dogmaEffect, nil
}

func CategoryByID(id *int) (*model.Category, error) {
	var category *model.Category = new(model.Category)
	if id == nil {
		return nil, nil
	}
	base_url := fmt.Sprintf("%s/universe/categories/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "CategoryByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return category, err
	}

	if err := json.Unmarshal(responseBytes, &category); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return category, err
	}

	return category, nil
}

func RegionByID(id *int) (*model.Region, error) {
	var region *model.Region = new(model.Region)
	if id == nil {
		return nil, nil
	}
	base_url := fmt.Sprintf("%s/universe/regions/%s/", baseUriESI, strconv.Itoa(*id))
	redis_key := "RegionByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return region, err
	}

	if err := json.Unmarshal(responseBytes, &region); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return region, err
	}

	return region, nil
}

func FactionByID(id *int) (*model.Faction, error) {
	var faction *model.Faction = new(model.Faction)
	if id == nil {
		return nil, nil
	}

	inCache, result := cache.CheckRedisCache("FactionByID:" + strconv.Itoa(*id))
	if !inCache {
		faction, err := factionByArray(id)
		if err != nil {
			return nil, err
		} else {
			return faction, nil
		}
	} else {
		if err := json.Unmarshal(result, &faction); err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
			return faction, err
		} else {
			return faction, nil
		}
	}
}

func factionByArray(id *int) (*model.Faction, error) {
	var factions []*model.Faction = make([]*model.Faction, 0)
	var returnFaction *model.Faction
	base_url := fmt.Sprintf("%s/universe/factions/", baseUriESI)
	redis_key := "FactionByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, headers, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(responseBytes, &factions); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return nil, err
	}
	for _, faction := range factions {
		log.Info(*faction.Name)
		if *faction.FactionID == *id {
			returnFaction = faction
			log.Info("Found Faction ID")
		}
		factionBytes, err := json.Marshal(*faction)
		if err == nil {
			cache.AddToRedisCache("FactionByID:"+strconv.Itoa(*faction.FactionID), factionBytes, helpers.ESI_ttl_to_millis(headers.Get("expires")))
		} else {
			log.Errorf("Failure Marshalling: %v", err)
		}
	}
	return returnFaction, nil
}

func AncestryByID(id *int) (*model.Ancestry, error) {
	var ancestry *model.Ancestry = new(model.Ancestry)
	var err error
	if id == nil {
		return nil, nil
	}

	inCache, result := cache.CheckRedisCache("AncestryByID:" + strconv.Itoa(*id))
	if !inCache {
		ancestry, err = ancestryByArray(id)
		if err != nil {
			return nil, err
		} else {
			return ancestry, nil
		}
	} else {
		if err := json.Unmarshal(result, &ancestry); err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
			return ancestry, err
		} else {
			return ancestry, nil
		}
	}
}

func ancestryByArray(id *int) (*model.Ancestry, error) {
	var ancestries []*model.Ancestry = make([]*model.Ancestry, 0)
	var returnAncestry *model.Ancestry
	var redis_key = "AncestryByID:" + strconv.Itoa(*id)
	base_url := fmt.Sprintf("%s/universe/ancestries/", baseUriESI)

	var buffer bytes.Buffer
	responseBytes, headers, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(responseBytes, &ancestries); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return nil, err
	}
	for _, ancestry := range ancestries {
		log.Info(*ancestry.Name)
		if *ancestry.ID == *id {
			returnAncestry = ancestry
			log.Info("Found Ancestry ID")
		}
		ancestryBytes, err := json.Marshal(*ancestry)
		if err == nil {
			cache.AddToRedisCache("AncestryByID:"+strconv.Itoa(*ancestry.ID), ancestryBytes, helpers.ESI_ttl_to_millis(headers.Get("expires")))
		} else {
			log.Errorf("Failure Marshalling: %v", err)
		}
	}
	return returnAncestry, nil
}

func BloodlineByID(id *int) (*model.Bloodline, error) {
	var bloodline *model.Bloodline = new(model.Bloodline)
	var err error
	if id == nil {
		return nil, nil
	}

	inCache, result := cache.CheckRedisCache("BloodlineByID:" + strconv.Itoa(*id))
	if !inCache {
		bloodline, err = bloodlineByArray(id)
		if err != nil {
			return nil, err
		} else {
			return bloodline, nil
		}
	} else {
		if err := json.Unmarshal(result, &bloodline); err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
			return bloodline, err
		} else {
			return bloodline, nil
		}
	}
}

func bloodlineByArray(id *int) (*model.Bloodline, error) {
	var bloodlines []*model.Bloodline = make([]*model.Bloodline, 0)
	var returnBloodline *model.Bloodline
	base_url := fmt.Sprintf("%s/universe/bloodlines/", baseUriESI)
	redis_key := "BloodlineByID"

	var buffer bytes.Buffer
	responseBytes, headers, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(responseBytes, &bloodlines); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return nil, err
	}
	for _, bloodline := range bloodlines {
		log.Info(*bloodline.Name)
		if *bloodline.BloodlineID == *id {
			returnBloodline = bloodline
			log.Info("Found Bloodline ID")
		}
		bloodlineBytes, err := json.Marshal(*bloodline)
		if err == nil {
			cache.AddToRedisCache("BloodlineByID:"+strconv.Itoa(*bloodline.BloodlineID), bloodlineBytes, helpers.ESI_ttl_to_millis(headers.Get("expires")))
		} else {
			log.Errorf("Failure Marshalling: %v", err)
		}
	}
	return returnBloodline, nil
}

func RaceByID(id *int) (*model.Race, error) {
	var race *model.Race = new(model.Race)
	var err error
	if id == nil {
		return nil, nil
	}

	inCache, result := cache.CheckRedisCache("RaceByID:" + strconv.Itoa(*id))
	if !inCache {
		race, err = raceByArray(id)
		if err != nil {
			return nil, err
		} else {
			return race, nil
		}
	} else {
		if err := json.Unmarshal(result, &race); err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
			return race, err
		} else {
			return race, nil
		}
	}
}

func raceByArray(id *int) (*model.Race, error) {
	var races []*model.Race = make([]*model.Race, 0)
	var returnRace *model.Race
	var headers http.Header = nil
	base_url := fmt.Sprintf("%s/universe/races/", baseUriESI)
	redis_key := "RaceByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, headers, err := makeRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(responseBytes, &races); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return nil, err
	}
	for _, race := range races {
		log.Info(*race.Name)
		if *race.RaceID == *id {
			returnRace = race
			log.Info("Found Race ID")
		}
		raceBytes, err := json.Marshal(*race)
		if err == nil {
			cache.AddToRedisCache("RaceByID:"+strconv.Itoa(*race.RaceID), raceBytes, helpers.ESI_ttl_to_millis(headers.Get("expires")))
		} else {
			log.Errorf("Failure Marshalling: %v", err)
		}
	}
	return returnRace, nil
}

func idForName(name *string, name_type string) (int, error) {
	var ids *model.Names = new(model.Names)
	base_url := fmt.Sprintf("%s/universe/ids/", baseUriESI)
	if name == nil {
		return 0, errors.New("nil name")
	}
	redis_key := "IDForName:" + *name
	singleItemArray := []string{*name}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(singleItemArray)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	responseBytes, _, err := makeRESTCall(base_url, http.MethodPost, buf, nil, redis_key)
	if err != nil {
		return 0, err
	}

	if err := json.Unmarshal(responseBytes, &ids); err != nil {
		log.WithFields(log.Fields{"name": *name}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return 0, err
	}

	switch name_type {
	case model.AGENTS:
		return *ids.Agents[0].ID, nil
	case model.ALLIANCES:
		return *ids.Alliances[0].ID, nil
	case model.CHARACTERS:
		return *ids.Characters[0].ID, nil
	case model.CONSTELLATIONS:
		return *ids.Constellations[0].ID, nil
	case model.CORPORATIONS:
		return *ids.Corporations[0].ID, nil
	case model.FACTIONS:
		return *ids.Factions[0].ID, nil
	case model.INVENTORY_TYPES:
		return *ids.InventoryTypes[0].ID, nil
	case model.REGIONS:
		return *ids.Regions[0].ID, nil
	case model.SYSTEMS:
		return *ids.Systems[0].ID, nil
	default:
		return 0, errors.New("all fields nil")
	}

}

func makeRESTCall(base_url string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
	inCache, result := cache.CheckRedisCache(redis_query_key)
	if !inCache {
		crest_url, err := url.Parse(base_url)
		if err != nil {
			log.WithFields(log.Fields{"base_url": base_url, "verb": verb}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, nil, err
		}
		queryParameters := crest_url.Query()
		for _, kv := range configuration.AppConfig.Esi.Default.Query_params {
			queryParameters.Add(kv.Key, kv.Value)
		}
		for _, kv := range additional_query_params {
			queryParameters.Add(kv.Key, kv.Value)
		}

		crest_url.RawQuery = queryParameters.Encode()
		url := crest_url.String()

		log.WithFields(log.Fields{"url": url}).Info("Making REST Call")
		request, err := http.NewRequest(verb, url, &body)
		if err != nil {
			log.WithFields(log.Fields{"url": url}).Errorf("Could not build request. : %v", err)
		}
		response, err := Client.Do(request)
		if err != nil {
			log.WithFields(log.Fields{"url": url}).Errorf("Could not make request. : %v", err)
			return make([]byte, 0), nil, err
		}

		h := response.Header
		responseBytes, err := ioutil.ReadAll(response.Body)
		if response.StatusCode != 200 {
			log.WithFields(log.Fields{"url": url, "status_code": response.StatusCode}).Errorf("Received bad status code. : %v", err)
			return make([]byte, 0), nil, err
		}
		if err != nil {
			log.WithFields(log.Fields{"url": url}).Errorf("Could not read response for body. : %v", err)
			return make([]byte, 0), nil, err
		}
		cache.AddToRedisCache(redis_query_key, responseBytes, helpers.ESI_ttl_to_millis(h.Get("expires")))
		return responseBytes, h, nil
	}
	return result, nil, nil
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	baseUriESI string
	Client     HTTPClient
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	Client = &http.Client{}
	baseUriESI = configuration.AppConfig.Esi.Default.Url
}
