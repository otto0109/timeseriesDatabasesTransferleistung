package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	days := 365
	sensors := 4

	fakeData := countData.GetFakeData(days, sensors)
	fakeDateRange := countData.GetFakeDateRange(days, sensors)

	fmt.Println("Data generated")

	creationMatrix := make([][]string, len(fakeData))
	selectionMatrix := make([][]string, len(fakeDateRange))

	postgresCreationTime, postgresDateRangeTime := functions.WriteDataToPostgre(fakeData, fakeDateRange)
	timescaleCreationTime, timescaleDateRangeTime := functions.WriteDataToTimescale(fakeData, fakeDateRange)
	questCreationTime, questDateRangeTime := functions.WriteDataToQuestDB(fakeData, fakeDateRange)
	//influxCreationTime, influxDateRangeTime := functions.WriteDataToInfluxDB(fakeData, fakeDateRange)

	for i := 0; i <= len(fakeData); i++ {
		array := make([]string, 1)
		array = append(array[:0],
			<-postgresCreationTime,
			<-postgresDateRangeTime,
			<-timescaleCreationTime,
			<-timescaleDateRangeTime,
			<-questCreationTime,
			<-questDateRangeTime)
		creationMatrix[i] = array
	}

	creationCSV, _ := os.Create("creation.csv")
	selectionCSV, _ := os.Create("selection.csv")

	defer creationCSV.Close()
	defer selectionCSV.Close()

	creationWriter := csv.NewWriter(creationCSV)
	creationWriter.Comma = ';'

	selectionWriter := csv.NewWriter(selectionCSV)
	selectionWriter.Comma = ';'

	fmt.Println("Write Results")
	err := creationWriter.WriteAll(creationMatrix) // calls Flush internally
	if err != nil {
		log.Fatal(err)
	}

	err = selectionWriter.WriteAll(selectionMatrix) // calls Flush internally
	if err != nil {
		log.Fatal(err)
	}

}
