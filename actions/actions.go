package actions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/bwynn/goalmanac/models"
	"github.com/bwynn/goalmanac/utils"
)

// GetWeather() - primary exposed function at init
//
// as set currently, the idea here would be to expose the GetWeather function
// to the client side by using the zip code as the parameter passed via
// request. for now, this value will be set to a string
//
// when invoked, GetWeather sends a GET request, to the url endpoint for the
// data/service retrieved. the response is in json format, and is marshaled
// and formatted, and written to a file as:
//
// <zipcode_timestampEpoch.json>
// where zipcode has been passed to handleWeatherRequest, immediately proceeding
// GetWeather() and the remaining filename assembly is handled via the
// utils.BuildFilename(<zipcode>) function
//
func GetWeather() {
	// build up query string
	var zipcode string = "95032"

	resp, err := handleWeatherRequest(zipcode)
	if err != nil {
		log.Fatal(err)
	}
	// generate filename to write
	formattedFilename := utils.BuildFilename(utils.TimeFormat(), zipcode)

	// re-marshall response
	setResponse, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	// write to json file
	ioutil.WriteFile(formattedFilename, []byte(setResponse), 0644)
}

// handleWeatherRequest()
//
// - generates url query string
// - sends http GET request
func handleWeatherRequest(zipcode string) (*models.CurrentObservations, error) {
	// get api key from env variables
	k := os.Getenv("WUNDERGROUND_API_FORECAST")
	// generate query string
	u := &url.URL{
		Scheme: "https",
		Host:   "api.wunderground.com",
		Path:   "/api/" + k + "/conditions/q/CA/" + zipcode + ".json",
	}
	// send request
	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// instantiate a new instance of Current Observations
	// map response to struct, return the new instance
	var o = new(models.CurrentObservations)
	err = json.Unmarshal(body, &o)
	if err != nil {
		log.Fatal(err)
	}

	return o, err
}
