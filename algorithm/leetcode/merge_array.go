package leetcode

//面试题 10.01. 合并排序的数组
//给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。
//
//初始化 A 和 B 的元素数量分别为 m 和 n。
//
//示例:
//
//输入:
//A = [1,2,3,0,0,0], m = 3
//B = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6]
//说明:
//
//A.length == n + m

func merge(A []int, m int, B []int, n int) {
	pa := m - 1        // 指向A数组最后一个有效的元素
	pb := n - 1        // 指向B数组最后一个有效的元素
	pTail := n + m - 1 // 指向A数组最后一个元素

	for pTail >= 0 {
		if pa >= 0 && pb >= 0 {
			if A[pa] > B[pb] {
				A[pTail] = A[pa]
				pa--
			} else {
				A[pTail] = B[pb]
				pb--
			}
		} else if pa < 0 && pb >= 0 {
			A[pTail] = B[pb]
			pb--
		} else {
			break
		}
		pTail--
	}
}
