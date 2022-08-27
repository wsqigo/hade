package main

import "github.com/test/framework"

func registerRouter(core *framework.Core) {
	// 设置控制器
	core.Get("foo", FooController)
}
