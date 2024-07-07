package dal

import "context"

func CreateOrUpdateCalendarSyncStatus(
	ctx context.Context,
	openUserID string,
	pageSize int,
	pageToken string,
	syncToken string) (err error) {
	panic("implement me")
	return nil
}

func GetCalendarSyncStatus(ctx context.Context, openUserID string) (pageSize int, pageToken string, syncToken string, err error) {
	panic("implement me")
	return
}
