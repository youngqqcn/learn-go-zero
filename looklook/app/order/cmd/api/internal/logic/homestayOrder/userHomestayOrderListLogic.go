package homestayOrder

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/order/cmd/rpc/order"
	"looklook/common/ctxdata"

	"looklook/app/order/cmd/api/internal/svc"
	"looklook/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHomestayOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderListLogic {
	return &UserHomestayOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomestayOrderListLogic) UserHomestayOrderList(req *types.UserHomestayOrderListReq) (resp *types.UserHomestayOrderListResp, err error) {

	userId := ctxdata.GetUidFromCtx(l.ctx)

	homestayOrderList, err := l.svcCtx.OrderRpc.UserHomestayOrderList(l.ctx, &order.UserHomestayOrderListReq{
		LastId:     req.LastId,
		PageSize:   req.PageSize,
		UserId:     userId,
		TradeState: req.TradeFilter,
	})
	if err != nil {
		return nil, errors.Errorf("l.svcCtx.OrderRpc.UserHomestayOrderList: %v", err.Error())
	}

	resp = new(types.UserHomestayOrderListResp)
	//copier.Copy(resp.List, homestayOrderList.List)

	for _, item := range homestayOrderList.List {
		var v types.UserHomestayOrderListView
		_ = copier.Copy(&v, item)
		resp.List = append(resp.List, v)
	}

	return
}
