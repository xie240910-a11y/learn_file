// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package article

import (
	"context"
	"router/internal/svc"
	"router/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByIdLogic {
	return &GetArticleByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleByIdLogic) GetArticleById(req *types.ArticleRequest) (resp *types.CommonResponse, err error) {
	id := req.Id
	article, err := l.svcCtx.ArticleModel.FindOne(l.ctx, int64(id))
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
		Data: types.Article{
			Id:          int(article.Id),
			Title:       article.Title,
			Content:     article.Content,
			Description: article.Description,
			Pic:         article.Pic,
		},
	}, err
}
