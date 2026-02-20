package Feb

import (
	"cmp"
	"slices"
	"strings"
)

/**
 * 761. 特殊的二进制字符串
 * <a href="https://leetcode.cn/problems/special-binary-string/description/"/>
 */

func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}
	substrings := []string{}
	diff := 0
	start := 0
	for i, ch := range s {
		if ch == '1' {
			diff++
		} else if diff--; diff == 0 {
			substrings = append(substrings, "1"+makeLargestSpecial(s[start+1:i])+"0")
			start = i + 1
		}
	}
	slices.SortFunc(substrings, func(a, b string) int { return cmp.Compare(b, a) })
	return strings.Join(substrings, "")
}
