package test_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func IsCancelCTX(ctx context.Context, i int) bool {
	select {
	case res, flag := <-ctx.Done(): // Done 返回的是channel
		fmt.Println("context: ", &ctx, "is canceled,", res, flag, ",error info:", ctx.Err(), ",", ctx.Value(i))
		return true
	default:
		return false
	}
}

func TestConText(t *testing.T) {
	/*
	   Background()：返回一个空且不可取消的context,
	   WithCancel：传入父context，创建一个可以取消的子context,
	   cancel： 一个闭包函数（里面包了数据）, 调用后取消该ctx
	*/
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) { // 让ctx控制这几个go routine

			task := fmt.Sprintf("task %d", i)
			vCtx := context.WithValue(ctx, i, task)
			tCtx, cancel := context.WithTimeout(vCtx, 2*time.Second) // 2秒超时context
			defer cancel()                                           // 只取消tCtx，他的父vCtx没被取消

			for {

				if IsCancelCTX(tCtx, i) {
					if IsCancelCTX(vCtx, i) == false { // vCtx 没被超时取消
						fmt.Println("vCtx:", &vCtx, "is not cancel")
					}
					break
				}
				fmt.Printf("%s is run...\n", task)
				time.Sleep(3 * time.Second) // 让超时发生
			}
			fmt.Println("go routine ", i, " is stop")
		}(i, ctx)
	}
	time.Sleep(4 * time.Second)

	fmt.Println("run root cancel:", cancel)

	cancel() // 闭包函数，取消所有的context
	cancel() // 可以重复cancel

	fmt.Println("end root cancel:", cancel)

	// 等待所有 go routine退出
	time.Sleep(2 * time.Second)
}
