package client

import (
	"context"
	"errors"
	"github.com/Nerzal/gocloak/v13"
	"github.com/h4x4d/go_hsse_hotels/pkg/models"
	"strconv"
)

func (c Client) CheckToken(token string) (user *models.User, err error) {
	ctx := context.Background()
	usrInfo, err := c.Client.GetUserInfo(ctx, token, c.Config.Realm)
	if err != nil {
		return nil, err
	}
	exact := true
	params := gocloak.GetUsersParams{
		Email: usrInfo.Email,
		Exact: &exact,
	}
	adminToken, err := c.GetAdminToken()
	if err != nil {
		return nil, err
	}
	users, err := c.Client.GetUsers(ctx, adminToken.AccessToken, c.Config.Realm, params)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("user not found")
	}
	userId := *users[0].ID
	tgId, err := strconv.Atoi((*users[0].Attributes)["telegram_id"][0])
	if err != nil {
		return nil, err
	}
	groups, err := c.Client.GetUserGroups(ctx, adminToken.AccessToken, c.Config.Realm, userId, gocloak.GetGroupsParams{})
	if err != nil {
		return nil, err
	}
	role := *groups[0].Name
	return &models.User{
		UserID:     userId,
		TelegramID: tgId,
		Role:       role,
	}, nil
}
