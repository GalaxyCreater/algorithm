package dynamic_plan

import "testing"

func TestDynamic01bagSimple(t *testing.T) {
	items := []int{5, 2, 4} // 添加一个辅助值0
	weight := 10
	Dynamic01bagSimple(items, weight)
	Dynamic01bagSimpleEx(items, weight)
}

func TestDynamic01bagPrice(t *testing.T) {
	Dynamic01bagPrice(30)
	Dynamic01bagPriceEx(30)
}
