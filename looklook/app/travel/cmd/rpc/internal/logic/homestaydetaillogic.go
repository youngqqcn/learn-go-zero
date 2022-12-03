package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"looklook/app/travel/cmd/rpc/internal/svc"
	"looklook/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomestayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayDetailLogic {
	return &HomestayDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// homestayDetail
func (l *HomestayDetailLogic) HomestayDetail(in *pb.HomestayDetailReq) (*pb.HomestayDetailResp, error) {
	// todo: add your logic here and delete this line

	hs, err := l.svcCtx.HomestayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("homestay findone error")
	}

	homestay := new(pb.Homestay)
	if hs != nil {
		_ = copier.Copy(homestay, hs)
	}

	return &pb.HomestayDetailResp{
		Homestay: homestay,
	}, nil
}
