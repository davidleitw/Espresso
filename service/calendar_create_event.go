package service

import (
	"Espresso/models"
	serial "Espresso/serialization"
	"net/http"
)

type CreateEventPoster struct {
	ID      string `json:"User_ID"`
	Ti      string `json: "Title"`
	St      string `json: "StartTime"`
	Et      string `json: "EndTime"`
	Rt      string `json: "RemindTime"`
	Context string `json: "Context"`
}

func (service *CreateEventPoster) CalendarCreateEvent() serial.Response {
	event := &models.Event{
		User_ID:    service.ID,
		Title:      service.Ti,
		StartTime:  service.St,
		EndTime:    service.Et,
		RemindTime: service.Rt,
		Context:    service.Context,
	}

	err := models.DB.Create(&event).Error

	if err == nil {
		return serial.BuildResponse(http.StatusOK, service.Ti, "新增事件成功")
	}
	return serial.BuildResponse(404, "null", "新增事件失敗")

}
