// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"ZEROREQUEST/internal/svc"
	"ZEROREQUEST/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFocusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFocusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFocusLogic {
	return &GetFocusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFocusLogic) GetFocus() (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	resp = &types.CommonResponse{
		Code:    200,
		Message: "success",
		Success: true,
		Data: []types.Focus{
			{
				Id:    "1",
				Title: "轮播图1",
				Image: "https://pisum/100/200",
				Link:  "https://www.baidu.com",
			},
			{
				Id:    "2",
				Title: "轮播图2",
				Link:  "https://www.baidu.com",
				Image: "https://pisum/100/200",
			},
		},
	}
	return
}
