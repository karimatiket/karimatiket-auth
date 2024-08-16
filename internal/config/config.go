package config

import "golang.org/x/oauth2"

type Config struct {
	GoogleLoginConfig  oauth2.Config
	DiscordLoginConfig oauth2.Config
}

var AppConfig Config
