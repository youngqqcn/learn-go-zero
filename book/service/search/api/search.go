package main

import (
	"flag"
	"fmt"
	"net/http"

	"book/service/search/api/internal/config"
	"book/service/search/api/internal/handler"
	"book/service/search/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/search-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 全局中间件(当接口请求时生效)
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func (w http.ResponseWriter, r *http.Request)  {
			logx.Info("=============global middleware")
			next(w, r)
		}
	})

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
