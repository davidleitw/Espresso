package service_test

import (
	"Espresso/models"
	"Espresso/server"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

type login struct {
	Ac string `json:"account"`
	Ps string `json:"password"`
}

func TestUserLoginCatcher_Login(t *testing.T) {
	models.ConnectDataBase("davidleitw:davidleitw0308@/calendardb?charset=utf8&parseTime=True&loc=Local")
	server := server.NewRouter()
	defer models.DB.Close()

	request := []struct {
		Req    login
		Status int
	}{
		{
			// 合法登入
			Req: login{
				Ac: "a001@gmail.com",
				Ps: "a001",
			},
			Status: 200,
		},
		{
			// 會判斷成帳號不存在
			Req: login{
				Ac: "",
				Ps: "",
			},
			Status: 403,
		},
		{
			// 密碼錯誤
			Req: login{
				Ac: "a001@gmail.com",
				Ps: "a002",
			},
			Status: 403,
		},
		{
			// 帳號不存在
			Req: login{
				Ac: "b001@gmail.com",
				Ps: "b001",
			},
			Status: 403,
		},
	}

	for _, req := range request {
		reqBody, _ := json.Marshal(req.Req)
		r := httptest.NewRequest(
			"POST", "/api/user/userLogin", bytes.NewReader(reqBody),
		)
		w := httptest.NewRecorder()
		server.ServeHTTP(w, r)
		assert.Equal(t, req.Status, w.Code)
	}

}

// type User struct {
// 	// Username user name
// 	Status string `json:"status"`
// 	// Password account password
// 	Message string `json:"message"`
// 	Data    string `json:"data"`
// }

// func TestUserLoginCatcher_Login(t *testing.T) {
// 	r := gofight.New()
// 	r.POST("/api/user/userLogin").
// 		SetJSONInterface(User{
// 			Status:  "foo",
// 			Message: "bar",
// 			Data:    "bar",
// 		}).
// 		Run(server.NewRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
// 			data := []byte(r.Body.String())
// 			u := gjson.GetBytes(data, "status")
// 			p := gjson.GetBytes(data, "message")
// 			d := gjson.GetBytes(data, "data")

// 			// assert.Equal(t,
// 			// fmt.Println(r)

// 			// fmt.Println(r.Body)
// 			assert.Equal(t, "foo", u.String())
// 			assert.Equal(t, "bar", p.String())
// 			assert.Equal(t, "bar", d.String())

// 			// assert.Equal(t, "Hello, appleboy", r.Body.String())
// 			assert.Equal(t, http.StatusOK, r.Code)
// 			assert.Equal(t, "application/json; charset=utf-8", r.HeaderMap.Get("Content-Type"))

// 		})

// }
