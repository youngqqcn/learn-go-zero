package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"looklook/app/order/cmd/rpc/internal/svc"
	"looklook/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserHomestayOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderListLogic {
	return &UserHomestayOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户订单列表
func (l *UserHomestayOrderListLogic) UserHomestayOrderList(in *pb.UserHomestayOrderListReq) (*pb.UserHomestayOrderListResp, error) {

	whereBuilder := l.svcCtx.HomestayOrderModel.RowBuilder().Where(squirrel.Eq{
		"user_id":     in.UserId,
		"trade_state": in.TradeState,
	})
	homestayOrderList, err := l.svcCtx.HomestayOrderModel.FindPageListByIdDESC(l.ctx, whereBuilder, 0, 10)
	if err != nil {
		return nil, err
	}

	var resp []*pb.HomestayOrder
	if len(homestayOrderList) > 0 {
		for _, item := range homestayOrderList {
			var order pb.HomestayOrder
			copier.Copy(&order, item)
			resp = append(resp, &order)
		}
	}

	return &pb.UserHomestayOrderListResp{
		List: resp,
	}, nil
}
