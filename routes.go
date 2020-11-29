package main

import (
	"encoding/json"
	"errors"
	"fmt"
	xj "github.com/basgys/goxml2json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

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

	routesPath := fmt.Sprintf("http://api.thebus.org/route/?key=%v&%v=%v", apiKey, routeType, url.QueryEscape(s))

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
