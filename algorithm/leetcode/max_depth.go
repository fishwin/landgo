package leetcode

//559. N叉树的最大深度
//给定一个 N 叉树，找到其最大深度。
//
//最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
//
//例如，给定一个 3叉树 :
//
//
//我们应返回其最大深度，3。
//
//说明:
//
//树的深度不会超过 1000。
//树的节点总不会超过 5000。

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1
	}

	max := 0
	for i := 0; i < len(root.Children); i++ {
		v := maxDepth(root.Children[i]) + 1
		if v > max {
			max = v
		}
	}

	return max
}
