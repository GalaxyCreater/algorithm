package heap_test

import (
	"fmt"
	"test_code/algorithm/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	lst := []int{3, 2, 4, 2, 1, 4, 3, 2, 1}
	h := heap.NewHeap(heap.MAX_HEAP)
	for i := 0; i < len(lst); i++ {
		h.Add(lst[i])
	}
	h.Print()

	// 测试弹出
	h.Pop()
	h.Print()
	h.Pop()
	h.Print()
	// h.Pop()
	// h.Print()
	// h.Pop()
	// h.Print()
	// h.Pop()
	// h.Print()
	// h.Pop()
	// h.Print()

	// 测试删除
	h.Delete(3)
	h.Print()

	// 测试堆排序
	heap.HeapSortMin(lst)
}

func TestFindTopK(t *testing.T) {
	k := 7
	lst := []interface{}{3, 2, 4, 2, 1, 4, 3, 2, 1}
	h := heap.NewHeap(heap.MIN_HEAP)
	h.FindTopK(lst, k)

	h.Print()
}

func TestGetTopPercent(t *testing.T) {
	percent := 0.5
	//lst := []int{3, 2, 4, 2, 1, 4, 3, 2}
	lst := []int{3, 2, 4, 2, 1, 4, 3, 2, 1, 3}
	bigh := heap.NewHeap(heap.MAX_HEAP)
	smallh := heap.NewHeap(heap.MIN_HEAP)

	// v := heap.GetTopPercent(smallh, bigh, lst, percent)
	// bigh.Print()
	// smallh.Print()
	// fmt.Println(v)

	percent = 0.7
	v := heap.GetTopPercent(smallh, bigh, lst, percent)
	bigh.Print()
	smallh.Print()
	fmt.Println(v)
}
