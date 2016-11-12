package client

import (
	"context"
	"errors"
	"log"
	"os"
	"reflect"
	"testing"

	handler "github.com/kopiczko/mikro/app"
	"github.com/kopiczko/mikro/app/apppb"
	"github.com/kopiczko/mikro/auth"
	"github.com/kopiczko/mikro/dbaccessor/dbaccessorpb"
	merrors "github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/registry/mock"
	"github.com/micro/go-micro/server"
)

const testErrorCode = 8855

type dbAccessorResp struct {
	TODOList dbaccessorpb.TODOListResponse
	Err      error
}

type dbAccessorMock struct {
	LastTODOListUsername string         // value of username argument from last TODOList call.
	DBAccessorResp       dbAccessorResp // Response to return on TODOList call.
}

func (d *dbAccessorMock) TODOList(ctx context.Context, username string) (dbaccessorpb.TODOListResponse, error) {
	d.LastTODOListUsername = username
	var err error
	if d.DBAccessorResp.Err != nil {
		e := merrors.Parse(d.DBAccessorResp.Err.Error())
		e.Code = testErrorCode
		err = e
	}
	return d.DBAccessorResp.TODOList, err
}

func (d *dbAccessorMock) User(ctx context.Context, name string) (user dbaccessorpb.UserResponse, ok bool, err error) {
	return dbaccessorpb.UserResponse{}, false, errors.New("dbAccessorMock.User not implemented")
}

var (
	reg        = mock.NewRegistry()
	dbAccessor = new(dbAccessorMock)
)

func TestMain(m *testing.M) {
	server.Init(
		server.Name(handler.ServiceName),
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

func TestAppTODOList(t *testing.T) {
	validToken, err := auth.CreateToken("pawel")
	if err != nil {
		t.Fatalf("unexpected CreateToken error = %v", err)
	}
	invalidToken := "token-abcd-invalid"
	testErr := errors.New("I don't have your to-do list")
	tests := []struct {
		Token          string
		DBAccessorResp dbAccessorResp
		TODOList       apppb.TODOListResponse
		ErrCode        int32 // non-negative means error
	}{
		{
			Token: validToken,
			DBAccessorResp: dbAccessorResp{
				TODOList: dbaccessorpb.TODOListResponse{
					Items: []string{"Sell all the things"},
				},
				Err: nil,
			},
			TODOList: apppb.TODOListResponse{
				Items: []string{"Sell all the things"},
			},
			ErrCode: -1,
		},
		{
			Token: validToken,
			DBAccessorResp: dbAccessorResp{
				TODOList: dbaccessorpb.TODOListResponse{},
				Err:      testErr,
			},
			TODOList: apppb.TODOListResponse{},
			ErrCode:  testErrorCode,
		},
		{
			Token:          invalidToken,
			DBAccessorResp: dbAccessorResp{},
			TODOList:       apppb.TODOListResponse{},
			ErrCode:        401,
		},
	}

	c := NewApp(reg)
	defer func() { dbAccessor.DBAccessorResp = dbAccessorResp{} }()

	for i, tt := range tests {
		dbAccessor.DBAccessorResp = tt.DBAccessorResp
		todoList, err := c.TODOList(context.TODO(), tt.Token)
		if err != nil || tt.ErrCode > 0 {
			code := int32(0)
			if err != nil {
				code = merrors.Parse(err.Error()).Code
			}
			if code != tt.ErrCode {
				t.Fatalf("#%d: got %v, want %v (%v)", i, code, tt.ErrCode, err)
			}
		}
		if dbAccessor.LastTODOListUsername != "pawel" {
			t.Errorf("#%d: got %s, want pawel", i, dbAccessor.LastTODOListUsername)
		}
		if !reflect.DeepEqual(todoList, tt.TODOList) {
			t.Errorf("#%d: got profile = %v, want %v", i, todoList, tt.TODOList)
		}
	}
}
