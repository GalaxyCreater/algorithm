package test_test

import (
	"fmt"
	"testing"
)

/*一、
输入：字符串由数字、小写字母、大写字母组成。输出：排序好的字符串。
排序的标准：
1. 数字>小写字母>大写字母。
2. 数字、字母间的相对顺序不变。
3. 额外存储空间：O（1）。
// Example
input: "abcd4312ABDC"
output: "4312abcdABDC"
*/
func isNum(v byte) bool {
	if v <= '9' && v >= '0' {
		return true
	}

	return false
}

func isCase(v byte) bool {
	if v <= 'Z' && v >= 'A' {
		return true
	}

	return false
}

func TestS(t *testing.T) {
	s := "3k4LdeDa9bAc"
	b := []byte(s)
	charIdx := -1
	numIdx := -1
	for i := 0; i < len(s); i++ {
		if isNum(b[i]) {
			if charIdx >= 0 {
				b[charIdx], b[i] = b[i], b[charIdx]
				numIdx = charIdx
				// 找下个需要替换的位置
				for {
					charIdx++
					if charIdx > len(s) {
						break
					}
					if !isNum(b[charIdx]) {
						break
					}
				}
			} else {
				numIdx = i
			}
		} else {
			if charIdx < 0 {
				charIdx = i
			}
		}
	}

	s = string(b)
	invalidIdx := -1
	for i := numIdx + 1; i < len(b); i++ {
		if isCase(b[i]) {
			if invalidIdx >= 0 {
				b[invalidIdx], b[i] = b[i], b[invalidIdx]
				// 找下个需要替换的位置
				for {
					invalidIdx++
					if invalidIdx >= len(b) {
						break
					}
					if !isCase(b[invalidIdx]) {
						break
					}
				}
			}
		} else {
			if invalidIdx == -1 {
				invalidIdx = i
			}
		}
	}

	fmt.Println(s)
}
