package leetcode

//102. 二叉树的层序遍历
//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//
//
//示例：
//二叉树：[3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回其层次遍历结果：
//
//[
//[3],
//[9,20],
//[15,7]
//]

func levelOrderCore(root *TreeNode, level int, ret *[][]int) {
	if root == nil {
		return
	}
	if len(*ret) < level {
		*ret = append(*ret, []int{})
	}
	(*ret)[level-1] = append((*ret)[level-1], root.Val)
	levelOrderCore(root.Left, level+1, ret)
	levelOrderCore(root.Right, level+1, ret)
}

// 记录层数
func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	levelOrderCore(root, 1, &ret)
	return ret
}
