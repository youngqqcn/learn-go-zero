package homestayBusiness

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/travel/cmd/api/internal/logic/homestayBusiness"
	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"
)

func GoodBossHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodBossReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := homestayBusiness.NewGoodBossLogic(r.Context(), svcCtx)
		resp, err := l.GoodBoss(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
