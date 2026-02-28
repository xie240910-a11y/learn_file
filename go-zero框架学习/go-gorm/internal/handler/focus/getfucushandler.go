// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package focus

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-gorm/internal/logic/focus"
	"go-gorm/internal/svc"
)

// 获取轮播图
func GetFucusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := focus.NewGetFucusLogic(r.Context(), svcCtx)
		resp, err := l.GetFucus()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
