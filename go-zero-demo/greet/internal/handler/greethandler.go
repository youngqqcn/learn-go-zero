package handler

import (
	"net/http"

	"github.com/go-zero-demo/greet/internal/logic"
	"github.com/go-zero-demo/greet/internal/svc"
	"github.com/go-zero-demo/greet/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GreetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGreetLogic(r.Context(), svcCtx)
		resp, err := l.Greet(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
