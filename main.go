package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	getWeather()
}

func buildString(s []string) string {
	return strings.Join(s, "")
}

type Observation struct {
	ObservationEpoch string    `json:"observation_epoch"`
	StationID        string    `json:"station_id"`
	TempC            float64   `json:"temp_c"`
	WindDegrees      int       `json:"wind_degrees"`
	WindKPH          float64   `json:"wind_kph"`
	WindGustKPH      string    `json:"wind_gust_kph"`
	PressureMB       string    `json:"pressure_mb"`
	PressureIN       string    `json:"pressure_in"`
	PressureTrend    string    `json:"pressure_trend"`
	DewpointC        int       `json:"dewpoint_c"`
	WindchillC       string    `json:"windchill_c"`
	Precip1HRMetric  string    `json:"precip_1hr_metric"`
	DisplayLocation  *Location `json:"display_location"`
}

type Location struct {
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Elevation string `json:"elevation"`
}

type CurrentObservations struct {
	CurrentObservation *Observation `json:"current_observation"`
}

func getWeather() {
	// build up query string
	var zipcode string = "95032"

	resp, err := handleWeatherRequest(zipcode)
	if err != nil {
		log.Fatal(err)
	}
	// generate filename to write
	formattedFilename := buildFilename(timeFormat(), zipcode)

	// re-marshall response
	setResponse, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	// write to json file
	ioutil.WriteFile(formattedFilename, []byte(setResponse), 0644)
}

func handleWeatherRequest(zipcode string) (*CurrentObservations, error) {
	// get api key from env variables
	k := os.Getenv("WUNDERGROUND_API_FORECAST")
	// generate query string
	wundergroundPath := []string{
		"https://api.wunderground.com/api/",
		k,
		"/conditions/q/CA/",
		zipcode,
		".json",
	}
	formattedQuery := buildString(wundergroundPath)
	// send request
	resp, err := http.Get(formattedQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var o = new(CurrentObservations)
	err = json.Unmarshal(body, &o)
	if err != nil {
		log.Fatal(err)
	}

	return o, err
}

// buildFilename
//
// the filename will be generated using a combination of
// timestamps and locations (postal code),
// allowing for queries to existing files prior to migration/seeding
// to a db instance
//
// e.g.: 95032_1535589784
//		 <postal_code>_<unix_timestamp>.json
func buildFilename(unixT int64, zipcode string) string {
	// format unix to string
	timestamp := strconv.FormatInt(unixT, 10)
	filename := []string{
		"data/",
		zipcode,
		"_",
		timestamp,
		".json",
	}
	// join the slice of strings w/ util method
	formattedFilename := buildString(filename)
	return formattedFilename
}

// returns unix timestamp and prints time of initialization
func timeFormat() int64 {
	// get the current time
	t := time.Now()
	// format and set timestamp to unix time
	// before writing to file
	setToUnixT := t.Unix()
	formatted := t.Format("Monday, Aug 1 15:04:05 -0800 PST 2006")
	// log time for posterity
	// ahoy!
	fmt.Printf("request sent at: %s", formatted)
	return setToUnixT
}
