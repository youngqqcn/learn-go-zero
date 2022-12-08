package user

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/ctxdata"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Detail 如果用Postman进行测试，需要在Postman的Authorization选择Bearer Token， 填写jwt的token
func (l *DetailLogic) Detail(req *types.UserInfoReq) (*types.UserInfoResp, error) {

	userId := ctxdata.GetUidFromCtx(l.ctx)

	userInfo, err := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		Id: userId,
	})
	logx.Infof("======User Info: %v", userInfo.User)
	if err != nil {
		return nil, err
	}

	var respUser types.User
	copier.Copy(&respUser, userInfo.User)

	return &types.UserInfoResp{
		UserInfo: respUser,
	}, nil
}
