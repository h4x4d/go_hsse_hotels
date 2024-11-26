package client

import (
	"auth/internal/config"
	"crypto/tls"
	"github.com/Nerzal/gocloak/v13"
)

type Client struct {
	Config *config.Config
	Client *gocloak.GoCloak
}

func NewClient() (*Client, error) {
	clientCfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	client := gocloak.NewClient("http://keycloak:" + clientCfg.Port)

	restyClient := client.RestyClient()
	restyClient.SetDebug(true)
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	return &Client{
		Client: client,
		Config: clientCfg,
	}, nil
}
