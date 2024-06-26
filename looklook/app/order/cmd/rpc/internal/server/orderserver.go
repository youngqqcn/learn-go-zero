// Code generated by goctl. DO NOT EDIT!
// Source: order.proto

package server

import (
	"context"

	"looklook/app/order/cmd/rpc/internal/logic"
	"looklook/app/order/cmd/rpc/internal/svc"
	"looklook/app/order/cmd/rpc/pb"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

// 下单
func (s *OrderServer) CreateHomestayOrder(ctx context.Context, in *pb.CreateHomestayOrderReq) (*pb.CreateHomestayOrderResp, error) {
	l := logic.NewCreateHomestayOrderLogic(ctx, s.svcCtx)
	return l.CreateHomestayOrder(in)
}

// 订单详情
func (s *OrderServer) HomestayOrderDetail(ctx context.Context, in *pb.HomestayOrderDetailReq) (*pb.HomestayOrderDetailResp, error) {
	l := logic.NewHomestayOrderDetailLogic(ctx, s.svcCtx)
	return l.HomestayOrderDetail(in)
}

// 更新订单
func (s *OrderServer) UpdateHomestayOrderTradeState(ctx context.Context, in *pb.UpdateHomestayOrderTradeStateReq) (*pb.UpdateHomestayOrderTradeStateResp, error) {
	l := logic.NewUpdateHomestayOrderTradeStateLogic(ctx, s.svcCtx)
	return l.UpdateHomestayOrderTradeState(in)
}

// 用户订单列表
func (s *OrderServer) UserHomestayOrderList(ctx context.Context, in *pb.UserHomestayOrderListReq) (*pb.UserHomestayOrderListResp, error) {
	l := logic.NewUserHomestayOrderListLogic(ctx, s.svcCtx)
	return l.UserHomestayOrderList(in)
}
