package main

import (
	"encoding/json"
	"fmt"
	xj "github.com/basgys/goxml2json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Vehicle struct {
	Driver         string `json:"driver"`
	Longitude      string `json:"longitude"`
	LastMessage    string `json:"last_message"`
	Headsign       string `json:"headsign"`
	RouteShortName string `json:"route_short_name"`
	Number         string `json:"number"`
	Trip           string `json:"trip"`
	Latitude       string `json:"latitude"`
	Adherence      string `json:"adherence"`
}

// Resolver
func (_ *Resolver) Vehicle(args struct {
	Num string
}) Vehicle {
	resp, err := fetchVehicle(args.Num)
	if err != nil {
		log.Print(err)
	}
	return resp.Vehicles.Vehicle
}

// API
type VehicleBody struct {
	Vehicles struct {
		Timestamp string  `json:"timestamp"`
		Vehicle   Vehicle `json:"vehicle"`
	} `json:"vehicles"`
}

func fetchVehicle(n string) (*VehicleBody, error) {
	vehiclePath := fmt.Sprintf("http://api.thebus.org/vehicle/?key=%v&num=%v", apiKey, n)
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

	vehicle := &VehicleBody{}

	jsonResp, err := xj.Convert(xmlString)
	if err != nil {
		log.Print(err)
	}

	if err := json.Unmarshal(jsonResp.Bytes(), vehicle); err != nil {
		log.Print(err)
	}

	return vehicle, nil
}
