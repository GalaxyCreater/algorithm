/*
 * 动态规划
 */
package dynamic_plan

import "fmt"

/*
desc:
	功能：给定一堆物品，放入可载重w的包中，要求重量尽可能逼近w
*/
func Dynamic01bagSimple(items []int, w int) {

	length := len(items)
	/*
		存放每阶段状态；
		一级索引：阶段；
		二级索引：当前重量；
		值：是否存在该方案
		把每次决定一个物品是否放入作为求最优解的一个阶段，一共有物品数量个阶段
	*/
	state := make([][]bool, length)
	for i := 0; i < length; i++ {
		state[i] = make([]bool, w+1) // 因为表示重量，所以要+1
	}

	/*
		为了方便索引使用，先处理第一个物品
	*/
	// 第一个物品不放
	state[0][0] = true
	// 第一个物品放入
	if items[0] <= w {
		state[0][items[0]] = true
	}

	for i := 1; i < length; i++ { // 从第二个物品开始
		for j := 0; j <= w; j++ {
			// items[i]不放入
			if state[i-1][j] == true { // 直接沿用上一阶段的值
				state[i][j] = true
			}

			// items[i]放入
			if state[i-1][j] == true {
				nextWeight := items[i] + j
				if nextWeight <= w { // 新加入的，重量要符合要求
					state[i][nextWeight] = true
				}
			}
		}
	}

	// 输出结果
	fmt.Print("	")
	for i := 0; i <= w; i++ {
		fmt.Print(i, "	")
	}
	fmt.Println("")
	for i := 0; i < length; i++ {
		for j := 0; j <= w; j++ {
			if j == 0 {
				fmt.Print(i, "	")
			}
			if state[i][j] == true {
				fmt.Print("O	")
			} else {
				fmt.Print("*	")
			}
		}
		fmt.Println("")
	}
}

/*
desc:
	功能：Dynamic01bagSimple的省内存版本
*/
func Dynamic01bagSimpleEx(items []int, w int) {

	length := len(items)
	/*
		存放每阶段状态；
		索引：当前重量；
		值：是否存在该方案
		把每次决定一个物品是否放入作为求最优解的一个阶段，一共有物品数量个阶段
	*/
	state := make([]bool, w+1)

	/*
		为了方便索引使用，先处理第一个物品
	*/
	// 第一个物品不放
	state[0] = true
	// 第一个物品放入
	if items[0] <= w {
		state[items[0]] = true
	}

	for i := 1; i < length; i++ { // 从第二个物品开始
		for j := w; j >= 0; j-- { // 必须由后往前，否则出现重复放入问题。如j=0放入设置了state[0]为true；后面j循环到j=2时，拿到这次询环的结果，而并非是上次的结果
			// items[i]不放入
			if state[j] == true { // 直接沿用上一阶段的值
				nextWeight := items[i] + j
				if nextWeight <= w { // 新加入的，重量要符合要求
					state[nextWeight] = true
				}
			}
		}
	}

	for i := w; i > 0; i-- {
		if state[i] {
			fmt.Println(i)
			break
		}
	}
}

/*
////////////////////////////// 01 背包 升级版 //////////////////////////////////////
问题：可装w重量的背包，每个物品都有重量和价格属性，求背包中可装的最大价格
*/
type Item struct {
	w     int
	price int
}

/*
@w:背包限制的重量
*/
func Dynamic01bagPrice(w int) {
	items := []Item{
		Item{15, 10}, Item{12, 8}, Item{4, 6},
	}
	length := len(items)

	/*
		存放每阶段状态；
		一级索引：阶段；
		二级索引：当前重量；
		值：-1:没有该方案；0：不放；>0:当前状态的最大价格
		把每次决定一个物品是否放入作为求最优解的一个阶段，一共有物品数量个阶段
	*/
	states := make([][]int, length)
	for i := 0; i < length; i++ {
		l := make([]int, w+1)
		for j := 0; j < len(l); j++ {
			l[j] = -1
		}
		states[i] = l
	}

	// 先处理第一个
	states[0][0] = 0 //不放
	if items[0].w <= w {
		states[0][items[0].w] = items[0].price //放
	}

	for i := 1; i < length; i++ { // 一个个阶段演进
		for j := 0; j <= w; j++ { // 一个个可能状态

			if states[i-1][j] >= 0 {
				// 第i个物品不放
				states[i][j] = states[i-1][j]

				// 第i个物品放
				nextWeight := items[i].w + j // 利用当前有效的重量，计算出可能的新重量
				nextPrice := states[i-1][j] + items[i].price
				if nextWeight <= w && states[i][nextWeight] < nextPrice {
					states[i][nextWeight] = nextPrice
				}
			}
		}
	}

	maxVal := 0
	for j := 0; j < w+1; j++ {
		// 最优解在最后的阶段
		if states[length-1][j] > maxVal {
			maxVal = states[length-1][j]
		}
	}
	fmt.Println("max price:", maxVal)
}

/*
desc:Dynamic01bagPrice的省内存版本
@w:背包限制的重量
*/
func Dynamic01bagPriceEx(w int) {
	items := []Item{
		Item{15, 10}, Item{12, 8}, Item{4, 6},
	}
	length := len(items)

	/*
		desc:存放当前阶段状态；把每次决定一个物品是否放入作为一个阶段，一共有物品数量个阶段，存放当前
		阶段的最优解
		索引：重量
		值：-1:没有改方案；0：不放；>0:当前状态的最大价格

	*/
	states := make([]int, w+1)
	for i := 0; i < len(states); i++ {
		states[i] = -1
	}

	// 先处理第一个
	states[items[0].w] = 0 // 不放
	if items[0].w <= w {
		states[items[0].w] = items[0].price //放
	}

	for i := 1; i < length; i++ {
		for j := w; j >= 0; j-- { // 必须从后往前
			if states[j] >= 0 {
				// 放
				nextWeight := j + items[i].w
				nextPrice := states[j] + items[i].price // 利用当前有效的重量，计算出可能的新重量
				if nextWeight <= w && nextPrice > states[nextWeight] {
					states[nextWeight] = nextPrice
				}
			}
		}
	}

	maxVal := 0
	for j := 0; j < w+1; j++ {
		if states[j] > maxVal {
			maxVal = states[j]
		}
	}
	fmt.Println("max price:", maxVal)

}
