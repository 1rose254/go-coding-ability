package Feb

import "slices"

/**
 * 3634. 使数组平衡的最少移除数目
 * <a href="https://leetcode.cn/problems/minimum-removals-to-balance-array/description/"/>
 */

func minRemoval(nums []int, k int) int {
	slices.Sort(nums)
	maxSave, left := 0, 0
	for i, mx := range nums {
		for nums[left]*k < mx {
			left++
		}
		maxSave = max(maxSave, i-left+1)
	}
	return len(nums) - maxSave
}
