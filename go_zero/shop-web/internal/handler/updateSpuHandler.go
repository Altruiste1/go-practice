package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-practice/go_zero/shop/internal/logic"
	"go-practice/go_zero/shop/internal/svc"
	"go-practice/go_zero/shop/internal/types"
)

func UpdateSpuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSpuRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateSpuLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSpu(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
