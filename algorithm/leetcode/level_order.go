package leetcode

//剑指 Offer 32 - II. 从上到下打印二叉树 II
//从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

//例如:
//给定二叉树: [3,9,20,null,null,15,7],
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
//
//
//提示：
//
//节点总数 <= 1000

func levelOrderCore(root *TreeNode, level int, ret *[][]int) {
	if root == nil {
		return
	}

	if len(*ret) <= level { // 如果ret中还未初始化当前行
		*ret = append(*ret, []int{root.Val})
	} else {
		(*ret)[level] = append((*ret)[level], root.Val)
	}

	levelOrderCore(root.Left, level + 1, ret)
	levelOrderCore(root.Right, level + 1, ret)
}

func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}

	levelOrderCore(root, 0, &ret)
	return ret
}