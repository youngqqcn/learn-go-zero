package homestayBusiness

import (
	"context"
	"github.com/jinzhu/copier"

	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBusinessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayBusinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBusinessListLogic {
	return &HomestayBusinessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBusinessListLogic) HomestayBusinessList(req *types.HomestayBusinessListReq) (resp *types.HomestayBusinessListResp, err error) {

	logx.Info("===========1111=====================")
	whereBuilder := l.svcCtx.HomestayBusinessModel.RowBuilder()
	list, err := l.svcCtx.HomestayBusinessModel.FindPageListByIdDESC(l.ctx, whereBuilder, req.LastId, req.PageSize)
	if err != nil {
		return nil, err
	}
	logx.Infof("=============== %d", len(list))

	//logx.Info("================================")
	var res []types.HomestayBusinessListInfo
	if len(list) > 0 {
		for _, item := range list {
			var typeHomestayBusinessListInfo types.HomestayBusinessListInfo
			copier.Copy(&typeHomestayBusinessListInfo, item)

			res = append(res, typeHomestayBusinessListInfo)
		}
	}

	resp = &types.HomestayBusinessListResp{
		List: res,
	}
	return
}
