package countData

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"strconv"
	"time"
)

type influxRepository struct {
	client influxdb2.Client
}

func NewInfluxRepository(client *influxdb2.Client) *influxRepository {

	return &influxRepository{
		client: *client,
	}
}

func (r influxRepository) CreateData(fakeData CountData, api api.WriteAPI) (executionTime string) {
	start := time.Now()

	p := influxdb2.NewPointWithMeasurement("stat").AddTag("sensor", strconv.Itoa(fakeData.SensorID)).AddField("value", fakeData.Value).SetTime(fakeData.Timestamp)

	errors := api.Errors()

	api.WritePoint(p)

	api.Flush()

	go func() {
		for err := range errors {
			fmt.Println(err.Error())
		}
	}()

	return time.Since(start).String()
}

func (r influxRepository) GetData(fakeDateRanges []SensorDateRanges) (timeArray string) {

	start := time.Now()
	for _, v := range fakeDateRanges {
		queryApi := r.client.QueryAPI("test")
		query := fmt.Sprintf("from(bucket:\"test\")|> range(start: %s, stop: %s) |> filter(fn:(r) => r.sensor == \"%d\")", v.begin.Format("2006-01-02T15:04:05Z"), v.end.Format("2006-01-02T15:04:05Z"), v.sensorId)

		_, err := queryApi.Query(context.Background(), query)

		if err != nil {
			fmt.Println(err)
		}
	}

	return (time.Since(start) / time.Duration(len(fakeDateRanges))).String()
}
