package countData

import (
	"github.com/brianvoe/gofakeit/v6"
	"math"
	"math/rand"
	"time"
)

func GetFakeData(days int, sensors int) (fakeData []CountData) {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day()-days, 0, 0, 0, 0, time.UTC)
	for k := 0; k < sensors; k++ {
		for i := 0; i < days; i++ {
			date = date.AddDate(0, 0, 1)
			for j := 0; j < 1440; j++ {
				date = date.Add(time.Minute)
				fakeData = append(fakeData, CountData{
					Timestamp: date,
					Value: int(math.Round(600/
						(math.Pow(math.E, float64((j-520)/100))+math.Pow(math.E, -1*float64((j-520)/100))))) + rand.Intn(10),
					SensorID: k,
				})
			}
		}
	}

	return fakeData
}

type SensorDateRanges struct {
	begin    time.Time
	end      time.Time
	sensorId int
}

func GetFakeDateRange(days int, sensors int) (fakeDateRangeData []SensorDateRanges) {

	for i := 0; i < sensors; i++ {
		for j := 0; j < 3; j++ {
			now := time.Now()

			timeRage := rand.Intn(days/2) + 1

			begin := time.Date(now.Year(), now.Month(), now.Day()-days, 0, 0, 0, 0, time.UTC)
			end := time.Date(now.Year(), now.Month(), now.Day()+1-timeRage, 0, 0, 0, 0, time.UTC)

			fakeBeginDate := gofakeit.DateRange(begin, end)

			fakeDateRangeData = append(fakeDateRangeData, SensorDateRanges{
				begin:    fakeBeginDate,
				end:      fakeBeginDate.AddDate(0, 0, timeRage),
				sensorId: i,
			})
		}
	}

	return
}
