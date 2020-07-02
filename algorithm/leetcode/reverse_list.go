package leetcode

//剑指 Offer 24. 反转链表
//定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
//
// 
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
// 
//
//限制：
//
//0 <= 节点个数 <= 5000

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 头结点原地不动，不断的将后面的元素扔到头部
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	p := head

	for p.Next != nil {
		// 将p后面的节点移到头部
		temp := p.Next.Next
		p.Next.Next = head
		head = p.Next
		p.Next = temp
	}

	return head
}
