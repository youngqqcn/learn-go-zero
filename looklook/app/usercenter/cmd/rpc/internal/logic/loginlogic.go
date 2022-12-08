package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/app/usercenter/model"
	"looklook/common/tool"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {

	var (
		userId int64
		err    error
	)
	// 获取用户Id
	switch in.AuthType {
	case model.UserAuthTypeSystem:
		userId, err = l.loginByMobile(in.AuthKey, in.Password)
	default:
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// 根据UserId 生成本次的jwt token
	tokenResp, err := NewGenerateTokenLogic(l.ctx, l.svcCtx).GenerateToken(&usercenter.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Errorf("NewGenerateTokenLogic  errror: %v", err)
	}

	return &pb.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) loginByMobile(mobile, password string) (int64, error) {
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	if err != nil {
		return 0, errors.Errorf("svcCtx.UserModel.FindOneByMobile error: %v", user)
	}

	if user == nil {
		return 0, errors.Errorf("user not exists")
	}

	logx.Infof("password: %v", user.Password)
	logx.Infof("req password: %v", tool.MdByString(password))
	if !(tool.MdByString(password) == user.Password) {
		//return 0, errors.Errorf("wrong username or password")
		return 0, errors.Errorf("wrong password")
	}

	return user.Id, nil
}
