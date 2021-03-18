package leetcode

//922. 按奇偶排序数组 II

//给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。
//
//对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。
//
//你可以返回任何满足上述条件的数组作为答案。
//
//
//
//示例：
//
//输入：[4,2,5,7]
//输出：[4,5,2,7]
//解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
//
//
//提示：
//
//2 <= A.length <= 20000
//A.length % 2 == 0
//0 <= A[i] <= 1000

func verify(i, v int) bool {
	if (i+v)%2 == 0 { // 奇数和偶数相加结果一定是奇数，否则是偶数
		return true
	}
	return false
}

func sortArrayByParityII(A []int) []int {
	n := len(A)

	// 遍历A数组，如果遇到不满足条件的值，则向后查找第一个不满足当前位置，但是满足这个位置的值，进行交换
	for i := 0; i < n; i++ {
		if !verify(i, A[i]) {
			j := i + 1
			for ; j < n; j++ {
				if !verify(j, A[j]) && verify(i, A[j]) {
					break
				}
			}
			A[i], A[j] = A[j], A[i]
		}
	}

	return A
}
