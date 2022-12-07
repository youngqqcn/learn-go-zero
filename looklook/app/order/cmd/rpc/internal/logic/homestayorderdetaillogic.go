package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"looklook/app/order/cmd/rpc/internal/svc"
	"looklook/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayOrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomestayOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayOrderDetailLogic {
	return &HomestayOrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 订单详情
func (l *HomestayOrderDetailLogic) HomestayOrderDetail(in *pb.HomestayOrderDetailReq) (*pb.HomestayOrderDetailResp, error) {

	order, err := l.svcCtx.HomestayOrderModel.FindOneBySn(l.ctx, in.Sn)
	if err != nil {
		return nil, err
	}

	var pbOrder pb.HomestayOrder
	if order != nil {
		copier.Copy(&pbOrder, order)

		// 保持原来的时间
		//pbOrder.CreateTime = order.CreateTime.Unix()
		//pbOrder.LiveStartDate = order.LiveStartDate.Unix()
		//pbOrder.LiveEndDate = order.LiveEndDate.Unix()
	}

	return &pb.HomestayOrderDetailResp{
		HomestayOrder: &pbOrder,
	}, nil
}
