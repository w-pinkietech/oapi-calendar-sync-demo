package event

import (
	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/calendar_sync"
	"github.com/larksuite/oapi-sdk-go/core"
	calendar "github.com/larksuite/oapi-sdk-go/service/calendar/v4"
	"github.com/sirupsen/logrus"
)

func CalendarChangeV4Handler(ctx *core.Context, event *calendar.CalendarChangedEvent) (err error) {
	logger := logrus.WithContext(ctx)
	if event == nil || event.Event == nil {
		logger.Warnf("calendar.CalendarChangedEvent invalid")
		return nil
	}
	for _, userId := range event.Event.UserIdList {
		if len(userId.OpenId) == 0 {
			continue
		}
		err := calendar_sync.CalendarIncrSync(ctx, userId.OpenId)
		if err != nil {
			logger.WithField("open_user_id", userId.OpenId).
				WithError(err).Errorf("calendar_sync.CalendarIncrSync failed")
		}
	}
	return nil
}
