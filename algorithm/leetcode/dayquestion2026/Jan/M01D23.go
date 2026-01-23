package Jan

import "container/heap"

/**
 * 3510. 移除最小数对使数组有序 II
 * <a href="https://leetcode.cn/problems/minimum-pair-removal-to-sort-array-ii/description/"/>
 */

func minimumPairRemoval2(nums []int) (ans int) {
	n := len(nums)
	h := make(hp, n-1)
	dec := 0
	for i := range n - 1 {
		x, y := nums[i], nums[i+1]
		if x > y {
			dec++
		}
		h[i] = pair{x + y, i}
	}
	heap.Init(&h)

	left := make([]int, n+1)
	right := make([]int, n)
	for i := range n {
		left[i] = i - 1
		right[i] = i + 1
	}
	remove := func(i int) {
		l, r := left[i], right[i]
		right[l] = r
		left[r] = l
		right[i] = n
	}

	for dec > 0 {
		ans++

		for right[h[0].i] >= n || nums[h[0].i]+nums[right[h[0].i]] != h[0].s {
			heap.Pop(&h)
		}
		p := heap.Pop(&h).(pair)
		s := p.s
		i := p.i

		nxt := right[i]
		if nums[i] > nums[nxt] {
			dec--
		}

		pre := left[i]
		if pre >= 0 {
			if nums[pre] > nums[i] {
				dec--
			}
			if nums[pre] > s {
				dec++
			}
			heap.Push(&h, pair{nums[pre] + s, pre})
		}

		nxt2 := right[nxt]
		if nxt2 < n {
			if nums[nxt] > nums[nxt2] {
				dec--
			}
			if s > nums[nxt2] {
				dec++
			}
			heap.Push(&h, pair{s + nums[nxt2], i})
		}

		nums[i] = s
		remove(nxt)
	}
	return
}

type pair struct{ s, i int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { a, b := h[i], h[j]; return a.s < b.s || a.s == b.s && a.i < b.i }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
