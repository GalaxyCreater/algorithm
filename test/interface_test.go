package test_test

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
	"time"
)

type Human interface {
	Say()
}

type Girl struct {
}

func (self Girl) Say() { // 实现Human的Say接口,不能用*Girl指针
	fmt.Println("girl say")
}

type Boy struct {
}

func (self Boy) Say() { // 实现Human的Say接口,不能用*Boy指针
	fmt.Println("boy say")
}

func WhoSay(h Human) { // 这里可以传递实现改接口的类的对象或指针
	h.Say()
}

func printAll(it interface{}) { // 错误代码，要改成vals []string；没必要用interface{}数组，
	vals := it.([]string)
	for _, val := range vals { // 因为interface{}已经表示任何类型
		fmt.Println(val)
	}
}

func TestInter(t *testing.T) {
	b := Boy{}
	g := Girl{}

	WhoSay(b)
	WhoSay(g)

	bPtr := &Boy{}
	gPtr := &Girl{}
	WhoSay(bPtr)
	WhoSay(gPtr)

	a := []string{"a", "b"}
	printAll(a)

	var i uint = 1
	f1(i)
}
func f1(i interface{}) {
	fmt.Println(i.(int)) // 程序运行到这里的时候会报错
	// interface conversion: interface {} is uint, not int
}

var ErrorA = errors.New("error A occur")
var ErrorB = errors.New("error B occur")

func Compare(a int) error {
	if a < 1 {
		return ErrorA
	}
	return ErrorB
}

func compute() {
	start := time.Now()
	defer func() {
		cost := time.Since(start)
		// caller info
		pc, file, line, ok := runtime.Caller(1)
		if ok == true {
			f := runtime.FuncForPC(pc)
			fmt.Printf("%s:%d  %s , cost=%s \n", file, line, f.Name(), cost)
		} else {
			fmt.Printf("cost = %d\n", cost)
		}

	}()
}

func NowBegin() {
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println("timeStr:", timeStr)
	t, _ := time.Parse("2006-01-02", timeStr)
	timeNumber := t.Unix()
	fmt.Println("timeNumber:", timeNumber, runtime.NumGoroutine())

}
func TestError(t *testing.T) {
	start := time.Now()
	druation := time.Since(start)
	fmt.Printf("get cost time :%s, %d \n", druation, druation)
	NowBegin()
}
