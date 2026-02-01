package Feb

import "math"

/**
 * 3010. 将数组分成最小总代价的子数组 I
 * <a href="https://leetcode.cn/problems/divide-an-array-into-subarrays-with-minimum-cost-i/description/"/>
 */

func minimumCost(nums []int) int {
	first, second := math.MaxInt, math.MaxInt
	for _, x := range nums[1:] {
		if x < first {
			second = first
			first = x
		} else if x < second {
			second = x
		}
	}
	return nums[0] + first + second
}
