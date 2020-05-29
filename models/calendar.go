package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Event struct {
	User_ID    string `gorm:"primary_key; not null"`
	Title      string `gorm:"size:50;"`
	StartTime  string `gorm:"not null"`
	EndTime    string `gorm:"not null"`
	RemindTime string `gorm:"not null"`
	Context    string `gorm:"size:450;"`
}

func CreateEventTable(db *gorm.DB) {
	db.CreateTable(&Event{})
}

func GetResultTime(t string, st time.Time) time.Time {
	t1, _ := time.ParseDuration(t)
	return st.Add(t1)
}

// change string to time.Time
func GetTimeValue(t string) time.Time {
	realtime, _ := time.Parse("2006-01-02 15:04:05", t)
	return realtime
}

// change time.Time to string
func GetTimeString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
