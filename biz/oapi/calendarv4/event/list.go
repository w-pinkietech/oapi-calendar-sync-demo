package event

import (
	"context"

	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/oapi"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/core"
	calendarV4 "github.com/larksuite/oapi-sdk-go/service/calendar/v4"
	"github.com/sirupsen/logrus"
)

func ListEvent(ctx context.Context, userAccessToken string, calendarID string, pageSize int, anchorTime, pageToken, syncToken string) (result *calendarV4.CalendarEventListResult, err error) {
	logger := logrus.WithContext(ctx)
	coreCtx := core.WrapContext(ctx)
	reqCall := oapi.CalendarService.CalendarEvents.List(coreCtx, request.SetUserAccessToken(userAccessToken))
	reqCall.SetCalendarId(calendarID)
	reqCall.SetPageSize(pageSize)
	reqCall.SetAnchorTime(anchorTime)
	reqCall.SetPageToken(pageToken)
	reqCall.SetSyncToken(syncToken)
	result, err = reqCall.Do()
	logger = logger.WithField("request_id", coreCtx.GetRequestID()).
		WithField("status_code", coreCtx.GetHTTPStatusCode())
	if err != nil {
		logger.WithError(err).Errorf("ListEvent call do failed")
		return nil, err
	}
	logger.Infof("ListEvent finish")
	return result, err
}
