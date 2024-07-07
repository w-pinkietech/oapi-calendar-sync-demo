package dal

import (
	"context"

	calendar "github.com/larksuite/oapi-sdk-go/service/calendar/v4"
)

func CreateOrUpdateCalendarEvent(ctx context.Context, calendarID string, eventList []*calendar.CalendarEvent) (err error) {
	panic("implement me")
	return err
}
