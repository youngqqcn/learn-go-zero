package logic

import (
	"context"

	"shuturl/api/internal/svc"
	"shuturl/api/internal/types"
	"shuturl/rpc/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenReq) (resp *types.ShortenResp, err error) {
	// todo: add your logic here and delete this line
	
	rsp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transform.ShortenReq{
		Url: req.Url,
	})

	if err != nil {
		return &types.ShortenResp{}, err
	}
	return &types.ShortenResp{
		Shorten: rsp.Shorten,
	}, nil
}
