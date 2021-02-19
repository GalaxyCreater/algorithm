package test_test

import (
	"fmt"
	"sync/atomic"
	"testing"
)

var count int32 = 0

func getCount1() int32 {
	a := atomic.LoadInt32(&count)
	//a := count
	return a
}

func TestAotm(t *testing.T) {
	fmt.Println(atomic.AddInt32(&count, 1))
	c := getCount1()
	fmt.Println(atomic.AddInt32(&count, 1))
	fmt.Println(c)
}
