package leetcode

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	rotate(nums, 2)
	fmt.Println(nums)
}
