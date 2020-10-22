/*
01背包问题：一堆不可分割的物品里找出背包最多可以装的重量
*/
package recall

import "fmt"

var (
	g_bagCapacity int = 100 //背包容量
	g_resWeight   int = 0   // 保存结果

	call_time int = 0

	hasCalc map[string]bool = map[string]bool{}
)

/*
desc:
@i：表示考察到哪个物品了
@curWeight：当前已经装进去的物品的重量和
@items：表示每个物品的重量
*/
func Bag01(curIdx int, curWeight int, items []int) {
	call_time++
	cnt := len(items) //物品个数

	if cnt == curIdx || curWeight == g_bagCapacity { //curIdx==n表示已经考察完所有的物品 或 装满了
		if curWeight > g_resWeight {
			g_resWeight = curWeight
		}
		return
	}

	/* 对递归算法的优化：已计算的结果，就不需要重复计算;
	 * 优化后，时间复杂度和动态规划差不多了
	 * 又叫备忘录方法
	 */
	k := fmt.Sprintf("%d_%d", curIdx, curWeight)
	if hasCalc[k] {
		return
	} else {
		hasCalc[k] = true
	}

	/*
		调用了两次Bag01是因为每次选择物品，都有装和不装选择，
		第一个Bag01表示不装，第二个Bag02表示装这种情况，两次的Bag01把这两种情况都列出来同时执行了，这样能穷举所有可能情况
	*/
	next := curIdx + 1
	Bag01(next, curWeight, items) // 表示不选择当前物品，直接考虑下一个

	nextWeight := curWeight + items[curIdx]
	if nextWeight <= g_bagCapacity { //还没装满背包
		Bag01(next, nextWeight, items) // 表示选择当前物品，考虑下一个时
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

type ValueBag struct {
	lst         []Item
	length      int
	bagCapacity int
	max_price   int

	call_time int
}

func (self *ValueBag) Init() {
	// self.lst = []Item{
	// 	Item{5, 10}, Item{200, 8}, Item{4, 12},
	// }

	self.lst = []Item{
		Item{15, 10}, Item{12, 8}, Item{4, 6},
	}

	self.bagCapacity = 30
	self.length = len(self.lst)
}

func (self *ValueBag) Calc(curIdx, curWeight, curPrice int) {
	if self.length == curIdx || curWeight >= self.bagCapacity { //curIdx==n表示已经考察完所有的物品 或 装满了
		if curPrice > self.max_price {
			self.max_price = curPrice
		}
		return
	}

	/* 对递归算法的优化：已计算的结果，就不需要重复计算;
	 * 优化后，时间复杂度和动态规划差不多了
	 * 又叫备忘录方法
	 */
	k := fmt.Sprintf("%d_%d_%d", curIdx, curWeight, curPrice)
	if hasCalc[k] {
		return
	} else {
		hasCalc[k] = true
	}

	self.call_time++

	/*
		调用了两次Calc是因为每次选择物品，都有装和不装选择，
		第一个Calc表示不装，第二个Calc表示装这种情况，两次的Calc把这两种情况都列出来同时执行了，这样能穷举所有可能情况
	*/
	next := curIdx + 1
	self.Calc(next, curWeight, curPrice) // 1.不选择当前物品，直接考虑下一个

	// 2.选择当前物品，考虑下一个时
	nextWeight := curWeight + self.lst[curIdx].w
	nextPrice := curPrice + self.lst[curIdx].price
	if nextWeight <= self.bagCapacity { //还没装满背包
		self.Calc(next, nextWeight, nextPrice) // 表示
	}
}

func (self *ValueBag) Print() {
	fmt.Println("price:", self.max_price, "call time:", self.call_time)
}
