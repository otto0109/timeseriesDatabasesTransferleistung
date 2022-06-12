package functions

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"timeseriesDatabasesTransferleistung/countData"
)

func WriteDataToQuestDB(fakeData []countData.CountData, fakeDateRange []countData.SensorDateRanges) (<-chan string, <-chan string) {
	creationTime := make(chan string)
	DateRangeTime := make(chan string)

	go func() {
		// Postgre
		connectionString := "host=localhost user=admin password=quest dbname=qdb port=5435 sslmode=disable"
		db, _ := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
		repo := countData.NewRepository(db)
		creationTime <- "Quest-Creation"
		DateRangeTime <- "Quest-TimeRange"

		for i, _ := range fakeData {
			creationTime <- repo.CreateData(fakeData[i])
			DateRangeTime <- repo.GetData(fakeDateRange)
		}
	}()

	return creationTime, DateRangeTime
}
