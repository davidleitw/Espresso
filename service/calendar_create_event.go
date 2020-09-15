package service

import (
	"Espresso/models"
	serial "Espresso/serialization"
	"net/http"
	"time"

	"github.com/rs/xid"
)

type CreateEventPoster struct {
	Title   string `json:"Title"`
	Start   string `json:"StartTime"`
	End     string `json:"EndTime"`
	Remind  string `json:"RemindTime"`
	Context string `json:"Context"`
	Rurl    string `json:"ReferenceUrl"`
}

func (service *CreateEventPoster) CalendarCreateEvent(userID string) serial.Response {
	// 需要的information
	// 直接傳入資料庫的值StartTime, EndTime, Title, UserID, Context, ReferenceUrl
	// 需要計算出來的值 RemindTime, CalendarID

	// 建立一個唯一的Calendar ID
	guid := xid.New()
	email := models.GetFullEmail(userID)
	rTime := models.GetTimeString(models.GetResultTime(service.Remind, models.GetTimeValue(service.Start)))
	createTime := time.Now()

	em := &models.EventMain{
		CalendarID:   guid.String(),                    // 唯一的calendarID
		CreateTime:   models.GetTimeString(createTime), // 開始時間以字串形式
		StartTime:    service.Start,                    // 開始時間 以string存入database
		EndTime:      service.End,                      // 結束時間 以string形式存入database
		Title:        service.Title,                    // 標題
		Context:      service.Context,                  // 內容
		ReferenceUrl: service.Rurl,                     // 參考網址
		Remind:       service.Remind,
	}

	ed := &models.EventDetail{
		CalendarID: guid.String(), // 唯一的calendarID
		UserID:     email,         // 用戶的電子信箱
		Title:      service.Title, // 標題
		Creator:    true,          // 是否為此行程的創建人
		RemindTime: rTime,         // 預計提醒時間
		Accept:     true,          // 是否接受邀約
	}

	// 避免重複查詢
	var count int = 0
	models.DB.Model(&models.EventDetail{}).Where(
		"user_id=? AND title=? AND remind_time=?",
		email, service.Title, rTime,
	).Count(&count)

	if count != 0 {
		// http.StatusBadRequest => 400 請求的內容有誤
		return serial.BuildResponse(http.StatusBadRequest, "null", "請勿重複新增事件")
	}

	err1 := models.DB.Create(&em).Error
	err2 := models.DB.Create(&ed).Error

	if err1 == nil && err2 == nil && count == 0 {
		// http.StatusOk => 200 請求成功
		return serial.BuildResponse(http.StatusOK, service.Title, "新增事件成功")
	}
	// http.StatusInternamServerError => 500
	return serial.BuildResponse(http.StatusInternalServerError, "null", "新增事件失敗")

}
