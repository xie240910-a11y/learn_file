// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"router/internal/config"
	"router/model/mysql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	FocusModel   mysql.FocusModel // 轮播图的模型类
	ArticleModel mysql.ArticleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库
	conn := sqlx.NewMysql(c.Mysql.DataSource) // 新建数据库连接
	return &ServiceContext{
		Config:       c,
		FocusModel:   mysql.NewFocusModel(conn),
		ArticleModel: mysql.NewArticleModel(conn),
	}
}
