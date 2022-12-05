package homestayBusiness

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/mr"
	"looklook/app/travel/model"
	"looklook/app/usercenter/cmd/rpc/usercenter"

	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodBossLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodBossLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodBossLogic {
	return &GoodBossLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodBossLogic) GoodBoss(req *types.GoodBossReq) (resp *types.GoodBossResp, err error) {
	// todo: add your logic here and delete this line

	whereBuilder := l.svcCtx.HomestayActivityModel.RowBuilder().Where(squirrel.Eq{
		"row_type":   model.HomestayActivityGoodBusiType,
		"row_status": model.HomestayActivityUpStatus,
	})

	homestayActivityList, err := l.svcCtx.HomestayActivityModel.FindPageListByPage(l.ctx, whereBuilder, 0, 10, "data_id DESC")
	if err != nil {
		return nil, errors.New("l.svcCtx.HomestayActivityModel.FindPageListByPage error")
	}

	var retList []types.HomestayBusinessBoss
	if len(homestayActivityList) > 0 {
		// 使用map-reduce
		mr.MapReduceVoid(func(source chan<- interface{}) {
			// 生成器
			for _, homestayActivity := range homestayActivityList {
				source <- homestayActivity.DataId
			}

		}, func(item interface{}, writer mr.Writer, cancel func(error)) {
			// mapper 函数
			id := item.(int64)

			//homestay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, id)
			userResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
				Id: id,
			})
			if err != nil && err != model.ErrNotFound {
				logx.WithContext(l.ctx).Errorf("l.svcCtx.UsercenterRpc.GetUserInfo error")
				return
			}

			if userResp.User != nil && userResp.User.Id >= 0 {
				writer.Write(userResp.User)
			}
		}, func(pipe <-chan interface{}, cancel func(error)) {
			// reducer 函数

			for item := range pipe {
				var typeHomestayBusinessBoss types.HomestayBusinessBoss
				copier.Copy(&typeHomestayBusinessBoss, item)

				retList = append(retList, typeHomestayBusinessBoss)
			}
		})
	}

	resp = &types.GoodBossResp{
		List: retList,
	}
	err = nil
	return
}
