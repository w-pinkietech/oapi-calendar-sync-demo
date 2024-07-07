package oapi

import (
	"code.byted.org/larkcalendar/oapi_calendar_sync_demo/utils"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
	"github.com/larksuite/oapi-sdk-go/core/log"
	calendarV4 "github.com/larksuite/oapi-sdk-go/service/calendar/v4"
)

var (
	OapiConfig      *config.Config
	CalendarService *calendarV4.Service
)

// see https://github.com/larksuite/oapi-sdk-go
func Init() {
	appSettingsOpts := []config.AppSettingsOpt{core.SetAppCredentials(utils.Config.AppSettings.AppCredentials.AppID,
		utils.Config.AppSettings.AppCredentials.AppSecret)}

	if utils.Config.AppSettings.AppEventKey.EnableEncrypt {
		appSettingsOpts = append(appSettingsOpts, core.SetAppEventKey(utils.Config.AppSettings.AppEventKey.VerificationToken,
			utils.Config.AppSettings.AppEventKey.EncryptKey))
	}

	if utils.Config.AppSettings.HelpDeskCredentials.EnableHelpDesk {
		appSettingsOpts = append(appSettingsOpts, core.SetHelpDeskCredentials(utils.Config.AppSettings.HelpDeskCredentials.HelpDeskID,
			utils.Config.AppSettings.HelpDeskCredentials.HelpDeskToken))
	}
	appSettings := core.NewInternalAppSettings(appSettingsOpts...)
	OapiConfig = core.NewConfig(core.DomainFeiShu, appSettings,
		core.SetLogger(log.NewDefaultLogger()),
		core.SetLoggerLevel(core.LoggerLevelInfo))

	CalendarService = calendarV4.NewService(OapiConfig)
}
