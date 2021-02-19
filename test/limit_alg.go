/*
各种限流算法实现
*/
package test

import (
	"fmt"
	"time"

	"github.com/chasex/redis-go-cluster"
)

type CurrentLimit interface {
	IsLimit() bool
	AddToken()
}

var maxQps int64 = 10000

//var LimitAlg = &CurrentLimit1{}
var LimitAlg = &CurrentLimitPub{}

func init() {
	LimitAlg.AddToken()
}

/*
固定时间窗口限流算法
*/
type CurrentLimit1 struct {
}

func (self CurrentLimit1) IsLimit() bool {
	cur, err := redis.Int64(RedisQuery("incr", "limit_key"))
	if err != nil {
		fmt.Println(err)
		return true
	}
	if cur > maxQps {
		return true
	}

	return false
}

func (self CurrentLimit1) AddToken() {
	go func() {
		for {
			RedisQuery("set", "limit_key", 0)
			time.Sleep(time.Second)
		}
	}()
}

/*令牌桶限流算法*/
type CurrentLimitPub struct {
}

func (self CurrentLimitPub) IsLimit() (lm bool) {
	lm = true
	cur, err := redis.Int64(RedisQuery("lpop", "limit_pub"))
	if err != nil {
		// fmt.Println("--------", err)
		return
	}
	if cur == 1 {
		lm = false
		return
	}

	return
}

func (self CurrentLimitPub) AddToken() {
	go func() {
		for {
			_, err := RedisQuery("lpush", "limit_pub", 1)
			if err != nil {
				fmt.Println("+++++", err)
			}
			inter := time.Duration(float64(1*time.Second) / float64(maxQps))
			time.Sleep(inter)
		}
	}()
}
