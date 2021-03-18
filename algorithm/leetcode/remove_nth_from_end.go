package leetcode

//19. 删除链表的倒数第N个节点
//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
//示例：
//
//给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//说明：
//
//给定的 n 保证是有效的。
//
//进阶：
//
//你能尝试使用一趟扫描实现吗？

func core2(head *ListNode, n int) (h *ListNode, reverseN int, remove bool) {
	if head.Next == nil {
		return head, 1, false
	}

	h, reverseN, remove = core2(head.Next, n)

	if reverseN == n {
		head.Next = h.Next
		remove = true
	} else {
		head.Next = h
	}
	reverseN++
	return head, reverseN, remove
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || (head.Next == nil && n == 1) {
		return nil
	}
	_, _, remove := core2(head, n)
	if !remove {
		return head.Next
	}
	return head
}
