package event_sync

type StartEventSyncReq struct {
	OpenUserID string `json:"open_user_id"`
	CalendarID string `json:"calendar_id"`
	AnchorTime string `json:"anchor_time,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	PageToken  string `json:"page_token,omitempty"`
}

func (p *StartEventSyncReq) ToCalendarEventSyncParam() *CalendarEventSyncParam {
	if p == nil {
		return nil
	}
	return &CalendarEventSyncParam{
		OpenUserID: p.OpenUserID,
		CalendarID: p.CalendarID,
		AnchorTime: p.AnchorTime,
		PageSize:   p.PageSize,
		PageToken:  p.PageToken,
	}
}

type CalendarEventSyncParam struct {
	OpenUserID      string `json:"open_user_id"`
	UserAccessToken string `json:"user_access_token"`
	CalendarID      string `json:"calendar_id"`
	AnchorTime      string `json:"anchor_time"`
	PageSize        int    `json:"page_size"`
	PageToken       string `json:"page_token"`
	SyncToken       string `json:"sync_token"`
}
