package functions

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"timeseriesDatabasesTransferleistung/countData"
)

func WriteDataToTimescale(fakeData []countData.CountData, fakeDateRange []countData.SensorDateRanges) (<-chan string, <-chan string) {
	timescaleCreationTime := make(chan string)
	timescaleDateRangeTime := make(chan string)

	go func() {
		// Postgre
		timescaleConnectionString := "host=localhost user=postgres password=timescale dbname=postgres port=5433 sslmode=disable TimeZone=UTC"
		timescaleDb, _ := gorm.Open(postgres.Open(timescaleConnectionString), &gorm.Config{})
		timescaleRepo := countData.NewRepository(timescaleDb)
		timescaleCreationTime <- "Timescale-Creation"
		timescaleDateRangeTime <- "Timescale-TimeRange"

		for i, _ := range fakeData {
			timescaleCreationTime <- timescaleRepo.CreateData(fakeData[i])
			timescaleDateRangeTime <- timescaleRepo.GetData(fakeDateRange)
		}
	}()

	return timescaleCreationTime, timescaleDateRangeTime
}
