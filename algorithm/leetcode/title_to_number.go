package leetcode

import (
	"math"
)

//171. Excel表列序号
//给定一个Excel表格中的列名称，返回其相应的列序号。
//
//例如，
//
//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28
//...
//示例 1:
//
//输入: "A"
//输出: 1
//示例 2:
//
//输入: "AB"
//输出: 28
//示例 3:
//
//输入: "ZY"
//输出: 701

func titleToNumber(s string) int {
	sm := make(map[byte]int)

	num := 1
	for c := 'A'; c <= 'Z'; c++ {
		sm[byte(c)] = num
		num++
	}

	// 26进制
	mi := 0
	ret := 0
	for i := len(s) - 1; i >= 0; i-- {
		ret += sm[s[i]] * int(math.Pow(26, float64(mi)))
		mi++
	}

	return ret
}
