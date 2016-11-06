package auth

import (
	"fmt"

	"github.com/go-micro/client"
	"github.com/go-micro/errors"
	"github.com/kopiczko/mikro/auth/authpb"
	"github.com/kopiczko/mikro/dbaccessor/dbaccessorpb"
	"golang.org/x/net/context"
)

const (
	id             = "mikro.auth"
	serviceName    = "mikro.userservice"
	serviceGetUser = "UserService.GetUser"
)

func getProfile(username string) (profile authpb.ProfileResponse, ok bool, err error) {
	req := client.NewRequest(serviceName, serviceGetUser, &dbaccessorpb.UserRequest{
		Username: username,
	})
	var rsp dbaccessorpb.UserResponse
	if err = client.Call(context.Background(), req, &rsp); err != nil {
		rpcErr := errors.Parse(err.Error())
		if rpcErr.Code == 404 {
			return profile, false, nil
		}
		return profile, false, err
	}
	profile.Name = rsp.Name
	profile.FullName = rsp.FullName
	return profile, true, nil
}

type Auth struct{}

func (*Auth) Profile(ctx context.Context, req *authpb.ProfileRequest, rsp *authpb.ProfileResponse) error {
	p, ok, err := getProfile(req.Username)
	if err != nil {
		return errors.InternalServerError(id, err.Error())
	}
	if !ok {
		return errors.NotFound(id, fmt.Sprintf("user %s not found", req.Username))
	}
	rsp.Name = p.Name
	rsp.FullName = p.FullName
	return nil
}
