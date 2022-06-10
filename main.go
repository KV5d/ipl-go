package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	var api string
	api = "http://ip-api.com/json/"

	var api_fields string
	api_fields = "?fields=1017"

	fmt.Println("")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Target: ")
	target, _ := reader.ReadString('\n')

	var ip string
	ip = strings.TrimSuffix(target, "\n")

	var url string
	url = api + ip + api_fields

	var http_client = http.Client{Timeout: 10 * time.Second}

	r, err := http_client.Get(url)
	if err != nil {
		fmt.Println("error: 001")
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	type Response struct {
		Country    string  `json:"country"`
		RegionName string  `json:"regionName"`
		City       string  `json:"city"`
		Zip        string  `json:"zip"`
		Lat        float64 `json:"lat"`
		Lon        float64 `json:"lon"`
		Timezone   string  `json:"timezone"`
		Isp        string  `json:"isp"`
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("error: 002")
	}

	fmt.Println("\nCountry: " + result.Country)
	fmt.Println("Region: " + result.RegionName)
	fmt.Println("City: " + result.City)
	fmt.Print("Lat/Lon: ")
	fmt.Print(result.Lat)
	fmt.Print(" | ")
	fmt.Print(result.Lon)
	fmt.Println("\nZip: " + result.Zip)
	fmt.Println("TZ: " + result.Timezone)
	fmt.Println("ISP: " + result.Isp + "\n")

}
