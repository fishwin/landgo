package leetcode

//面试题 02.02. 返回倒数第 k 个节点
//实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
//
//注意：本题相对原题稍作改动
//
//示例：
//
//输入： 1->2->3->4->5 和 k = 2
//输出： 4
//说明：
//
//给定的 k 保证是有效的。

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLastCore(head *ListNode, k int) (index int, v int) {
	if head == nil {
		return 0, -1
	}

	index, v = kthToLastCore(head.Next, k)
	index++
	if index == k {
		v = head.Val
	}

	return index, v
}

func kthToLast(head *ListNode, k int) int {
	_, v := kthToLastCore(head, k)
	return v
}