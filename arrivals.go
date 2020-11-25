package main

import (
	// "fmt"
	"log"
)

type Arrival struct {
	Headsign string
}

type ArrivalResolver struct {
	arrival Arrival
}

func (a *ArrivalResolver) Headsign() *string {
	return &a.arrival.Headsign
}

func (_ *Resolver) Arrivals(args struct {
	Stop string
}) *[]*ArrivalResolver {
	resp, err := fetchArrivals(args.Stop)
	if err != nil {
		log.Print(err)
	}
	arrivals := resp.StopTimes.Arrival

	ar := make([]*ArrivalResolver, len(arrivals))
	al := &ar
	for i, v := range arrivals {
		item := &Arrival{
			Headsign: v.Headsign,
		}
		ar[i] = &ArrivalResolver{
			arrival: *item,
		}
	}
	return al
}
