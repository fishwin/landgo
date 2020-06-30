package leetcode

//面试题 10.05. 稀疏数组搜索

//稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。
//
//示例1:
//
//输入: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
//输出：-1
//说明: 不存在返回-1。
//示例2:
//
//输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
//输出：4
//提示:
//
//words的长度在[1, 1000000]之间

func FindString(words []string, s string) int {
	left := 0
	right := len(words) - 1
	for left <= right {
		mid := (left + right) / 2

		temp := mid
		compareWord := words[mid]
		if compareWord == "" {
			for temp = mid; temp <= right; temp++ {
				if words[temp] != "" {
					compareWord = words[temp]
					break
				}
			}
		}

		if compareWord == "" {
			right = mid - 1
		} else {
			if compareWord == s {
				return temp
			} else if compareWord < s {
				left = temp + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
