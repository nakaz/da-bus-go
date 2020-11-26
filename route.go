package main

import (
	"log"
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
