package sort

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	nums := []int{1,4,5,6,2,6,7,2,3,5,7,0,3,4,20,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,4,4,4,4,4,4}
	quickSort(nums)
	fmt.Println(nums)
}
