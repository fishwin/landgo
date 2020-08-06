package sort

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	return merge(MergeSort(nums[:mid]), MergeSort(nums[mid:]))
}

func merge(nums1 []int, nums2 []int) []int {
	if len(nums1) <= 0 {
		return nums2
	}
	if len(nums2) <= 0 {
		return nums1
	}

	var ret []int
	p1 := 0
	p2 := 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] < nums2[p2] {
			ret = append(ret, nums1[p1])
			p1++
		} else {
			ret = append(ret, nums2[p2])
			p2++
		}
	}

	if p1 != len(nums1) {
		ret = append(ret, nums1[p1:]...)
	}
	if p2 != len(nums2) {
		ret = append(ret, nums2[p2:]...)
	}

	return ret
}
