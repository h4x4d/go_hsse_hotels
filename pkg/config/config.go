package config

import (
	"errors"
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

func NewConfig() (*Config, error) {
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
