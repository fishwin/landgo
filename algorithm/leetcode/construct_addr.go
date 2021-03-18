package leetcode

//剑指 Offer 66. 构建乘积数组
//给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
//
//示例:
//
//输入: [1,2,3,4,5]
//输出: [120,60,40,30,24]
//
//
//提示：
//
//所有元素乘积之和不会溢出 32 位整数
//a.length <= 100000

func constructArr(a []int) []int {
	if len(a) == 0 {
		return nil
	}

	zeroCount := 0 // 0的个数
	temp := 1      // 非0所有数字的乘积
	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			zeroCount++
		} else {
			temp *= a[i]
		}
	}

	if zeroCount == 0 { // 如果元素都非0,则直接乘积除以当前元素即为新元素值
		for i := 0; i < len(a); i++ {
			a[i] = temp / a[i]
		}
	} else if zeroCount == 1 { // 如果有一个0元素,那么除0元素的新元素非0外,其余都是0
		for i := 0; i < len(a); i++ {
			if a[i] == 0 {
				a[i] = temp
			} else {
				a[i] = 0
			}
		}
	} else { // 如果有一个以上的0元素,那么新元素全是0
		a = make([]int, len(a))
	}

	return a
}
