package main

import (
	"github.com/test/framework"
)

func SubjectAddController(ctx *framework.Context) error {
	ctx.Json(200, "ok, SubjectAddController")
	return nil
}

func SubjectListController(ctx *framework.Context) error {
	ctx.Json(200, "ok, SubjectListController")
	return nil
}

func SubjectDelController(ctx *framework.Context) error {
	ctx.Json(200, "ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.Json(200, "ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(ctx *framework.Context) error {
	ctx.Json(200, SubjectGetController)
	return nil
}

func SubjectNameController(c *framework.Context) error {
	c.Json(200, "ok, SubjectNameController")
	return nil
}
