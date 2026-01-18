package Jan

/**
 * 1895. 最大的幻方
 * <a href="https://leetcode.cn/problems/largest-magic-square/description/"/>
 */
func largestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rowSum := make([][]int, m)    // → 前缀和
	colSum := make([][]int, m+1)  // ↓ 前缀和
	diagSum := make([][]int, m+1) // ↘ 前缀和
	antiSum := make([][]int, m+1) // ↙ 前缀和
	for i := range m + 1 {
		colSum[i] = make([]int, n)
		diagSum[i] = make([]int, n+1)
		antiSum[i] = make([]int, n+1)
	}

	for i, row := range grid {
		rowSum[i] = make([]int, n+1)
		for j, x := range row {
			rowSum[i][j+1] = rowSum[i][j] + x
			colSum[i+1][j] = colSum[i][j] + x
			diagSum[i+1][j+1] = diagSum[i][j] + x
			antiSum[i+1][j] = antiSum[i][j+1] + x
		}
	}

	// k×k 子矩阵的左上角为 (i−k, j−k)，右下角为 (i−1, j−1)
	for k := min(m, n); ; k-- {
		for i := k; i <= m; i++ {
		next:
			for j := k; j <= n; j++ {
				// 子矩阵主对角线的和
				sum := diagSum[i][j] - diagSum[i-k][j-k]

				// 子矩阵反对角线的和
				if antiSum[i][j-k]-antiSum[i-k][j] != sum {
					continue
				}

				// 子矩阵每行的和
				for _, rowS := range rowSum[i-k : i] {
					if rowS[j]-rowS[j-k] != sum {
						continue next
					}
				}

				// 子矩阵每列的和
				for c := j - k; c < j; c++ {
					if colSum[i][c]-colSum[i-k][c] != sum {
						continue next
					}
				}

				return k
			}
		}
	}
}
