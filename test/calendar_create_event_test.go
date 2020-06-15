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

/* Go test 小細節, 如果你是用同一個test資料夾的其他檔案function,
打指令的時候必須加上 go test -v calendar_create_event_test.go Tool_test.go
															^ Tool_test內部放了所有test所需要的共通工具
*/

type event struct {
	Title   string `json:"Title"`
	Start   string `json:"StartTime"`
	End     string `json:"EndTime"`
	Remind  string `json:"RemindTime"`
	Context string `json:"Context"`
	Rurl    string `json:"ReferenceUrl"`
}

func Test_CalendarCreateEvent(t *testing.T) {
	//models.ConnectDataBase("davidleitw:davidleitw0308@/calendardb?charset=utf8&parseTime=True&loc=Local")
	models.ConnectDataBase("root:@(database)/calendardb?charset=utf8&parseTime=True&loc=Local")
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

	request := []struct {
		Event  event
		Status int
	}{
		{
			Event: event{
				Title:   "0123" + RandString(6),
				Start:   "2014-01-21 12:45:32",
				End:     "2014-01-21 12:55:32",
				Remind:  "-3m",
				Context: "Test event 001",
				Rurl:    "gamer.com.tw",
			},
			Status: 200,
		},
		{
			Event: event{
				Title:   "0123",
				Start:   "2014-01-21 12:45:32",
				End:     "2014-01-21 12:55:32",
				Remind:  "-3m",
				Context: "Test event 001",
				Rurl:    "gamer.com.tw",
			},
			Status: 400,
		},
	}

	for _, req := range request {
		// login, 取得登入狀態的session
		loginBody, _ := json.Marshal(user)
		Reqlogin := httptest.NewRequest("POST", "/api/user/userLogin", bytes.NewReader(loginBody))
		Reslogin := httptest.NewRecorder()
		server.ServeHTTP(Reslogin, Reqlogin)

		// 先把登入之後獲得的response print出來, 取得裡面mysession的value
		// 以便於後續測試登入過後的api使用
		// create event if login

		reqBody, _ := json.Marshal(req.Event)
		r := httptest.NewRequest("POST", "/api/calendar/a001/createNewEvent",
			bytes.NewReader(reqBody))
		// 設置登入過後狀態cookie, 因為本專案是使用cookie base 的 session.
		r.Header.Set("Content-Type", "application/json; charset=utf-8")
		r.Header.Set("Cookie", "mysession=MTU5MjIwODAyNnxEdi1CQkFFQ180SUFBUkFCRUFBQVF2LUNBQUlHYzNSeWFXNW5EQXNBQ1d4dloybHVkWE5sY2daemRISnBibWNNQmdBRVlUQXdNUVp6ZEhKcGJtY01DUUFIYVhOc2IyZHBiZ1JpYjI5c0FnSUFBUT09fPJ1meA-rBjslCFThdcgabnOSJdhYgE8OBls0HBdh6VJ")

		w := httptest.NewRecorder()
		server.ServeHTTP(w, r)
		assert.Equal(t, req.Status, w.Code)
	}
}
