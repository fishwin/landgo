package sort

// 插入排序：将序列分为有序和无序两部分，依次遍历无序序列中的每个元素，并将他们插入到有序序列的合适位置
func InsertSort(src []int) []int {
	if len(src) <= 1 {
		return src
	}

	for i := 1; i < len(src); i++ {
		for j := i; j > 0; j-- {
			if src[j] < src[j-1] { // 交换
				src[j], src[j-1] = src[j-1], src[j]
			} else {
				break
			}
		}
	}

	return src
}
