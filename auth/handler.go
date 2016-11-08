package auth

import (
	"fmt"

	"github.com/kopiczko/mikro/auth/authpb"
	"github.com/kopiczko/mikro/dbaccessor/client"
	"github.com/micro/go-micro/errors"
	"golang.org/x/net/context"
)

const (
	ServiceName = "mikro.auth"
)

type Auth struct {
	dbAccessor client.DBAccessor
}

func New(dbAccessor client.DBAccessor) *Auth {
	return &Auth{
		dbAccessor: dbAccessor,
	}
}

func (a *Auth) Profile(ctx context.Context, req *authpb.ProfileRequest, rsp *authpb.ProfileResponse) error {
	p, ok, err := a.dbAccessor.User(ctx, req.Username)
	if err != nil {
		return errors.InternalServerError(ServiceName, err.Error())
	}
	if !ok {
		return errors.NotFound(ServiceName, fmt.Sprintf("user %s not found", req.Username))
	}
	rsp.Name = p.Name
	rsp.FullName = p.FullName
	return nil
}
