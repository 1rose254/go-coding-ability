package Feb

/**
 * 1382. 将二叉搜索树变平衡
 * <a href="https://leetcode.cn/problems/balance-a-binary-search-tree/description/"/>
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	m := len(nums) / 2
	return &TreeNode{
		Val:   nums[m],
		Left:  sortedArrayToBST(nums[:m]),
		Right: sortedArrayToBST(nums[m+1:]),
	}
}

func balanceBST(root *TreeNode) *TreeNode {
	nums := inorderTraversal(root)
	return sortedArrayToBST(nums)
}
