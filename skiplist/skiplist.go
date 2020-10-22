package skiplist

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	//最高层数
	MAX_LEVEL = 16
)

type NodeData interface {
	Value()
}

type StringData struct {
	Val string
}

func (self *StringData) Value() {

}

//跳表节点结构体
type skipListNode struct {
	//跳表保存的值,比如人名
	key NodeData
	//用于排序的分值
	val int
	//用有的层高
	level int
	//每层的前进指针, 指向当前层下一个元素，数组索引即层数
	next []*skipListNode
}

type SkipList struct {
	//跳表头结点
	head *skipListNode
	//跳表当前层数
	level int
	//跳表长度
	length int
}

func ConsturctSkipList() *SkipList {
	return &SkipList{
		head:   newNode(nil, math.MinInt32, MAX_LEVEL),
		level:  1,
		length: 0,
	}
}

func newNode(key NodeData, v, level int) *skipListNode {
	return &skipListNode{
		key:   key,
		val:   v,
		level: level,
		next:  make([]*skipListNode, level, level),
	}
}

func Rand() int32 {
	rand.Seed(int64(time.Now().Nanosecond() / 1000))
	return rand.Int31()
}

func (self *SkipList) Find(key NodeData, val int) *skipListNode {
	if self.length == 0 || key == nil {
		return nil
	}

	cur := self.head
	for i := MAX_LEVEL - 1; i >= 0; i-- { // 逐层查找
		for cur.next[i] != nil {
			if cur.next[i].val < val {
				cur = cur.next[i]
			} else if cur.next[i].val == val && cur.next[i].key == key {
				return cur.next[i]
			} else {
				break
			}
		}
	}

	return nil
}

func (self *SkipList) AddNode(key NodeData, val int) bool {
	if key == nil {
		return false
	}

	// 记录val在跳表每层的路径
	path := make([]*skipListNode, MAX_LEVEL, MAX_LEVEL)
	cur := self.head
	for i := MAX_LEVEL - 1; i >= 0; i-- { // 逐层查找
		for cur.next[i] != nil { // 1.下一级节点不到末尾
			if cur.next[i].val < val {
				cur = cur.next[i]
			} else if cur.next[i].val == val && cur.next[i].key == key {
				return true
			} else { // >
				path[i] = cur
				break
			}
		}

		//2.下一级节点已到末尾
		if cur.next[i] == nil {
			path[i] = cur
		}
	}

	// 随机一个高度，按高度构造这个节点
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if Rand()%7 == 1 {
			level++
		}
	}

	node := newNode(key, val, level)

	// 插入节点（此时，node是一个level层高的列表）
	for i := 0; i < level; i++ {
		node.next[i] = path[i].next[i]
		path[i].next[i] = node
	}

	if self.level < level {
		self.level = level
	}

	self.length += 1

	return true
}

func (self *SkipList) Print() {
	if self.length <= 0 {
		return
	}

	for i := self.level - 1; i >= 0; i-- {
		cur := self.head.next[i]
		for cur != nil {
			fmt.Printf("[%v, %v]	", cur.key, cur.val)
			cur = cur.next[i]
		}
		if cur == nil {
			fmt.Println("nil")
		}
		if cur != nil {
			fmt.Println("")
		}
	}
	fmt.Printf("length:%d, level:%d\n", self.length, self.level)
}

func (self *SkipList) DelNode(key NodeData, val int) bool {
	if key == nil {
		return true
	}

	//  val 在每层的路径
	path := make([]*skipListNode, MAX_LEVEL, MAX_LEVEL)
	// 从上层到下层遍历
	cur := self.head
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		for cur.next[i] != nil {
			if cur.next[i].val > val {
				break
			} else if cur.next[i].val == val && cur.next[i].key == key {
				path[i] = cur
				break
			} else {
				cur = cur.next[i]
			}
		}
	}

	lv := 1
	cur = path[0].next[0]
	for i := 0; i < MAX_LEVEL-1; i++ {
		if path[i] == nil {
			continue
		}

		path[i].next[i] = path[i].next[i].next[i]

		if path[i].next[i] != nil && i != 0 {
			lv++
		}
	}

	self.level = lv
	self.length--

	return true
}
