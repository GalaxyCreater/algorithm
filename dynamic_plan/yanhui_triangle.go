package dynamic_plan

import "fmt"

/*
在下面的杨辉三角中，节点的值可以任意填写，从顶点到最下层所经过的节点值的和为路径长度，求最短路径

	      o
	o			o
o		o	 o		o
*/

/*
完全二叉树可以用数组表示


func TriangleShortestpath(triangle []int) {
	length := len(triangle)
	//height := int(match.Floor(match.Log2(float64(length))) + 1)



	states := make([]map[int]int, length, length)
	for i := 0; i < length; i++ {
		states[i] = map[int]bool{}
	}

	// 顶点必须经过
	states[0][triangle[0]] = true

	for i := 1; i < length; i++ {
		father := i / 2

		// 跳过
		states[i] = states[father]
		// 经过
		nextVal := states[i-1][father] + triangle[i]

		if _, ok := states[father][]; ok { //父节点必须被选中
			// 跳过
			states[i][i] = states[i-1][father]

			// 经过
			nextVal := states[i-1][father] + triangle[i]
			states[i][i] = nextVal
		}
	}

	min := 0
	m := states[length-1]
	for _, v := range m {
		if min <= 0 || v < min {
			min = v
		}
	}

	fmt.Println("min: ", min)

	for i := 0; i < length; i++ {

	}

}
*/

func yanghuiTriangle(matrix [][]int) {
	length := len(matrix)
	// 用于存储每一层的状态
	min := make([]int, length, length)
	for i := length - 1; i >= 0; i-- {

		rawNums := matrix[i]
		rowLength := len(rawNums)
		for j := 0; j < rowLength; j++ {

			m := min[j]
			if min[j+1] < m {
				m = min[j+1]
			}

			min[j] = m + rawNums[j]
		}
	}

	fmt.Println(min[0])
}
