package countData

import (
	"time"
)

type CountData struct {
	Timestamp time.Time
	Value     int
	SensorID  int
}
