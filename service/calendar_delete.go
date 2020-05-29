package service

import (
	"Espresso/models"
	serial "Espresso/serialization"
)

type CalendarDelete struct {
	Title     string `json:"title"`
	StartTime string `json:"start_time"`
}

func (service *CalendarDelete) Delete(userID string) serial.Response {
	var event models.Event

	models.DB.Where("user_id = ? AND title = ? AND start_time = ?", models.GetFullEmail(userID), service.Title, service.StartTime).Delete(&event)
	return serial.BuildResponse(200, "null", "刪除成功")

}
