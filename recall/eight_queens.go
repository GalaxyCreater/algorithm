/*
desc: 回溯法解决八皇后问题（放入的棋子八个方向不能有其他棋子）
g_size * g_size棋盘
*/

package recall

import "fmt"

var g_size int = 8

// 存储queen的位置，索引为行，值为列，这种方式可以让这个g_size * g_size棋盘不用重置值
var g_result []int = make([]int, g_size, g_size)

var g_total int = 0

/*
@row:当前行数
*/
func Cal8Queens(row int) { // 调用：Cal8Queens(0)
	if row == 8 { // 已到棋盘结尾
		Print8Queens()
		return
	}

	for col := 0; col < g_size; col++ {
		if !IsValidQueen(row, col) { // 该位置要满足条件
			continue
		}
		g_result[row] = col //记录queen的位置
		Cal8Queens(row + 1) // 移到下一行
	}

}

func Print8Queens() {
	for row := 0; row < g_size; row++ {
		for col := 0; col < g_size; col++ {
			if g_result[row] == col {
				fmt.Printf("q	")
				//g_result[row] = -1
			} else {
				fmt.Printf("*	")
			}
		}
		fmt.Println("")
	}
	fmt.Printf("\n\n\n")
	g_total++
}

/*
因为是一步步往下的，所以只需要判断左右上角是否合法
*/
func IsValidQueen(row int, col int) bool {
	left := col - 1
	right := col + 1
	for r := row - 1; r >= 0; r-- { // 逐行往上考察每一行
		if g_result[r] == col { // 上方
			return false
		}

		if left >= 0 { //左上角
			if g_result[r] == left {
				return false
			}
		}

		if right < g_size { // 右上角
			if g_result[r] == right {
				return false
			}
		}

		left--
		right++
	}

	return true
}
