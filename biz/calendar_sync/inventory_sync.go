package calendar_sync

import (
	"context"
	"net/http"

	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/calendar_sync/dal"
	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/oapi/calendarv4/calendar"
	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/biz/oapi/token"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func StartCalendarSync(c *gin.Context) {
	req := &StartCalendarSyncReq{}
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

	param := req.ToCalendarSyncParam()

	userAccessToken, err := token.GetUserAccessToken(c, param.OpenUserID)
	if err != nil {
		logrus.WithError(err).Errorf("token.GetUserAccessToken failed")
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	param.UserAccessToken = userAccessToken

	if err := inventoryCalendarSync(c, param); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"err":        err.Error(),
			"msg":        "inventoryCalendarSync failed",
			"page_size":  param.PageSize,
			"page_token": param.PageToken,
			"sync_token": param.SyncToken,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":        "success",
		"page_size":  param.PageSize,
		"page_token": param.PageToken,
		"sync_token": param.SyncToken,
	})
	return
}

func inventoryCalendarSync(ctx context.Context, req *CalendarSyncParam) (err error) {
	logger := logrus.WithContext(ctx)

	// 1、初次获取日历列表，page_token和sync_token传空即可
	// 2、再次获取日历列表，传入上一次获取的page_token和sync_token
	var hasMore = true
	for hasMore || len(req.SyncToken) == 0 {
		hasMore, err = ListCalendar(ctx, req)
		if err != nil {
			logger.WithError(err).Errorf("calendar_sync.ListCalendar failed")
			return
		}
	}
	// 3、直到has_more为false，sync_token返回值不为空时，完成存量日历列表的同步
	// 4、订阅日历变更事件
	go func() {
		err := calendar.AddCalendarSubscription(ctx, req.UserAccessToken)
		if err != nil {
			logger.WithError(err).Errorf("AddCalendarSubscription failed")
		}
	}()
	return nil
}

func ListCalendar(ctx context.Context, req *CalendarSyncParam) (hasMore bool, err error) {
	logger := logrus.WithContext(ctx).
		WithField("page_size", req.PageSize).
		WithField("page_token", req.PageToken).
		WithField("sync_token", req.SyncToken)

	result, err := calendar.ListCalendar(ctx, req.UserAccessToken, req.PageSize, req.PageToken, req.SyncToken)
	if err != nil {
		logger.WithError(err).Errorf("ListCalendar failed")
		return
	}
	err = dal.CreateOrUpdateCalendar(ctx, req.OpenUserID, result.CalendarList)
	if err != nil {
		logger.WithError(err).Errorf("CreateOrUpdateCalendar failed")
		return false, err
	}
	err = dal.CreateOrUpdateCalendarSyncStatus(ctx, req.OpenUserID, req.PageSize, result.PageToken, result.SyncToken)
	if err != nil {
		logger.WithError(err).Errorf("CreateOrUpdateCalendar failed")
		return false, err
	}
	req.PageToken = result.PageToken
	req.SyncToken = result.SyncToken
	return result.HasMore, nil
}
