package main

import (
	"log"
)

type Arrivals struct {
	StopTimes struct {
		Timestamp string     `json:"timestamp"`
		Arrival   []*Arrival `json:"arrival"`
		Stop      string     `json:"stop"`
	} `json:"stopTimes"`
}

func (_ *Resolver) Arrivals(args struct {
	Stop string
}) []*Arrival {
	resp, err := fetchArrivals(args.Stop)
	if err != nil {
		log.Print(err)
	}
	return resp.StopTimes.Arrival
}
