package homestay

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"looklook/app/travel/cmd/rpc/travel"

	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayDetailLogic {
	return &HomestayDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayDetailLogic) HomestayDetail(req *types.HomestayDetailReq) (resp *types.HomestayDetailResp, err error) {

	// 直接查数据库
	//homeStay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, req.Id)

	// 使用 RPC获取详情
	homeStay, err := l.svcCtx.TravelRpc.HomestayDetail(l.ctx, &travel.HomestayDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.New("FindOne error")
	}
	resp = new(types.HomestayDetailResp)
	copier.Copy(&resp.Homestay, homeStay)
	err = nil
	return
}
