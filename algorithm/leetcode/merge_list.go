package leetcode

//21. 合并两个有序链表
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
//示例：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 将pp的链表依次遍历插入到pBeInserted中
	pDest := l1  // 被插入的链表
	pRange := l2 // 需遍历的链表

	// 为了排除边界情况，将第一个值更小的链表最为被插入链表
	if l1.Val > l2.Val {
		pDest = l2
		pRange = l1
	}

	ret := pDest
	// -9,3
	// 5,7

	// 遍历pRange将各节点插入到pDest中合适的位置
	for pRange != nil {
		for pDest.Next != nil {
			if pDest.Next.Val >= pRange.Val {
				break
			}
			pDest = pDest.Next
		}

		pRangeNext := pRange.Next
		pRange.Next = nil
		if pDest.Next != nil {
			pRange.Next = pDest.Next
		}
		pDest.Next = pRange
		pDest = pRange
		pRange = pRangeNext
	}

	return ret
}
