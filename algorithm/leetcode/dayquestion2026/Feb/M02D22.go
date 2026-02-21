package Feb

import "math"

/**
 * 868. 二进制间距
 * <a href="https://leetcode.cn/problems/binary-gap/description/"/>
 */

func binaryGap(n int) int {
	ans := 0
	for i, j := 31, -1; i >= 0; i-- {
		if ((n >> i) & 1) == 1 {
			if j != -1 {
				ans = int(math.Max(float64(ans), float64(j-i)))
			}
			j = i
		}
	}
	return ans
}
