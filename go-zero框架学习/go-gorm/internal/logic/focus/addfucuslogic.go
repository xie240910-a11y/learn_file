// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package focus

import (
	"context"

	"go-gorm/internal/svc"
	"go-gorm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFucusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增轮播图
func NewAddFucusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFucusLogic {
	return &AddFucusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFucusLogic) AddFucus(req *types.AddFocusRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
