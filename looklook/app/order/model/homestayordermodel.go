package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/globalkey"
)

var _ HomestayOrderModel = (*customHomestayOrderModel)(nil)

type (
	// HomestayOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayOrderModel.
	HomestayOrderModel interface {
		homestayOrderModel
		RowBuilder() squirrel.SelectBuilder
		UpdateWithVersion(ctx context.Context, data *HomestayOrder) error
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayOrder, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayOrder, error)
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

func (m *defaultHomestayOrderModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayOrder, error) {
	if orderBy == "" {
		rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder.OrderBy(orderBy)
	}

	// 分页处理
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayOrder
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}

	return []*HomestayOrder{}, nil
}

// export to logic use
func (m *defaultHomestayOrderModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(homestayOrderRows).From(m.table)
}

func (m *defaultHomestayOrderModel) FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayOrder, error) {

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayOrder
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}

	return []*HomestayOrder{}, nil
}
