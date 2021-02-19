package sort

func MyQuickSort(arr []int) {

	myQuickSort(arr, 0, len(arr)-1)
}

func myQuickSort(arr []int, beg int, end int) {
	if beg >= end {
		return
	}
	part := MyParttion(arr, end)
	myQuickSort(arr, beg, part-1)
	myQuickSort(arr, part+1, end)
}

func MyParttion(arr []int, part int) int {
	i := 0 // 第一个需要交换的位置
	j := 0
	tmp := arr[part]
	for {
		if arr[j] < tmp { // 找小的放到前面
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j++
		} else {
			j++
		}
		if j == part {
			arr[i], arr[part] = arr[part], arr[i]
			break
		}
	}

	return i
}
