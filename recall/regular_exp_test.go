package recall

import (
	"fmt"
	"testing"
)

func TestRegular(t *testing.T) {
	reg := SimpleRegular{}
	reg.Init("你?谁")
	fmt.Println(reg.Find("你谁"))

	reg.Init("你?谁")
	fmt.Println(reg.Find("你是谁"))

	reg.Init("你*谁")
	fmt.Println(reg.Find("你abcd谁"))

	reg.Init("你*谁")
	fmt.Println(reg.Find("你谁"))

	reg.Init("你*")
	fmt.Println(reg.Find("你谁"))
}
