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
	orderFieldNames          = builder.RawFieldNames(&Order{})
	orderRows                = strings.Join(orderFieldNames, ",")
	orderRowsExpectAutoSet   = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`"), ",")
	orderRowsWithPlaceHolder = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`"), "=?,") + "=?"

	cacheOrderIdPrefix = "cache:order:id:"
)

type (
	orderModel interface {
		Insert(ctx context.Context, data *Order) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Order, error)
		Update(ctx context.Context, data *Order) error
		Delete(ctx context.Context, id int64) error
	}

	defaultOrderModel struct {
		sqlc.CachedConn
		table string
	}

	Order struct {
		Id         int64     `db:"id"`
		Uid        int64     `db:"uid"`    // 用户ID
		Pid        int64     `db:"pid"`    // 产品ID
		Amount     int64     `db:"amount"` // 订单金额
		Status     int64     `db:"status"` // 订单状态
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newOrderModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultOrderModel {
	return &defaultOrderModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`order`",
	}
}

func (m *defaultOrderModel) Delete(ctx context.Context, id int64) error {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, orderIdKey)
	return err
}

func (m *defaultOrderModel) FindOne(ctx context.Context, id int64) (*Order, error) {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, id)
	var resp Order
	err := m.QueryRowCtx(ctx, &resp, orderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderRows, m.table)
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

func (m *defaultOrderModel) Insert(ctx context.Context, data *Order) (sql.Result, error) {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, orderRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uid, data.Pid, data.Amount, data.Status)
	}, orderIdKey)
	return ret, err
}

func (m *defaultOrderModel) Update(ctx context.Context, data *Order) error {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uid, data.Pid, data.Amount, data.Status, data.Id)
	}, orderIdKey)
	return err
}

func (m *defaultOrderModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheOrderIdPrefix, primary)
}

func (m *defaultOrderModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultOrderModel) tableName() string {
	return m.table
}
