package leetcode

//回文链表
//编写一个函数，检查输入的链表是否是回文的。
//
//示例 1：
//
//输入： 1->2
//输出： false
//示例 2：
//
//输入： 1->2->2->1
//输出： true
//
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 快慢指针查找中间节点，执行结束后，slow指向的节点即为中间节点
	slow := head
	fast := head.Next.Next
	for slow != nil && fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			fast = nil
		} else {
			fast = fast.Next.Next
		}
	}

	mid := slow
	q := mid.Next // 指向反转的第一个节点

	// 将中间节点后面的后半段链表进行反转
	for q.Next != nil {
		temp := q.Next
		q.Next = q.Next.Next
		temp.Next = mid.Next
		mid.Next = temp
	}

	// 指向前半段链表起始位置
	p1 := head
	// 指向后半段链表起始位置
	p2 := mid.Next
	for p2 != nil {
		if p2.Val != p1.Val {
			return false
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}
