package test_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func RegulateString(cont string, regStr string) (ret string) {
	ret = ""
	reg := regexp.MustCompile(regStr) //start开始，end结束，中间全是数字

	/*
		在 s 中查找 re 中编译好的正则表达式，并返回所有匹配的内容
		同时返回子表达式匹配的内容
		{
		{完整匹配项, 子匹配项, 子匹配项, ...},
		{完整匹配项, 子匹配项, 子匹配项, ...},
		...
		}
		只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
	*/
	lst := reg.FindAllStringSubmatch(cont, -1)

	fmt.Println(len(lst), lst)

	return
}

func ReplaceTempByStar(src string) string {
	exp := regexp.MustCompile(`\{[a-zA-Z]+\.{0,1}[a-zA-Z]+\}`)
	corpu := exp.ReplaceAllString(src, "${n}*")
	// 处理连续出现的****
	for strings.Contains(corpu, "**") {
		corpu = strings.Replace(corpu, "**", "*", -1)
	}

	return corpu
}

func MatchTemp(text string, temp string) string {
	corpu := ReplaceTempByStar(temp)
	// 模板匹配
	res := ""
	//reg := "^" + strings.Replace(corpu, "*", "(.{1,})", -1) + "$"
	// 完全模板匹配已经找不到，所以这里不用完全匹配，移除^ $
	reg := strings.Replace(corpu, "*", "(.{1,})", -1)
	exp := regexp.MustCompile(reg)
	if len(exp.FindAllString(text, -1)) > 0 {
		match_result := exp.FindAllStringSubmatch(text, -1)
		res = strings.Join(match_result[0][1:], ";")
	} else {
		return ""
	}

	return res
}

func TestReg(t *testing.T) {
	fmt.Println(RegulateString("模板開始1子1結束開始1子2結束，模板開始2子1結束開始2子2結束，", `開始(.*?)結束開始(.*?)結束`))

	fmt.Println(MatchTemp("帮我看看从广州这里到北京的飞机票吧", "从{FromCity}到{ToCity}的飞机票"))
}
