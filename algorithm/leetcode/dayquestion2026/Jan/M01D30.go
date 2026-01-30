package Jan

import "math"

/**
 * 2977. 转换字符串的最小成本 II
 * <a href="https://leetcode.cn/problems/minimum-cost-to-convert-string-ii/description/"/>
 */

func minimumCost2(source, target string, original, changed []string, cost []int) int64 {
	const inf = math.MaxInt / 2
	type node struct {
		son [26]*node
		sid int
	}
	root := &node{}
	sid := 0
	put := func(s string) int {
		o := root
		for _, b := range s {
			b -= 'a'
			if o.son[b] == nil {
				o.son[b] = &node{sid: -1}
			}
			o = o.son[b]
		}
		if o.sid < 0 {
			o.sid = sid
			sid++
		}
		return o.sid
	}

	m := len(cost)
	dis := make([][]int, m*2)
	for i := range dis {
		dis[i] = make([]int, m*2)
		for j := range dis[i] {
			if j != i {
				dis[i][j] = inf
			}
		}
	}
	for i, c := range cost {
		x := put(original[i])
		y := put(changed[i])
		dis[x][y] = min(dis[x][y], c)
	}

	for k := 0; k < sid; k++ {
		for i := 0; i < sid; i++ {
			if dis[i][k] == inf {
				continue
			}
			for j := 0; j < sid; j++ {
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}

	n := len(source)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i >= n {
			return 0
		}
		ptr := &memo[i]
		if *ptr != -1 {
			return *ptr
		}
		res := inf
		if source[i] == target[i] {
			res = dfs(i + 1)
		}
		p, q := root, root
		for j := i; j < n; j++ {
			p = p.son[source[j]-'a']
			q = q.son[target[j]-'a']
			if p == nil || q == nil {
				break
			}
			if p.sid >= 0 && q.sid >= 0 {
				res = min(res, dis[p.sid][q.sid]+dfs(j+1))
			}
		}
		*ptr = res
		return res
	}
	ans := dfs(0)
	if ans == inf {
		return -1
	}
	return int64(ans)
}
