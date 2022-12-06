package homestayComment

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"
	"looklook/app/travel/model"
	"looklook/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListReq) (*types.CommentListResp, error) {

	whereBuilder := l.svcCtx.HomestayCommentModel.RowBuilder().Where(squirrel.Eq{
		"del_state": model.HomestayCommentNotDeleted, // no_deleted
	})

	homestayCommentList, err := l.svcCtx.HomestayCommentModel.FindPageListByIdDESC(l.ctx, whereBuilder, req.LastId, req.PageSize)
	if err != nil {
		return nil, err
	}

	var ret []types.HomestayComment
	if len(homestayCommentList) > 0 {
		for _, item := range homestayCommentList {
			var homestayComment types.HomestayComment
			copier.Copy(&homestayComment, item)

			// 获取business的信息
			userInfo, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
				Id: homestayComment.UserId,
			})
			if err != nil {
				return nil, err
			}

			copier.Copy(&homestayComment, userInfo)
			ret = append(ret, homestayComment)
		}
	}

	return &types.CommentListResp{
		List: ret,
	}, nil
}
