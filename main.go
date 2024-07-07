package main

import (
	"fmt"

	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/event"
	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/oapi"
	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := utils.InitConfig(); err != nil {
		logrus.Fatalf("InitConfig failed")
	}
	oapi.Init()
	r := gin.Default()

	// event callback address
	event.AddEventWebhook(r)

	// other feature
	Router(r)

	// start http server
	if err := r.Run(fmt.Sprintf("0.0.0.0:%v", utils.Config.HttpServerPort)); err != nil {
		logrus.WithError(err).Errorf("http server start failed")
		return
	}
	logrus.Warnf("http server exit")
}
