package user

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/usercenter/cmd/api/internal/logic/user"
	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	logx.Error("11111111111111111111")
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		logx.Error("222222222222222")
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		logx.Error("3333333333333333333")
		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
