package user

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	wx "github.com/silenceper/wechat/v2"
	cache "github.com/silenceper/wechat/v2/cache"
	wxconfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/app/usercenter/model"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxMiniAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxMiniAuthLogic {
	return &WxMiniAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxMiniAuthLogic) WxMiniAuth(req *types.WXMiniAuthReq) (resp *types.WXMiniAuthResp, err error) {
	// 微信小程序授权？

	miniProgram := wx.NewWechat().GetMiniProgram(&wxconfig.Config{
		AppID:     l.svcCtx.Config.WxConf.AppId,
		AppSecret: l.svcCtx.Config.WxConf.Secret,
		Cache:     cache.NewMemory(),
	})

	// 获取
	authResp, err := miniProgram.GetAuth().Code2Session(req.Code)
	if err != nil || authResp.ErrCode != 0 || authResp.OpenID == "" {
		return nil, errors.Errorf("获取授权失败: %v", err.Error())
	}

	// 解析
	userData, err := miniProgram.GetEncryptor().Decrypt(authResp.SessionKey, req.EncryptedData, req.IV)
	if err != nil {
		return nil, errors.Errorf("解析数据失败: %v", err.Error())
	}

	rpcRsp, err := l.svcCtx.UserCenterRpc.GetUserAuthByAuthKey(l.ctx, &usercenter.GetUserAuthByAuthKeyReq{
		AuthType: model.UserAuthTypeSmallWX,
		AuthKey:  authResp.OpenID,
	})
	if err != nil {
		return nil, errors.Errorf("l.svcCtx.UserCenterRpc.GetUserAuthByAuthKey error: %v", err.Error())
	}

	if rpcRsp.UserAuth == nil || rpcRsp.UserAuth.Id == 0 {
		// 如果已经是没有注册的，进行绑定（注册）

		mobile := userData.PhoneNumber
		nickName := fmt.Sprintf("gooook%s", mobile[4:])

		// 密码不填
		registerResp, err := l.svcCtx.UserCenterRpc.Register(l.ctx, &usercenter.RegisterReq{
			Mobile:   mobile,
			Nickname: nickName,
			AuthType: model.UserAuthTypeSmallWX,
			AuthKey:  authResp.OpenID,
		})
		if err != nil {
			return nil, errors.Errorf("l.svcCtx.UserCenterRpc.Register rpc error: %v", err.Error())
		}

		// 返回jwt token
		return &types.WXMiniAuthResp{
			AccessToken:  registerResp.AccessToken,
			AccessExpire: registerResp.AccessExpire,
			RefreshAfter: registerResp.RefreshAfter,
		}, nil
	}

	// 如果已经绑定（注册），直接返回jwt token即可
	tokenResp, err := l.svcCtx.UserCenterRpc.GenerateToken(l.ctx, &usercenter.GenerateTokenReq{
		UserId: rpcRsp.UserAuth.UserId,
	})
	if err != nil {
		return nil, errors.Errorf("l.svcCtx.UserCenterRpc.GenerateToken: %v", err)
	}
	return &types.WXMiniAuthResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
