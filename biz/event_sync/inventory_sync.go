package event_sync

import (
	"context"
	"net/http"

	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/event_sync/dal"
	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/oapi/calendarv4/event"
	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/oapi/token"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func StartEventSync(c *gin.Context) {
	req := &StartEventSyncReq{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "request param invalid",
		})
		return
	}
	if req.PageSize < 50 {
		req.PageSize = 50
	}
	if req.PageSize > 500 {
		req.PageSize = 500
	}

	param := req.ToCalendarEventSyncParam()

	userAccessToken, err := token.GetUserAccessToken(c, param.OpenUserID)
	if err != nil {
		logrus.WithError(err).Errorf("token.GetUserAccessToken failed")
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	param.UserAccessToken = userAccessToken

	if err := inventoryEventSync(c, param); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"err":         err.Error(),
			"msg":         "inventoryEventSync failed",
			"calendar_id": param.CalendarID,
			"anchor_time": param.AnchorTime,
			"page_size":   param.PageSize,
			"page_token":  param.PageToken,
			"sync_token":  param.SyncToken,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":         "success",
		"calendar_id": param.CalendarID,
		"anchor_time": param.AnchorTime,
		"page_size":   param.PageSize,
		"page_token":  param.PageToken,
		"sync_token":  param.SyncToken,
	})
	return
}

func inventoryEventSync(ctx context.Context, req *CalendarEventSyncParam) (err error) {
	logger := logrus.WithContext(ctx)

	// 1、初次获取日程列表，page_token和sync_token传空
	// 2、再次获取日程列表，带上第一次返回的sync_token和page_token
	var hasMore = true
	for hasMore || len(req.SyncToken) == 0 {
		hasMore, err = ListEvent(ctx, req)
		if err != nil {
			logger.WithError(err).Errorf("event_sync.ListEvent failed")
			return
		}
	}
	// 3、直到has_more为false，完成存量日程同步
	// 4、订阅日程变更事件
	go func() {
		err := event.AddCalendarEventSubscription(ctx, req.UserAccessToken)
		if err != nil {
			logger.WithError(err).Errorf("AddCalendarSubscription failed")
		}
	}()
	return nil
}

func ListEvent(ctx context.Context, req *CalendarEventSyncParam) (hasMore bool, err error) {
	logger := logrus.WithContext(ctx).
		WithField("calendar_id", req.CalendarID).
		WithField("page_size", req.PageSize).
		WithField("page_token", req.PageToken).
		WithField("sync_token", req.SyncToken)

	result, err := event.ListEvent(ctx, req.UserAccessToken, req.CalendarID, req.PageSize, req.AnchorTime, req.PageToken, req.SyncToken)
	if err != nil {
		logger.WithError(err).Errorf("ListEvent failed")
		return
	}
	err = dal.CreateOrUpdateCalendarEvent(ctx, req.CalendarID, result.Items)
	if err != nil {
		logger.WithError(err).Errorf("CreateOrUpdateCalendarEvent failed")
		return false, err
	}
	err = dal.CreateOrUpdateCalendarEventSyncStatus(ctx, req.OpenUserID, req.CalendarID, req.AnchorTime, req.PageSize, result.PageToken, result.SyncToken)
	if err != nil {
		logger.WithError(err).Errorf("CreateOrUpdateCalendarEventSyncStatus failed")
		return false, err
	}
	req.PageToken = result.PageToken
	req.SyncToken = result.SyncToken
	return result.HasMore, nil
}
