// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package focus

import (
	"context"

	"router/internal/svc"
	"router/internal/types"

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
	// 获取轮播图数据
	focus, err := l.svcCtx.FocusModel.FindAll(l.ctx)
	if err != nil {
		return &types.CommonResponse{
			Code:    500,
			Message: "查询失败",
		}, err
	}
	var focusList []*types.Focus
	for _, data := range focus {
		focusList = append(focusList, &types.Focus{
			Id:       int(data.Id),
			Title:    data.Title,
			Pic:      data.Pic,
			Link:     data.Link,
			Position: data.Position,
		})
	}
	return &types.CommonResponse{
		Code:    200,
		Message: "查询成功",
		Data:    focusList,
		Success: true,
	}, nil
}
