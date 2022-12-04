package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/globalkey"
)

var _ HomestayActivityModel = (*customHomestayActivityModel)(nil)

type (
	// HomestayActivityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayActivityModel.
	HomestayActivityModel interface {
		homestayActivityModel
		RowBuilder() squirrel.SelectBuilder
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayActivity, error)
	}

	customHomestayActivityModel struct {
		*defaultHomestayActivityModel
	}
)

// NewHomestayActivityModel returns a model for the database table.
func NewHomestayActivityModel(conn sqlx.SqlConn, c cache.CacheConf) HomestayActivityModel {
	return &customHomestayActivityModel{
		defaultHomestayActivityModel: newHomestayActivityModel(conn, c),
	}
}

func (m *defaultHomestayActivityModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayActivity, error) {
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

	resp := new([]*HomestayActivity)
	m.QueryRowsNoCacheCtx(ctx, resp, query, values...)
	switch err {
	case nil:
		return *resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultHomestayActivityModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(homestayActivityRows).From(m.table)
}
