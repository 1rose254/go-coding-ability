package Feb

import (
	"go-coding-ability/algorithm/leetcode/base"
)

/**
 * 110. 平衡二叉树
 * <a href="https://leetcode.cn/problems/balanced-binary-tree/description/"/>
 */

func getHeight(node *base.TreeNode) int {
	if node == nil {
		return 0
	}
	leftH := getHeight(node.Left)
	if leftH == -1 {
		return -1 // 提前退出，不再递归
	}
	rightH := getHeight(node.Right)
	if rightH == -1 || abs(leftH-rightH) > 1 {
		return -1
	}
	return max(leftH, rightH) + 1
}

func isBalanced(root *base.TreeNode) bool {
	return getHeight(root) != -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
