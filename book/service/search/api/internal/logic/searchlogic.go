package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"book/service/search/api/internal/svc"
	"book/service/search/api/internal/types"
	"book/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (resp *types.SearchReply, err error) {
	// todo: add your logic here and delete this line
	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("userId: %s", userIdNumber)
	userId, err := userIdNumber.Int64()
	if err != nil {
        return nil, err
    }

	userInfo, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{
		Id: userId,
	})
	logx.Info("userInfo is %v", userInfo)
	if err != nil {
		return nil, err
	}

	resp = &types.SearchReply {
		Name: req.Name,
		Count: 1000,
	}

	return
}
