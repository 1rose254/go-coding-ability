package Feb

/**
 * 3714. 最长的平衡子串 II
 * <a href="https://leetcode.cn/problems/longest-balanced-substring-ii/description/"/>
 */

func longestBalanced4(s string) (ans int) {
	n := len(s)

	for i := 0; i < n; {
		start := i
		for i++; i < n && s[i] == s[i-1]; i++ {
		}
		ans = max(ans, i-start)
	}

	f := func(x, y byte) {
		for i := 0; i < n; i++ {
			pos := map[int]int{0: i - 1}
			d := 0
			for ; i < n && (s[i] == x || s[i] == y); i++ {
				if s[i] == x {
					d++
				} else {
					d--
				}
				if j, ok := pos[d]; ok {
					ans = max(ans, i-j)
				} else {
					pos[d] = i
				}
			}
		}
	}
	f('a', 'b')
	f('a', 'c')
	f('b', 'c')

	type pair struct{ diffAB, diffBC int }
	pos := map[pair]int{{}: -1}
	cnt := [3]int{}
	for i, b := range s {
		cnt[b-'a']++
		p := pair{cnt[0] - cnt[1], cnt[1] - cnt[2]}
		if j, ok := pos[p]; ok {
			ans = max(ans, i-j)
		} else {
			pos[p] = i
		}
	}
	return
}
