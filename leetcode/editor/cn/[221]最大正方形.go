package main

//在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
//
//
//
// 示例 1：
//
//
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
//,["1","0","0","1","0"]]
//输出：4
//
//
// 示例 2：
//
//
//输入：matrix = [["0","1"],["1","0"]]
//输出：1
//
//
// 示例 3：
//
//
//输入：matrix = [["0"]]
//输出：0
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 300
// matrix[i][j] 为 '0' 或 '1'
//
//
// Related Topics 数组 动态规划 矩阵 👍 1503 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func maximalSquare(matrix [][]byte) int {
	res := 0
	if matrix == nil || len(matrix) == 0 {
		return res
	}
	n, m := len(matrix), len(matrix[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				res = 1
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if dp[i][j] == 0 {
				continue
			}
			dp[i][j] = min221(min221(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			res = max221(res, dp[i][j])
		}
	}

	return res * res
}

func max221(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min221(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
