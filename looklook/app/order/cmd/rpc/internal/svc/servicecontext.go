package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/order/cmd/rpc/internal/config"
	"looklook/app/order/model"
	"looklook/app/travel/cmd/rpc/travel"
)

type ServiceContext struct {
	Config config.Config

	TravelRpc          travel.Travel
	HomestayOrderModel model.HomestayOrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,

		TravelRpc:          travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
		HomestayOrderModel: model.NewHomestayOrderModel(sqlConn, c.Cache),
	}
}
