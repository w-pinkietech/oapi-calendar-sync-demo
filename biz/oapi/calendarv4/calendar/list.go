package calendar

import (
	"context"

	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/oapi"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/core"
	calendarV4 "github.com/larksuite/oapi-sdk-go/service/calendar/v4"
	"github.com/sirupsen/logrus"
)

func ListCalendar(ctx context.Context, userAccessToken string, pageSize int, pageToken string, syncToken string) (result *calendarV4.CalendarListResult, err error) {
	logger := logrus.WithContext(ctx)
	coreCtx := core.WrapContext(ctx)
	reqCall := oapi.CalendarService.Calendars.List(coreCtx, request.SetUserAccessToken(userAccessToken))
	reqCall.SetPageSize(pageSize)
	reqCall.SetPageToken(pageToken)
	reqCall.SetSyncToken(syncToken)
	result, err = reqCall.Do()
	logger = logger.WithField("request_id", coreCtx.GetRequestID()).
		WithField("status_code", coreCtx.GetHTTPStatusCode())
	if err != nil {
		logger.WithError(err).Errorf("ListCalendar call do failed")
		return nil, err
	}
	logger.Infof("ListCalendar finish")
	return result, err
}
