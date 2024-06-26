package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/globalkey"
)

var _ HomestayModel = (*customHomestayModel)(nil)

type (
	// HomestayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayModel.
	HomestayModel interface {
		homestayModel
		RowBuilder() squirrel.SelectBuilder
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Homestay, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Homestay, error)
	}

	customHomestayModel struct {
		*defaultHomestayModel
	}
)

// NewHomestayModel returns a model for the database table.
func NewHomestayModel(conn sqlx.SqlConn, c cache.CacheConf) HomestayModel {
	return &customHomestayModel{
		defaultHomestayModel: newHomestayModel(conn, c),
	}
}

// export to logic use
func (m *defaultHomestayModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(homestayRows).From(m.table)
}

func (m *defaultHomestayModel) FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Homestay, error) {
	if preMinId > 0 {
		rowBuilder = rowBuilder.Where(" id < ? ", preMinId)
	}

	querySql, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Homestay
	err = m.QueryRowsNoCacheCtx(ctx, &resp, querySql, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Homestay, error) {
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

	var resp []*Homestay
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}

}
