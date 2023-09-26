package leetcode

import (
	"math"
)

//209. 长度最小的子数组
//给定一个含有 n 个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其总和大于等于 target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
//
//
//
//示例 1：
//
//输入：target = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//示例 2：
//
//输入：target = 4, nums = [1,4,4]
//输出：1
//示例 3：
//
//输入：target = 11, nums = [1,1,1,1,1,1,1,1]
//输出：0
//
//
//提示：
//
//1 <= target <= 109
//1 <= nums.length <= 105
//1 <= nums[i] <= 105
//
//
//进阶：
//
//如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。

func minSubArrayLen(target int, nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	exist := false
	window := [2]int{0, -1}
	windowSum, minRes := 0, math.MaxInt32

	for window[0] < len(nums) && window[1] < len(nums) {
		if windowSum >= target {
			exist = true
			minRes = min(minRes, lenWindow(window))
			windowSum -= nums[window[0]]
			window[0]++
		} else {
			window[1]++
			if window[1] < len(nums) {
				windowSum += nums[window[1]]
			}
		}
	}
	if !exist {
		return 0
	}

	return minRes
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lenWindow(window [2]int) int {
	return window[1] - window[0] + 1
}
