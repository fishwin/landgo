package sort

func reset(nums []int, left, right int) int {
	ref := left
	p := ref + 1
	for i := p; i <= right; i++ {
		if nums[i] < nums[left] {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}

	// 基准值复位
	nums[p - 1], nums[ref] = nums[ref], nums[p - 1]
	return p - 1
}

func _quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	resetIndex := reset(nums, left, right)
	_quickSort(nums, left, resetIndex-1)
	_quickSort(nums, resetIndex+1, right)
}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	_quickSort(nums, 0, len(nums)-1)
}
