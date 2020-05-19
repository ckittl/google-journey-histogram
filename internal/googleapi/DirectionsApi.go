package googleapi

import (
	"encoding/json"
	"googleJourneyHistogram/internal/model"
	"net/http"
)

// Queries the Google Directions api with the given origin, destination and api key
func QueryApi(origin string, destination string, apiKey string) (model.ApiResponse, error) {
	var apiResponse model.ApiResponse

	queryString := "https://maps.googleapis.com/maps/api/directions/json?" +
		"origin=" + origin +
		"&destination=" + destination +
		"&mode=driving" +
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
