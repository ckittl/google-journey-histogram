package main

import "googleJourneyHistogram/internal/googleapi"

func main() {
	apiResponse, err := googleapi.QueryApi("Hamburg", "Berlin", "YOUR-API-KEY")
	if err != nil {
		panic(err)
	}
	println(apiResponse.Status)
}
