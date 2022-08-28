// In this example we will be using some data from the Citi Bike API
// ( https:/​/www.​citibikenyc.​com/​system-​data ), a bike-sharing
// service operating in New York City.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// citiBikeURL provides the station statuses of CitiBike bike sharing
// stations.
const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

// stationData is used to unmarshal the JSON document returned form
// citiBikeURL.
type stationData struct {
	LastUpdated int `json:"last_updated"`
	TTL         int `json:"ttl"`
	Data        struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

// station is used to unmarshal each of the station documents in
// stationData.
type station struct {
	ID                string `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisabled  int    `json:"num_bike_disabled"`
	NumDocksAvailable int    `json:"num_docks_available"`
	NumDocksDisabled  int    `json:"num_docks_disabled"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
	HasAvailableKeys  bool   `json:"eightd_has_available_keys"`
}

func main() {
	resp, err := http.Get(citiBikeURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the body of the response into bytes
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Declare a var of type stationData.
	var sd stationData

	// Unmarshal the JSON data into the sd var.
	err = json.Unmarshal(b, &sd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Marshal the data.
	data, err := json.Marshal(sd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create file to save marshalled data.
	f, err := os.Create("citibike.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Save the marshalled data to a file.
	w := bufio.NewWriter(f)
	n, err := w.Write(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Written %d bytes\n", n)
}
