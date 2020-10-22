package heap

import (
	"fmt"
	"math"
)

/*
父节点：i/2
左子节点：i*2
右子节点：i*2+1
*/

var (
	MAX_HEAP int8 = 0 // 大顶堆
	MIN_HEAP int8 = 1 //小顶堆
)

type HeapNode struct {
	v interface{}
}

func (self *HeapNode) Value() interface{} {
	return self.v
}

type Heap struct {
	ty   int8 //0:大顶堆	1:小顶堆
	arry []*HeapNode
	size int
}

/*
@ty:堆类型
*/
func NewHeap(ty int8) *Heap {
	return &Heap{
		ty:   ty,
		arry: make([]*HeapNode, 1, 1), // 0索引位置保留
		size: 0,
	}
}

/*
用数组创建堆
*/
func NewHeapArry(ty int8, lst []int) *Heap {
	h := &Heap{
		ty:   ty,
		size: len(lst),
	}
	h.arry = make([]*HeapNode, 1, 1)
	for i := 0; i < len(lst); i++ {
		n := &HeapNode{
			v: lst[i],
		}
		h.arry = append(h.arry, n)
	}

	// 向下堆化
	start := h.size / 2
	for i := start; i > 0; i-- {
		h.HeapDown(i, h.size)
	}

	return h
}

func (self *Heap) Size() int {
	return self.size
}

func (self *Heap) Top() *HeapNode {
	return self.arry[1]
}

func (self *Heap) Print() {
	lv := 1
	var tmp float64 = 1
	for i := 1; i <= self.size; i++ {
		fmt.Printf("%v	", self.arry[i].v)
		if float64(i) == tmp {
			fmt.Println()
			tmp += math.Pow(float64(2), float64(lv))
			lv++
		}
	}

	fmt.Printf("\n-------------------------------------\n")
}

/*
符合当前堆规则返回true
*/
func (self *Heap) compare(fatherVal interface{}, childVal interface{}) bool {
	if self.ty == MAX_HEAP {
		return fatherVal.(int) >= childVal.(int)
	} else {
		return fatherVal.(int) <= childVal.(int)
	}
}

func (self *Heap) exchange(lhs int, rhs int) {
	self.arry[lhs], self.arry[rhs] = self.arry[rhs], self.arry[lhs]
}

func (self *Heap) LeftChild(idx int) (p *HeapNode, tar int) {
	tar = 0
	p = nil
	if self.size == 0 {
		return
	}
	tar = idx * 2
	if tar > self.size {
		return
	}
	p = self.arry[tar]
	return
}

func (self *Heap) RightChild(idx int) (p *HeapNode, tar int) {
	p = nil
	tar = 0
	if self.size == 0 {
		return
	}
	tar = idx*2 + 1
	if tar > self.size {
		return
	}
	p = self.arry[tar]
	return
}

func (self *Heap) Father(idx int) *HeapNode {
	if self.size == 0 {
		return nil
	}
	tar := idx / 2
	if tar == 0 {
		return nil
	}
	return self.arry[tar]
}

/*
从下向上堆化
*/
func (self *Heap) HeapUp(idx int) {
	curIdx := idx
	for curIdx > 0 {
		father := self.Father(curIdx)
		if father == nil {
			break
		}
		if self.compare(father.v, self.arry[curIdx].v) {
			break
		} else {
			fIdx := curIdx / 2
			self.exchange(fIdx, curIdx)
		}

		curIdx = curIdx / 2
	}
}

/*
从上往下堆化
*/
func (self *Heap) HeapDown(beg_idx int, end_idx int) {
	curIdx := beg_idx
	for curIdx < end_idx {
		// 找出三者中的极值
		peak := self.arry[curIdx]
		pIdx := curIdx
		left, lIdx := self.LeftChild(curIdx)
		if left == nil || lIdx > end_idx {
			break
		}
		if !self.compare(peak.v, left.v) {
			peak = left
			pIdx = lIdx
		}

		right, rIdx := self.RightChild(curIdx)
		if right != nil &&
			rIdx <= end_idx &&
			!self.compare(peak.v, right.v) {
			pIdx = rIdx
		}

		if pIdx == curIdx { // 左右节点都符合堆规则
			break
		} else { // 交换极值
			self.exchange(pIdx, curIdx)
		}
		curIdx = pIdx
	}
}

func (self *Heap) Add(v interface{}) {
	if v == nil {
		return
	}
	node := &HeapNode{
		v: v,
	}

	// 先放入数组最后位置
	self.arry = append(self.arry, node)
	self.size++

	if self.size == 1 {
		return
	}

	// 从下往上堆化
	self.HeapUp(self.size)
}

/*
弹出堆顶元素
*/
func (self *Heap) Pop() (res interface{}) {
	res = nil
	if self.size <= 0 {
		return
	}

	res = self.arry[1].v

	if self.size == 1 {
		self.arry = self.arry[:1]
		self.size--
		return
	}

	// 首尾交换
	self.exchange(1, self.size)
	// 移除
	self.arry = self.arry[:self.size]
	self.size--

	// 向下堆化
	self.HeapDown(1, self.size)

	return
}

func (self *Heap) Find(v interface{}) (res *HeapNode, idx int) {
	res = nil
	idx = 0
	for i := 1; i <= self.size; i++ {
		if self.arry[i].v == v {
			idx = i
			res = self.arry[i]
			return
		}
	}

	return
}

/*
删除某值
*/
func (self *Heap) Delete(v interface{}) (res *HeapNode) {
	res, idx := self.Find(v)
	if res == nil {
		return
	}

	// 首尾交换
	self.exchange(idx, self.size)
	// 移除
	self.arry = append(self.arry, self.arry[:self.size-1]...)
	self.size--

	// 向下堆化
	self.HeapDown(1, self.size)

	return
}

/*
堆排序
*/
func (self *Heap) Sort() {
	lst := []int{}
	for i := self.size; i > 1; {
		self.exchange(i, 1)
		lst = append(lst, self.arry[i].v.(int))
		i--
		self.HeapDown(1, i)
	}

	fmt.Println(lst)
}

func HeapSortMin(lst []int) {
	h := NewHeapArry(MAX_HEAP, lst)
	h.Print()

	h.Sort()
	h.Print()
}
