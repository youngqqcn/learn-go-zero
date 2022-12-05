package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/travel/cmd/api/internal/config"
	"looklook/app/travel/cmd/rpc/travel"
	"looklook/app/travel/model"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config config.Config

	HomestayModel         model.HomestayModel
	HomestayActivityModel model.HomestayActivityModel
	TravelRpc             travel.Travel
	UsercenterRpc         usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:                c,
		HomestayModel:         model.NewHomestayModel(sqlConn, c.Cache),
		HomestayActivityModel: model.NewHomestayActivityModel(sqlConn, c.Cache),

		// RPC
		TravelRpc:     travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
