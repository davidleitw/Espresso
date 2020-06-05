package service

import (
	"Espresso/models"
	serial "Espresso/serialization"
	"net/http"
)

type CalendarDeletePoster struct {
	Title     string `json:"title"`
	StartTime string `json:"start_time"`
}

func (service *CalendarDeletePoster) Delete(userID string) serial.Response {
	// 刪除 先藉由 UserID, title, starttime 找到唯一的事件
	// 拿到ID之後刪除兩個表中有同ID的事件
	var emain models.EventMain
	//var edetail models.EventDetail

	var delmain models.EventMain
	var deldetail models.EventDetail

	email := models.GetFullEmail(userID)

	models.DB.Where(
		"user_id=? AND title=? AND start_time=?",
		email, service.Title, service.StartTime,
	).First(&email)

	calendarID := emain.CalendarID
	err1 := models.DB.Where("calendar_id=?", calendarID).Delete(&delmain).Error
	err2 := models.DB.Where("calendar_id=?", calendarID).Delete(&deldetail).Error

	if err1 == nil && err2 == nil {
		// http.StatusNoContent => 204 刪除成功 沒有回傳值
		return serial.BuildResponse(http.StatusNoContent, "null", "刪除成功.")
	} else {
		// http.StatusInternamServerError => 500 伺服器有某些地方出問題了
		return serial.BuildResponse(http.StatusInternalServerError, "null", "刪除失敗")
	}
}
