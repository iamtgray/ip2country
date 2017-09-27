package config

import (
	cfg "github.com/infinityworks/go-common/config"
)

// Config stores all of the runtime configuration for the blueprint library
type Config struct {
	*cfg.BaseConfig
	updateURL string
}

// Init configuration using environment variables
func Init() Config {
	ac := cfg.Init()

	defaultUpdateURL := "https://www.maxmind.com/app/update_secure?editionId=GeoLite2-Country"

	appConfig := Config{
		&ac,
		defaultUpdateURL,
	}

	return appConfig
}
