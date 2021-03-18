package leetcode

//剑指 Offer 38. 字符串的排列
//输入一个字符串，打印出该字符串中字符的所有排列。
//
//
//
//你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
//
//
//
//示例:
//
//输入：s = "abc"
//输出：["abc","acb","bac","bca","cab","cba"]
//
//
//限制：
//
//1 <= s 的长度 <= 8

// remain 剩余未使用过的字符数组
// last 上一次得到的结果

func core(remain []byte, last []byte) []string {
	if len(remain) == 1 {
		return []string{string(append(last, remain[0]))}
	}

	var ret []string
	var remainBytes []byte
	m := make(map[byte]bool)

	for i := 0; i < len(remain); i++ {
		if m[remain[i]] { // 同一层中如果已经计算过相同字符，就不再重复计算
			continue
		}
		m[remain[i]] = true
		// 在剩余字符中删除当前字符，因为已经使用过
		remainBytes = append(remainBytes, remain[:i]...)
		remainBytes = append(remainBytes, remain[i+1:]...)

		// 添加当前字符到结果中
		//res := append(last, remain[i])
		ret = append(ret, core(remainBytes, append(last, remain[i]))...)
		remainBytes = remainBytes[0:0]
	}

	return ret
}

func permutation(s string) []string {
	if len(s) == 0 {
		return nil
	}
	return core([]byte(s), []byte{})
}
