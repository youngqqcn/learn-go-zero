package homestayBusiness

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/travel/cmd/api/internal/logic/homestayBusiness"
	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"
)

func HomestayBusinessDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayBusinessDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := homestayBusiness.NewHomestayBusinessDetailLogic(r.Context(), svcCtx)
		resp, err := l.HomestayBusinessDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
