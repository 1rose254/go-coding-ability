package Jan

/**
 * 3314. 构造最小位运算数组 I
 * <a href="https://leetcode.cn/problems/construct-the-minimum-bitwise-array-i/description/"/a>
 */

func minBitwiseArray(nums []int) []int {
	for i, x := range nums {
		if x == 2 {
			nums[i] = -1
		} else {
			t := ^x
			nums[i] ^= t & -t >> 1
		}
	}
	return nums
}
