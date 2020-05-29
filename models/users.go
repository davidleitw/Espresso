package models

import (
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	User_Id    string `gorm:"primary_key;"`
	User_Name  string `gorm:"size:20; not null"`
	PassWord   string `gorm:"size:60; "`
	Mobile_Rid string `gorm:"size:25;"`
}

func CreateUserTable(db *gorm.DB) {
	db.CreateTable(&Users{})
}

func (user *Users) CheckPassWord(password string) bool {
	// user.Password => 60byte string
	result := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password))
	return result == nil
}

// davidleitw@gmail.com => return "davidleitw"
func GetPrefixEmail(email string) string {
	return strings.Split(email, "@")[0]
}

// davidleitw => davidleitw@gmail.com 方便找到資料庫所在
func GetFullEmail(ID string) string {
	return ID + "@gmail.com"
}
