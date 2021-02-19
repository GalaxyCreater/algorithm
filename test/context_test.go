package test_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Monitor(ctx context.Context, name string) {
	go func() {
		n := name
		for {
			select {
			case <-ctx.Done():
				pre := ctx.Value(1)
				fmt.Println(pre, n, "end... because:", ctx.Err())
				return
			default:
				pre := ctx.Value(1)
				fmt.Println(pre, n, "monitor...")
				time.Sleep(2 * time.Second)
			}
		}
	}()
}

func TestCtx(t *testing.T) {
	//func main() {
	bg := context.Background()
	ctx, cancel := context.WithCancel(bg)
	// 让ctx能携带值;传的值是线程安全的,只传递必要的数据,不要什么数据都传
	ctx = context.WithValue(ctx, 1, "监控")

	Monitor(ctx, "1")
	Monitor(ctx, "2")
	Monitor(ctx, "3")

	time.Sleep(4 * time.Second)

	cancel()

	time.Sleep(1 * time.Second)
}

type B1 struct {
	x int
	y int
}

type D1 struct {
	B1

	xx int
}

func TestDev(t *testing.T) {
	b := B1{1, 2}
	d := D1{B1: b}

	fmt.Println(d, d.B1.x, d.x)

}
