package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	// 增加自定的配置

	// 数据库
	MySql struct {
		DataSource string
	}

	CacheRedis cache.CacheConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
