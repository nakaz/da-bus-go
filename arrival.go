package main

import (
	"log"
	"strconv"
)

type Arrival struct {
	Headsign  string `json:"headsign"`
	Vehicle   string `json:"vehicle"`
	StopTime  string `json:"stopTime"`
	Estimated string `json:"estimated"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Shape     string `json:"shape"`
	ID        string `json:"id"`
	Trip      string `json:"trip"`
	Canceled  string `json:"canceled"`
	Date      string `json:"date"`
	Route     string `json:"route"`
	Direction string `json:"direction"`
}

func (a *Arrival) VehicleID() int32 {
	i, err := strconv.Atoi(a.Vehicle)
	if err != nil {
		// TODO: Handle '???' values. Currently defaults to 0.
		log.Print(err)
	}
	return int32(i)
}

func (a *Arrival) LatLng() []float64 {
	lat, err := strconv.ParseFloat(a.Latitude, 64)
	if err != nil {
		log.Print(err)
	}
	lng, err := strconv.ParseFloat(a.Longitude, 64)
	if err != nil {
		log.Print(err)
	}
	return []float64{float64(lat), float64(lng)}
}
