package homestayBusiness

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/usercenter"

	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBusinessDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayBusinessDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBusinessDetailLogic {
	return &HomestayBusinessDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBusinessDetailLogic) HomestayBusinessDetail(req *types.HomestayBusinessDetailReq) (*types.HomestayBusinessDetailResp, error) {

	businessDetail, err := l.svcCtx.HomestayBusinessModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if businessDetail == nil {
		return nil, errors.New("business not exists")
	}

	var retHomestayBusinessBoss types.HomestayBusinessBoss
	if businessDetail != nil {

		logx.Infof("===> userId is %v", businessDetail.UserId)
		userInfo, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
			Id: businessDetail.UserId,
		})
		if err != nil {
			return nil, err
		}

		if userInfo.User != nil && userInfo.User.Id > 0 {
			//copier.Copy(&retHomestayBusinessBoss, userInfo.User)
			retHomestayBusinessBoss = types.HomestayBusinessBoss{
				Id:       userInfo.User.Id,
				UserId:   userInfo.User.Id,
				NickName: userInfo.User.Nickname,
				Avatar:   userInfo.User.Avatar, // 为什么是空的？？？？  注意Redis缓存
				Info:     userInfo.User.Info,
				Rank:     userInfo.User.Id,
			}
		}
	}

	return &types.HomestayBusinessDetailResp{
		Boss: retHomestayBusinessBoss,
	}, nil
}
