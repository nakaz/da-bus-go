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

// Resolver
func (_ *Resolver) Arrivals(args struct {
	Stop int32
}) []*Arrival {
	resp, err := fetchArrivals(args.Stop)
	if err != nil {
		log.Print(err)
	}
	return resp.StopTimes.Arrival
}

// API
type ArrivalsBody struct {
	StopTimes struct {
		Timestamp string     `json:"timestamp"`
		Arrival   []*Arrival `json:"arrival"`
		Stop      string     `json:"stop"`
	} `json:"stopTimes"`
}

func fetchArrivals(s int32) (*ArrivalsBody, error) {
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

	arrivals := &ArrivalsBody{}

	jsonResp, err := xj.Convert(xmlString)
	if err != nil {
		log.Print(err)
	}

	if err := json.Unmarshal(jsonResp.Bytes(), arrivals); err != nil {
		log.Print(err)
	}

	return arrivals, nil
}
