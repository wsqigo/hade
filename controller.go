package main

import (
	"context"
	"fmt"
	"github.com/test/framework"
	"log"
	"time"
)

func FooController(ctx *framework.Context) error {
	// 这个channel负责通知结束
	finish := make(chan struct{}, 1)
	// 这个channel负责通知panic异常
	panicChan := make(chan interface{}, 1)

	// 在业务逻辑处理前，创建有定时器功能的 context
	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), 1*time.Second)
	defer cancel()

	go func() {
		// 这里增加异常处理
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// Do real action
		time.Sleep(10 * time.Second)
		ctx.Json(200, "ok")

		// 新的goroutine结束的时候通过一个finish通道告知负goroutine
		finish <- struct{}{}
	}()

	select {
	// 监听panic
	case p := <-panicChan:
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		log.Println(p)
		ctx.Json(500, "panic")
	// 监听结束事件
	case <-finish:
		fmt.Println("finish")
	// 监听超时事件
	case <-durationCtx.Done():
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		ctx.Json(500, "time out")
		ctx.SetHasTimeout()
	}
	return nil
}
