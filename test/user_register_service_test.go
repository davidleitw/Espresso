package service_test

import (
	"Espresso/models"
	"Espresso/server"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

type register struct {
	Mail string `json:"UserMail" binding:"required"`
	Name string `json:"UserName" binding:"required"`
	Pas1 string `json:"UserPass" binding:"required"`
	Pas2 string `json:"UserPassConfirm" binding:"required"`
	Rid  string `json:"UserRid" binding:"required"`
}

func TestUserRegisterCatcher_Register(t *testing.T) {
	models.ConnectDataBase("davidleitw:davidleitw0308@/calendardb?charset=utf8&parseTime=True&loc=Local")
	//models.ConnectDataBase("root:@(database)/calendardb?charset=utf8&parseTime=True&loc=Local")
	server := server.NewRouter()
	defer models.DB.Close()

	request := []struct {
		Req    register
		Status int
	}{
		{
			// super user
			Req: register{
				Mail: "a001@gmail.com",
				Name: "a001",
				Pas1: "a001a001",
				Pas2: "a001a001",
				Rid:  "00000000",
			},
			Status: 200,
		},
		{
			// 理論上應該也要是可以的
			Req: register{
				Mail: "a8763" + RandString(4) + "@gmail.com",
				Name: "Kirito" + RandString(2),
				Pas1: "20221106",
				Pas2: "20221106",
				Rid:  "00000000",
			},
			Status: 200,
		},
		{
			Req: register{
				Mail: "a8763@gmail.com",
				Name: "Kirito",
				Pas1: "20221106",
				Pas2: "20221106",
				Rid:  "00000000",
			},
			Status: 200,
		},
		{
			// 密碼小於八碼
			Req: register{
				Mail: "a004@gmail.com",
				Name: "BadRequest",
				Pas1: "0308",
				Pas2: "0308",
				Rid:  "00000000",
			},
			Status: 403,
		},
		{
			// 兩次輸入密碼不同
			Req: register{
				Mail: "a005@gmail.com",
				Name: "Pas2Error",
				Pas1: "01230123",
				Pas2: "01240124",
				Rid:  "00000000",
			},
			Status: 403,
		},
	}

	for idx, req := range request {
		reqBody, _ := json.Marshal(req.Req)
		r := httptest.NewRequest(
			"POST", "/api/user/userRegister", bytes.NewReader(reqBody),
		)
		w := httptest.NewRecorder()
		server.ServeHTTP(w, r)
		fmt.Printf("Test Case%d: ReqStatus=%d <=> wStatus:%d\n", idx, req.Status, w.Code)
		assert.Equal(t, req.Status, w.Code)
	}
}

// func TestPostJSONData(t *testing.T) {
// 	models.ConnectDataBase("root:rootroot@/calendardb?charset=utf8mb4&parseTime=True&loc=Local")
// 	defer models.DB.Close()

// 	r := gofight.New()

// 	r.POST("/api/user/userRegister").
// 		SetJSON(gofight.D{
// 			"UserMail":        "a002@gmail.com",
// 			"UserName":        "a002",
// 			"UserPass":        "a002",
// 			"UserPassConfirm": "a002",
// 			"USerRid":         "1234562",
// 		}).
// 		Run(server.NewRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
// 			// data := []byte(r.Body.String())

// 			// a, _ := jsonparser.GetInt(data, "a")
// 			// b, _ := jsonparser.GetInt(data, "b")

// 			// assert.Equal(t, 1, int(a))
// 			// assert.Equal(t, 2, int(b))
// 			assert.Equal(t, http.StatusOK, r.Code)
// 			assert.Equal(t, "application/json; charset=utf-8", r.HeaderMap.Get("Content-Type"))
// 		})
// }

/*
	models.ConnectDataBase("root:rootroot@/calendardb?charset=utf8mb4&parseTime=True&loc=Local")
	defer models.DB.Close()

	// ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	r := server.NewRouter()

	reqData := []struct {
		UserMail        string `json:"UserMail" binding:"required"`
		UserName        string `json:"UserName" binding:"required"`
		UserPass        string `json:"UserPass" binding:"required"`
		UserPassConfirm string `json:"UserPassConfirm" binding:"required"`
		UserRid         string `json:"UserRid" binding:"required"`
	}{
		{UserMail: "a004@gmail.com", UserName: "a004", UserPass: "a004", UserPassConfirm: "a004", UserRid: "01102"},
		{UserMail: "a005@gmail.com", UserName: "a005", UserPass: "a005", UserPassConfirm: "a005", UserRid: "1"},
	}
	// var jsonStr = []byte(`{"account":"123","password":"123"}`)

	// bytes.NewReader(reqBody)
	for _, c := range reqData {
		reqBody, _ := json.Marshal(c)
		fmt.Println("input:", string(reqBody))

		req := httptest.NewRequest(
			"POST", "/api/user/userRegister", bytes.NewReader(reqBody))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		fmt.Println(w.Body)

		// result := rr.Result()
		// api.UserLogin(ctx)

		assert.Equal(t, 403, w.Code)

	}
*/
