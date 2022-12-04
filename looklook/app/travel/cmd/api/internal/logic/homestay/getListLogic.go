package homestay

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListLogic {
	return &GetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListLogic) GetList(req *types.GuessListReq) (resp *types.GuessListResp, err error) {

	list, err := l.svcCtx.HomestayModel.FindPageListByIdDESC(l.ctx, l.svcCtx.HomestayModel.RowBuilder(), 0, 5)
	if err != nil {
		return nil, errors.New("FindPageListByIdDESC")
	}

	var rsp []types.Homestay
	if len(list) > 0 {
		for _, homestay := range list {
			typeHomestay := new(types.Homestay)
			copier.Copy(typeHomestay, homestay)

			// TODO, 分转成元
			rsp = append(rsp, *typeHomestay)
		}
	}

	resp = &types.GuessListResp{
		List: rsp,
	}
	err = nil
	return
}
