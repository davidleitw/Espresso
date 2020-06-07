package main

import (
	"Espresso/models"
	"fmt"
	"time"
)

func main() {

	// models.ConnectDataBase("davidleitw:davidleitw0308@/calendardb?charset=utf8&parseTime=True&loc=Local")
	// models.CreateEventMainTable(models.DB)
	// models.CreateEventDetailTable(models.DB)
	//models.CreateUserTable(models.DB)
	//models.CreateEventTable(models.DB)
	start := "2020-01-03 17:24:15"
	rem := "-15m"
	r := models.GetRemindTime(start, rem)
	fmt.Println(r)
}

func GetTimeValue(t string) time.Time {
	realtime, _ := time.Parse("2006-01-02 15:04:05", t)
	return realtime
}
