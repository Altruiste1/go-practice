package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-practice/go_zero/shop/internal/logic"
	"go-practice/go_zero/shop/internal/svc"
	"go-practice/go_zero/shop/internal/types"
)

func ShopHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewShopLogic(r.Context(), svcCtx)
		resp, err := l.Shop(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
