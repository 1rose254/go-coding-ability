package Jan

import (
	"math"
	"slices"
)

/**
 * 1984. 学生分数的最小差值
 * <a href="https://leetcode.cn/problems/minimum-difference-between-highest-and-lowest-of-k-scores/description/"/>
 */

func minimumDifference(nums []int, k int) int {
	slices.Sort(nums)
	ans := math.MaxInt
	for i := k - 1; i < len(nums); i++ {
		ans = min(ans, nums[i]-nums[i-k+1])
	}
	return ans
}
