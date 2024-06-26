package logic

import (
	"book/service/user/model"
	"context"
	"strings"
	"time"

	"book/common/errorx"
	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"

	"github.com/golang-jwt/jwt"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
	}

	// number学号作为username
	userInfo, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, req.Username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("用户不存在")
	default:
		return nil, err
	}

	if userInfo.Password != req.Password {
		return nil, errorx.NewDefaultError("用户名或密码错误")
	}

	// ---start---
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
	if err != nil {
		return nil, err
	}
	// ---end---

	return &types.LoginResponse{
		Id:           userInfo.Id,
		Name:         userInfo.Name,
		Gender:       userInfo.Gender,
		AccessSecret: jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
