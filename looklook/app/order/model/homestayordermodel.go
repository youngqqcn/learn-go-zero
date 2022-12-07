package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HomestayOrderModel = (*customHomestayOrderModel)(nil)

type (
	// HomestayOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayOrderModel.
	HomestayOrderModel interface {
		homestayOrderModel
		UpdateWithVersion(ctx context.Context, data *HomestayOrder) error
	}

	customHomestayOrderModel struct {
		*defaultHomestayOrderModel
	}
)

// NewHomestayOrderModel returns a model for the database table.
func NewHomestayOrderModel(conn sqlx.SqlConn, c cache.CacheConf) HomestayOrderModel {
	return &customHomestayOrderModel{
		defaultHomestayOrderModel: newHomestayOrderModel(conn, c),
	}
}

func (m *defaultHomestayOrderModel) UpdateWithVersion(ctx context.Context, newData *HomestayOrder) error {

	oldVersion := newData.Version
	data, err := m.FindOne(ctx, newData.Id)

	data.Version += 1 // 版本号 +1

	if err != nil {
		return err
	}

	homestayOrderIdKey := fmt.Sprintf("%s%v", cacheHomestayOrderIdPrefix, data.Id)
	homestayOrderSnKey := fmt.Sprintf("%s%v", cacheHomestayOrderSnPrefix, data.Sn)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ?", m.table, homestayOrderRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Version, newData.Sn, newData.UserId, newData.HomestayId, newData.Title, newData.SubTitle, newData.Cover, newData.Info, newData.PeopleNum, newData.RowType, newData.NeedFood, newData.FoodInfo, newData.FoodPrice, newData.HomestayPrice, newData.MarketHomestayPrice, newData.HomestayBusinessId, newData.HomestayUserId, newData.LiveStartDate, newData.LiveEndDate, newData.LivePeopleNum, newData.TradeState, newData.TradeCode, newData.Remark, newData.OrderTotalPrice, newData.FoodTotalPrice, newData.HomestayTotalPrice, newData.Id, oldVersion)
	}, homestayOrderIdKey, homestayOrderSnKey)
	return err
}
