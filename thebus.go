package main

import (
	"encoding/json"
	"fmt"
	xj "github.com/basgys/goxml2json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	API_KEY = ""
)

func fetchArrivals(s string) ([]byte, error) {
	arrivalsPath := fmt.Sprintf("http://api.thebus.org/arrivals/?key=%v&stop=%v", API_KEY, s)
	fmt.Printf("%s \n", arrivalsPath)
	resp, err := http.Get(arrivalsPath)
	if err != nil {
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	respString := string(respData)
	xmlString := strings.NewReader(respString)
	jsonResp, err := xj.Convert(xmlString)
	if err != nil {
	}
	r, _ := json.Marshal(jsonResp.String())
	return r, nil
}
