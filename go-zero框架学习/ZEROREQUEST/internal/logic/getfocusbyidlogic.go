// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"ZEROREQUEST/internal/svc"
	"ZEROREQUEST/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFocusByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFocusByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFocusByIdLogic {
	return &GetFocusByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFocusByIdLogic) GetFocusById(req *types.FocusRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	l.Info("获取：", "GetFocusById", req.Id)
	resp = &types.CommonResponse{
		Code:    200,
		Message: "success",
		Success: true,
		Data: types.Focus{
			Id:    req.Id,
			Title: "轮播图2",
			Link:  "https://www.baidu.com",
			Image: "https://pisum/100/200",
		},
	}
	return
}
