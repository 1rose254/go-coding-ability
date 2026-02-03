package Feb

import "math"

/**
 * 3640. 三段式数组 II
 * <a href="https://leetcode.cn/problems/trionic-array-ii/description/"/>
 */

func maxSumTrionic(nums []int) int64 {
	const negInf = math.MinInt / 2
	ans, f1, f2, f3 := negInf, negInf, negInf, negInf
	for i := 1; i < len(nums); i++ {
		x, y := nums[i-1], nums[i]
		if x < y {
			f3 = max(f3, f2) + y
			ans = max(ans, f3)
			f2 = negInf
			f1 = max(f1, x) + y
		} else if x > y {
			f2 = max(f2, f1) + y
			f1, f3 = negInf, negInf
		} else {
			f1, f2, f3 = negInf, negInf, negInf
		}
	}
	return int64(ans)
}
