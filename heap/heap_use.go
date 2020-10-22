package heap

import "math"

/*
找出Top K
desc:因为数组不是已排序的，只满足二叉树的大小关系，最大堆前k个元素不一定是最大的k个元素
*/
func (self *Heap) FindTopK(lst []interface{}, k int) (resLst []*HeapNode) {
	resLst = []*HeapNode{}

	for i := 0; i < len(lst); i++ {
		if self.size < k {
			self.Add(lst[i])
			continue
		}
		// lst[i]在堆顶上面，跳过 (lst[i]要比堆顶元素大)
		if self.compare(lst[i], self.Top().v) {
			continue
		}

		self.arry[1].v = lst[i]
		// 向下堆化
		self.HeapDown(1, self.size)
	}

	resLst = append(resLst, self.arry[1:]...)

	return
}

/*
在动态数据中求在某个百分比的数
desc：如1，2，3，4。。。100 , 70在70%， 90在90%，中位数在50%
*/
func GetTopPercent(smallHeap *Heap, // 保存最大那部分元素的小顶堆
	bigHeap *Heap, // 保存最小那部分元素的大顶堆
	lst []int, percent float64) interface{} {

	for i := 0; i < len(lst); i++ {
		// 判断该放到哪个堆
		if smallHeap.size == 0 || lst[i] >= smallHeap.Top().Value().(int) {
			smallHeap.Add(lst[i]) // 大的放到小顶堆
		} else if bigHeap.size == 0 || lst[i] <= bigHeap.Top().Value().(int) {
			bigHeap.Add(lst[i]) // 小的放到大顶堆
		}

		total := float64(smallHeap.size + bigHeap.size)
		mid := math.Ceil(total * percent)
		if float64(bigHeap.size) < mid { // 数量超了就给对方
			// 调整大小,将小顶堆的堆顶给大顶堆
			h := smallHeap.Pop()
			bigHeap.Add(h)
		} else if float64(smallHeap.size) < (total - mid) { // 数量超了就给对方
			h := bigHeap.Pop()
			smallHeap.Add(h)
		}
	}

	h := bigHeap.Pop()
	if h != nil {
		return h
	}
	return nil
}
