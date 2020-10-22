package array

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestArry(t *testing.T) {
	a := [][]int{{1}}
	fmt.Println(a)
}

func TestLoop(t *testing.T) {
	i := 0
	fmt.Println(&i)
	arry := [3]int{0}
	for ; i < 3; i++ {
		fmt.Println("hello ", &arry[i])
	}

	fmt.Println(&i)
}

func Test1(t *testing.T) {

	rd := bufio.NewReader(os.Stdin)
	rd.

}
