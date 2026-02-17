// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"firstdemo/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Mysql  string
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Mysql:  c.Mysql.DataSource, // 数据库连接
	}
}
