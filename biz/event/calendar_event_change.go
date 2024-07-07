package event

import (
	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/event_sync"
	"github.com/larksuite/oapi-sdk-go/core"
	calendar "github.com/larksuite/oapi-sdk-go/service/calendar/v4"
	"github.com/sirupsen/logrus"
)

func CalendarEventChangeV4Handler(ctx *core.Context, event *calendar.CalendarEventChangedEvent) (err error) {
	logger := logrus.WithContext(ctx)
	if event == nil || event.Event == nil {
		logger.Warnf("calendar.CalendarChangedEvent invalid")
		return nil
	}
	if len(event.Event.CalendarId) == 0 {
		logger.Warnf("calendar.CalendarChangedEvent calendarID invalid")
		return nil
	}

	logger = logger.WithField("calendar_id", event.Event.CalendarId)
	for _, userId := range event.Event.UserIdList {
		if len(userId.OpenId) == 0 {
			continue
		}
		err := event_sync.CalendarEventIncrSync(ctx, userId.OpenId, event.Event.CalendarId)
		if err != nil {
			logger.WithField("open_user_id", userId.OpenId).
				WithError(err).Errorf("event_sync.CalendarEventIncrSync failed")
		}
	}
	return nil
}
