// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package focus

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-gorm/internal/logic/focus"
	"go-gorm/internal/svc"
	"go-gorm/internal/types"
)

// 根据id获取轮播图
func GetFucusByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FocusRequestById
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := focus.NewGetFucusByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetFucusById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
