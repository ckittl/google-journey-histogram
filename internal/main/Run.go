package main

import (
	"googleJourneyHistogram/internal/googleapi"
	"googleJourneyHistogram/internal/io"
	"log"
	"strconv"
	"time"
)

func main() {
	/* ==== Configure the behaviour here ==== */
	outputFile := "./output/journeyTime.csv"
	/* ==== End of config section ==== */

	/* Create a csv file with headline to later write content to */
	err := io.CreateCsvFile(outputFile, []string{"date", "dayType", "hourOfDay", "travelTime"})
	if err != nil {
		log.Fatalf("Error during creation of csv file '%s'. Error: %s", outputFile, err)
	}

	/* Get the travel time from google api */
	travelTime, err := googleapi.GetTravelTime("Hamburg", "Berlin", "YOUR-API-KEY")
	if err != nil {
		log.Fatalf("Error during API request. Error: '%s'", err)
	}

	/* Write the travel time to csv file */
	timeStamp := time.Now()
	entry := io.BuildCsvEntry(timeStamp, travelTime)
	err = io.AppendToFile(outputFile, entry)
	if err != nil {
		log.Fatalf("Error during appending of travel time for %s to output file '%s'. Error: '%s'", timeStamp.Format("2006-01-02 15:04:05"), outputFile, err)
	}
	log.Printf("The travel time at '%s' is %s minutes.", timeStamp.Format("2006-01-02 15:04:05"), strconv.Itoa(travelTime))
}
