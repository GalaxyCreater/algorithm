package list

import (
	"fmt"
	"testing"
)

/*
反转链表
*/

/*
{1,2}
1）	p 	mid	last
	nil	1	2		nil
2）	p 	mid	last
	1	2	nil		nil
*/
func (self *Mylist) Rerverse() {
	p := self.head.next
	if p == nil {
		fmt.Println("list is empty")
		return
	}

	if p.next == nil {
		fmt.Println("only one ele")
		return
	}

	pmid := p
	plast := p.next
	p = nil
	for pmid != nil {
		pmid.next = p
		// 后面已经没元素了，结束移动
		if plast == nil {
			break
		}
		// 3个指针都移位
		p = pmid
		pmid = plast
		plast = plast.next

	}

	self.head.next = pmid
}

func TestListRerverse(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	obj := CreateList(arr)
	obj.Print()

	//obj.RerverseTest()
	obj.Rerverse()
	obj.Print()
}

func (self *Mylist) RerverseTest() {
	pre := self.head.next

	var worker = (*Node)(nil)
	if pre != nil {
		worker = pre.next
		self.tail = pre
		self.tail.next = nil
	} else {
		return
	}

	for worker != nil {
		last := worker.next
		worker.next = pre
		pre = worker

		if last == nil {
			self.head.next = worker
			break
		}
		worker = last
	}
}
