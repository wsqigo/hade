package main

import (
	"github.com/test/framework"
	"github.com/test/framework/middleware"
)

// 注册路由规则
func registerRouter(core *framework.Core) {
	// 在core中使用 middleware.Test3() 为单个路由增加中间件
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	// 需求3: 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test3())
		// 需求4: 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		// 在group中使用middleware.Test3()为单个路由增加中间件
		subjectApi.Get("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
	}
}
