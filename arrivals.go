package main

import (
	"log"
)

type ArrivalResolver struct {
	arrival *Arrival
}

func (a *ArrivalResolver) Headsign() *string {
	return &a.arrival.Headsign
}

func (a *ArrivalResolver) Vehicle() *string {
	return &a.arrival.Vehicle
}

func (a *ArrivalResolver) StopTime() *string {
	return &a.arrival.StopTime
}

func (a *ArrivalResolver) Estimated() *string {
	return &a.arrival.Estimated
}

func (a *ArrivalResolver) Latitude() *string {
	return &a.arrival.Latitude
}

func (a *ArrivalResolver) Longitude() *string {
	return &a.arrival.Longitude
}

func (a *ArrivalResolver) Shape() *string {
	return &a.arrival.Shape
}

func (a *ArrivalResolver) ID() *string {
	return &a.arrival.ID
}

func (a *ArrivalResolver) Trip() *string {
	return &a.arrival.Trip
}

func (a *ArrivalResolver) Canceled() *string {
	return &a.arrival.Canceled
}

func (a *ArrivalResolver) Date() *string {
	return &a.arrival.Date
}

func (a *ArrivalResolver) Route() *string {
	return &a.arrival.Route
}

func (a *ArrivalResolver) Direction() *string {
	return &a.arrival.Direction
}

func (_ *Resolver) Arrivals(args struct {
	Stop string
}) []*ArrivalResolver {
	resp, err := fetchArrivals(args.Stop)
	if err != nil {
		log.Print(err)
	}
	arrivals := resp.StopTimes.Arrival

	ar := make([]*ArrivalResolver, len(arrivals))
	for i, arrival := range arrivals {
		ar[i] = &ArrivalResolver{
			arrival: arrival,
		}
	}
	return ar
}
