package service

import (
	"github.com/karimatiket/karimatiket-auth/internal/config"
	"golang.org/x/oauth2"
)

type Auth struct {
	Config *config.Config
}

func NewAuth(config *config.Config) *Auth {
	return &Auth{
		Config: config,
	}
}

func (a *Auth) GetGoogleAuthURL() string {
	url := a.Config.GoogleLoginConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return url
}

func (a *Auth) GetDiscordAuthURL() string {
	url := a.Config.DiscordLoginConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return url
}
