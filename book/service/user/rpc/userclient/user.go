// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"book/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdReq            = user.IdReq
	UserInfoResponse = user.UserInfoResponse

	User interface {
		GetUser(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserInfoResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUser(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}