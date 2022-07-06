package svc

import (
	"shuturl/rpc/internal/config"
	"shuturl/rpc/transform/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	c config.Config
	Model model.ShorturlModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := &ServiceContext{
		c: c,
		Model: model.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
	return svc
}
