package leetcode

import "sort"

//剑指 Offer 61. 扑克牌中的顺子
//从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。
//
//
//
//示例 1:
//
//输入: [1,2,3,4,5]
//输出: True
//
//
//示例 2:
//
//输入: [0,0,1,2,5]
//输出: True
//
//
//限制：
//
//数组长度为 5
//
//数组的数取值为 [0, 13] .

func isStraight(nums []int) bool {
	sort.Ints(nums)
	zeroCount := 0 // 王牌的个数
	for i := 0; i < len(nums);i++ {
		if nums[i] == 0 {
			zeroCount++
		} else {
			if i > 0 {
				if nums[i] == nums[i-1] { // 有重复的
					return false
				}
			}
		}
	}

	if zeroCount == 0 {
		if nums[4] - nums[0] == 4 {
			return true
		}
	} else if zeroCount == 1 {
		if nums[4] - nums[1] < 5 {
			return true
		}
	} else if zeroCount == 2 {
		if nums[4] - nums[2] < 5 {
			return true
		}
	} else if zeroCount == 3 {
		if nums[4] - nums[3] < 5 {
			return true
		}
	} else { // 4, 5
		return true
	}

	return false
}