package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/usercenter/cmd/rpc/internal/config"
	"looklook/app/usercenter/model"
)

type ServiceContext struct {
	Config config.Config

	UserModel     model.UserModel
	UserAuthModel model.UserAuthModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,

		UserModel:     model.NewUserModel(sqlConn, c.Cache),
		UserAuthModel: model.NewUserAuthModel(sqlConn, c.Cache),
	}

}
