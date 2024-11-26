package config

import (
	"auth/internal/restapi/operations"
	"context"
	"crypto/tls"
	"errors"
	"github.com/Nerzal/gocloak/v13"
	"os"
)

const (
	KEYCLOAK_PORT           = "KEYCLOAK_INNER_PORT"
	KEYCLOAK_CLIENT         = "KEYCLOAK_CLIENT"
	KEYCLOAK_REALM          = "KEYCLOAK_REALM"
	KEYCLOAK_CLIENT_SECRET  = "KEYCLOAK_CLIENT_SECRET"
	KEYCLOAK_ADMIN          = "KEYCLOAK_ADMIN"
	KEYCLOAK_ADMIN_PASSWORD = "KEYCLOAK_ADMIN_PASSWORD"
	KEYCLOAK_MASTER_REALM   = "master"
)

type Config struct {
	Client        string
	Realm         string
	ClientSecret  string
	Port          string
	Admin         string
	AdminPassword string
	MasterRealm   string
}

func LoadEnv() (*Config, error) {
	keycloakClient := os.Getenv(KEYCLOAK_CLIENT)
	if keycloakClient == "" {
		return nil, errors.New("keycloak client not set")
	}
	keycloakRealm := os.Getenv(KEYCLOAK_REALM)
	if keycloakRealm == "" {
		return nil, errors.New("keycloak realm not set")
	}
	keycloakSecret := os.Getenv(KEYCLOAK_CLIENT_SECRET)
	if keycloakSecret == "" {
		return nil, errors.New("keycloak secret not set")
	}
	keycloakPort := os.Getenv(KEYCLOAK_PORT)
	if keycloakPort == "" {
		return nil, errors.New("keycloak port not set")
	}
	keycloakAdmin := os.Getenv(KEYCLOAK_ADMIN)
	if keycloakAdmin == "" {
		return nil, errors.New("keycloak admin not set")
	}
	keycloakAdminPassword := os.Getenv(KEYCLOAK_ADMIN_PASSWORD)
	if keycloakAdminPassword == "" {
		return nil, errors.New("keycloak admin password not set")
	}
	return &Config{
		Client:        keycloakClient,
		Realm:         keycloakRealm,
		ClientSecret:  keycloakSecret,
		Port:          keycloakPort,
		Admin:         keycloakAdmin,
		AdminPassword: keycloakAdminPassword,
		MasterRealm:   KEYCLOAK_MASTER_REALM,
	}, nil
}

type Client struct {
	Config *Config
	Client *gocloak.GoCloak
}

func NewClient() (*Client, error) {
	config, err := LoadEnv()
	if err != nil {
		return nil, err
	}
	client := gocloak.NewClient("http://keycloak:" + KEYCLOAK_PORT)

	restyClient := client.RestyClient()
	restyClient.SetDebug(true)
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	return &Client{
		Client: client,
		Config: config,
	}, nil
}

func (c Client) GetAdminToken() (*gocloak.JWT, error) {
	ctx := context.Background()
	token, err := c.Client.LoginAdmin(ctx, c.Config.Admin, c.Config.AdminPassword, c.Config.MasterRealm)

	if err != nil {
		return nil, err
	}
	return token, nil
}

func (c Client) CreateUser(fields operations.PostRegisterBody) (*string, error) {
	ctx := context.Background()
	user := gocloak.User{
		Email:    fields.Email,
		Enabled:  gocloak.BoolP(true),
		Username: fields.Login,
		Groups:   &[]string{*fields.Role},
	}
	token, err := c.GetAdminToken()
	if err != nil {
		return nil, err
	}
	userId, err := c.Client.CreateUser(ctx, token.AccessToken, c.Config.Realm, user)
	if err != nil {
		return nil, err
	}
	err = c.Client.SetPassword(ctx, token.AccessToken, userId, c.Config.Realm, *fields.Password, false)
	if err != nil {
		return nil, err
	}
	userToken, err := c.Client.Login(ctx, c.Config.Client, c.Config.ClientSecret, c.Config.Realm, *fields.Login, *fields.Password)
	if err != nil {
		return nil, err
	}
	return &userToken.AccessToken, nil
}

func (c Client) LoginUser(fields operations.PostLoginBody) (*string, error) {
	ctx := context.Background()
	userToken, err := c.Client.Login(ctx, c.Config.Client, c.Config.ClientSecret, c.Config.Realm, *fields.Login, *fields.Password)
	if err != nil {
		return nil, err
	}
	return &userToken.AccessToken, nil
}
