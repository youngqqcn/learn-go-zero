package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/order/cmd/api/internal/config"
	"looklook/app/order/cmd/rpc/order"
	"looklook/app/travel/cmd/rpc/travel"
)

type ServiceContext struct {
	Config config.Config

	TravelRpc travel.Travel
	OrderRpc  order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		TravelRpc: travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
		OrderRpc:  order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
	}
}
