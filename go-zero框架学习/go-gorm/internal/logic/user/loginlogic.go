// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"
	"time"

	"go-gorm/internal/biz"
	"go-gorm/internal/svc"
	"go-gorm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取轮播图
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	token, _ := biz.GetJwtToken(l.svcCtx.Config.Auth.Secret, time.Now().Unix(), l.svcCtx.Config.Auth.Expire, req.UserName)
	return biz.Success(token), nil
}
