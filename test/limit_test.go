/*
测试限流算法
*/

package test_test

import (
	"alg/test"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var preCount uint64 = 0
var nCount uint64 = 0 // 访问次数
var rwMutex sync.RWMutex

func addCount() {
	atomic.AddUint64(&nCount, 1)
}

func getCount() uint64 {
	cnt := atomic.LoadUint64(&nCount)
	return cnt
}

func PrintCount() {
	go func() {
		for {
			cnt := getCount()
			fmt.Println("req count:", cnt, ", increse count:", cnt-preCount)
			preCount = cnt
			time.Sleep(time.Second)
		}
	}()
}

func Req(l test.CurrentLimit) {
	for i := 0; i < 50; i++ {
		go func() {
			for {
				flag := l.IsLimit()
				if false == flag {
					addCount()
				}

				time.Sleep(10 * time.Microsecond)
			}
		}()
	}
}

func TestRds(t *testing.T) {
	test.RedisQuery("lpush", "limit_pub", 1)
}

// 测试redis令牌桶算法
func TestAddToken(t *testing.T) {
	PrintCount()
	//Req(test.LimitAlg)

	time.Sleep(100 * time.Second)
}
