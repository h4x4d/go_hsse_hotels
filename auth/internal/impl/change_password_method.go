package impl

import (
	"auth/internal/restapi/operations"
	"context"
	"github.com/Nerzal/gocloak/v13"
	"github.com/h4x4d/go_hsse_hotels/internal/client"
)

func ChangePasswordUser(c client.Client, fields operations.PostChangePasswordBody) (*string, error) {
	ctx := context.Background()

	// Check Old Password
	_, err := c.Client.Login(ctx, c.Config.Client, c.Config.ClientSecret,
		c.Config.Realm, *fields.Login, *fields.OldPassword)
	if err != nil {
		return nil, err
	}

	token, err := c.GetAdminToken()
	if err != nil {
		return nil, err
	}
	params := gocloak.GetUsersParams{
		Username: fields.Login,
	}
	user, _ := c.Client.GetUsers(ctx, token.AccessToken, c.Config.Realm, params)
	err = c.Client.SetPassword(ctx, token.AccessToken, *user[0].ID,
		c.Config.Realm, *fields.NewPassword, false)
	if err != nil {
		return nil, err
	}

	userToken, err := c.Client.Login(ctx, c.Config.Client, c.Config.ClientSecret,
		c.Config.Realm, *fields.Login, *fields.NewPassword)
	if err != nil {
		return nil, err
	}
	return &userToken.AccessToken, nil
}
