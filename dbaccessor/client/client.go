package client

import (
	"context"

	"github.com/kopiczko/mikro/dbaccessor/dbaccessorpb"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
)

const (
	Service        = "mikro.dbaccessor"
	TODOListMethod = "DBAccessor.TODOList"
	UserMethod     = "DBAccessor.User"
)

type DBAccessor interface {
	// TODOList returns the user's to-do list itemes. When user is not
	// found error is returned.
	TODOList(ctx context.Context, username string) (dbaccessorpb.TODOListResponse, error)
	// User returns UserResponse for a user's name. ok is set when user is found.
	User(ctx context.Context, name string) (user dbaccessorpb.UserResponse, ok bool, err error)
}

func NewDBAccessor(opt ...client.Option) DBAccessor {
	return &dbAccessor{
		Client: client.NewClient(opt...),
	}
}

type dbAccessor struct {
	client.Client
}

func (c *dbAccessor) TODOList(ctx context.Context, username string) (dbaccessorpb.TODOListResponse, error) {
	req := c.NewRequest(Service, TODOListMethod, &dbaccessorpb.TODOListRequest{
		Username: username,
	})
	var rsp dbaccessorpb.TODOListResponse
	err := c.Call(ctx, req, &rsp)
	return rsp, err
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
