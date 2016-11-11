package dbaccessor

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/kopiczko/mikro/dbaccessor/dbaccessorpb"
	"github.com/micro/go-micro/errors"
)

// Service name registered in micro.
const ServiceName = "mikro.dbaccessor"

type user struct {
	Name, FullName string
	todoList       []string
}

// TODOList returns a copy of the user's to-do list.
func (u *user) TODOList() []string {
	list := make([]string, len(u.todoList))
	for i, e := range u.todoList {
		list[i] = e
	}
	return list
}

var fixtures = map[string]user{
	"pawel": {
		Name:     "pawel",
		FullName: "Paweł Kopiczko",
		todoList: []string{
			"Book the flight",
			"Water the plants",
		},
	},
}

func getUser(username string) (u user, ok bool) {
	u, ok = fixtures[username]
	return
}

type DBAccessor struct{}

func (*DBAccessor) TODOList(ctx context.Context, req *dbaccessorpb.TODOListRequest, rsp *dbaccessorpb.TODOListResponse) error {
	u, ok := getUser(req.Username)
	if !ok {
		return errors.NotFound("", fmt.Sprintf("user %s not found", req.Username))
	}
	rsp.Items = u.TODOList()
	return nil
}

func (*DBAccessor) User(ctx context.Context, req *dbaccessorpb.UserRequest, rsp *dbaccessorpb.UserResponse) error {
	u, ok := getUser(req.Username)
	if !ok {
		return errors.NotFound("", fmt.Sprintf("user %s not found", req.Username))
	}
	rsp.Name = u.Name
	rsp.FullName = u.FullName
	return nil
}
