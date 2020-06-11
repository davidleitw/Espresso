package api

import (
	"Espresso/models"
	"Espresso/serialization"
	"Espresso/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventItem struct {
	Title     string
	StartTime string
}

// @Summary 登入時該用戶所有資料的概要
// @Tags Calendar
// @version 1.0
// @accept application/json
// @produce application/json
// @Router /api/calendar/{ID}/getAllEvent [get]
// @Success 200 {object} serialization.Response{}
// @Failure 500 {object} serialization.Response
func CalendarGetAllEvent(ctx *gin.Context) {
	var EventSet []EventItem
	var EventDetailSet []models.EventDetail
	email := models.GetFullEmail(ctx.Param("ID"))

	err := models.DB.Model(&models.EventDetail{}).Where("user_id=?", email).Find(&EventDetailSet).Error
	for _, item := range EventDetailSet {
		cID := item.CalendarID
		var EventInfo models.EventMain
		models.DB.Model(&models.EventMain{}).Where("calendar_id=?", cID).First(&EventInfo)
		item := EventItem{Title: EventInfo.Title, StartTime: EventInfo.StartTime}
		EventSet = append(EventSet, item)
	}

	if err == nil {
		// return 一個陣列 有著該使用者所有事件的title, start time.
		ctx.JSON(
			http.StatusOK,
			serialization.BuildResponse(http.StatusOK, EventSet, "獲得所有資料"))
	} else {
		// 不明原因無法查詢 也許是因為找不到該使用者
		ctx.JSON(
			http.StatusInternalServerError,
			serialization.BuildResponse(http.StatusInternalServerError, "null", "獲得資料失敗"))
	}
}

// @Summary 獲得某個特定事件的資料
// @Tags Calendar
// @version 1.0
// @accept application/json
// @produce application/json
// @param title header string true "填入想要查詢事件的標題"
// @param start_time header string true "填入想要查詢事件的開始時間"
// @param remind_time header string true "填入想要查詢事件需要提前幾分鐘提醒 以-3h這種形式填入(創建事件要填的那個寫法)"
// @Router /api/calendar/{ID}/getEventInfo  [post]
// @Success 200
// @Failure 500
func CalendarGetEventInfo(ctx *gin.Context) {
	var servicer service.GetEventInfoPoster
	if err := ctx.ShouldBind(&servicer); err == nil {
		res := servicer.CalendarGetEventInfo(ctx.Param("ID"))
		ctx.JSON(serialization.GetResStatus(res).(int), res)
	}
}

// @Summary 更新某個事件的內容
// @Tags Calendar
// @version 1.0
// @accept application/json
// @produce application/json
// @param OldTitle header string true "填入修改前的事件標題"
// @param OldStart header string true "填入修改前的開始時間"
// @param OldRemind header string true "填入修改前的欲提醒時間"
// @param Title header string true "填入修改後的新標題"
// @param StartTime header string true "填入修改後的新開始時間"
// @param EndTime header string true "填入修改後的新結束時間"
// @param RemindTime header string true "填入修改後的新提醒時間 以-3h這種形式"
// @param Context header string true "填入修改後的事件內容"
// @param ReferenceUrl header string true "填入修改後的參考網址"
// @Router /api/calendar/{ID}/updateEvent [put]
// @Success 204
// @Failure 500
func CalendarUpdateEvent(ctx *gin.Context) {
	var servicer service.UpdateEventPoster
	if err := ctx.ShouldBind(&servicer); err == nil {
		res := servicer.CalendarUpdateEvent(ctx.Param("ID"))
		ctx.JSON(serialization.GetResStatus(res).(int), res)
	}
}

// @Summary 創建事件
// @Tags Calendar
// @version 1.0
// @accept application/json
// @produce application/json
// @param Title header string true "填入事件標題"
// @param StartTime header string true "填入事件開始時間"
// @param EndTime header string true "填入事件的結束時間"
// @param RemindTime header string true "填入事件想要提前幾分鐘提醒 以-3h這種形式"
// @param Context header string true "填入事件內容"
// @parma ReferenceUrl header string true "填入事件的參考網址"
// @Router /api/calendar/{ID}/createNewEvent [post]
// @Success 200
// @Failure 400
// @Failure 500
func CalendarCreateEvent(ctx *gin.Context) {
	var servicer service.CreateEventPoster
	if err := ctx.ShouldBind(&servicer); err == nil {
		res := servicer.CalendarCreateEvent(ctx.Param("ID"))
		ctx.JSON(serialization.GetResStatus(res).(int), res)
	}
}

// @Summary 刪除指定的事件
// @Tags Calendar
// @version 1.0
// @accept application/json
// @produce application/json
// @param title header string true "填入想要刪除事件的標題"
// @param start_time header string true "填入想要刪除事件的開始時間"
// @param remind_time header string true "填入想要刪除事件需要提前幾分鐘提醒 以-3h這種形式填入(創建事件要填的那個寫法)"
// @Router /api/calendar/{ID}/deleteEvent [delete]
// @Success 204
// @Failure 500
func CalendarDeleteEvent(ctx *gin.Context) {
	var servicer service.CalendarDeletePoster
	//session := sessions.Default(ctx)

	if err := ctx.ShouldBind(&servicer); err == nil {
		res := servicer.Delete(ctx.Param("ID"))
		ctx.JSON(serialization.GetResStatus(res).(int), res)
	}
}
