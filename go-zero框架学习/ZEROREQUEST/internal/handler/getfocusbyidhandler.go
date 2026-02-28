// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"ZEROREQUEST/internal/logic"
	"ZEROREQUEST/internal/svc"
	"ZEROREQUEST/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetFocusByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FocusRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetFocusByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetFocusById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
