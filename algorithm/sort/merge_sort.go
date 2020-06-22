package sort

func MergeSort(src []int) []int {
	if len(src) <= 1 {
		return src
	}

	return mergeSort(src, 0, len(src) - 1)
}

func mergeSort(src []int, left int, right int) []int {
	if left == right {
		return []int{src[left]}
	}
	mid := (left + right) / 2
	ret1 := mergeSort(src, left, mid)
	ret2 := mergeSort(src, mid+1, right)
	return merge(ret1, ret2)
}

func merge(ret1 []int, ret2 []int) []int {
	var ret []int

	i, j := 0, 0
	for i < len(ret1) && j < len(ret2) {
		if ret1[i] < ret2[j] {
			ret = append(ret, ret1[i])
			i++
		} else {
			ret = append(ret, ret2[j])
			j++
		}
	}

	if i == len(ret1) && j < len(ret2) {
		for ; j < len(ret2);j++ {
			ret = append(ret, ret2[j])
		}
	} else if j == len(ret2) && i < len(ret1) {
		for ; i < len(ret1);i++ {
			ret = append(ret, ret1[i])
		}
	}

	return ret
}
