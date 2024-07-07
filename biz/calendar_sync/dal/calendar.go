package dal

import (
	"context"

	calendarV4 "github.com/larksuite/oapi-sdk-go/service/calendar/v4"
)

func CreateOrUpdateCalendar(ctx context.Context, openUserID string, calendarList []*calendarV4.Calendar) (err error) {
	panic("implement me")
	return nil
}
