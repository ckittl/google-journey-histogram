package googleapi

import (
	"encoding/json"
	"errors"
	"googleJourneyHistogram/internal/model"
	thisTime "googleJourneyHistogram/internal/time"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Query the API and calculate the travel time in minutes
func GetTravelTime(origin string, destination string, timeStamp time.Time, apiKey string) (int, error) {
	apiResponse, err := queryApi(origin, destination, timeStamp, apiKey)
	if err != nil {
		return -1, err
	}

	if apiResponse.Status == "OK" {
		travelTime := thisTime.CalcDurationInMinutes(apiResponse.Routes[0].Legs[0].DurationInTraffic.Text)
		return travelTime, nil
	} else {
		return -1, errors.New("invalid response from API")
	}
}

// Queries the Google Directions api with the given origin, destination and api key
func queryApi(origin string, destination string, timeStamp time.Time, apiKey string) (model.ApiResponse, error) {
	var apiResponse model.ApiResponse

	cleanOrigin := strings.ReplaceAll(origin, " ", "%20")
	cleanDestination := strings.ReplaceAll(destination, " ", "%20")

	queryString := "https://maps.googleapis.com/maps/api/directions/json?" +
		"origin=" + cleanOrigin +
		"&destination=" + cleanDestination +
		"&mode=driving" +
		"&departure_time=" + strconv.FormatInt(timeStamp.Unix(), 10) +
		"&key=" + apiKey
	resp, err := http.Get(queryString)
	if err != nil {
		return apiResponse, err
	}
	defer resp.Body.Close()
	jsonDecoder := json.NewDecoder(resp.Body)
	err = jsonDecoder.Decode(&apiResponse)
	if err != nil {
		return apiResponse, err
	}
	return apiResponse, nil
}
