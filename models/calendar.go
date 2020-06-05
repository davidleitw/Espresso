package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type JsonTime struct {
	time.Time
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JsonTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type Event struct {
	User_ID    string `gorm:"primary_key; not null"`
	Title      string `gorm:"size:50;"`
	StartTime  string `gorm:"not null"`
	EndTime    string `gorm:"not null"`
	RemindTime string `gorm:"not null"`
	Context    string `gorm:"size:450;"`
}

type EventMain struct {
	CalendarID   string `gorm:"primary_key; not null; unique;"`
	CreateTime   string `gorm not null`
	StartTime    string `gorm:"not null"`
	EndTime      string `gorm:"not null"`
	Title        string `gorm:"size:50;"`
	Context      string
	ReferenceUrl string
}

type EventDetail struct {
	CalendarID string `gorm:"primary_key; not null; unique;"`
	UserID     string `gorm:"not null"` // 使用者ID, 有可能是創建者 或者是獲邀的使用者
	Creator    bool   `gorm:"not null"` // 判斷此userID是否為創建者, 預設為true
	RemindTime string `gorm:"not null"` // 需要提醒的時間點
	Accept     bool
}

func CreateEventTable(db *gorm.DB) {
	db.CreateTable(&Event{})
}

func CreateEventMainTable(db *gorm.DB) {
	db.CreateTable(&EventMain{})
}

func CreateEventDetailTable(db *gorm.DB) {
	db.CreateTable(&EventDetail{})
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
