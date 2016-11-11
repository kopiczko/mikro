package client

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"

	handler "github.com/kopiczko/mikro/auth"
	"github.com/kopiczko/mikro/auth/authpb"
	cmd "github.com/kopiczko/mikro/auth/cmd/auth"
	"github.com/kopiczko/mikro/dbaccessor/dbaccessorpb"
	"github.com/micro/go-micro/registry/mock"
	"github.com/micro/go-micro/server"
)

type dbAccessorResp struct {
	User dbaccessorpb.UserResponse
	OK   bool
	Err  error
}

type dbAccessorMock struct {
	LastUserName string         // value of name argument from last User call.
	UserResp     dbAccessorResp // Response to return on User call.
}

func (d *dbAccessorMock) User(ctx context.Context, name string) (user dbaccessorpb.UserResponse, ok bool, err error) {
	d.LastUserName = name
	return d.UserResp.User, d.UserResp.OK, d.UserResp.Err
}

var (
	reg        = mock.NewRegistry()
	dbAccessor = new(dbAccessorMock)
)

func TestMain(m *testing.M) {
	server.Init(
		server.Name(cmd.ServiceName),
		server.Registry(reg),
	)
	server.Handle(server.NewHandler(
		handler.New(dbAccessor),
	))
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
	server.Register()

	code := m.Run()

	server.Stop()

	os.Exit(code)
}

func TestAuthLogin(t *testing.T) {
	c := NewAuth(reg)
	token, err := c.Login(context.TODO(), "pawel", "currently_not_checked")
	if err != nil {
		t.Fatalf("unexpected Login error = %v", err)
	}
	if token == "" {
		t.Errorf("want non-empty token")
	}
}

func TestAuthProfile(t *testing.T) {
	tests := []struct {
		DBAccessorResp dbAccessorResp
		Username       string
		OK             bool
		Profile        authpb.ProfileResponse
	}{
		{
			DBAccessorResp: dbAccessorResp{
				User: dbaccessorpb.UserResponse{
					Name:     "pawel",
					FullName: "Paweł Kopiczko",
				},
				OK:  true,
				Err: nil,
			},
			Username: "pawel",
			OK:       true,
			Profile: authpb.ProfileResponse{
				Name:     "pawel",
				FullName: "Paweł Kopiczko",
			},
		},
		{
			Username: "you_do_not_know_who",
			OK:       false,
		},
	}

	c := NewAuth(reg)
	defer func() { dbAccessor.UserResp = dbAccessorResp{} }()

	for i, tt := range tests {
		dbAccessor.UserResp = tt.DBAccessorResp
		profile, ok, err := c.Profile(context.TODO(), tt.Username)
		if err != nil {
			t.Fatalf("#%d: unexpected Profile error = %v", i, err)
		}
		if dbAccessor.LastUserName != tt.Username {
			t.Errorf("got %s, want %s", dbAccessor.LastUserName, tt.Username)
		}
		if ok != tt.OK {
			t.Errorf("#%d: got ok = %v, want %v", i, ok, tt.OK)
		}
		if !reflect.DeepEqual(profile, tt.Profile) {
			t.Errorf("#%d: got profile = %v, want %v", i, profile, tt.Profile)
		}
	}
}
