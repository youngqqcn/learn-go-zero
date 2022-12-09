package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/xerrors"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/app/usercenter/model"
	"looklook/common/tool"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// 注册逻辑

	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerrors.New("not found"), "not found")
	}

	// 如果用户已经存在， 则报错？
	if user != nil {
		return nil, errors.New("already exits")
	}

	// 开启事务
	var userId int64
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user := new(model.User)
		user.Mobile = in.Mobile
		user.Avatar = in.Nickname

		if len(user.Nickname) == 0 {
			user.Nickname = tool.Krand(8, tool.KC_RAND_KIND_ALL)
		}

		if len(in.Password) > 0 {
			user.Password = tool.MdByString(in.Password)
		}

		// 插入user表
		insertResult, err := l.svcCtx.UserModel.Insert(l.ctx, user)
		if err != nil {
			return errors.New("insert user to db error")
		}

		lastId, err := insertResult.LastInsertId()
		if err != nil {
			return errors.New("insertResult error")
		}

		// 用户Id
		userId = lastId

		userAuth := new(model.UserAuth)
		userAuth.UserId = lastId
		userAuth.AuthKey = in.AuthKey
		userAuth.AuthType = in.AuthType

		// 插入user_auth表
		if _, err := l.svcCtx.UserAuthModel.Insert(ctx, userAuth); err != nil {
			return errors.New("UserAuthModel insert error")
		}

		return nil
	}); err != nil {
		return nil, err
	}

	// 生成token
	genTokenLogin := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := genTokenLogin.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: userId,
	})

	if err != nil {
		return nil, errors.New("GenerateToken error")
	}

	return &pb.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
