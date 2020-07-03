package leetcode

//92. 反转链表 II
//反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//
//说明:
//1 ≤ m ≤ n ≤ 链表长度。
//
//示例:
//
//输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n || head == nil || head.Next == nil {
		return head
	}
	// 构造一个临时头节点p0
	tempHead := &ListNode{
		Val:  -1,
		Next: head,
	}

	// 查找第m-1个节点地址
	index := 0
	p0 := tempHead
	for p0 != nil {
		if index == m - 1 {
			break
		}
		index++
		p0 = p0.Next
	}

	// 从第m个节点依次遍历n-m个，m原地不动,依次将其后面的节点放到p0后面
	times := n - m
	pm := p0.Next
	for pm.Next != nil {
		temp := pm.Next.Next
		p0Next := p0.Next
		p0.Next = pm.Next
		p0.Next.Next = p0Next
		pm.Next = temp

		times--
		if times == 0 {
			break
		}
	}

	return tempHead.Next
}