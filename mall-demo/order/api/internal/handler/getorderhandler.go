package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mall-demo/order/api/internal/logic"
	"mall-demo/order/api/internal/svc"
	"mall-demo/order/api/internal/types"
)

func getOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetOrderLogic(r.Context(), svcCtx)
		resp, err := l.GetOrder(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}