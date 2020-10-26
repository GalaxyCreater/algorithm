/*
测试限流算法
*/

package test_test

import (
	"alg/test"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/chasex/redis-go-cluster"
)

var nCount uint64 = 0 // 访问次数

func addCount() {
	atomic.AddUint64(&nCount, 1)
}

func TestRds(t *testing.T) {
	v, ok := redis.String(test.RedisQuery("get", "k"))
	fmt.Println(v, ok)
}

func Req() {
	for i := 0; i < 10000; i++ {
		go func() {
			for {
				addCount()
				time.Sleep(1 * time.Microsecond)
			}
		}()
	}
}

func AddToken() {
	go func() {

	}()
}

// 测试redis令牌桶算法
func TestAddToken(t *testing.T) {

}
