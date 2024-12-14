package impl

import (
	"context"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/restapi/operations"
	"github.com/h4x4d/go_hsse_hotels/pkg/client"
)

func LoginUser(clt *client.Client, fields operations.PostLoginBody) (*string, error) {
	ctx := context.Background()
	userToken, err := clt.Client.Login(ctx, clt.Config.Client, clt.Config.ClientSecret,
		clt.Config.Realm, *fields.Login, *fields.Password)
	if err != nil {
		return nil, err
	}
	return &userToken.AccessToken, nil
}
