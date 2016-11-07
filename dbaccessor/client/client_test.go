package client

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/go-micro/registry/mock"
	"github.com/go-micro/server"
	"github.com/kopiczko/mikro/dbaccessor"
	cmd "github.com/kopiczko/mikro/dbaccessor/cmd/dbaccessor"
	"github.com/kopiczko/mikro/dbaccessor/dbaccessorpb"
)

var reg = mock.NewRegistry()

func TestMain(m *testing.M) {
	server.Init(
		server.Name(cmd.ServiceName),
		server.Registry(reg),
	)
	server.Handle(server.NewHandler(
		new(dbaccessor.DBAccessor),
	))
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
	server.Register()

	code := m.Run()

	server.Stop()

	os.Exit(code)
}

func TestDBAccessor(t *testing.T) {
	tests := []struct {
		Username string
		OK       bool
		User     dbaccessorpb.UserResponse
	}{
		{
			Username: "pawel",
			OK:       true,
			User: dbaccessorpb.UserResponse{
				Name:     "pawel",
				FullName: "Paweł Kopiczko",
			},
		},
		{
			Username: "you_do_not_know_who",
			OK:       false,
		},
	}

	c := NewDBAccessor(reg)

	for i, tt := range tests {
		user, ok, err := c.User(context.TODO(), tt.Username)
		if err != nil {
			t.Fatalf("#%d: unexpected User error = %v", i, err)
		}
		if ok != tt.OK {
			t.Errorf("#%d: got ok = %v, want %v", i, ok, tt.OK)
		}
		if !reflect.DeepEqual(user, tt.User) {
			t.Errorf("#%d: got user = %v, want %v", i, user, tt.User)
		}
	}
}
