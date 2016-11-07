package client

import (
	"context"

	"github.com/go-micro/client"
	"github.com/go-micro/errors"
	"github.com/kopiczko/mikro/dbaccessor/dbaccessorpb"
	"github.com/micro/go-micro/registry"
)

const (
	Service    = "mikro.userservice"
	UserMethod = "UserService.GetUser"
)

type DBAccessor interface {
	// User returns UserResponse for a user's name. ok is set when user is found.
	User(ctx context.Context, name string) (user dbaccessorpb.UserResponse, ok bool, err error)
}

func NewDBAccessor(r registry.Registry) DBAccessor {
	return &dbAccessor{
		Client: client.NewClient(client.Registry(r)),
	}
}

type dbAccessor struct {
	client.Client
}

func (c *dbAccessor) User(ctx context.Context, name string) (user dbaccessorpb.UserResponse, ok bool, err error) {
	req := c.NewRequest(Service, UserMethod, &dbaccessorpb.UserRequest{
		Username: name,
	})
	var rsp dbaccessorpb.UserResponse
	if err = c.Call(ctx, req, &rsp); err != nil {
		rpcErr := errors.Parse(err.Error())
		if rpcErr.Code == 404 {
			return rsp, false, nil
		}
		return rsp, false, rpcErr
	}
	return rsp, true, nil
}
