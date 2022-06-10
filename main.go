package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	var api string
	api = "http://ip-api.com/json/"

	var api_fields string
	api_fields = "?fields=1082361"

	if len(os.Args) < 2 {
		fmt.Println("Provide an IP address")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("Too many Arguments")
		os.Exit(2)
	}

	var ip string
	ip = os.Args[1]

	var url string
	url = api + ip + api_fields

	fmt.Println("\u001b[36m\nTarget \u001b[0m" + ip)

	var http_client = http.Client{Timeout: 10 * time.Second}

	r, err := http_client.Get(url)
	if err != nil {
		fmt.Println("error: 3")
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	type Response struct {
		Message    string  `json:"message"`
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
		fmt.Println("error: 4")
	}

	fmt.Println("\n\u001b[36mCountry \u001b[0m" + result.Country)
	fmt.Println("\u001b[36mRegion \u001b[0m" + result.RegionName)
	fmt.Println("\u001b[36mCity \u001b[0m" + result.City)
	fmt.Print("\u001b[36mLat/Lon \u001b[0m")
	fmt.Print(result.Lat)
	fmt.Print("\u001b[36m | \u001b[0m")
	fmt.Print(result.Lon)
	fmt.Println("\n\u001b[36mZip \u001b[0m" + result.Zip)
	fmt.Println("\u001b[36mTimezone \u001b[0m" + result.Timezone)
	fmt.Println("\u001b[36mISP \u001b[0m" + result.Isp)
	fmt.Println("\n\u001b[36mInfo \u001b[0m" + result.Message + "\n")

}
