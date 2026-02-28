// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package focus

import (
	"context"

	"go-gorm/internal/svc"
	"go-gorm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFucusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改轮播图
func NewUpdateFucusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFucusLogic {
	return &UpdateFucusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFucusLogic) UpdateFucus(req *types.UpdateFocusRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
