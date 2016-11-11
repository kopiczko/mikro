package client

import (
	"context"

	"github.com/kopiczko/mikro/auth/authpb"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/registry"
)

const (
	Service       = "mikro.auth"
	ProfileMethod = "Auth.Profile"
)

type Auth interface {
	// User returns UserResponse for a user's name. ok is set when user is found.
	Profile(ctx context.Context, name string) (user authpb.ProfileResponse, ok bool, err error)
}

func NewAuth(r registry.Registry) Auth {
	return &auth{
		Client: client.NewClient(client.Registry(r)),
	}
}

type auth struct {
	client.Client
}

func (c *auth) Profile(ctx context.Context, name string) (user authpb.ProfileResponse, ok bool, err error) {
	req := c.NewRequest(Service, ProfileMethod, &authpb.ProfileRequest{
		Username: name,
	})
	var rsp authpb.ProfileResponse
	if err = c.Call(ctx, req, &rsp); err != nil {
		rpcErr := errors.Parse(err.Error())
		if rpcErr.Code == 404 {
			return rsp, false, nil
		}
		return rsp, false, rpcErr
	}
	return rsp, true, nil
}
