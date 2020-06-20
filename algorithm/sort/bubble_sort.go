package sort

// 冒泡排序：两两相互比较把较大值移到尾部
func BubbleSort(src []int) []int {
	for i := 0; i < len(src) - 1; i ++ {
		for j := 0; j < len(src) - i - 1; j ++ {
			if src[j] > src[j+1] {
				src[j], src[j+1] = src[j+1], src[j]
			}
		}
	}
	return src
}
