package countData

import (
	"gorm.io/gorm"
	"time"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	db.AutoMigrate(&CountData{})

	return &repository{
		db: db,
	}
}

func (r repository) CreateData(fakeData CountData) (executionTime string) {
	start := time.Now()
	r.db.Create(fakeData)

	return time.Since(start).String()
}

func (r repository) GetData(fakeDateRanges []SensorDateRanges) (timeArray string) {

	start := time.Now()
	for _, v := range fakeDateRanges {
		var countData []CountData

		r.db.Where("sensor_id = ? and timestamp BETWEEN  ? and ? ", v.sensorId, v.begin, v.end).Find(&countData)
	}

	return (time.Since(start) / time.Duration(len(fakeDateRanges))).String()
}
