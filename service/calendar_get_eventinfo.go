package service

import (
	"Espresso/models"
	serial "Espresso/serialization"
	"net/http"
)

type GetEventInfoPoster struct {
	Title      string `json: "title"`
	StartTime  string `json: "start_time"`
	RemindTime string `json: "remind_time"`
}

type EventInfo struct {
	Title        string
	Context      string
	StartTime    string
	RemindTime   string
	EndTime      string
	ReferenceUrl string
}

func (service *GetEventInfoPoster) CalendarGetEventInfo(userID string) serial.Response {
	var Info models.EventDetail
	var Em models.EventMain
	var Ed models.EventDetail

	email := models.GetFullEmail(userID)
	remindtime := models.GetRemindTime(service.StartTime, service.RemindTime)

	models.DB.Model(&models.EventDetail{}).Where(
		"user_id=? AND title=? AND remind_time=?",
		email, service.Title, remindtime,
	).First(&Info)

	err1 := models.DB.Where(&models.EventMain{}).Where("calendar_id=?", Info.CalendarID).First(&Em).Error
	err2 := models.DB.Where(&models.EventDetail{}).Where("calendar_id=?", Info.CalendarID).First(&Ed).Error

	if err1 == nil && err2 == nil {
		// http.StatusOk => 200
		return serial.BuildResponse(
			http.StatusOK,
			EventInfo{
				Title:        Em.Title,
				Context:      Em.Context,
				StartTime:    Em.StartTime,
				RemindTime:   Ed.RemindTime,
				EndTime:      Em.EndTime,
				ReferenceUrl: Em.ReferenceUrl,
			},
			"獲得資訊成功",
		)
	} else {
		return serial.BuildResponse(
			http.StatusInternalServerError,
			"null",
			"事件不存在或者存在某些問題",
		)
	}

}
