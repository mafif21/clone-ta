package helper

import "time"

func GetValidSidangTime() time.Time {
	changeTimeFormat := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.UTC)
	validTime := changeTimeFormat.Add(2 * time.Hour)

	return validTime
}
