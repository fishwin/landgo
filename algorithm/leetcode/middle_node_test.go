package leetcode

import (
	"fmt"
	"testing"
)

func Test_middleNode(t *testing.T) {
	n5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	n4 := &ListNode{
		Val:  4,
		Next: n5,
	}
	n3 := &ListNode{
		Val:  3,
		Next: n4,
	}
	n2 := &ListNode{
		Val:  2,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}
	//head := &ListNode{
	//	Val:  -1,
	//	Next: n1,
	//}

	fmt.Println(middleNode(n1))

}