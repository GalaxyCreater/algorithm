/*
回溯法实现正则表达式的 * 和 ？(0或1)
*/

package recall

type SimpleRegular struct {
	exp     []rune // 处理中文用rune
	exp_len int
	find    bool
}

func (self *SimpleRegular) Init(exp string) {
	self.exp = []rune(exp)
	self.exp_len = len(self.exp)
	self.find = false

}

func (self *SimpleRegular) Find(tar string) bool {
	t := []rune(tar)
	tar_len := len(t)
	self._find(t, tar_len, 0, 0)
	return self.find
}

func (self *SimpleRegular) _find(tar []rune, tar_len, exp_idx, tar_idx int) {
	if self.find == true {
		return
	}
	if exp_idx == self.exp_len { // 遍历完正则表达式
		if tar_idx >= tar_len-1 { // 遍历完原字符串
			self.find = true
		}
		return
	}

	if string(self.exp[exp_idx]) == "?" {
		self._find(tar, tar_len, exp_idx+1, tar_idx)   // 没字符匹配情况
		self._find(tar, tar_len, exp_idx+1, tar_idx+1) // 匹配到一个字符情况
	} else if string(self.exp[exp_idx]) == "*" {
		for i := tar_idx; i < tar_len; i++ { // 列出后面匹配0个到tar_len-tar_idx个的所有情况
			self._find(tar, tar_len, exp_idx+1, i) // 用表达式的下一个值对比
		}
	} else if exp_idx < self.exp_len && self.exp[exp_idx] == tar[tar_idx] {
		self._find(tar, tar_len, exp_idx+1, tar_idx+1)
	}
}
