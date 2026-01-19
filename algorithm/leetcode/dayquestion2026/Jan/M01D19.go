package Jan

/**
 * 1292. 元素和小于等于阈值的正方形的最大边长
 * <a href="https://leetcode.cn/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/description/"/>
 */
func maxSideLength(mat [][]int, threshold int) (ans int) {
	n, m := len(mat), len(mat[0])
	sum := make([][]int, n+1)
	sum[0] = make([]int, m+1)
	for i, row := range mat {
		sum[i+1] = make([]int, m+1)
		for j, x := range row {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + x
		}
	}

	for i := range n {
		for j := range m {
			// 边长为 ans+1 的正方形，左上角在 (i, j)，右下角在 (i+ans, j+ans)
			for i+ans < n && j+ans < m && query(sum, i, j, i+ans, j+ans) <= threshold {
				ans++
			}
		}
	}
	return
}

// 返回左上角在 (r1, c1)，右下角在 (r2, c2) 的子矩阵元素和
func query(sum [][]int, r1, c1, r2, c2 int) int {
	return sum[r2+1][c2+1] - sum[r2+1][c1] - sum[r1][c2+1] + sum[r1][c1]
}
