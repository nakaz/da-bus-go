package main

type Arrival struct {
	Headsign  string `json:"headsign"`
	Vehicle   string `json:"vehicle"`
	StopTime  string `json:"stopTime"`
	Estimated string `json:"estimated"`
	Latitude  string `json:"latitude"`
	Shape     string `json:"shape"`
	ID        string `json:"id"`
	Trip      string `json:"trip"`
	Canceled  string `json:"canceled"`
	Date      string `json:"date"`
	Longitude string `json:"longitude"`
	Route     string `json:"route"`
	Direction string `json:"direction"`
}
