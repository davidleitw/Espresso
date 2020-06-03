package server

import (
	"Espresso/api"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	config.AllowOrigins = []string{"http://127.0.0.1:8080", "http://espresso.nctu.me:8080"}
	config.AllowCredentials = true
	return cors.New(config)
}

func NewRouter() *gin.Engine {
	server := gin.Default()

	//store := cookie.NewStore([]byte("loginuser"), []byte("islogin"))
	store := cookie.NewStore([]byte("secret"))
	// session active time = 30 minutes .
	store.Options(sessions.Options{
		MaxAge: int(30 * time.Minute),
	})
	server.Use(setCors())
	server.Use(sessions.Sessions("mysession", store))
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/", api.HomeGet) // 回首頁
		user := apiRoutes.Group("/user")
		{
			user.GET("usercheckLogin", api.UserCheckLogin)                  // 檢查登入狀態
			user.POST("userRegister", api.UserRegister)                     // 註冊
			user.POST("userLogin", api.UserLogin)                           // 登入
			user.POST("userLogout", api.UserLogout)                         // 登出
			user.PATCH("userchangeName", api.AuthSessionMiddle(), api.Test) // 更換用戶名稱 x
			user.PATCH("userchangePass", api.AuthSessionMiddle(), api.Test) // 更換密碼    x
		}
		calendar := apiRoutes.Group("/calendar")
		{
			calendar.GET(":ID/getAllEvent", api.AuthSessionMiddle(), api.CalendarGetAllEvent)     // 登入成功之後抓資料進入前端
			calendar.GET(":ID/getEventInfo", api.AuthSessionMiddle(), api.Test)                   // 點開單一事件, 得到詳細的資訊
			calendar.PUT(":ID/updateEvent", api.AuthSessionMiddle(), api.Test)                    // 更新單一事件
			calendar.POST(":ID/createNewEvent", api.AuthSessionMiddle(), api.CalendarCreateEvent) // 新增一筆事件
			calendar.DELETE(":ID/deleteEvent", api.AuthSessionMiddle(), api.CalendarDeleteEvent)  // 刪除一筆事件
		}
		apiRoutes.POST("test", api.Test) // 測試function
	}

	return server
}
