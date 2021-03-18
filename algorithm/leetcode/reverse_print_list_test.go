package leetcode

import (
	"fmt"
	"testing"
)

func Test_reversePrint(t *testing.T) {
	n14 := &ListNode{
		Val:  4,
		Next: nil,
	}
	n12 := &ListNode{
		Val:  3,
		Next: n14,
	}
	n11 := &ListNode{
		Val:  1,
		Next: n12,
	}

	fmt.Println(reversePrint(n11))
}
