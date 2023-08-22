package config

import "os"

type AppConfig struct {
	AppName          string
	AppVersion       string
	Env              string
	AppWebserverPort string
	AppSecret        string
	AppUIPort        string
	AppLocation      string
	NodeBinLocation  string
}

func LoadAppConfig() AppConfig {
	return AppConfig{
		AppName:          os.Getenv("APP_NAME"),
		AppVersion:       os.Getenv("APP_VERSION"),
		Env:              os.Getenv("APP_ENV"),
		AppWebserverPort: os.Getenv("APP_WEBSERVER_PORT"),
		AppSecret:        os.Getenv("APP_SECRET"),
		AppLocation:      os.Getenv("APP_LOCATION"),
	}
}
