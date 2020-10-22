package tree

/*
用递归树分析该算法复杂度
*/

import "fmt"

func PermutationAll(lst []int) {
	permutationAll(lst, len(lst)-1)
}

/*
desc : lst里所有元素的全排列
流程：先排列idx-1位置，再排列下一位，直到全部排列完。从尾部开始方便写代码。
@idx : 要排列的位置
*/
func permutationAll(lst []int, idx int) {
	// 只剩一个元素，打印数组
	if idx == 0 {
		for i := 0; i < len(lst); i++ {
			fmt.Print(lst[i])
		}
		fmt.Println("")
	}

	for i := 0; i <= idx; i++ {
		lst[idx], lst[i] = lst[i], lst[idx]

		permutationAll(lst, idx-1)

		// 把交换了的值还原回去
		lst[idx], lst[i] = lst[i], lst[idx]
	}
}
