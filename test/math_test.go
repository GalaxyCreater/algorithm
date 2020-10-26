package test_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Circle struct {
	radius float64
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func DeferFunc() int {
	a := 1
	defer func() {
		fmt.Println("first a is", a)
	}()

	defer func() {
		fmt.Println("second a is", a)
	}()

	return a + 1
}
func TestT(t *testing.T) {
	alist := []int{1, 2, 3}
	func(ll []int) {
		blist := ll[:2] // [1,2]
		for i := range blist {
			blist[i] = 4 // [4,4],[4,4,3]
		}

		fmt.Println("ll : ")
		for i := range ll {
			fmt.Println(ll[i])
		}
	}(alist)
}

func TestConvert(t *testing.T) {
	var c uint32 = 0x04030201
	var d [4]byte
	p := unsafe.Pointer(&c)
	q := (*[4]byte)(p)
	copy(d[0:], (*q)[0:])
	fmt.Println(d) // 输出：[1 2 3 4]

}

func ChangeVal(p []Pro) {
	for i := 0; i < len(p); i++ {
		p[i].name = "aaaa"
	}
}

type Pro struct {
	val  int
	name string
}

func ChangeValPtr(p []*Pro) {
	for i := 0; i < len(p); i++ {
		p[i].name = "aaaa"
	}
}

func TestSlic(t *testing.T) {
	l := []int{1, 2, 3, 4}
	insertIdx := 2
	insertVal := 10
	l1 := append([]int{}, l[:insertIdx]...)
	l1 = append(l1, insertVal)
	l = append(l1, l[insertIdx:]...)
	fmt.Println(l) // 输出 [1 2 10 3 4]
}

func TestSlic2(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6}
	insertIdx := 2
	insertVal := 10
	// 错误插入做法
	l1 := append(l[:insertIdx], insertVal, insertVal) // 插入元素一定不能这么写,会修改到原列表
	fmt.Println(l)                                    // 输出:[1 2 10 4]
	l = append(l1, l[insertIdx:]...)                  // l1为:[1 2 10]
	fmt.Println(l)                                    // 输出: [1 2 10 3 4]

	a1 := []int{1, 2}
	a2 := append(a1, 3, 4)
	fmt.Println(a1, a2)

}

func TestM(t *testing.T) {
	m := map[string]int{}
	fmt.Println(m["a"]) // 输出:0(key不存在,返回默认值0, 并不会在map中创建"a")
	v, ok := m["a"]
	fmt.Println(v, ok) // 输出 0 false
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

// 这里方法的每次调用都直接使用Books实例同一块内存
func (self *Books) Read() {
	fmt.Println("Read Books")
}

// 新开辟了Books的一块内存，这里会Books实例的所有成员会进行值复制
func (self Books) Buy() {
	fmt.Println("Buy Books")
}

type GolangBook struct {
	Books //组合方式替代继承
}

func (self *GolangBook) Read() { // 重新实现了Books的Read
	fmt.Println("read GolangBook")
}

func TestStruct(t *testing.T) {

	c1 := &Books{ //创建方式1
		title: "a",
	}
	c2 := new(Books) //创建方式2
	c3 := Books{}    //创建方式3
	fmt.Println(c1, c2, c3)

	ptr := &GolangBook{}
	ptr.Buy()
	ptr.Read()

}

func TestPr(t *testing.T) {
	fmt.Println(rune('a'))
}
