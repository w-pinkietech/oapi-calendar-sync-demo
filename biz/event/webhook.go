package event

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/oapi"
	"github.com/gin-gonic/gin"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/event"
	calendar "github.com/larksuite/oapi-sdk-go/service/calendar/v4"
)

func AddEventWebhook(r *gin.Engine) {
	setEventHandler()

	r.POST("/webhook/event", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		oapi.OapiConfig.GetLogger().Info(c, "body = ", string(body))
		params := &struct {
			Challenge string `json:"challenge,omitempty" form:"challenge"`
		}{}
		err = json.Unmarshal(body, params)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		if len(params.Challenge) != 0 {
			c.JSON(http.StatusOK, params)
			return
		}
		//eventhttp.Handle(oapi.OapiConfig, c.Request, c.Writer)

		oapiRequest := &core.OapiRequest{
			Ctx:    c.Request.Context(),
			Uri:    c.Request.RequestURI,
			Header: core.NewOapiHeader(c.Request.Header),
			Body:   string(body),
		}
		err = event.Handle(oapi.OapiConfig, oapiRequest).WriteTo(c.Writer)
		if err != nil {
			oapi.OapiConfig.GetLogger().Error(oapiRequest.Ctx, err)
		}
	})
}

func setEventHandler() {
	// set calendar.calendar.changed_v4 handler
	calendar.SetCalendarChangedEventHandler(oapi.OapiConfig, CalendarChangeV4Handler)

	// set calendar.calendar.event.changed_v4 handler
	calendar.SetCalendarEventChangedEventHandler(oapi.OapiConfig, CalendarEventChangeV4Handler)
}
