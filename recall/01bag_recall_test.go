package recall

import (
	"fmt"
	"testing"
)

func TestBag01(t *testing.T) {
	var items []int = []int{55, 10, 2, 4, 5, 6, 44}
	Bag01(0, 0, items)
	fmt.Println("max weight:", g_resWeight)
}

func TestValueBag(t *testing.T) {
	bag := ValueBag{}
	bag.Init()
	bag.Calc(0, 0, 0)
	bag.Print()
}
