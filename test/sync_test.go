package test_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

func TestSyncPool(t *testing.T) {
	p := sync.Pool{
		New: func() interface{} { // 初始化New函数
			fmt.Println("create ...")
			return "default"
		},
	}

	fmt.Println(p.Get().(string)) // 池为空时，Get会执行New
	p.Put("a")
	p.Put("b")
	fmt.Println(p.Get().(string)) // 先进先出
	fmt.Println(p.Get().(string))
	fmt.Println("--------------")
	fmt.Println(p.Get().(string))
	fmt.Println(p.Get().(string))
}

type ObjInfo struct {
	Name string
}

var once sync.Once // 只执行一次,线程安全
var G_obj *ObjInfo

func GetSingleObj() *ObjInfo {
	once.Do(
		func() {
			fmt.Println("create obj")
			G_obj = &ObjInfo{
				Name: "aaa",
			}
		})
	return G_obj
}

func TestSingleObj(t *testing.T) {
	wg := sync.WaitGroup{}
	count := 5
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			obj := GetSingleObj()
			fmt.Printf("%d obj address: %x\n", i, unsafe.Pointer(obj))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestSyMap(t *testing.T) {

}
