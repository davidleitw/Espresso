package api

import (
	"Espresso/models"
	"Espresso/serialization"
	"Espresso/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CalendarGetAllEvent(ctx *gin.Context) {
	var EventSet []models.Event
	// session := sessions.Default(ctx)
	// userID := models.GetFullEmail(session.Get("loginuser").(string))
	userID := models.GetFullEmail(ctx.Param("ID"))
	err := models.DB.Where("user_id = ?", userID).Find(&EventSet).Error

	if err == nil {
		ctx.JSON(http.StatusOK, serialization.BuildResponse(http.StatusOK, EventSet, "Get all event information."))
	} else {
		ctx.JSON(404, serialization.BuildResponse(404, "null", "獲得全部資料失敗"))
	}
}

func CalendarGetEventInfo(ctx *gin.Context) {

}

func CalendarUpdateEvent(ctx *gin.Context) {

}

func CalendarCreateEvent(ctx *gin.Context) {
	var servicer service.CreateEventPoster
	if err := ctx.ShouldBind(&servicer); err == nil {
		res := servicer.CalendarCreateEvent()
		if serialization.GetResStatus(res) == 200 {
			ctx.JSON(http.StatusOK, res)
		} else {
			ctx.JSON(404, res)
		}
	}
}

func DeleteEvent(ctx *gin.Context) {
	var servicer service.CalendarDelete
	//session := sessions.Default(ctx)
	username := ctx.Param("ID")

	if err := ctx.ShouldBind(&servicer); err == nil {
		res := servicer.Delete(username)
		ctx.JSON(http.StatusForbidden, res)
	}
}
