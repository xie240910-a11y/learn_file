// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"
	"fmt"

	"go-gorm/internal/config"
	"go-gorm/internal/handler"
	"go-gorm/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/orm-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	// server.Start()
	fmt.Printf("Starting server at 0.0.0.0:%d...\n", c.Port)
	server.Start() // 在 0.0.0.0:8888 上启动
}
