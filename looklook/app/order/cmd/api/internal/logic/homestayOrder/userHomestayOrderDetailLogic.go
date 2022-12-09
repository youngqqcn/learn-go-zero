package homestayOrder

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/order/cmd/rpc/order"

	"looklook/app/order/cmd/api/internal/svc"
	"looklook/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHomestayOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderDetailLogic {
	return &UserHomestayOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomestayOrderDetailLogic) UserHomestayOrderDetail(req *types.UserHomestayOrderDetailReq) (resp *types.UserHomestayOrderDetailResp, err error) {

	orderDetailResp, err := l.svcCtx.OrderRpc.HomestayOrderDetail(l.ctx, &order.HomestayOrderDetailReq{
		Sn: req.Sn,
	})

	if err != nil {
		return nil, errors.Errorf("svcCtx.OrderRpc.HomestayOrderDetail, error: %v", err.Error())
	}
	if orderDetailResp == nil {
		return nil, errors.Errorf("order %v not exists", req.Sn)
	}

	resp = new(types.UserHomestayOrderDetailResp)
	copier.Copy(resp, orderDetailResp.HomestayOrder)
	return
}
