package leetcode

import (
	"testing"
)

func Test_levelOrder(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 7,
				Left: nil,
				Right: nil,
			},
		},
	}

	t.Log(levelOrder(root))
}
