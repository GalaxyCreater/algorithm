package list

import (
	"fmt"
	"testing"
)

/*
创建环形链表
*/
func CreateLoopList(lst []int) *Mylist {
	obj := &Mylist{}
	obj.Init()
	for _, v := range lst {
		obj.AddTail(v)
	}

	obj.tail.next = obj.head.next

	return obj
}

func (self *Mylist) LoopListPrint() {
	fmt.Println("------------ loop list ----------------")
	if self.head == nil {
		return
	}

	cur := self.head.next

	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next

		if cur == self.head.next {
			fmt.Println("second round:", cur.data, ",", cur.next.data)
			break
		}
	}
}

/*
是否为环形链表
@res:和环相交的环节点
@leng:环长度,>0表示存在环
*/

var step1 int = 0
var step2 int = 0

func (self *Mylist) CheckIsLoopList() (res *Node, leng int) {
	p1 := self.head.next
	p2 := p1
	res = nil
	leng = 0
	for {
		p2 = p2.next
		step2++
		if p2 == nil {
			return
		} else if p2 == p1 {
			break
		}
		p2 = p2.next
		step2++
		if p2 == nil {
			return
		} else if p2 == p1 {
			break
		}

		p1 = p1.next
	}

	// 存在环
	if p1 == p2 && p1 != nil {
		ringMap := map[*Node]bool{} // 记录环部分节点
		flagP := p1
		worker := p1
		ringMap[worker] = true
		for {
			worker = worker.next
			ringMap[worker] = true
			if worker == flagP {
				break
			}
		}

		leng = len(ringMap) // 环长度

		// 找出和环相交的节点,从链表头重新开始遍历
		worker = self.head.next
		if ringMap[worker] { // 整个链表都是环, 不存相交的节点
			return
		}
		for {
			if ringMap[worker.next] {
				res = worker
				return
			} else {
				worker = worker.next
			}
		}
	}

	return
}

func TestListLoop(t *testing.T) {
	//arr := []int{1, 2, 3, 4}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	obj := CreateLoopList(arr)
	//obj := CreateList(arr)
	//obj.LoopListPrint()
	res, l := obj.CheckIsLoopList()
	fmt.Println(step1, step2)
	if l == 0 {
		fmt.Println("no loop, tar:")
	} else {
		if res == nil {
			fmt.Printf("is ring, ring len:%d\n", l)
		} else {
			fmt.Printf("is line + ring,ring len:%d\n", l)
		}
	}
}

// 12 3123

/*
2

*/
