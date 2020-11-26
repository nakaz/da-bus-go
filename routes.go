package main

import (
	"encoding/json"
	"errors"
	"fmt"
	xj "github.com/basgys/goxml2json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Route struct {
	Headsign  string `json:"headsign"`
	RouteNum  string `json:"routeNum"`
	ShapeID   string `json:"shapeID"`
	FirstStop string `json:"firstStop"`
}

func (_ *Resolver) Route(args struct {
	Route string
}) []*Route {
	resp, err := fetchRoutes(args.Route, "route")
	if err != nil {
		log.Print(err)
	}
	return resp.Routes.Route
}

func (_ *Resolver) Headsign(args struct {
	Headsign string
}) []*Route {
	resp, err := fetchRoutes(args.Headsign, "headsign")
	if err != nil {
		log.Print(err)
	}
	return resp.Routes.Route
}

type RoutesBody struct {
	Routes struct {
		RouteName string   `json:"routeName"`
		RouteID   string   `json:"routeID"`
		Route     []*Route `json:"route"`
	} `json:"routes"`
}

func fetchRoutes(s string, t string) (*RoutesBody, error) {
	routeType := ""
	switch t {
	case "route":
		routeType = t
	case "headsign":
		routeType = t
	default:
		err := errors.New("Invalid route type")
		return nil, err
	}

	routesPath := fmt.Sprintf("http://api.thebus.org/route/?key=%v&%v=%v", apiKey, routeType, s)
	resp, err := http.Get(routesPath)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}

	respString := string(respData)
	xmlString := strings.NewReader(respString)

	routes := &RoutesBody{}

	jsonResp, err := xj.Convert(xmlString)
	if err != nil {
		log.Print(err)
	}

	if err := json.Unmarshal(jsonResp.Bytes(), routes); err != nil {
		log.Print(err)
	}

	return routes, nil
}
