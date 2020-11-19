package main

import (
	"fmt"
	"net/http"
)

const (
	API_KEY = ""
)

func fetchArrivals(s string) {
	arrivalsPath := fmt.Sprintf("http://api.thebus.org/arrivals/?key=%v&stop=%v", API_KEY, s)
	fmt.Printf(arrivalsPath)
	resp, err := http.Get(arrivalsPath)
	if err != nil {
	}
	fmt.Print(resp)
}
