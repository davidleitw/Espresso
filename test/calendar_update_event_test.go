package service_test

import (
	"Espresso/models"
	"Espresso/server"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"gopkg.in/go-playground/assert.v1"
)

type updateEvent struct {
	OldTitle  string `json:"OldTitle"`
	OldStart  string `json:"OldStart"`
	OldRemind string `json:"OldRemind"`
	Title     string `json:"Title"`
	Start     string `json:"StartTime"`
	End       string `json:"EndTime"`
	Remind    string `json:"RemindTime"`
	Context   string `json:"Context"`
	Rurl      string `json:"ReferenceUrl"`
}

func Test_CalendarUpdateEvent(t *testing.T) {
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
		Info   updateEvent
		Status int
	}{
		{
			Info: updateEvent{
				OldTitle:  "TestEvent",
				OldStart:  "2019-04-08 12:45:32",
				OldRemind: "-3m",
				Title:     "TestEvent_update",
				Start:     "2014-01-21 12:45:32",
				End:       "2018-07-07 05-24-40",
				Remind:    "-3m",
				Context:   RandString(25) + "Test",
				Rurl:      RandString(10),
			},
			Status: 204,
		},
	}

	for _, req := range requests {
		// Get login session information
		loginBody, _ := json.Marshal(user)
		Reqlogin := httptest.NewRequest("POST", "/api/user/userLogin", bytes.NewReader(loginBody))
		Reslogin := httptest.NewRecorder()
		server.ServeHTTP(Reslogin, Reqlogin)

		reqBody, _ := json.Marshal(req.Info)
		request := httptest.NewRequest("PUT", "/api/calendar/a001/updateEvent",
			bytes.NewReader(reqBody))
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		request.Header.Set("Cookie", "mysession=MTU5MjIwODAyNnxEdi1CQkFFQ180SUFBUkFCRUFBQVF2LUNBQUlHYzNSeWFXNW5EQXNBQ1d4dloybHVkWE5sY2daemRISnBibWNNQmdBRVlUQXdNUVp6ZEhKcGJtY01DUUFIYVhOc2IyZHBiZ1JpYjI5c0FnSUFBUT09fPJ1meA-rBjslCFThdcgabnOSJdhYgE8OBls0HBdh6VJ")

		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assert.Equal(t, req.Status, response.Code)
	}

}
