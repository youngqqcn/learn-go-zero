package homestay

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBusinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessListLogic {
	return &BusinessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BusinessListLogic) BusinessList(req *types.BusinessListReq) (*types.BusinessListResp, error) {
	// todo: add your logic here and delete this line

	logx.Info("=======================")
	whereBuilder := l.svcCtx.HomestayModel.RowBuilder().Where(squirrel.Eq{"homestay_business_id": req.HomestayBusinessId})
	list, err := l.svcCtx.HomestayModel.FindPageListByIdDESC(l.ctx, whereBuilder, req.LastId, req.PageSize)
	if err != nil {
		logx.Error(err.Error())
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

	//homeStay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, req.LastId)
	//if err != nil {
	//	return nil, errors.New("FindOne error")
	//}
	//
	//var rsp []types.Homestay
	//typeHomestay := new(types.Homestay)
	//copier.Copy(typeHomestay, homeStay)
	//rsp = append(rsp, *typeHomestay)

	return &types.BusinessListResp{
		List: rsp,
	}, nil
}
