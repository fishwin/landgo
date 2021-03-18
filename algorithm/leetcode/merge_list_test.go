package leetcode

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	//n14 := &ListNode{
	//	Val:  4,
	//	Next: nil,
	//}
	n12 := &ListNode{
		Val:  3,
		Next: nil,
	}
	n11 := &ListNode{
		Val:  -9,
		Next: n12,
	}

	//n24 := &ListNode{
	//	Val:  4,
	//	Next: nil,
	//}
	n23 := &ListNode{
		Val:  7,
		Next: nil,
	}
	n21 := &ListNode{
		Val:  5,
		Next: n23,
	}

	head := mergeTwoLists(n11, n21)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}
