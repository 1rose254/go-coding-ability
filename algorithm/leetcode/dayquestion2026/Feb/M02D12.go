package Feb

/**
 * 3713. 最长的平衡子串 I
 * <a href="https://leetcode.cn/problems/longest-balanced-substring-i/description/"/>
 */

func longestBalanced3(s string) (ans int) {
	for i := range s {
		cnt := [26]int{}
		mx, kinds := 0, 0
		for j := i; j < len(s); j++ {
			b := s[j] - 'a'
			if cnt[b] == 0 {
				kinds++
			}
			cnt[b]++
			mx = max(mx, cnt[b])
			if mx*kinds == j-i+1 {
				ans = max(ans, j-i+1)
			}
		}
	}
	return
}
