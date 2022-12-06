package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/globalkey"
)

var _ HomestayCommentModel = (*customHomestayCommentModel)(nil)

type (
	// HomestayCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayCommentModel.
	HomestayCommentModel interface {
		homestayCommentModel
		RowBuilder() squirrel.SelectBuilder
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayComment, error)
	}

	customHomestayCommentModel struct {
		*defaultHomestayCommentModel
	}
)

// NewHomestayCommentModel returns a model for the database table.
func NewHomestayCommentModel(conn sqlx.SqlConn, c cache.CacheConf) HomestayCommentModel {
	return &customHomestayCommentModel{
		defaultHomestayCommentModel: newHomestayCommentModel(conn, c),
	}
}

func (m *defaultHomestayCommentModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(homestayCommentRows).From(m.table)
}

func (m *defaultHomestayCommentModel) FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayComment, error) {

	if preMinId > 0 {
		rowBuilder = rowBuilder.Where(" id < ? ", preMinId)
	}

	//query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	//if err != nil {
	//	return nil, err
	//}
	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	logx.Infof("=== %v", query)

	var resp []*HomestayComment
	//logx.Info("&&&&&&&&&&&&&&&&&&&&&&&&&")
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		logx.Info("xxxxxxxxxxxxxxxxx, %v", err.Error())
		return nil, err
	}
}
