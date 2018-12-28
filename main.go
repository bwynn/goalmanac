package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bwynn/goalmanac/actions"
	"github.com/gorilla/mux"
)

// default msg to send
func serveClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("response ok - 200")
	w.Write([]byte("Hello! Welcome to the Goalmanac Weather Service"))
}

// when pinged, sends request to wunderground, and writes response to /data
func fetchWeatherService(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("weather service dispatched"))
	actions.GetWeather() // write file to data
}

// should create and run server
//
// - fetch weather service
//		- sends GET request to wunderground api
//		- writes response to file in /data
//
// - retrieve weather data
//		- handle GET request for weather entries between specified
//		  timestamp ranges
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveClient)
	r.HandleFunc("/fetch-weather", fetchWeatherService)
	log.Fatal(http.ListenAndServe(":8080", r))
}
