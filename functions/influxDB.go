package functions

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/http"
	"timeseriesDatabasesTransferleistung/countData"
)

func WriteDataToInfluxDB(fakeData []countData.CountData, fakeDateRange []countData.SensorDateRanges) (<-chan string, <-chan string) {
	creationTime := make(chan string)
	DateRangeTime := make(chan string)

	client := influxdb2.NewClient("http://localhost:8086", "CxAiME0fgSooNZ35Sv7GuOJ_Af9mAwtbn8qv7I58Cgu9XzeEuurETya04RrlHOshNfgsb5iDlEJbr9uEA8hFXg==")
	defer client.Close()
	repo := countData.NewInfluxRepository(&client)

	_, err := client.Health(context.Background())

	if err != nil {
		panic(err)
	}

	go func() {

		for i, _ := range fakeData {

			creationTime <- "Quest-Creation"
			DateRangeTime <- "Quest-TimeRange"

			writeApi := client.WriteAPI("test", "test")

			writeApi.SetWriteFailedCallback(func(batch string, error http.Error, retryAttempts uint) bool {
				fmt.Println(batch, error, retryAttempts)

				return true
			})
			creationTime <- repo.CreateData(fakeData[i], writeApi)
			DateRangeTime <- repo.GetData(fakeDateRange)
		}
	}()

	return creationTime, DateRangeTime
}
