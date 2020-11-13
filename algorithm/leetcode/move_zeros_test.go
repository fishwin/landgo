package leetcode

import "testing"

func Test_moveZeroes(t *testing.T) {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	t.Log(nums)
}