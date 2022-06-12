package functions

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"timeseriesDatabasesTransferleistung/countData"
)

func WriteDataToPostgre(fakeData []countData.CountData, fakeDateRange []countData.SensorDateRanges) (<-chan string, <-chan string) {
	postgresCreationTime := make(chan string)
	postgresDateRangeTime := make(chan string)

	go func() {
		// Postgre
		postgresConnectionString := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
		postgresDb, _ := gorm.Open(postgres.Open(postgresConnectionString), &gorm.Config{})
		postgreRepo := countData.NewRepository(postgresDb)
		postgresCreationTime <- "Postgre-Creation"
		postgresDateRangeTime <- "Postgre-TimeRange"

		for i, _ := range fakeData {
			postgresCreationTime <- postgreRepo.CreateData(fakeData[i])
			postgresDateRangeTime <- postgreRepo.GetData(fakeDateRange)
		}
	}()

	return postgresCreationTime, postgresDateRangeTime
}
