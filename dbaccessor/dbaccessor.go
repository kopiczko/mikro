package dbaccessor

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/go-micro/errors"
	"github.com/kopiczko/mikro/dbaccessor/dbaccessorpb"
)

var fixtures = map[string]dbaccessorpb.UserResponse{
	"pawel": dbaccessorpb.UserResponse{
		Name:     "pawel",
		FullName: "Paweł Kopiczko",
	},
}

func getUser(username string) (user dbaccessorpb.UserResponse, ok bool) {
	user, ok = fixtures[username]
	return
}

type UserService struct{}

func (*UserService) GetUser(ctx context.Context, req *dbaccessorpb.UserRequest, rsp *dbaccessorpb.UserResponse) error {
	u, ok := getUser(req.Username)
	if !ok {
		return errors.NotFound("", fmt.Sprintf("user %s not found", req.Username))
	}
	rsp.Name = u.Name
	rsp.FullName = u.FullName
	return nil
}