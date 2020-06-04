package service

import (
	"Espresso/models"
	serial "Espresso/serialization"
	"net/http"
)

type UpdateEventPoster struct {
	ID      string `json:"User_ID"`
	Title   string `json:"Title"`
	Start   string `json:"StartTime"`
	End     string `json:"EndTime"`
	Remind  string `json:"RemindTime"`
	Context string `json:"Context"`
	Rurl    string `json:"ReferenceUrl"`
}

func (service *UpdateEventPoster) CalendarUpdateEvent() serial.Response {
	email := models.GetFullEmail(service.ID)
	rtime := models.GetResultTime(service.Remind, models.GetTimeValue(service.Start))
	var Em models.EventMain
	var Ed models.EventDetail

	models.DB.Where(
		"user_id=? AND title=? AND start_time=?",
		email, service.Title, service.Start,
	).First(&Em)
	models.DB.Where(
		"user_id=? AND title=? AND start_time=?",
		email, service.Title, service.Start,
	).First(&Ed)

	Em.StartTime = service.Start
	Em.EndTime = service.End
	Em.Title = service.Title
	Em.Context = service.Context
	Em.ReferenceUrl = service.Rurl

	Ed.UserID = models.GetFullEmail(service.ID)
	Ed.RemindTime = models.GetTimeString(rtime)

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
