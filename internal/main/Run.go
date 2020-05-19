package main

import (
	"googleJourneyHistogram/internal/googleapi"
	"googleJourneyHistogram/internal/io"
	"googleJourneyHistogram/internal/model"
	"log"
	"strconv"
	"time"
)

func main() {
	/* ==== Configure the behaviour here ==== */
	apiKey := "YOUR-API-KEY"
	outputFile := "./output/journeyTime.csv"
	journeys := []model.Journey{
		{"Hamburg", "Berlin"},
		{"Hamburg", "München"},
	}
	/* ==== End of config section ==== */

	/* Create a csv file with headline to later write content to */
	err := io.CreateCsvFile(outputFile, []string{"date", "dayType", "hourOfDay", "origin", "destination", "travelTime"})
	if err != nil {
		log.Fatalf("Error during creation of csv file '%s'. Error: %s", outputFile, err)
	}

	/* Go through all journeys */
	for _, journey := range journeys {
		origin := journey.Origin
		destination := journey.Destination

		/* Get the travel time from google api */
		travelTime, err := googleapi.GetTravelTime(origin, destination, apiKey)
		if err != nil {
			log.Fatalf("Error during API request. Error: '%s'", err)
		}

		/* Write the travel time to csv file */
		timeStamp := time.Now()
		entry := io.BuildCsvEntry(timeStamp, origin, destination, travelTime)
		err = io.AppendToFile(outputFile, entry)
		if err != nil {
			log.Fatalf("Error during appending of travel time from %s to %s at %s to output file '%s'. Error: '%s'",
				origin,
				destination,
				timeStamp.Format("2006-01-02 15:04:05"),
				outputFile,
				err)
		}
		log.Printf("The travel time from %s to %s at '%s' is %s minutes.",
			origin,
			destination,
			timeStamp.Format("2006-01-02 15:04:05"),
			strconv.Itoa(travelTime))
	}
}
