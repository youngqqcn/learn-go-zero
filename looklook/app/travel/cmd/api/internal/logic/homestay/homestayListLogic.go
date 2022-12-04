package homestay

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"
	"looklook/app/travel/model"

	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayListLogic {
	return &HomestayListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayListLogic) HomestayList(req *types.HomestayListReq) (resp *types.HomestayListResp, err error) {

	whereBuilder := l.svcCtx.HomestayActivityModel.RowBuilder().Where(squirrel.Eq{
		"row_type":   model.HomestayActivityPreferredType,
		"row_status": model.HomestayActivityUpStatus,
	})

	homestayActivityList, err := l.svcCtx.HomestayActivityModel.FindPageListByPage(l.ctx, whereBuilder, req.Page, req.PageSize, "data_id DESC")
	if err != nil {
		return nil, errors.New("l.svcCtx.HomestayActivityModel.FindPageListByPage error")
	}

	var retList []types.Homestay
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

			homestay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, id)
			if err != nil && err != model.ErrNotFound {
				logx.WithContext(l.ctx).Errorf("l.svcCtx.HomestayModel.FindOne error")
				return
			}
			writer.Write(homestay)
		}, func(pipe <-chan interface{}, cancel func(error)) {
			// reducer 函数

			for item := range pipe {
				homestay := item.(*model.Homestay)

				var typeHomestay types.Homestay
				copier.Copy(&typeHomestay, homestay)

				retList = append(retList, typeHomestay)
			}
		})
	}

	resp = &types.HomestayListResp{
		List: retList,
	}
	err = nil
	return
}
