package service

import (
	"Espresso/models"
	serial "Espresso/serialization"
	"net/http"
)

type CalendarDelete struct {
	Title     string `json:"title"`
	StartTime string `json:"start_time"`
}

func (service *CalendarDelete) Delete(userID string) serial.Response {
	var event models.Event

	err := models.DB.Where("user_id = ? AND title = ? AND start_time = ?", models.GetFullEmail(userID), service.Title, service.StartTime).Delete(&event).Error
	if err == nil {
		// 204 刪除成功 沒有回傳值
		return serial.BuildResponse(http.StatusNoContent, "null", "刪除成功")
	} else {
		// 500 伺服器有某些地方出問題了
		return serial.BuildResponse(http.StatusInternalServerError, "null", "刪除失敗")
	}

}
