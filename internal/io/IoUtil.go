package io

import (
	"encoding/csv"
	thisTime "googleJourneyHistogram/internal/time"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

// Creates a file, writes the headline and closes it
func CreateCsvFile(filePath string, headline []string) error {
	/* Determine file path and file name*/
	regex, err := regexp.Compile("(.*)" + string(filepath.Separator) + "(.*\\..*)")
	if err != nil {
		return err
	}
	matches := regex.FindSubmatch([]byte(filePath))
	folderPath := string(matches[1])

	/* Create missing folders */
	err = os.MkdirAll(folderPath, 0766)
	if err != nil {
		return err
	}

	/* Create the file */
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	csvWriter := csv.NewWriter(file)

	writeErr := csvWriter.Write(headline)
	if writeErr != nil {
		log.Fatalf("Error during writing to file: %s", writeErr)
		return err
	}
	csvWriter.Flush()
	return nil

}

// Opens a file and appends the given input. The file is closed on exit
func AppendToFile(path string, content []string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	csvWriter := csv.NewWriter(file)

	writeErr := csvWriter.Write(content)
	if writeErr != nil {
		return err
	}
	csvWriter.Flush()
	return nil
}

// Builds an array of strings as entry to a csv file
func BuildCsvEntry(timeStamp time.Time, travelTime int) []string {
	return []string{timeStamp.Format("2006-01-02 15:04:05"), thisTime.GetDayType(timeStamp).String(), strconv.Itoa(timeStamp.Hour()), strconv.Itoa(travelTime)}
}
