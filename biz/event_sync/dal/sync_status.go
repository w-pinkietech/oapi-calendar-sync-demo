package dal

import "context"

func CreateOrUpdateCalendarEventSyncStatus(
	ctx context.Context,
	openUserID string,
	calendarID string,
	anchorTime string,
	pageSize int,
	pageToken string,
	syncToken string) (err error) {
	panic("implement me")
	return err
}

func GetCalendarEventSyncStatus(ctx context.Context, openUserID string, calendarID string) (
	anchorTime string, pageSize int, pageToken string, syncToken string, err error) {
	panic("implement me")
	return
}
