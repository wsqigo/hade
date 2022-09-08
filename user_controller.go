package main

import "github.com/test/framework"

func UserLoginController(ctx *framework.Context) error {
	// 打印控制器名字
	ctx.Json(200, "ok, UserLoginController")
	return nil
}
