package bserach

import (
	"fmt"
	"strconv"
)

/*
** 二分查找
 */
func Bsearch(lst []int, val int) int {
	length := len(lst)
	beg := 0
	end := length - 1
	return bsearch(lst, val, beg, end)
}
func bsearch(lst []int, val, beg, end int) int {
	for beg <= end {
		//mid := (beg + end) / 2
		mid := beg + ((end - beg) >> 1)
		if lst[mid] == val {
			return mid
		}

		if lst[mid] > val {
			end = mid - 1
		} else {
			beg = mid + 1
		}
	}

	return -1
}

func BsearchRecursive(lst []int, val int) int {
	length := len(lst)
	if length == 0 {
		return -1
	}
	beg := 0
	end := length - 1
	return bsearchRecursive(lst, val, beg, end)
}
func bsearchRecursive(lst []int, val, beg, end int) int {
	if beg > end {
		return -1
	}

	mid := beg + ((end - beg) >> 1)
	if lst[mid] == val {
		return mid
	} else if lst[mid] > val {
		end = mid - 1
	} else {
		beg = mid + 1
	}

	return bsearch(lst, val, beg, end)
}

/*
** 计算某值的平方,精确到小数点后6位
 */
func CalcSqrt(v float64) float64 {
	var rate float64 = 0.000001
	var diff float64 = 0.0000001 // 这里要比精确度小一位
	var end float64 = v
	var beg float64 = rate
	mid := 0.0
	for beg <= end {
		mid = (beg + end) / 2
		sq := mid * mid
		if (sq-v) > rate && (sq-v) < rate {

			s := fmt.Sprintf("%.6f", mid)
			f, _ := strconv.ParseFloat(s, 64)
			return f
		}

		if sq > v {
			end = mid - diff
		} else {
			beg = mid + diff
		}
	}

	s := fmt.Sprintf("%.6f", mid)
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

/*
**查找等于某值的第一个位置
 */
func FindFirstEqual(lst []int, val int) int {
	length := len(lst)
	beg := 0
	end := length - 1
	for beg <= end {
		//mid := (beg + end) / 2
		mid := beg + ((end - beg) >> 1)
		if lst[mid] == val {
			tmp := mid
			for {
				if tmp == 0 {
					return tmp
				}
				if lst[tmp-1] != val {
					return tmp
				}
				tmp--
			}
		}

		if lst[mid] > val {
			end = mid - 1
		} else {
			beg = mid + 1
		}
	}

	return -1
}

func FindFirstEqualEx(lst []int, val int) int {
	length := len(lst)
	beg := 0
	end := length - 1
	for beg <= end {
		//mid := (beg + end) / 2
		mid := beg + ((end - beg) >> 1)
		if lst[mid] > val {
			end = mid - 1
		} else if lst[mid] < val {
			beg = mid + 1
		} else {
			tmp := mid
			if tmp == 0 || lst[tmp-1] != val {
				return tmp
			} else {
				end = tmp - 1 // 如果相等元素多了，那这个效率更高
			}
		}
	}

	return -1
}

/*
**查找等于某值的最后一个位置
 */
func FindLastEqual(lst []int, val int) int {
	length := len(lst)
	beg := 0
	end := length - 1
	for beg <= end {
		//mid := (beg + end) / 2
		mid := beg + ((end - beg) >> 1)
		if lst[mid] > val {
			end = mid - 1
		} else if lst[mid] < val {
			beg = mid + 1
		} else {
			tmp := mid
			if tmp == length || lst[tmp+1] != val {
				return tmp
			} else {
				beg = tmp + 1
			}
		}
	}

	return -1
}

/*
**查找第一个大于等于val的位置
 */
func FindFirstGreaterEqual(lst []int, val int) int {
	length := len(lst)
	beg := 0
	end := length - 1
	for beg <= end {
		//mid := (beg + end) / 2
		mid := beg + ((end - beg) >> 1)
		if lst[mid] >= val {
			if mid == 0 || lst[mid-1] < val {
				return mid
			} else {
				end = mid - 1
			}
		} else {
			beg = mid + 1
		}
	}

	return -1
}

/*
**查找最后一个小于等于val的位置
 */
func FindLastLessEqual(lst []int, val int) int {
	length := len(lst)
	beg := 0
	end := length - 1
	for beg <= end {
		//mid := (beg + end) / 2
		mid := beg + ((end - beg) >> 1)
		if lst[mid] <= val {
			if mid == length-1 || lst[mid+1] > val {
				return mid
			} else {
				beg = mid + 1
			}
		} else {
			end = mid - 1
		}
	}

	return -1
}

/*
**有序循环数组里使用二分查找

 如果首元素小于 mid，说明前半部分是有序的，后半部分是循环有序数组；
 如果首元素大于 mid，说明后半部分是有序的，前半部分是循环有序的数组；
 如果目标元素在有序数组范围中，使用二分查找；
 如果目标元素在循环有序数组中，设定数组边界后，使用以上方法继续查找。
*/
func BinsearchCircular(lst []int, val int) int {
	length := len(lst)
	beg := 0
	end := length - 1
	for beg <= end {

		mid := beg + ((end - beg) >> 1)
		if lst[mid] == val {
			return mid
		} else if lst[mid] > lst[beg] { //前半部分有序，后半部分无序
			if val < lst[mid] && val >= lst[beg] { // 在有序的前半部分
				return bsearch(lst, val, beg, mid-1)
			} else { // 目标值再次落在无序数组中
				beg = mid + 1
			}
		} else { // 前半部分无序，后半部分有序
			if val > lst[mid] && val <= lst[end] { // 在有序的后半部分
				return bsearch(lst, val, mid+1, end)
			} else { // 目标值再次落在无序数组中
				end = mid - 1
			}
		}
	}

	return -1
}
