package main

import (
	_ "Espresso/docs"
	"Espresso/models"
	"Espresso/server"
)
 
// @title Espresso Example API File
// @version 0.0.1
// @contact.name davidleitw
// @contact.email davidleitw@gmail.com
// @description Expreesso calendar的api文檔, 方便串接前後端
// @license.name Apache 2.0
// @host http://espresso.nctu.me:8080/
// @BasePath /api/
func main() {
	models.ConnectDataBase("davidleitw:davidleitw0308@/calendardb?charset=utf8&parseTime=True&loc=Local")
	//models.ConnectDataBase("root:davidleitw@(davidleitw)/calendardb?charset=utf8&parseTime=True&loc=Local")
	defer models.DB.Close()
	r := server.NewRouter()
	_ = r.Run(":3000")
}
