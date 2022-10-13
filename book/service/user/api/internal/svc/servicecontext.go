package svc

import (
	"book/service/user/api/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"book/service/user/model"
)

type ServiceContext struct {
	Config config.Config

	// 增加自定义
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
