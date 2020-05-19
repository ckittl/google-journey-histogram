package main

import (
	"googleJourneyHistogram/internal/googleapi"
	"strconv"
)

func main() {
	travelTime, err := googleapi.GetTravelTime("Hamburg", "Berlin", "YOUR-API-KEY")
	if err != nil {
		panic(err)
	}

	println("The current journey would take you " + strconv.Itoa(travelTime) + " minutes.")
}
