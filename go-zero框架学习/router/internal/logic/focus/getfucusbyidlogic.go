// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package focus

import (
	"context"

	"router/internal/svc"
	"router/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFucusByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据id获取轮播图
func NewGetFucusByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFucusByIdLogic {
	return &GetFucusByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFucusByIdLogic) GetFucusById(req *types.FocusRequestById) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	id := req.Id
	focus, err := l.svcCtx.FocusModel.FindOne(l.ctx, int64(id))
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
		Data: types.Focus{
			Id:       int(focus.Id),
			Title:    focus.Title,
			Pic:      focus.Pic,
			Link:     focus.Link,
			Position: focus.Position,
		},
	}, err
}
