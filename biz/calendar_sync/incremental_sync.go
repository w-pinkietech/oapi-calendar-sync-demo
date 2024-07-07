package calendar_sync

import (
	"context"

	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/calendar_sync/dal"
	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/oapi/token"
	"github.com/sirupsen/logrus"
)

func CalendarIncrSync(ctx context.Context, openUserID string) (err error) {
	logger := logrus.WithContext(ctx)

	pageSize, pageToken, syncToken, err := dal.GetCalendarSyncStatus(ctx, openUserID)
	if err != nil {
		logger.WithError(err).Errorf("dal.GetCalendarSyncStatus failed")
		return err
	}

	userAccessToken, err := token.GetUserAccessToken(ctx, openUserID)
	if err != nil {
		logger.WithError(err).Errorf("token.GetUserAccessToken failed")
		return err
	}

	req := &CalendarSyncParam{
		UserAccessToken: userAccessToken,
		OpenUserID:      openUserID,
		PageSize:        pageSize,
		PageToken:       pageToken,
		SyncToken:       syncToken,
	}

	var hasMore = true
	for hasMore {
		hasMore, err = ListCalendar(ctx, req)
		if err != nil {
			logger.WithError(err).Errorf("calendar_sync.ListCalendar failed")
			return
		}
	}
	return nil
}
