// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package focus

import (
	"context"

	"go-gorm/internal/svc"
	"go-gorm/internal/types"
	"go-gorm/modul/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFucusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取轮播图
func NewGetFucusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFucusLogic {
	return &GetFucusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFucusLogic) GetFucus() (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	focus := []gorm.Focus{}
	err = l.svcCtx.DB.Find(&focus).Error
	if err != nil {
		return &types.CommonResponse{
			Code:    500,
			Message: "查询失败",
		}, err
	}
	return &types.CommonResponse{
		Code:    200,
		Message: "查询成功",
		Success: true,
		Data:    focus,
	}, nil
}
