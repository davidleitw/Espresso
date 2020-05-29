package api

import (
	"Espresso/serialization"
	"Espresso/service"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// @Summary 用戶註冊
// @Tags User
// @version 1.0
// @accept application/json
// @produce application/json
// @param UserMail header string true "填入用戶的電子信箱"
// @param UserName header string true "填入用戶的使用者名稱 可隨時更改"
// @param UserPass header string true "填入用戶的密碼 需要八碼以上"
// @param UserPassConfirm header string true "密碼確認, 需要UserPass參數相同"
// @param USerRid header string true "用戶的電話號碼 如果使用者沒填 前端會補上00000000"
// @Router /api/user/userRegister [post]
// @Success 200 {object} serialization.Response
// @Failure 404 {object} serialization.Response
func UserRegister(ctx *gin.Context) {
	var servicer service.UserRegisterCatcher
	if err := ctx.ShouldBind(&servicer); err == nil {
		if res := servicer.Register(); serialization.GetResStatus(res) != 200 {
			// http code 403 => Forbidden 用戶端無訪問權限
			ctx.JSON(http.StatusForbidden, res)
		} else {
			// create user successful.
			ctx.JSON(http.StatusOK, res)
			// http.StatusOk = 200
		}
	} else {
		// waiting for completion
		ctx.JSON(http.StatusNotFound, serialization.BuildResponse(404, "null", "不明原因錯誤, 請稍後再試."))
	}
}

// @Summary 用戶登入
// @Tags User
// @version 1.0
// @accept application/json
// @produce application/json
// @param account header string true "用戶電子信箱(帳號)"
// @param password header string true "用戶密碼"
// @Router /api/user/userLogin [post]
// @Success 200 {object} serialization.Response{data=string}
// @Failure 404 {object} serialization.Response{}
func UserLogin(ctx *gin.Context) {
	var servicer service.UserLoginCatcher
	if err := ctx.ShouldBind(&servicer); err == nil {
		// json get successfully
		if res := servicer.Login(); serialization.GetResStatus(res) != 200 {
			// 收到json, 但是登入失敗
			// http code 403 => Forbidden 用戶端無訪問權限
			ctx.JSON(http.StatusForbidden, res)
		} else {
			// 登入成功, 將使用者登入資料存入session中
			session := sessions.Default(ctx)
			//session.Set("loginuser", res.GetMessage().(string))
			session.Set("loginuser", serialization.GetResData(res).(string))
			session.Set("islogin", true)
			err := session.Save()

			if err == nil {
				// show session information to debug.
				if gin.Mode() == "debug" {
					login := session.Get("islogin")
					username := session.Get("loginuser")
					log.Println("login status is ", login, ", and username is ", username)
				}
				ctx.JSON(http.StatusOK, res)
			}
		}

	} else {
		ctx.JSON(http.StatusNotFound, serialization.BuildResponse(404, "null", "不明原因錯誤, 請稍後再試."))
	}
}

// @Summary 用戶登出
// @Tags User
// @version 1.0
// @accept application/json
// @produce application/json
// @Router /api/user/userLogout [post]
// @Success 200 {string} string "登出成功"
// @Failure 404 {object} serialization.Response
func UserLogout(ctx *gin.Context) {
	// 清空session
	session := sessions.Default(ctx)
	session.Delete("loginuser")
	session.Delete("islogin")
	session.Clear()
	err := session.Save()
	if err == nil {
		ctx.JSON(http.StatusOK, serialization.BuildResponse(http.StatusOK, "null", "登出成功"))
	} else {
		ctx.JSON(http.StatusNotFound, serialization.BuildResponse(http.StatusNotFound, "null", "不明原因登出失敗, 請稍後再試一次."))
	}
}

// @Summary 確認登入狀態
// @Tags User
// @version 1.0
// @accept application/json
// @produce application/json
// @Router /api/user/usercheckLogin [get]
// @Success 200 {string} string "true"
// @Failure 200 {string} string "false"
func UserCheckLogin(ctx *gin.Context) {
	session := sessions.Default(ctx)
	login := session.Get("islogin")
	username := session.Get("loginuser")
	if gin.Mode() == "debug" {
		log.Println("login status is ", login, ", and username is ", username)
	}
	if login != nil && username != nil {
		ctx.JSON(200, serialization.BuildResponse(http.StatusOK, username, "ok"))
	}
}

func UserChangeName(ctx *gin.Context) {
	var servicer service.Name
	if err := ctx.ShouldBind(&servicer); err == nil {
		ctx.JSON(http.StatusOK, servicer.ChangeName())
	} else {
		ctx.JSON(http.StatusNotFound, serialization.BuildResponse(404, "null", "不明原因錯誤, 請稍後再試."))
	}
}

func AuthSessionMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//login := sessions.Default(ctx).Get("islogin")
		_, login := ctx.Request.Cookie("islogin")
		if login == nil {
			// http.StatusUnauthorized => 400 未認證，可能需要登入或 Token
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Not login",
			})
			ctx.JSON(http.StatusUnauthorized, serialization.BuildResponse(
				http.StatusUnauthorized,
				"null",
				"Not login, please login!"))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
