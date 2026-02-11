package Feb

import "math"

/**
 * 3721. 最长平衡子数组 II
 * <a href="https://leetcode.cn/problems/longest-balanced-subarray-ii/description/"/>
 */

func longestBalanced2(nums []int) (ans int) {
	n := len(nums)
	B := int(math.Sqrt(float64(n+1)))/2 + 1
	sum := make([]int, n+1)

	type block struct {
		l, r int
		todo int
		pos  map[int]int
	}
	blocks := make([]block, n/B+1)
	calcPos := func(l, r int) map[int]int {
		pos := map[int]int{}
		for j := r - 1; j >= l; j-- {
			pos[sum[j]] = j
		}
		return pos
	}
	for i := 0; i <= n; i += B {
		r := min(i+B, n+1)
		pos := calcPos(i, r)
		blocks[i/B] = block{i, r, 0, pos}
	}

	rangeAdd := func(l, r, v int) {
		for i := range blocks {
			b := &blocks[i]
			if b.r <= l {
				continue
			}
			if b.l >= r {
				break
			}
			if l <= b.l && b.r <= r {
				b.todo += v
			} else {
				for j := b.l; j < b.r; j++ {
					sum[j] += b.todo
					if l <= j && j < r {
						sum[j] += v
					}
				}
				b.pos = calcPos(b.l, b.r)
				b.todo = 0
			}
		}
	}

	findFirst := func(r, v int) int {
		for i := range blocks {
			b := &blocks[i]
			if b.r <= r {
				if j, ok := b.pos[v-b.todo]; ok {
					return j
				}
			} else {
				for j := b.l; j < r; j++ {
					if sum[j] == v-b.todo {
						return j
					}
				}
				break
			}
		}
		return n
	}
	last := map[int]int{}
	for i := 1; i <= n; i++ {
		x := nums[i-1]
		v := x%2*2 - 1
		if j := last[x]; j == 0 {
			rangeAdd(i, n+1, v)
		} else {
			rangeAdd(j, i, -v)
		}
		last[x] = i

		s := sum[i] + blocks[i/B].todo
		ans = max(ans, i-findFirst(i-ans, s))
	}
	return
}
