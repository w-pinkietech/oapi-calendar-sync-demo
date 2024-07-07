package utils

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var (
	Config   cfg
	loadPath = "./conf/config.yml"
)

type cfg struct {
	HttpServerPort int32       `yaml:"HttpServerPort"`
	AppSettings    AppSettings `yaml:"AppSettings"`
}

type AppSettings struct {
	AppCredentials      AppCredentials      `yaml:"AppCredentials"`
	AppEventKey         AppEventKey         `yaml:"AppEventKey"`
	HelpDeskCredentials HelpDeskCredentials `yaml:"HelpDeskCredentials"`
}

type AppCredentials struct {
	AppID     string `yaml:"AppID"`
	AppSecret string `yaml:"AppSecret"`
}

type AppEventKey struct {
	EnableEncrypt     bool   `yaml:"EnableEncrypt"`
	VerificationToken string `yaml:"VerificationToken"`
	EncryptKey        string `yaml:"EncryptKey"`
}

type HelpDeskCredentials struct {
	EnableHelpDesk bool   `yaml:"EnableHelpDesk"`
	HelpDeskID     string `yaml:"HelpDeskID"`
	HelpDeskToken  string `yaml:"HelpDeskToken"`
}

func InitConfig() (err error) {
	configByte, err := ioutil.ReadFile(loadPath)
	if err != nil {
		logrus.WithError(err).Errorf("read config file failed")
		return err
	}
	err = yaml.Unmarshal(configByte, &Config)
	if err != nil {
		logrus.WithError(err).Errorf("yaml unmarshal failed")
		return err
	}
	return nil
}
