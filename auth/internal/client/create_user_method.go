package client

import (
	"auth/internal/restapi/operations"
	"context"
	"github.com/Nerzal/gocloak/v13"
	"strconv"
)

func (c Client) CreateUser(fields operations.PostRegisterBody) (*string, error) {
	ctx := context.Background()
	user := gocloak.User{
		Email:    fields.Email,
		Enabled:  gocloak.BoolP(true),
		Username: fields.Login,
		Attributes: &map[string][]string{
			"telegram_id": {strconv.FormatInt(*fields.TelegramID, 10)},
		},
		Groups: &[]string{*fields.Role},
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
