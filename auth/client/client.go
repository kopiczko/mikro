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
	LoginMethod   = "Auth.Login"
	ProfileMethod = "Auth.Profile"
)

type Auth interface {
	// Login returns a token for the username authenticated with the
	// password. Currently the password is not checked.
	Login(ctx context.Context, username, password string) (token string, err error)
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

func (c *auth) Login(ctx context.Context, username, password string) (token string, err error) {
	req := c.NewRequest(Service, LoginMethod, &authpb.LoginRequest{
		Username: username,
		Password: password,
	})
	var rsp authpb.LoginResponse
	err = c.Call(ctx, req, &rsp)
	return rsp.Token, err
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
