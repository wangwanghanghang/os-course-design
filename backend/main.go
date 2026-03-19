package main

import (
	"os-course-design/api"
)

func main() {
	// 初始化路由
	r := api.SetupRouter()

	// 启动服务，监听 8080 端口
	println("后端服务启动成功，监听端口 8080...")
	r.Run(":8080")
}

// 这是我为了测试 Git 提交新加的注释
