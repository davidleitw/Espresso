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

func CalendarGetAllEvent(ctx *gin.Context) {
	var EventSet []EventItem
	var EventMainSet []models.EventMain
	email := models.GetFullEmail(ctx.Param("ID"))

	err := models.DB.Where("user_id=?", email).Find(&EventMainSet).Error
	for _, item := range EventMainSet {
		item := EventItem{Title: item.Title, StartTime: item.StartTime}
		EventSet = append(EventSet, item)
	}

	if err != nil {
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

func CalendarGetEventInfo(ctx *gin.Context) {
	var servicer service.GetEventInfoPoster
	if err := ctx.ShouldBind(&servicer); err == nil {
		res := servicer.CalendarGetEventInfo(ctx.Param("ID"))
		ctx.JSON(serialization.GetResStatus(res).(int), res)
	}
}

func CalendarUpdateEvent(ctx *gin.Context) {
	var servicer service.UpdateEventPoster
	if err := ctx.ShouldBind(&servicer); err == nil {
		res := servicer.CalendarUpdateEvent()
		ctx.JSON(serialization.GetResStatus(res).(int), res)
	}
}

func CalendarCreateEvent(ctx *gin.Context) {
	var servicer service.CreateEventPoster
	if err := ctx.ShouldBind(&servicer); err == nil {
		res := servicer.CalendarCreateEvent()
		ctx.JSON(serialization.GetResStatus(res).(int), res)
	}
}

func CalendarDeleteEvent(ctx *gin.Context) {
	var servicer service.CalendarDeletePoster
	//session := sessions.Default(ctx)

	if err := ctx.ShouldBind(&servicer); err == nil {
		res := servicer.Delete(ctx.Param("ID"))
		ctx.JSON(serialization.GetResStatus(res).(int), res)
	}
}
