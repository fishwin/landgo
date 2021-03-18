package leetcode

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	// [1,3,2,4,3,2,1]

	n4 := &ListNode{
		Val:  1,
		Next: nil,
	}
	n3 := &ListNode{
		Val:  2,
		Next: n4,
	}
	n2 := &ListNode{
		Val:  3,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  4,
		Next: n2,
	}
	n11 := &ListNode{
		Val:  2,
		Next: n1,
	}
	n12 := &ListNode{
		Val:  3,
		Next: n11,
	}
	n13 := &ListNode{
		Val:  1,
		Next: n12,
	}

	fmt.Println(isPalindrome(n13))
}
