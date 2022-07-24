package functions

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func WriteDataToTimescale(fakeData []countData.CountData, fakeDateRange []countData.SensorDateRanges) (<-chan string, <-chan string) {
	timescaleCreationTime := make(chan string)
	timescaleDateRangeTime := make(chan string)
	// Postgre
	timescaleConnectionString := "host=localhost user=postgres password=timescale dbname=postgres port=5433 sslmode=disable TimeZone=UTC"
	timescaleDb, _ := gorm.Open(postgres.Open(timescaleConnectionString), &gorm.Config{})
	timescaleRepo := countData.NewRepository(timescaleDb)

	go func() {
		timescaleCreationTime <- "Timescale-Creation"
		timescaleDateRangeTime <- "Timescale-TimeRange"

		for i, _ := range fakeData {
			max := i + 99999

			if max > len(fakeData) {
				max = len(fakeData) - 1
			}

			timescaleCreationTime <- timescaleRepo.CreateData(fakeData[i:max])
			timescaleDateRangeTime <- timescaleRepo.GetData(fakeDateRange)
		}
	}()

	return timescaleCreationTime, timescaleDateRangeTime
}
