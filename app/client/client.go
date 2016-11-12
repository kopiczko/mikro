package client

import (
	"context"

	"github.com/kopiczko/mikro/app/apppb"
	"github.com/kopiczko/mikro/auth"
	"github.com/micro/go-micro/client"
)

const (
	Service        = "mikro.app"
	TODOListMethod = "App.TODOList"
)

type App interface {
	// TODOList retruns a to-do list for an authenticated user, i.e. sending valid JWT
	// token in authorization metadata. Token should be obtained with Auth.Login.
	// WithToken can be used to fill context with the token.
	TODOList(ctx context.Context, token string) (apppb.TODOListResponse, error)
}

func NewApp(opt ...client.Option) App {
	return &app{
		Client: client.NewClient(opt...),
	}
}

type app struct {
	client.Client
}

func (c *app) TODOList(ctx context.Context, token string) (apppb.TODOListResponse, error) {
	req := c.NewRequest(Service, TODOListMethod, new(apppb.TODOListRequest))
	var rsp apppb.TODOListResponse
	err := c.Call(auth.WithToken(ctx, token), req, &rsp)
	return rsp, err
}
