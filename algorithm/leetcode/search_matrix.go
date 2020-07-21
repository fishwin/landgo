package leetcode

//面试题 10.09. 排序矩阵查找
//给定M×N矩阵，每一行、每一列都按升序排列，请编写代码找出某元素。
//
//示例:
//
//现有矩阵 matrix 如下：
//
//[
//[1,   4,  7, 11, 15],
//[2,   5,  8, 12, 19],
//[3,   6,  9, 16, 22],
//[10, 13, 14, 17, 24],
//[18, 21, 23, 26, 30]
//]
//给定 target = 5，返回 true。
//
//给定 target = 20，返回 false。

//思路：逐行遍历，查找首元素小于target和尾元素大于target的行，然后对每一行进行二分查找
func searchMatrix(matrix [][]int, target int) bool {
	for i := 0 ; i < len(matrix); i++ {
		if len(matrix[i]) <= 0 {
			return false
		}
		if matrix[i][0] == target || matrix[i][len(matrix[i])-1] == target { // 边界元素
			return true
		}
		if matrix[i][0] > target || matrix[i][len(matrix[i]) - 1] < target { // 首元素大于目标值或者尾元素小于目标值，直接跳过
			continue
		}

		// 二分查找
		left, right, mid := 0, len(matrix[i]) - 1, 0
		for left <= right {
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
			mid = (left + right) / 2
		}
	}
	return false
}
