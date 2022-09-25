package util

import (
	"fmt"
	"time"
)

const (
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
)

type Time time.Time

func GetDateNow() string {
	return time.Now().Format(DateFormat)
}

func GetTimeNow() string {
	return time.Now().Format(TimeFormat)
}

func GetLocationTime(timeStr string) (time.Time, error) {
	location, err := time.LoadLocation("Asia/Taipei")
	locTime, err := time.ParseInLocation(TimeFormat, timeStr, location)

	if err != nil {
		fmt.Println("error:", err)
		return locTime, err
	}
	return locTime, nil
}

func GetLocationDate(timeStr string) (time.Time, error) {
	location, err := time.LoadLocation("Asia/Taipei")
	locDate, err := time.ParseInLocation(DateFormat, timeStr, location)

	if err != nil {
		fmt.Println("error:", err)
		return locDate, err
	}
	return locDate, nil
}
