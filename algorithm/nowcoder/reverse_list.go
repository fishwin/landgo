package nowcoder

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}

	p, head, temp := pHead, pHead, pHead


	fmt.Println("test")

	// 第一个节点不动，不断的将其后的节点放到头部
	for p.Next != nil {
		temp = p.Next // 指向将要移到头部的节点
		p.Next = temp.Next
		temp.Next = head
		head = temp
	}

	return head
}
