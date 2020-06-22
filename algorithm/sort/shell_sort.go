package sort

// 希尔排序：将序列按照步长分组，每组进行插入排序，步长每次除2，直到等于0结束
func ShellSort(src []int) []int {
	if len(src) <= 1 {
		return src
	}
	for step := len(src) / 2; step > 0; step /= 2 {
		for i := step; i < len(src); i += step {
			for j := i; j > 0; j -= step {
				if j-step >= 0 && src[j] < src[j-step] {
					src[j], src[j-step] = src[j-step], src[j]
				} else {
					break
				}
			}
		}
	}

	return src
}
