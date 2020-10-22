package dynamic_plan

import "fmt"

/*
有一堆商品在购物车，有个满200减50的活动，尽挑选总价最低的商品组合，极限“褥羊毛”,
得出方案
*/

/*
@items:
	索引：	商品索引
	值：	商品价格
@min:刚好满足满减条件
@max:总价上限（可以是items价格总和，也可以设定一个值）
*/
func BuyMinPrice(items []int, min int, max int) {
	length := len(items)
	/*
		存放每阶段状态；
		一级索引：阶段；
		二级索引：当前价格
		值：是否存在该方案
	*/
	states := make([][]bool, length)
	for i := 0; i < length; i++ {
		l := make([]bool, max+1)
		for j := 0; j < len(l); j++ {
			l[j] = false
		}
		states[i] = l // 上限值+1
	}

	minPrice := 0

	// 先处理第一个物品
	states[0][0] = true //不买
	if items[0] <= max {
		states[0][items[0]] = true //买
		minPrice = items[0]
	}

	for i := 1; i < length; i++ {
		for j := 0; j < max+1; j++ {
			if states[i-1][j] == true {
				// 不买
				states[i][j] = states[i-1][j]

				// 买
				nextPrice := j + items[i]
				if nextPrice <= max { // 找出符合所有条件的方案
					if minPrice < min { // 还没买够，放入所有符合上限的方案
						states[i][nextPrice] = true
						minPrice = nextPrice
					} else { // 买够，只选择最少值方案
						if nextPrice >= min && nextPrice < minPrice {
							states[i][nextPrice] = true
							minPrice = nextPrice
						}
					}
				}

			}
		}
	}

	if minPrice > max {
		fmt.Println("more money")
		return
	}
	if minPrice < min {
		fmt.Println("not enough")
		return
	}

	fmt.Println(minPrice, states)

	// 求方案
	lst := []int{}
	pre := minPrice
	for i := length - 1; i >= 0; i-- {
		tmp := pre - items[i]
		if tmp >= 0 && states[i][tmp] == true {
			pre = tmp
			lst = append(lst, items[i])
		} // 不满组if条件的，表示当前这个物品不被选中
	}
	fmt.Println(lst)
}

/*
课程中的实现方式
*/
func BuyMinPriceOther(items []int, min int, max int) {
	length := len(items)
	/*
		存放每阶段状态；
		一级索引：阶段；
		二级索引：当前价格
		值：是否存在该方案
	*/
	states := make([][]bool, length)
	for i := 0; i < length; i++ {
		l := make([]bool, max+1)
		for j := 0; j < len(l); j++ {
			l[j] = false
		}
		states[i] = l // 上限值+1
	}

	// 先处理第一个物品
	states[0][0] = true //不买
	if items[0] <= max {
		states[0][items[0]] = true //买
	}

	for i := 1; i < length; i++ {
		for j := 0; j < max+1; j++ {
			if states[i-1][j] == true {
				// 不买
				states[i][j] = states[i-1][j]

				// 买
				nextPrice := j + items[i]
				if nextPrice <= max { // 先找出符合所有上限条件的方案
					states[i][nextPrice] = true
				}
			}
		}
	}

	j := min // 保证满足最低条件
	for ; j <= max; j++ {
		if states[length-1][j] == true {
			break
		}
	}
	if j > max {
		fmt.Println("not find")
		return
	}
	fmt.Println("min: ", j, states)

	// 求方案
	lst := []int{}
	pre := j
	for i := length - 1; i >= 0; i-- {
		tmp := pre - items[i]
		if tmp >= 0 && states[i][tmp] == true {
			pre = tmp
			lst = append(lst, items[i])
		} // 不满组if条件的，表示当前这个物品不被选中
	}
	fmt.Println(lst)
}
