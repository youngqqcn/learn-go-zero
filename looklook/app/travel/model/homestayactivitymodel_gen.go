// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	homestayActivityFieldNames          = builder.RawFieldNames(&HomestayActivity{})
	homestayActivityRows                = strings.Join(homestayActivityFieldNames, ",")
	homestayActivityRowsExpectAutoSet   = strings.Join(stringx.Remove(homestayActivityFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	homestayActivityRowsWithPlaceHolder = strings.Join(stringx.Remove(homestayActivityFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"

	cacheHomestayActivityIdPrefix = "cache:homestayActivity:id:"
)

type (
	homestayActivityModel interface {
		Insert(ctx context.Context, data *HomestayActivity) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*HomestayActivity, error)
		Update(ctx context.Context, data *HomestayActivity) error
		Delete(ctx context.Context, id int64) error
	}

	defaultHomestayActivityModel struct {
		sqlc.CachedConn
		table string
	}

	HomestayActivity struct {
		Id         int64     `db:"id"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		DeleteTime time.Time `db:"delete_time"`
		DelState   int64     `db:"del_state"`
		RowType    string    `db:"row_type"`   // 活动类型
		DataId     int64     `db:"data_id"`    // 业务表id（id跟随活动类型走）
		RowStatus  int64     `db:"row_status"` // 0:下架 1:上架
		Version    int64     `db:"version"`    // 版本号
	}
)

func newHomestayActivityModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultHomestayActivityModel {
	return &defaultHomestayActivityModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`homestay_activity`",
	}
}

func (m *defaultHomestayActivityModel) Delete(ctx context.Context, id int64) error {
	homestayActivityIdKey := fmt.Sprintf("%s%v", cacheHomestayActivityIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, homestayActivityIdKey)
	return err
}

func (m *defaultHomestayActivityModel) FindOne(ctx context.Context, id int64) (*HomestayActivity, error) {
	homestayActivityIdKey := fmt.Sprintf("%s%v", cacheHomestayActivityIdPrefix, id)
	var resp HomestayActivity
	err := m.QueryRowCtx(ctx, &resp, homestayActivityIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", homestayActivityRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHomestayActivityModel) Insert(ctx context.Context, data *HomestayActivity) (sql.Result, error) {
	homestayActivityIdKey := fmt.Sprintf("%s%v", cacheHomestayActivityIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, homestayActivityRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.RowType, data.DataId, data.RowStatus, data.Version)
	}, homestayActivityIdKey)
	return ret, err
}

func (m *defaultHomestayActivityModel) Update(ctx context.Context, data *HomestayActivity) error {
	homestayActivityIdKey := fmt.Sprintf("%s%v", cacheHomestayActivityIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, homestayActivityRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.RowType, data.DataId, data.RowStatus, data.Version, data.Id)
	}, homestayActivityIdKey)
	return err
}

func (m *defaultHomestayActivityModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheHomestayActivityIdPrefix, primary)
}

func (m *defaultHomestayActivityModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", homestayActivityRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultHomestayActivityModel) tableName() string {
	return m.table
}
