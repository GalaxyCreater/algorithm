package dynamic_plan

import "testing"

func TestBuyMinPrice(t *testing.T) {
	items := []int{3, 4, 5, 11, 4}
	BuyMinPrice(items, 10, 20)
	BuyMinPriceOther(items, 10, 20)
}
