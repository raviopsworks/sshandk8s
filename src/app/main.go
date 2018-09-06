package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func weatherAPI(w http.ResponseWriter, r *http.Request) {
	var apiURL string
	// http://localhost:80/endPath
	endPath := r.URL.Path
	endPath = strings.TrimPrefix(message, "/")

	// apiURL := "http://api.openweathermap.org/data/2.5/forecast?id=524901&APPID=*******************************" + message
	// api.openweathermap.org/data/2.5/weather?zip={zip code},{country code}
	// api.openweathermap.org/data/2.5/weather?q={city name}

	// apiURL := "google.com"
	// response, err := http.Get("http://www.google.com")

	// check if endpoint is a zipcode
	if _, err := strconv.Atoi(endPath); err == nil {
		// apirequest for weather info of zipcode in us with APPID passed from env
		apiURL = "http://api.openweathermap.org/data/2.5/weather?zip=" + endPath + ",us" + "&appid=" + os.Getenv("APPID")
	} else {
		// if endpoint is city then api request for city weather with APPID passed from env
		apiURL = "http://api.openweathermap.org/data/2.5/weather?q=" + endPath + "&appid=" + os.Getenv("APPID")
	}

	response, err := http.Get(apiURL)

	if err != nil {
		fmt.Printf("The weather api HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		w.Write([]byte(data))
	}
}

func main() {

	// pass the http handle to weatherAPI func
	http.HandleFunc("/", weatherAPI)

	// listen for web server traffic on port from env variable
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), nil)

}
