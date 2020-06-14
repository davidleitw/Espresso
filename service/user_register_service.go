package service

import (
	"Espresso/models"
	serial "Espresso/serialization"

	"golang.org/x/crypto/bcrypt"
)

// catch user register json
type UserRegisterCatcher struct {
	UserMail        string `json:"UserMail" binding:"required"`
	UserName        string `json:"UserName" binding:"required"`
	UserPass        string `json:"UserPass" binding:"required"`
	UserPassConfirm string `json:"UserPassConfirm" binding:"required"`
	UserRid         string `json:"UserRid" binding:"required"`
}

func (service *UserRegisterCatcher) Register() serial.Response {
	//var user *models.Users
	// var res *serial.Response
	if service.UserPass != service.UserPassConfirm {
		return serial.BuildResponse(403, nil, "密碼驗證錯誤")
	}

	user := &models.Users{
		User_Id:    service.UserMail,
		User_Name:  service.UserName,
		PassWord:   service.UserPass,
		Mobile_Rid: service.UserRid,
	}

	if len(user.PassWord) < 8 {
		return serial.BuildResponse(403, "null", "密碼的長度不能小於八碼")
	}

	var count int = 0
	models.DB.Model(&models.Users{}).Where("user_id = ?", service.UserMail).Count(&count)
	if count > 0 {
		res := serial.BuildResponse(403, "null", "此電子信箱已經被註冊過.")
		return res
	}

	// 加密
	pass, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), bcrypt.DefaultCost)
	if err == nil {
		user.PassWord = string(pass)
	}

	if err := models.DB.Create(&user).Error; err == nil {
		return serial.BuildResponse(200, user.User_Id, "註冊成功")
	}

	res := serial.BuildResponse(403, "null", "註冊失敗, 請再試一次.")
	return res
}
