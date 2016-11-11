package app

import (
	"github.com/kopiczko/mikro/app/apppb"
	"github.com/kopiczko/mikro/auth"
	dbaccessor "github.com/kopiczko/mikro/dbaccessor/client"
	"golang.org/x/net/context"
)

// Service name registered in micro.
const ServiceName = "mikro.app"

var (
	ErrMissingMetadata               = "missing metadata"
	ErrAuthorizationMetadataNotFound = "authorization metadata not found"
)

type App struct {
	authorizer auth.Authorizer
	dbAccessor dbaccessor.DBAccessor
}

func New(dbAccessor dbaccessor.DBAccessor) *App {
	return &App{
		authorizer: auth.NewAuthorizer(ServiceName),
		dbAccessor: dbAccessor,
	}
}

func (a *App) TODOList(ctx context.Context, req *apppb.TODOListRequest, rsp *apppb.TODOListResponse) error {
	user, err := a.authorizer.FromContext(ctx)
	if err != nil {
		return err
	}
	todoList, err := a.dbAccessor.TODOList(ctx, user)
	if err != nil {
		return err
	}
	rsp.Items = todoList.Items
	return nil
}
