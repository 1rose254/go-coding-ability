package Jan

/**
 * 3507. 移除最小数对使数组有序 I
 * <a href="https://leetcode.cn/problems/minimum-pair-removal-to-sort-array-i/description/"/>
 */
func minimumPairRemoval(nums []int) int {
	list := make([]int, len(nums))
	copy(list, nums)
	op := 0
	for {
		if check(list) {
			break
		}
		minSum := int(^uint(0) >> 1)
		idx := -1

		for i := 0; i < len(list)-1; i++ {
			sum := list[i] + list[i+1]
			if sum < minSum {
				minSum = sum
				idx = i
			}
		}

		val := list[idx] + list[idx+1]
		list = append(list[:idx+1], list[idx+2:]...)
		list[idx] = val
		op++
	}

	return op
}

func check(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}
	return true
}
