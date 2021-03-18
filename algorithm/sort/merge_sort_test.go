package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {

	t.Log(MergeSort([]int{1, 2, 4, 1, 2, 9, 8, 9, 4, 3, 6, 7, 5, 3, 7, 0, 4, 4, 4, 3, 2, 1, 1, 1, 1, 4, 5, 6, 7, 8, 9}))
}
