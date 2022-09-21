package leetcode

func oddEvenArray(nums []int) int {

	odd, even := -1, -1
	numOdd, numEven := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] & 1 == 0 {
			numEven++
			even = nums[i]
		} else {
			numOdd++
			odd = nums[i]
		}
		if numEven > 1 && odd != -1 {
			return odd
		}
		if numOdd > 1 && even != -1 {
			return even
		}
	}
	//if numEven > 1 {
	//	return odd
	//}
	//if numOdd > 1 {
	//	return even
	//}
	return -1
}
