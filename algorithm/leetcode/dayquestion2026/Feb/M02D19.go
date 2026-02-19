package Feb

/**
 * 696. 计数二进制子串
 * <a href="https://leetcode.cn/problems/count-binary-substrings/description/"/>
 */

func countBinarySubstrings(s string) (ans int) {
	n := len(s)
	pre, cur := 0, 0
	for i := range n {
		cur++
		if i == n-1 || s[i] != s[i+1] {
			ans += min(pre, cur)
			pre = cur
			cur = 0
		}
	}
	return
}
