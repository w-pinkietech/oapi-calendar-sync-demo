package main

import (
	"net/http"

	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/calendar_sync"
	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/event_sync"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// ルートパスのハンドラーを追加
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Calendar Sync API",
		})
	})

	// 既存のルート
	r.POST("/calendar_sync", calendar_sync.StartCalendarSync)
	r.POST("/event_sync", event_sync.StartEventSync)
}
