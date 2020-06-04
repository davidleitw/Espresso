package main

import (
	"Espresso/models"
	"time"
)

func main() {

	models.ConnectDataBase("davidleitw:davidleitw0308@/calendardb?charset=utf8&parseTime=True&loc=Local")
	// guid := xid.New()
	// fmt.Println(guid.Time())
	// fmt.Println(guid.String())
	models.CreateEventMainTable(models.DB)
	models.CreateEventDetailTable(models.DB)
	//models.CreateUserTable(models.DB)
	//models.CreateEventTable(models.DB)
	// // 运行一定时间后退出
	// st := time.Now()
	// fmt.Println(st)

	// a, _ := time.ParseDuration("1m")
	// st = st.Add(a)
	// fmt.Println(st)
	// fmt.Println(st.String())
	// fmt.Println(st.Format("2006-01-02 15:04:05"))
	// fmt.Printf("type of st is %T\n", st.Format("2006-01-02 15:04:05"))

	// nt, _ := time.Parse("2006-01-02 15:04:05", st.Format("2006-01-02 15:04:05"))
	// fmt.Println(nt)
}

func GetTimeValue(t string) time.Time {
	realtime, _ := time.Parse("2006-01-02 15:04:05", t)
	return realtime
}
