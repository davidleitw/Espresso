package service

import (
	"Espresso/models"
	serial "Espresso/serialization"
	"html/template"
)

type Name struct {
	NewName string `json: "NewName"`
	UserID  string `json: "UserID"`
}

func (n *Name) ChangeName() serial.Response {
	// 防止sql injection
	var user models.Users
	newName := template.HTMLEscapeString(n.NewName)
	models.DB.Model(&user).Where("user_id = ?", models.GetFullEmail(n.UserID)).Update("user_name", newName)
	return serial.BuildResponse(200, user, "change name sucessfully.")
}
