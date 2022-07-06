package logic

import (
	"context"

	"shuturl/api/internal/svc"
	"shuturl/api/internal/types"
	"shuturl/rpc/transformer"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req *types.ExpandReq) ( *types.ExpandResp,  error) {
	// todo: add your logic here and delete this line
	rsp, err := l.svcCtx.Transformer.Expand(l.ctx, &transformer.ExpandReq{
		Shorten: req.Shorten,
	})

	if err != nil {
		return &types.ExpandResp{}, err
	}


	return &types.ExpandResp{
		Url: rsp.Url,
	}, nil
}
