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

type delevent struct {
	Title      string `json:"title"`
	StartTime  string `json:"start_time"`
	RemindTime string `json:"remind_time"`
}

func Test_CalendarDeleteEvent(t *testing.T) {
	models.ConnectDataBase("davidleitw:davidleitw0308@/calendardb?charset=utf8&parseTime=True&loc=Local")
	//models.ConnectDataBase("root:@(database)/calendardb?charset=utf8&parseTime=True&loc=Local")
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
		Ps: "a001a001",
	}

	requests := []struct {
		Event  delevent
		Status int
	}{
		{
			// 應該會被刪除的一筆資料
			Event: delevent{
				Title:      "EventOK to change",
				StartTime:  "2000-00-00 12:00:00",
				RemindTime: "-4m",
			},
			Status: 204,
		},
		{
			// 沒有這筆資料, 空做刪除的動作
			Event: delevent{
				Title:      "No Exist event delete",
				StartTime:  "2014-01-06 14:06:24",
				RemindTime: "-4m",
			},
			Status: 204,
		},
	}

	for idx, req := range requests {
		// Get login session information
		loginBody, _ := json.Marshal(user)
		Reqlogin := httptest.NewRequest("POST", "/api/user/userLogin", bytes.NewReader(loginBody))
		Reslogin := httptest.NewRecorder()
		server.ServeHTTP(Reslogin, Reqlogin)

		reqBody, _ := json.Marshal(req.Event)
		request := httptest.NewRequest("DELETE", "/api/calendar/a001/deleteEvent",
			bytes.NewReader(reqBody))
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		request.Header.Set("Cookie", "mysession=MTU5MjIwODAyNnxEdi1CQkFFQ180SUFBUkFCRUFBQVF2LUNBQUlHYzNSeWFXNW5EQXNBQ1d4dloybHVkWE5sY2daemRISnBibWNNQmdBRVlUQXdNUVp6ZEhKcGJtY01DUUFIYVhOc2IyZHBiZ1JpYjI5c0FnSUFBUT09fPJ1meA-rBjslCFThdcgabnOSJdhYgE8OBls0HBdh6VJ")

		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		fmt.Printf("Delete unit test Case%d=> reqStatus %d <> resStatus%d\n", idx, req.Status, response.Code)
		assert.Equal(t, req.Status, response.Code)

	}
}
