package test

import (
	"log"
	"testing"

	"github.com/karimatiket/karimatiket-auth/internal/config"
	"github.com/karimatiket/karimatiket-auth/pkg/v1/service"
)

func TestAuth(t *testing.T) {
	config.Google()

	authService := service.NewAuth(&config.AppConfig)
	url := authService.GetGoogleAuthURL()
	log.Println(url)
}
