package sort

// 每次从后面选出最小的值放到前面
func SelectionSort(src []int) []int {
	length := len(src)

	for i := 0; i < length - 1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if src[j] < src[min] {
				min = j
			}
		}
		if min != i {
			src[i], src[min] = src[min], src[i]
		}
	}
	return src
}
