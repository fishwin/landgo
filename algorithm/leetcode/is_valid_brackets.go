package leetcode

//20. 有效的括号
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//
//示例 1:
//
//输入: "()"
//输出: true
//示例 2:
//
//输入: "()[]{}"
//输出: true
//示例 3:
//
//输入: "(]"
//输出: false
//示例 4:
//
//输入: "([)]"
//输出: false
//示例 5:
//
//输入: "{[]}"
//输出: true

func isValid(s string) bool {
	temp := make([]byte, len(s)) // 使用数组模拟栈，用于存放左括号
	tail := -1                   // 指向temp最后一个元素

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' { // 左括号入栈
			tail++
			temp[tail] = s[i]
		} else {
			if tail == -1 {
				return false
			}
			if (s[i] == ')' && temp[tail] == '(') || (s[i] == ']' && temp[tail] == '[') || (s[i] == '}' && temp[tail] == '{') {
				tail--
			} else {
				return false
			}
		}
	}

	// 最后如果tail不指向-1则表明左括号没有被匹配完，故返回false
	if tail != -1 {
		return false
	}
	return true
}
