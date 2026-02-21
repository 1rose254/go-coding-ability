package Feb

import (
	"math/bits"
	"slices"
)

/**
 * 762. 二进制表示中质数个计算置位
 * <a href="https://leetcode.cn/problems/prime-number-of-set-bits-in-binary-representation/description/"/>
 */

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19}

func countPrimeSetBits(left, right int) (ans int) {
	for x := left; x <= right; x++ {
		if slices.Contains(primes, bits.OnesCount(uint(x))) {
			ans++
		}
	}
	return
}
