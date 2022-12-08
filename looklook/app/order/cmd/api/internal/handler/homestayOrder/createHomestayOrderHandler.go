package homestayOrder

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/order/cmd/api/internal/logic/homestayOrder"
	"looklook/app/order/cmd/api/internal/svc"
	"looklook/app/order/cmd/api/internal/types"
)

func CreateHomestayOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateHomestayOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := homestayOrder.NewCreateHomestayOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateHomestayOrder(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
