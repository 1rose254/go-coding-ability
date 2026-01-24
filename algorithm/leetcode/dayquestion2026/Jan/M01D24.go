package Jan

import "slices"

/**
 * 1877. 数组中最大数对和的最小值
 * <a href="https://leetcode.cn/problems/minimize-maximum-pair-sum-in-array/description/"/>
 */

func minPairSum(nums []int) (ans int) {
	slices.Sort(nums)
	n := len(nums)
	for i, x := range nums[:n/2] {
		ans = max(ans, x+nums[n-1-i])
	}
	return
}
