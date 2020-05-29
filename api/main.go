package api

import (
	"log"
	"runtime"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetSession(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	login := session.Get("loginuser")
	return login != nil
}

func HomeGet(ctx *gin.Context) {
	islogin := GetSession(ctx)
	ctx.JSON(200, gin.H{
		"islogin": islogin,
	})
}

func Test(ctx *gin.Context) {
	session := sessions.Default(ctx)
	log.Printf("[session]=> login user name is %s\n", session.Get("loginuser"))
	ctx.JSON(200, "ok")
}

// 獲取當前所在的function名稱
func GetFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
