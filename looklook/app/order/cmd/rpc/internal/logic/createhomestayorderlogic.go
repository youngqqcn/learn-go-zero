package logic

import (
	"context"
	"looklook/app/order/cmd/rpc/internal/svc"
	"looklook/app/order/cmd/rpc/pb"
	"looklook/app/order/model"
	"looklook/app/travel/cmd/rpc/travel"
	"looklook/common/tool"
	"looklook/common/uniqueid"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHomestayOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateHomestayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHomestayOrderLogic {
	return &CreateHomestayOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 下单
func (l *CreateHomestayOrderLogic) CreateHomestayOrder(in *pb.CreateHomestayOrderReq) (*pb.CreateHomestayOrderResp, error) {
	h, err := l.svcCtx.TravelRpc.HomestayDetail(l.ctx, &travel.HomestayDetailReq{
		Id: in.HomestayId,
	})
	if err != nil {
		return nil, err
	}

	// TODO: 参数校验

	var order model.HomestayOrder
	order.DeleteTime = time.Unix(0, 0) // 注意时间

	order.Sn = uniqueid.GenSn(uniqueid.SN_PREFIX_HOMESTAY_ORDER) // "12345" + fmt.Sprintf("%d", in.UserId)
	order.UserId = in.UserId
	order.HomestayId = in.HomestayId
	order.Title = h.Homestay.Title
	order.SubTitle = h.Homestay.SubTitle
	order.Cover = h.Homestay.Banner
	order.Info = h.Homestay.Info
	order.PeopleNum = h.Homestay.PeopleNum
	order.RowType = h.Homestay.RowType
	order.FoodInfo = h.Homestay.FoodInfo
	order.FoodPrice = h.Homestay.FooPrice
	order.MarketHomestayPrice = h.Homestay.MarketHomestayPrice
	order.HomestayBusinessId = h.Homestay.HomestayBusinessId
	order.HomestayUserId = h.Homestay.UserId
	order.LiveStartDate = time.Unix(in.LiveStartTime, 0)
	order.LiveEndDate = time.Unix(in.LiveEndTime, 0)
	order.LivePeopleNum = in.LivePeopleNum // 实际住几人
	order.TradeState = model.HomestayOrderTradeStateWaitPay
	order.TradeCode = tool.Krand(8, tool.KC_RAND_KIND_ALL)
	order.Remark = in.Remark
	order.OrderTotalPrice = 0
	order.FoodTotalPrice = 0
	order.HomestayTotalPrice = 0
	order.NeedFood = model.HomestayOrderNeedFoodNo

	if in.IsFood {
		order.NeedFood = model.HomestayOrderNeedFoodYes
	}

	liveDays := int64(order.LiveEndDate.Sub(order.LiveStartDate).Seconds() / 86400) //Stayed a few days in total
	order.HomestayTotalPrice = int64(h.Homestay.HomestayPrice * liveDays)           //Calculate the total price of the B&B
	if in.IsFood {
		order.NeedFood = model.HomestayOrderNeedFoodYes
		order.FoodTotalPrice = int64(h.Homestay.FooPrice * in.LivePeopleNum * liveDays)
	}
	order.OrderTotalPrice = order.HomestayTotalPrice + order.FoodTotalPrice //Calculate total order price.

	_, err = l.svcCtx.HomestayOrderModel.Insert(l.ctx, &order)
	if err != nil {
		return nil, err
	}

	// TODO: Send Msg to MQ

	return &pb.CreateHomestayOrderResp{
		Sn: order.Sn,
	}, nil
}
