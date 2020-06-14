package service

import (
	"Espresso/models"
	serial "Espresso/serialization"
	"log"
)

type UserLoginCatcher struct {
	LoginAccount  string `json:"account"`
	LoginPassword string `json:"password"`
}

func (service *UserLoginCatcher) Login() serial.Response {
	//var user *models.Users
	//user := new(models.Users)
	var user models.Users

	if err := models.DB.Where("user_id = ?", service.LoginAccount).First(&user).Error; err != nil {
		return serial.BuildResponse(403, "null", "帳號不存在, 請重新確認輸入的帳號是否正確.")
	}

	if !user.CheckPassWord(service.LoginPassword) {
		log.Println("Pass word error!")
		return serial.BuildResponse(403, "null", "密碼輸入錯誤.")
	}
	// davidleitw@gmail.com 取得 davidleitw 當作唯一辨識
	return serial.BuildResponse(200, models.GetPrefixEmail(user.User_Id), "登入成功")
}
