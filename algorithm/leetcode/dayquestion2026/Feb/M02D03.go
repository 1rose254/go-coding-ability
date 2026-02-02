package Feb

/**
 * 3637. 三段式数组 I
 * <a href="https://leetcode.cn/problems/trionic-array-i/description/"/>
 */

func isTrionic(nums []int) bool {
	if nums[0] >= nums[1] {
		return false
	}
	cnt := 1
	for i := 2; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return false
		}
		if (nums[i-2] < nums[i-1]) != (nums[i-1] < nums[i]) {
			cnt++
		}
	}
	return cnt == 3
}
