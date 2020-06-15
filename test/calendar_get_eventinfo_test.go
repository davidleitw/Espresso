package service_test

import (
	"Espresso/models"
	"Espresso/server"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"gopkg.in/go-playground/assert.v1"
)

type eventInfo struct {
	Title      string `json:"title"`
	StartTime  string `json:"start_time"`
	RemindTime string `json:"remind_time"`
}

func Test_CalendarGetInfo(t *testing.T) {
	models.ConnectDataBase("davidleitw:davidleitw0308@/calendardb?charset=utf8&parseTime=True&loc=Local")
	store := cookie.NewStore([]byte("secret"))
	server := server.NewRouter()
	server.Use(SetCors())
	server.Use(sessions.Sessions("mysession", store))
	defer models.DB.Close()

	user := struct {
		Ac string `json:"account"`
		Ps string `json:"password"`
	}{
		Ac: "a001@gmail.com",
		Ps: "a001",
	}

	requests := []struct {
		Info   eventInfo
		Status int
	}{
		{
			Info: eventInfo{
				Title:      "0123icSdTO",
				StartTime:  "2014-01-21 12:45:32",
				RemindTime: "-3m",
			},
			Status: 200,
		},
		{
			Info: eventInfo{
				Title:      "Not exist event test",
				StartTime:  "2017-08-04 14:28:36",
				RemindTime: "-7m",
			},
			Status: 500,
		},
	}
	for _, req := range requests {
		// Get login session information
		loginBody, _ := json.Marshal(user)
		Reqlogin := httptest.NewRequest("POST", "/api/user/userLogin", bytes.NewReader(loginBody))
		Reslogin := httptest.NewRecorder()
		server.ServeHTTP(Reslogin, Reqlogin)

		reqBody, _ := json.Marshal(req.Info)
		request := httptest.NewRequest("POST", "/api/calendar/a001/getEventInfo",
			bytes.NewReader(reqBody))
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		request.Header.Set("Cookie", "mysession=MTU5MjIwODAyNnxEdi1CQkFFQ180SUFBUkFCRUFBQVF2LUNBQUlHYzNSeWFXNW5EQXNBQ1d4dloybHVkWE5sY2daemRISnBibWNNQmdBRVlUQXdNUVp6ZEhKcGJtY01DUUFIYVhOc2IyZHBiZ1JpYjI5c0FnSUFBUT09fPJ1meA-rBjslCFThdcgabnOSJdhYgE8OBls0HBdh6VJ")

		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assert.Equal(t, req.Status, response.Code)

		// 解析buffer to map[string]interface的方法
		var tmp interface{}
		err := json.Unmarshal(response.Body.Bytes(), &tmp)
		if err == nil {
			t := tmp.(map[string]interface{})
			fmt.Println(t["data"])
		}
	}
}
