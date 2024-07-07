package event_sync

import (
	"context"

	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/event_sync/dal"
	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/oapi/token"
	"github.com/sirupsen/logrus"
)

func CalendarEventIncrSync(ctx context.Context, openUserID string, calendarID string) (err error) {
	logger := logrus.WithContext(ctx)

	anchorTime, pageSize, pageToken, syncToken, err := dal.GetCalendarEventSyncStatus(ctx, openUserID, calendarID)
	if err != nil {
		logger.WithError(err).Errorf("dal.GetCalendarEventSyncStatus failed")
		return err
	}

	userAccessToken, err := token.GetUserAccessToken(ctx, openUserID)
	if err != nil {
		logger.WithError(err).Errorf("token.GetUserAccessToken failed")
		return err
	}

	req := &CalendarEventSyncParam{
		OpenUserID:      openUserID,
		UserAccessToken: userAccessToken,
		CalendarID:      calendarID,
		AnchorTime:      anchorTime,
		PageSize:        pageSize,
		PageToken:       pageToken,
		SyncToken:       syncToken,
	}

	var hasMore = true
	for hasMore || len(req.SyncToken) == 0 {
		hasMore, err = ListEvent(ctx, req)
		if err != nil {
			logger.WithError(err).Errorf("event_sync.ListEvent failed")
			return
		}
	}

	return err
}
