package leetcode

//107. 二叉树的层次遍历 II
//给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//
//例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回其自底向上的层次遍历为：
//
//[
//[15,7],
//[9,20],
//[3]
//]


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func _core(root *TreeNode, level int, ret *[][]int) {
	if root == nil {
		return
	}
	if len(*ret) < level {
		*ret = append(*ret, []int{})
	}
	(*ret)[level-1] = append((*ret)[level-1], root.Val)
	_core(root.Left, level+1, ret)
	_core(root.Right, level+1, ret)
}

func reverse(ret *[][]int) {
	n := len(*ret)
	for i := 0 ; i < n / 2; i++ {
		(*ret)[i], (*ret)[n-i-1] =  (*ret)[n-i-1],(*ret)[i]
	}
}

func levelOrderBottom(root *TreeNode) [][]int {
	var ret [][]int
	_core(root, 1, &ret)
	reverse(&ret)
	return ret
}
