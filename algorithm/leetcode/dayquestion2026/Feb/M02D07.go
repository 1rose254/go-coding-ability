package Feb

import "strings"

/**
 * 1653. 使字符串平衡的最少删除次数
 * <a href="https://leetcode.cn/problems/minimum-deletions-to-make-string-balanced/description/"/>
 */

func minimumDeletions(s string) int {
	del := strings.Count(s, "a")
	ans := del
	for _, c := range s {
		del += int((c-'a')*2 - 1)
		if del < ans {
			ans = del
		}
	}
	return ans
}
