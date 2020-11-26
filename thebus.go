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

type Vehicle struct {
	Vehicles struct {
		Timestamp string `json:"timestamp"`
		Vehicle   struct {
			Driver         string `json:"driver"`
			Longitude      string `json:"longitude"`
			LastMessage    string `json:"last_message"`
			Headsign       string `json:"headsign"`
			RouteShortName string `json:"route_short_name"`
			Number         string `json:"number"`
			Trip           string `json:"trip"`
			Latitude       string `json:"latitude"`
			Adherence      string `json:"adherence"`
		} `json:"vehicle"`
	} `json:"vehicles"`
}

type Routes struct {
	Routes struct {
		RouteName string `json:"routeName"`
		RouteID   string `json:"routeID"`
		Route     []struct {
			Headsign  string `json:"headsign"`
			RouteNum  string `json:"routeNum"`
			ShapeID   string `json:"shapeID"`
			FirstStop string `json:"firstStop"`
		} `json:"route"`
	} `json:"routes"`
}

func fetchArrivals(s string) (*Arrivals, error) {
	arrivalsPath := fmt.Sprintf("http://api.thebus.org/arrivals/?key=%v&stop=%v", apiKey, s)
	resp, err := http.Get(arrivalsPath)
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

	arrivals := &Arrivals{}

	jsonResp, err := xj.Convert(xmlString)
	if err != nil {
		log.Print(err)
	}

	if err := json.Unmarshal(jsonResp.Bytes(), arrivals); err != nil {
		log.Print(err)
	}

	return arrivals, nil
}

func fetchVehicle(s string) (*Vehicle, error) {
	vehiclePath := fmt.Sprintf("http://api.thebus.org/vehicle/?key=%v&num=%v", apiKey, s)
	resp, err := http.Get(vehiclePath)
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

	vehicle := &Vehicle{}

	jsonResp, err := xj.Convert(xmlString)
	if err != nil {
		log.Print(err)
	}

	if err := json.Unmarshal(jsonResp.Bytes(), vehicle); err != nil {
		log.Print(err)
	}

	return vehicle, nil
}

func fetchRoutes(s string, t string) (*Routes, error) {
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

	routes := &Routes{}

	jsonResp, err := xj.Convert(xmlString)
	if err != nil {
		log.Print(err)
	}

	if err := json.Unmarshal(jsonResp.Bytes(), routes); err != nil {
		log.Print(err)
	}

	return routes, nil
}
