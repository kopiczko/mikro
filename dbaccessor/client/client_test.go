package client

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/kopiczko/mikro/dbaccessor"
	"github.com/kopiczko/mikro/dbaccessor/dbaccessorpb"
	"github.com/micro/go-micro/registry/mock"
	"github.com/micro/go-micro/server"
)

var reg = mock.NewRegistry()

func TestMain(m *testing.M) {
	server.Init(
		server.Name(dbaccessor.ServiceName),
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

func TestDBAccessorTODOList(t *testing.T) {
	tests := []struct {
		Username string
		Found    bool
	}{
		{
			Username: "pawel",
			Found:    true,
		},
		{
			Username: "you_do_not_know_who",
			Found:    false,
		},
	}

	c := NewDBAccessor(reg)

	for i, tt := range tests {
		todoList, err := c.TODOList(context.TODO(), tt.Username)
		if tt.Found {
			if err != nil {
				t.Fatalf("#%d: unexpected User error = %v", i, err)
			}
			if todoList.Items == nil {
				t.Errorf("#%d: got items = nil, want non-nil slice", i)
			}
		} else {
			if err == nil {
				t.Errorf("#%d: got error = nil, want non-nil", i)
			}
		}
	}
}

func TestDBAccessorUser(t *testing.T) {
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
