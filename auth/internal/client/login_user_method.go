package client

import (
	"auth/internal/restapi/operations"
	"context"
)

func (c Client) LoginUser(fields operations.PostLoginBody) (*string, error) {
	ctx := context.Background()
	userToken, err := c.Client.Login(ctx, c.Config.Client, c.Config.ClientSecret,
		c.Config.Realm, *fields.Login, *fields.Password)
	if err != nil {
		return nil, err
	}
	return &userToken.AccessToken, nil
}
