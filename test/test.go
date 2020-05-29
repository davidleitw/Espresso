package main

import (
	"Espresso/models"
	"fmt"
)

// 拿來測試一些gorm的語法
func main() {
	models.ConnectDataBase("davidleitw:davidleitw0308@/calendardb?charset=utf8&parseTime=True&loc=Local")
	name := "fr@gmail.com"
	name2 := "davidllleii@gmail.com"

	var u []models.Users
	var user1 models.Users
	var user2 models.Users

	models.DB.Where("user_id = ?", name).First(&user1)
	models.DB.Where("user_id = ?", name2).First(&user2)

	models.DB.Where("user_name = ?", "davidleitw").Find(&u)

	fmt.Println(user1.Mobile_Rid)
	fmt.Println(user2.Mobile_Rid)

	//fmt.Println(u)
	for _, j := range u {
		fmt.Println(j)
	}
}
