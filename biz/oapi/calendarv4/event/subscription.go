package event

import (
	"context"

	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/oapi"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/sirupsen/logrus"
)

func AddCalendarEventSubscription(ctx context.Context, userAccessToken string) (err error) {
	logger := logrus.WithContext(ctx)
	coreCtx := core.WrapContext(ctx)
	reqCall := oapi.CalendarService.CalendarEvents.Subscription(coreCtx, request.SetUserAccessToken(userAccessToken))
	_, err = reqCall.Do()
	logger = logger.WithField("request_id", coreCtx.GetRequestID()).
		WithField("status_code", coreCtx.GetHTTPStatusCode())
	if err != nil {
		logger.WithError(err).Errorf("AddCalendarEventSubscription call do failed")
		return err
	}
	logger.Infof("AddCalendarEventSubscription finish")
	return err
}
