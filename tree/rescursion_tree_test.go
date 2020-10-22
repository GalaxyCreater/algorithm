package tree_test

import (
	"fmt"
	"testing"
)

/*
递归树章节
*/

/*
1 个细胞的生命周期是 3 小时，1 小时分裂一次，求 n 小时后，容器内有多少细胞？
@t:t小时后
*/
func CellSplit(num, t int) int {
	if t == 0 {
		return num
	}

	// 第t个小时的数量 - t的前3小时的数量
	return AddCell(num, t) - AddCell(num, t-3)
}

func AddCell(num, t int) int {
	if t < 0 {
		return 0
	}
	if t == 0 {
		return num
	}

	return AddCell(num, t-1) * 2
}

func TestCellSplit(t *testing.T) {
	fmt.Println(CellSplit(1, 2))
	fmt.Println(CellSplit(1, 3))
	fmt.Println(CellSplit(1, 4))
	fmt.Println(CellSplit(1, 24))
}
