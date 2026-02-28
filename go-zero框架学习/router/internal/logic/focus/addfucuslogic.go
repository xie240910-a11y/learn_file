// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package focus

import (
	"context"

	"router/internal/svc"
	"router/internal/types"
	"router/model/mysql"

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
	focus := mysql.Focus{
		Title:    req.Title,
		Pic:      req.Pic,
		Link:     req.Link,
		Position: req.Position,
	}
	_, err = l.svcCtx.FocusModel.Insert(l.ctx, &focus)
	if err != nil {
		return &types.CommonResponse{
			Code:    500,
			Message: "增加数据失败",
		}, err
	}
	return &types.CommonResponse{
		Code:    200,
		Message: "增加数据成功",
		Success: true,
	}, nil

}
