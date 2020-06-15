package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func ConnectDataBase(dbname string) {
	db, err := gorm.Open("mysql", dbname)
	CheckError(err)
	fmt.Println("Successfully connect database.")

	// set db to debug mode
	if gin.Mode() != "release" {
		db.LogMode(true)
	}

	DB = db
	//CreateUserTable(db)
	migration()
}

func migration() {
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&Users{}, &EventMain{}, &EventDetail{})
}
