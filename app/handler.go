package app

import (
	"github.com/kopiczko/mikro/app/apppb"
	"github.com/kopiczko/mikro/auth"
	dbaccessor "github.com/kopiczko/mikro/dbaccessor/client"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/context"
)

// Service name registered in micro.
const ServiceName = "mikro.app"

var (
	ErrMissingMetadata               = "missing metadata"
	ErrAuthorizationMetadataNotFound = "authorization metadata not found"
)

type App struct {
	dbAccessor dbaccessor.DBAccessor
}

func New(dbAccessor dbaccessor.DBAccessor) *App {
	return &App{
		dbAccessor: dbAccessor,
	}
}

func (a *App) TODOList(ctx context.Context, req *apppb.TODOListRequest, rsp *apppb.TODOListResponse) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.Unauthorized(ServiceName, ErrMissingMetadata)
	}
	token, ok := md["authorization"]
	if !ok {
		return errors.Unauthorized(ServiceName, ErrAuthorizationMetadataNotFound)
	}
	user, err := auth.ReadUser(token)
	if err != nil {
		return errors.Unauthorized(ServiceName, err.Error())
	}
	todoList, err := a.dbAccessor.TODOList(ctx, user)
	if err != nil {
		return err
	}
	rsp.Items = todoList.Items
	return nil
}
