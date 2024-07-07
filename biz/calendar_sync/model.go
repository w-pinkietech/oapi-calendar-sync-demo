package calendar_sync

type StartCalendarSyncReq struct {
	OpenUserID string `json:"open_user_id"`
	PageSize   int    `json:"page_size,omitempty"`
	PageToken  string `json:"page_token,omitempty"`
}

func (p *StartCalendarSyncReq) ToCalendarSyncParam() *CalendarSyncParam {
	if p == nil {
		return nil
	}
	return &CalendarSyncParam{
		OpenUserID: p.OpenUserID,
		PageSize:   p.PageSize,
		PageToken:  p.PageToken,
	}
}

type CalendarSyncParam struct {
	UserAccessToken string `json:"user_access_token"`
	OpenUserID      string `json:"open_user_id"`
	PageSize        int    `json:"page_size"`
	PageToken       string `json:"page_token"`
	SyncToken       string `json:"sync_token"`
}
