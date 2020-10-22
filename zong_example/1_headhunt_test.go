package zong_example

import (
	"fmt"
	"testing"
)

func TestHeadHunt(t *testing.T) {
	p1 := HuntHead{id: 1}
	p2 := HuntHead{id: 1}
	fmt.Println(p1 == p2)
}
