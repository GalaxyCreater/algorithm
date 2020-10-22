package sort

import "math"

/*
** 冒泡排序
 */
func BubbleSort(lst []int) {
	if len(lst) < 2 {
		return
	}

	for i := 0; i < len(lst); i++ {
		flag := false
		for j := 0; j < len(lst)-i-1; j++ {
			if lst[j] > lst[j+1] {
				lst[j+1], lst[j] = lst[j], lst[j+1]
				flag = true
			}
		}

		if flag == false {
			break
		}
	}
}

/*
** 插入排序
 */
func InsertionSort(lst []int) {
	if len(lst) <= 1 {
		return
	}

	for i := 1; i < len(lst); i++ {
		val := lst[i]
		j := i - 1
		for ; j >= 0; j-- {
			if val < lst[j] {
				lst[j+1] = lst[j]
			} else {
				break
			}
		}
		lst[j+1] = val // j+1 != i,在for里面多j--,所以这要加1
	}
}

/*
** 归并排序
 */
func MergeSort(lst []int) {
	mergeSort(lst, 0, len(lst)-1)
}

func mergeSort(lst []int, beg, end int) {
	if beg >= end {
		return
	}

	mid := (beg + end) / 2

	mergeSort(lst, beg, mid)
	mergeSort(lst, mid+1, end)

	merge(lst, beg, mid, end)
}

func merge(lst []int, beg, mid, end int) {
	if beg >= end {
		return
	}
	tmp := make([]int, end-beg+1)

	idx := 0
	lftIdx := beg
	rigIdx := mid + 1
	for ; lftIdx <= mid && rigIdx <= end; idx++ { // 因为传参时，已经传len(lst)-1了，不会越界
		if lst[lftIdx] < lst[rigIdx] {
			tmp[idx] = lst[lftIdx]
			lftIdx++
		} else {
			tmp[idx] = lst[rigIdx]
			rigIdx++
		}

	}

	// 把剩余有序数组添加到tmp
	for ; lftIdx <= mid; lftIdx++ {
		tmp[idx] = lst[lftIdx]
		idx++
	}
	for ; rigIdx <= end; rigIdx++ {
		tmp[idx] = lst[rigIdx]
		idx++
	}

	copy(lst[beg:end+1], tmp)
}

/*
** 快速排序
 */
func QuickSort(lst []int) {
	quickSort(lst, 0, len(lst)-1)
}

func quickSort(lst []int, beg, end int) {
	if beg >= end {
		return
	}

	part := partition(lst, beg, end)
	quickSort(lst, beg, part-1) // 中间节点左边一个
	quickSort(lst, part+1, end) // 中间节点右边一个
}

func partition(lst []int, beg, end int) (part int) {

	part = lst[end]
	i := beg
	j := beg
	for ; j < end; j++ {
		if lst[j] < part {
			lst[i], lst[j] = lst[j], lst[i]
			i++
		}
	}

	lst[i], lst[end] = lst[end], lst[i]

	return i
}

/*
** 利用快排思想，找出第k大的数, 复杂度O(n)
 */
func FindTopK(lst []int, k int) (ret int) {
	ret = 0
	length := len(lst)
	if k > length {
		return
	}
	tar := length - k
	idx := findTopK(lst, 0, length-1, tar)
	return lst[idx]
}
func findTopK(lst []int, beg, end, tar int) int {
	// 目标在数组中的索引位置
	p := partition(lst, beg, end)

	if p == tar {
		return tar
	}

	if p > tar {
		end = p - 1
	} else {
		beg = p + 1
	}

	return findTopK(lst, beg, end, tar)
}

// 非递归方式实现
func findTopKLoop(lst []int, k int) int {

	length := len(lst)
	if k > length {
		return 0
	}
	tar := length - k
	beg := 0
	end := length - 1

	for {
		// 目标在数组中的索引位置
		p := partition(lst, beg, end)

		if p == tar {
			return lst[tar]
		}

		if p > tar {
			end = p - 1
		} else {
			beg = p + 1
		}
	}
}

/*
** 桶排序
 */
func BucketSort(lst []int) {
	length := len(lst)
	if length <= 1 {
		return
	}

	// 获取lst中最大元素
	max := lst[0]
	for i := 1; i < length; i++ {
		if max < lst[i] {
			max = lst[i]
		}
	}

	// 按lst元素个数创建桶列表
	bucketList := make([][]int, length, length)

	// 将所有元素放到桶中
	for i := 0; i < length; i++ {
		buckIdx := lst[i] * (length - 1) / max // 实际上是 lst[i] / max * (length-1);这里表示占索引的百分比,所以要减1
		bucketList[buckIdx] = append(bucketList[buckIdx], lst[i])
	}

	// 对桶列表的所有桶排序
	lstIdx := 0
	for i := 0; i < len(bucketList); i++ {
		buckLen := len(bucketList[i])

		if buckLen == 0 {
			continue
		}
		MergeSort(bucketList[i]) // 不适用quicksort,用稳定排序

		// 将排序好的桶放到lst中
		copy(lst[lstIdx:lstIdx+buckLen], bucketList[i][:])
		lstIdx += buckLen
	}
}

/*
** 计数排序
 */
func CountingSort(lst []int) {
	length := len(lst)
	if length <= 1 {
		return
	}
	max := lst[0]
	for i := 1; i < length; i++ {
		if max < lst[i] {
			max = lst[i]
		}
	}
	// 统计数量
	cntLst := make([]int, max+1) // 用值来做索引数组，所以要加1
	for i := 0; i < length; i++ {
		cntLst[lst[i]]++
	}
	// 记录小于等于某值的数量, 实际上是他们的排名
	tmp := 0
	for i := 0; i < len(cntLst); /*max*/ i++ {
		tmp += cntLst[i]
		cntLst[i] = tmp
	}

	// 将原列表按cntLst顺序放入新列表
	resLst := make([]int, length)
	for _, v := range lst {
		idx := cntLst[v]
		resLst[idx-1] = v
		cntLst[v]--
	}
	copy(lst, resLst)
}

/*
**基数排序
 */
func RadixSort(lst []int) {
	bitLen := getMaxBitLen(lst) //要排序的数位数，如103，位数位3
	// 将lst所有元素由低位到高位比较，依次放入arr
	for i := 1; i <= bitLen; i++ {
		insertionSortByBit(lst, i, compareBit)
	}
}

func getMaxBitLen(lst []int) int {
	getBitlen := func(tmp int) int {
		b := 1
		for {
			tmp = tmp / 10
			if tmp <= 0 {
				break
			} else {
				b++
			}
		}
		return b
	}

	max := 1
	for i := 0; i < len(lst); i++ {
		b := getBitlen(lst[i])
		if b > max {
			max = b
		}
	}

	return max
}

func compareBit(a, b, bit int) bool {
	div := math.Pow(10, float64(bit))
	if float64(a)/div > float64(b)/div {
		return false
	}
	return true
}

func insertionSortByBit(lst []int, bit int, compareFunc func(a, b, bit int) bool) {
	if len(lst) <= 1 {
		return
	}

	for i := 1; i < len(lst); i++ {
		val := lst[i]
		j := i - 1
		for ; j >= 0; j-- {
			if compareFunc(val, lst[j], bit) {
				lst[j+1] = lst[j]
			} else {
				break
			}
		}
		lst[j+1] = val // j+1 != i,在for里面多j--,所以这要加1
	}
}
