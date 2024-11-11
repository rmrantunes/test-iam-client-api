package config

import (
	"os"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

func NewCasdoorClient() *casdoorsdk.Client {
	return casdoorsdk.NewClientWithConf(&casdoorsdk.AuthConfig{
		Endpoint:         os.Getenv("CASDOOR_ENDPOINT"),
		ClientId:         os.Getenv("CASDOOR_CLIENT_ID"),
		ClientSecret:     os.Getenv("CASDOOR_CLIENT_SECRET"),
		Certificate:      os.Getenv("CASDOOR_CERTIFICATE"),
		OrganizationName: os.Getenv("CASDOOR_ORG_NAME"),
	})
}
