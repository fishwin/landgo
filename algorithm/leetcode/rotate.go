package leetcode

//
//189. 轮转数组
//给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
//
//
//
//示例 1:
//
//输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右轮转 1 步: [7,1,2,3,4,5,6]
//向右轮转 2 步: [6,7,1,2,3,4,5]
//向右轮转 3 步: [5,6,7,1,2,3,4]
//示例 2:
//
//输入：nums = [-1,-100,3,99], k = 2
//输出：[3,99,-1,-100]
//解释:
//向右轮转 1 步: [99,-1,-100,3]
//向右轮转 2 步: [3,99,-1,-100]
//
//
//提示：
//
//1 <= nums.length <= 105
//-231 <= nums[i] <= 231 - 1
//0 <= k <= 105
//
//
//进阶：
//
//尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
//你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？

func rotate(nums []int, k int) {
	if len(nums) == 0 || k <= 0 {
		return
	}
	if k > len(nums) {
		k = k % len(nums)
	}

	ret := nums[len(nums)-k:]
	ret = append(ret, nums[:len(nums)-k]...)
	//nums = append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, ret)
	// fmt.Println(nums)
}

// 给你一个整数数组 arr ，数组中的每个整数 互不相同 。另有一个由整数数组构成的数组 pieces，其中的整数也 互不相同 。
// 请你以 任意顺序 连接 pieces 中的数组以形成 arr 。但是，不允许 对每个数组 pieces[i] 中的整数重新排序。
//
//如果可以连接 pieces 中的数组形成 arr ，返回 true ；否则，返回 false 。
//
//
//示例 1：
//
//输入：arr = [15,88], pieces = [[88],[15]]
//输出：true
//解释：依次连接 [15] 和 [88]
//示例 2：
//
//输入：arr = [49,18,16], pieces = [[16,18,49]]
//输出：false
//解释：即便数字相符，也不能重新排列 pieces[0]
//示例 3：
//
//输入：arr = [91,4,64,78], pieces = [[78],[4,64],[91]]
//输出：true
//解释：依次连接 [91]、[4,64] 和 [78]

func canFormArray(arr []int, pieces [][]int) bool {
	if len(arr) == 0 {
		return false
	}

	m := make(map[int]bool)
	pArr, pPrev := 0, 0

	for pArr < len(arr) {
		x := 0
		for ; x < len(pieces); x++ {
			if m[x] {
				continue
			}
			pPrev = pArr
			for y := 0; y < len(pieces[x]); y++ {
				if pieces[x][y] != arr[pArr] {
					break
				}
				pArr++
			}
			if pArr-pPrev != len(pieces[x]) {
				pArr = pPrev
			} else {
				m[x] = true
				break
			}
		}
		if x == len(pieces) {
			return false
		}
	}

	return true
}

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//
//
//
//示例 1:
//
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]
//示例 2:
//
//输入: nums = [0]
//输出: [0]

func moveZeroes2(nums []int) {
	if len(nums) == 0 {
		return
	}

	c := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			c++
		} else {
			if c > 0 {
				nums[i-c] = nums[i]
			}
		}
	}

	for i := 1; i <= c; i++ {
		nums[len(nums)-i] = 0
	}
}

func twoSum(numbers []int, target int) []int {
	return nil
}
