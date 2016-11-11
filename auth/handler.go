package auth

import (
	"fmt"

	"github.com/kopiczko/mikro/auth/authpb"
	"github.com/kopiczko/mikro/dbaccessor/client"
	"github.com/micro/go-micro/errors"
	"golang.org/x/net/context"
)

// Service name registered in micro.
const ServiceName = "mikro.auth"

type Auth struct {
	dbAccessor client.DBAccessor
}

func New(dbAccessor client.DBAccessor) *Auth {
	return &Auth{
		dbAccessor: dbAccessor,
	}
}

func (a *Auth) Login(ctx context.Context, req *authpb.LoginRequest, rsp *authpb.LoginResponse) error {
	// TODO: challenge req.Password
	token, err := CreateToken(req.Username)
	if err != nil {
		return err
	}
	rsp.Token = token
	return nil
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
