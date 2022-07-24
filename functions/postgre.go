package functions

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func WriteDataToPostgre(fakeData []countData.CountData, fakeDateRange []countData.SensorDateRanges) (<-chan string, <-chan string) {
	postgresCreationTime := make(chan string)
	postgresDateRangeTime := make(chan string)
	// Postgre
	postgresConnectionString := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	postgresDb, _ := gorm.Open(postgres.Open(postgresConnectionString), &gorm.Config{})
	postgreRepo := countData.NewRepository(postgresDb)

	go func() {
		postgresCreationTime <- "Postgre-Creation"
		postgresDateRangeTime <- "Postgre-TimeRange"

		for i := 0; i < len(fakeData)/10000; i++ {

			max := i + 99999

			if max > len(fakeData) {
				max = len(fakeData) - 1
			}

			postgresCreationTime <- postgreRepo.CreateData(fakeData[i:max])
			postgresDateRangeTime <- postgreRepo.GetData(fakeDateRange)
		}
	}()

	return postgresCreationTime, postgresDateRangeTime
}
