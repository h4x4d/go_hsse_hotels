package client

import (
	"context"
	"errors"
	"github.com/Nerzal/gocloak/v13"
)

func (c Client) CheckToken(token string) (userId string, err error) {
	ctx := context.Background()
	usrInfo, err := c.Client.GetUserInfo(ctx, token, c.Config.Realm)
	if err != nil {
		return "", err
	}
	params := gocloak.GetUsersParams{
		Email: usrInfo.Email,
	}
	adminToken, err := c.GetAdminToken()
	if err != nil {
		return "", err
	}
	usr, err := c.Client.GetUsers(ctx, adminToken.AccessToken, c.Config.Realm, params)
	if err != nil {
		return "", err
	}
	if len(usr) == 0 {
		return "", errors.New("user not found")
	}
	return *usr[0].ID, nil
}
