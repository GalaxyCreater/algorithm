package string_match

import "math"

/*
desc:坏字符规则
@idx:在原字符串中开始比较索引
@src:原字符串
@mod:模式串
@return:需要移动的位数,坏字符索引
*/
func BmBadSearch(idx int, src string, mod string, bad_array []int) (length int, badIdx int) {
	length = 0
	badIdx = -1
	for i := len(mod) - 1; i >= 0; i-- {
		if mod[i] != src[idx+i] { //出现坏字符
			findIdx := bad_array[src[idx+i]]
			length = i - findIdx
			badIdx = idx + i
			return
		}

	}

	return
}

func constructBadArray(mod string) []int {
	bad_array := make([]int, 256, 256)
	for i := 0; i < len(bad_array); i++ {
		bad_array[i] = -1
	}

	for i := 0; i < len(mod); i++ {
		bad_array[mod[i]] = i
	}

	return bad_array
}

/*
desc:好后缀规则
@idx:在原字符串中开始比较索引
@return:需要移动的位数
*/
func constructGoodArray(mod string) (suffix []int, prefix []bool) {
	length := len(mod)
	suffix = make([]int, length, length)
	prefix = make([]bool, length, length)
	for i := 0; i < length; i++ {
		suffix[i] = -1
		prefix[i] = false
	}

	// 从前往后
	for i := 0; i < length-1; i++ { // i为前缀的结束位置，后缀的结束位置;因为是前缀，所以-1
		j := i // j为匹配到的子串的起始位置
		matchLen := 0
		for j >= 0 && mod[j] == mod[length-matchLen-1] {
			matchLen++
			suffix[matchLen] = j
			j--
		}

		if j < 0 {
			prefix[matchLen] = true
		}

	}

	return
}

func BmSearch(src string, mod string) int {
	bad_array := constructBadArray(mod)
	suffix, prefix := constructGoodArray(mod)

	srcLen := len(src)
	modLen := len(mod)
	i := 0
	for i < len(src) {
		if i > srcLen-modLen {
			break
		}

		badMoveLen, badIdx := BmBadSearch(i, src, mod, bad_array)
		if badIdx == -1 { // 没有坏字符串，匹配
			return i
		}

		// 好后缀规则
		goodLen := i + modLen - badIdx - 1 // 好字符串长度
		goodMoveLen := 0
		// 有匹配的子串
		if suffix[goodLen] != -1 {
			goodMoveLen = badIdx - suffix[goodLen]
		} else {
			for gi := goodLen; gi >= 1; gi-- {
				if prefix[gi] {
					// 有匹配的前缀子串
					goodMoveLen = modLen - goodLen + 1
					break
				}
			}
		}

		mv := int(math.Max(float64(badMoveLen), float64(goodMoveLen)))

		i += mv
	}

	return -1
}
