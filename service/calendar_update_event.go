package service

import (
	"Espresso/models"
	serial "Espresso/serialization"
	"net/http"
)

type UpdateEventPoster struct {
	Title   string `json:"Title"`
	Start   string `json:"StartTime"`
	End     string `json:"EndTime"`
	Remind  string `json:"RemindTime"`
	Context string `json:"Context"`
	Rurl    string `json:"ReferenceUrl"`
}

func (service *UpdateEventPoster) CalendarUpdateEvent(userID string) serial.Response {
	var Info models.EventDetail
	var Em models.EventMain
	var Ed models.EventDetail

	email := models.GetFullEmail(userID)
	remindtime := models.GetRemindTime(service.Start, service.Remind)

	models.DB.Model(&models.EventDetail{}).Where(
		"user_id=? AND title=? AND remind_time=?",
		email, service.Title, remindtime,
	).First(&Info)

	models.DB.Where(&models.EventMain{}).Where("calendar_id=?", Info.CalendarID).First(&Em)
	models.DB.Where(&models.EventDetail{}).Where("calendar_id=?", Info.CalendarID).First(&Ed)

	Em.StartTime = service.Start
	Em.EndTime = service.End
	Em.Title = service.Title
	Em.Context = service.Context
	Em.ReferenceUrl = service.Rurl

	Ed.UserID = email
	Ed.RemindTime = remindtime

	err1 := models.DB.Save(&Em).Error
	err2 := models.DB.Save(&Ed).Error

	if err1 == nil && err2 == nil {
		return serial.BuildResponse(
			// 204
			http.StatusNoContent,
			"null",
			"更新成功",
		)
	} else {
		return serial.BuildResponse(
			http.StatusInternalServerError,
			"null",
			"更新失敗",
		)
	}
}
