package main

import (
	"terraform-manager/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 注册路由
	routes.RegisterRoutes(r)

	// 启动服务
	r.Run(":8080") // 默认监听 8080 端口
}
