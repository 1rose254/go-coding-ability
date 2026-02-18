package Feb

/**
 * 693. 交替位二进制数
 * <a href="https://leetcode.cn/problems/binary-number-with-alternating-bits/description/"/>
 */

func hasAlternatingBits(n int) bool {
	x := n>>1 ^ n
	return (x+1)&x == 0
}
