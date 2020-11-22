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

var (
	apiKey = ""
)

type Arrivals struct {
	StopTimes struct {
		Timestamp string `json:"timestamp"`
		Arrival   []struct {
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
		} `json:"arrival"`
		Stop string `json:"stop"`
	} `json:"stopTimes"`
}

func fetchArrivals(s string) (*Arrivals, error) {
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

	arrivals := &Arrivals{}

	jsonResp, err := xj.Convert(xmlString)
	if err != nil {
		log.Print(err)
	}

	if err := json.Unmarshal(jsonResp.Bytes(), arrivals); err != nil {
		log.Print(err)
	}

	return arrivals, nil
}
