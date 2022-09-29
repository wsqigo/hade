package main

import (
	"github.com/test/framework"
	"github.com/test/framework/middleware"
	"net/http"
)

func main() {
	core := framework.NewCore()
	// core 中使用use注册中间件
	//core.Use(
	//	middleware.Test1(),
	//	middleware.Test2())
	core.Use(middleware.Recovery())

	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
