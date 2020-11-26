package main

import (
	"log"
)

func (_ *Resolver) Arrivals(args struct {
	Stop string
}) []*Arrival {
	resp, err := fetchArrivals(args.Stop)
	if err != nil {
		log.Print(err)
	}
	return resp.StopTimes.Arrival
}
