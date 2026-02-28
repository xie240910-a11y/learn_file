// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package focus

import (
	"context"

	"go-gorm/internal/svc"
	"go-gorm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFucusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除轮播图
func NewDeleteFucusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFucusLogic {
	return &DeleteFucusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFucusLogic) DeleteFucus(req *types.FocusRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
