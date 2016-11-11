// Code generated by protoc-gen-go.
// source: app.proto
// DO NOT EDIT!

/*
Package apppb is a generated protocol buffer package.

It is generated from these files:
	app.proto

It has these top-level messages:
	TODOListRequest
	TODOListResponse
*/
package apppb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TODOListRequest struct {
}

func (m *TODOListRequest) Reset()                    { *m = TODOListRequest{} }
func (m *TODOListRequest) String() string            { return proto.CompactTextString(m) }
func (*TODOListRequest) ProtoMessage()               {}
func (*TODOListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TODOListResponse struct {
	Items []string `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *TODOListResponse) Reset()                    { *m = TODOListResponse{} }
func (m *TODOListResponse) String() string            { return proto.CompactTextString(m) }
func (*TODOListResponse) ProtoMessage()               {}
func (*TODOListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*TODOListRequest)(nil), "apppb.TODOListRequest")
	proto.RegisterType((*TODOListResponse)(nil), "apppb.TODOListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for App service

type AppClient interface {
	// Retruns a to-do list for an authenticated user, i.e. sending valid JWT
	// token in authorization metadata.
	TODOList(ctx context.Context, in *TODOListRequest, opts ...client.CallOption) (*TODOListResponse, error)
}

type appClient struct {
	c           client.Client
	serviceName string
}

func NewAppClient(serviceName string, c client.Client) AppClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "apppb"
	}
	return &appClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *appClient) TODOList(ctx context.Context, in *TODOListRequest, opts ...client.CallOption) (*TODOListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "App.TODOList", in)
	out := new(TODOListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for App service

type AppHandler interface {
	// Retruns a to-do list for an authenticated user, i.e. sending valid JWT
	// token in authorization metadata.
	TODOList(context.Context, *TODOListRequest, *TODOListResponse) error
}

func RegisterAppHandler(s server.Server, hdlr AppHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&App{hdlr}, opts...))
}

type App struct {
	AppHandler
}

func (h *App) TODOList(ctx context.Context, in *TODOListRequest, out *TODOListResponse) error {
	return h.AppHandler.TODOList(ctx, in, out)
}

func init() { proto.RegisterFile("app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0x28, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0x28, 0x28, 0x48, 0x52, 0x12, 0xe4, 0xe2,
	0x0f, 0xf1, 0x77, 0xf1, 0xf7, 0xc9, 0x2c, 0x2e, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51,
	0xd2, 0xe0, 0x12, 0x40, 0x08, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x89, 0x70, 0xb1, 0x66,
	0x96, 0xa4, 0xe6, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x06, 0x41, 0x38, 0x46, 0x2e, 0x5c,
	0xcc, 0x8e, 0x05, 0x05, 0x42, 0xb6, 0x5c, 0x1c, 0x30, 0x0d, 0x42, 0x62, 0x7a, 0x60, 0x73, 0xf5,
	0xd0, 0x0c, 0x95, 0x12, 0xc7, 0x10, 0x87, 0x98, 0xac, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0x90, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x39, 0xf1, 0x08, 0xf7, 0x9d, 0x00, 0x00, 0x00,
}
