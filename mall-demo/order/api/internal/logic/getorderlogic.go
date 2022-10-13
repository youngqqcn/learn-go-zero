package logic

import (
	"context"

	"mall-demo/mall/user/rpc/types/user"
	"mall-demo/order/api/internal/svc"
	"mall-demo/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
		Id: "1",
	})
	if err != nil {
		return nil, err
	}

	// if user.Name != "test" {
	// 	return nil, errors.New("user not exists")
	// }

	return &types.OrderReply{
		Id:   req.Id,
		Name: user.Name,
	}, nil
}
