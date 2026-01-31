package Jan

/**
 * 744. 寻找比目标字母大的最小字母
 * <a href="https://leetcode.cn/problems/find-smallest-letter-greater-than-target/description/"/>
 */

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	left, right := -1, n
	for left+1 < right {
		mid := left + (right-left)/2
		if letters[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}
	if right == n {
		return letters[0]
	}
	return letters[right]
}
