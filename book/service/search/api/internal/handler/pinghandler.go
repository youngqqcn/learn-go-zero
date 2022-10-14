package handler

import (
	"book/response" // 自定义响应
	"book/service/search/api/internal/logic"
	"book/service/search/api/internal/svc"
	"net/http"
)

func pingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewPingLogic(r.Context(), svcCtx)
		err := l.Ping()
		response.Response(w, nil, err)

	}
}
