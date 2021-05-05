package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IPINFO struct {
	IP           string `json:"ip"`
	Hostname     string `json:"hostname"`
	City         string `json:"city"`
	Region       string `json:"region"`
	Country      string `json:"country"`
	Location     string `json:"loc"`
	Organisation string `json:"org"`
	Postal       string `json:"postal"`
	Timezone     string `json:"timezone"`
}

func main() {
	url := "https://ipinfo.io/?token=YOUR_TOKEN"
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))
	var ipinfo IPINFO
	err = json.Unmarshal(body, &ipinfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("IP           : " + ipinfo.IP)
	fmt.Println("Hostname     : " + ipinfo.Hostname)
	fmt.Println("City         : " + ipinfo.City)
	fmt.Println("Region       : " + ipinfo.Region)
	fmt.Println("Country      : " + ipinfo.Country)
	fmt.Println("Location     : " + ipinfo.Location)
	fmt.Println("Organisation : " + ipinfo.Organisation)
	fmt.Println("Postal       : " + ipinfo.Postal)
	fmt.Println("Timezone     : " + ipinfo.Timezone)
}
