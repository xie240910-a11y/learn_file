// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"
	"router/internal/svc"
	"router/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	l.Info("新增用户：", req.Name, req.Age, req.Sex)

	return &types.CommonResponse{
		Code:    200,
		Message: "success",
		Data: types.AddUserRequest{
			Name: req.Name,
			Age:  req.Age,
			Sex:  req.Sex,
		},
		Success: true,
	}, nil
}
