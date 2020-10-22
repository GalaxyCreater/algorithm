package list

import "fmt"

type Node struct {
	data int
	next *Node
}

type Mylist struct {
	head *Node // (一般不存有效元素,但不为nil)
	tail *Node
}

func (self *Mylist) Init() {

	self.head = &Node{
		data: 0,
		next: nil,
	}
	self.tail = self.head
}

func (self *Mylist) AddTail(data int) {
	p := &Node{
		data: data,
		next: nil,
	}

	self.tail.next = p
	self.tail = p
}

func (self *Mylist) AddHead(data int) {
	p := &Node{
		data: data,
		next: nil,
	}
	// 第一插入
	if self.head == self.tail && self.head.next == nil {
		self.tail.next = p
		self.tail = p
	} else {
		p.next = self.head.next
		self.head.next = p
	}
}

func (self *Mylist) AddHeadPtr(p *Node) {
	// 第一插入
	if self.head == self.tail && self.head.next == nil {
		self.tail.next = p
		self.tail = p
	} else {
		p.next = self.head.next
		self.head.next = p
	}
}

func CreateList(lst []int) *Mylist {
	obj := &Mylist{}
	obj.Init()
	for _, v := range lst {
		obj.AddTail(v)
	}

	return obj
}

func (self *Mylist) Print() {
	fmt.Println("------------ list ----------------")
	if self.head == nil {
		return
	}

	cur := self.head.next
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}
}
