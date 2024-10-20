package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-practice/go_zero/go_im/user/api/internal/logic"
	"go-practice/go_zero/go_im/user/api/internal/svc"
	"go-practice/go_zero/go_im/user/api/internal/types"
)

func userHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.User(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
