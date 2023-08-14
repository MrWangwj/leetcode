package main

//给你一个下标从 0 开始、大小为 n x n 的整数矩阵 grid ，返回满足 Ri 行和 Cj 列相等的行列对 (Ri, Cj) 的数目。
//
// 如果行和列以相同的顺序包含相同的元素（即相等的数组），则认为二者是相等的。
//
//
//
// 示例 1：
//
//
//
//
//输入：grid = [[3,2,1],[1,7,6],[2,7,7]]
//输出：1
//解释：存在一对相等行列对：
//- (第 2 行，第 1 列)：[2,7,7]
//
//
// 示例 2：
//
//
//
//
//输入：grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
//输出：3
//解释：存在三对相等行列对：
//- (第 0 行，第 0 列)：[3,1,2,2]
//- (第 2 行, 第 2 列)：[2,4,2,2]
//- (第 3 行, 第 2 列)：[2,4,2,2]
//
//
//
//
// 提示：
//
//
// n == grid.length == grid[i].length
// 1 <= n <= 200
// 1 <= grid[i][j] <= 10⁵
//
//
// Related Topics 数组 哈希表 矩阵 模拟 👍 31 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	Nums  int
	Nodes map[int]*Trie
}

func equalPairs(grid [][]int) int {
	//root := &Trie{Nodes: map[int]*Trie{}}
	//n := len(grid)
	//for i := 0; i < n; i++ {
	//	tmp := root
	//	for j := 0; j < n; j++ {
	//		n := grid[i][j]
	//		if _, ok := tmp.Nodes[n]; !ok {
	//			tmp.Nodes[n] = &Trie{Nodes: map[int]*Trie{}}
	//		}
	//		tmp = tmp.Nodes[n]
	//	}
	//	tmp.Nums++
	//}
	//
	//result := 0
	//for i := 0; i < n; i++ {
	//	tmp := root
	//	for j := 0; j < n; j++ {
	//		n := grid[j][i]
	//		if _, ok := tmp.Nodes[n]; !ok {
	//			break
	//		}
	//		tmp = tmp.Nodes[n]
	//	}
	//	result += tmp.Nums
	//}
	//return result

	n := len(grid)
	m := make(map[[200]int]int, n)
	for i := 0; i < n; i++ {
		arr := [200]int{}
		for j := 0; j < n; j++ {
			arr[j] = grid[j][i]
		}
		m[arr]++
	}

	result := 0
	for i := 0; i < n; i++ {
		arr := [200]int{}
		for j := 0; j < n; j++ {
			arr[j] = grid[i][j]
		}
		result += m[arr]
	}
	return result

	//help := map[[200]int]int{}
	//toKey := func(in []int) [200]int {
	//	temp := [200]int{}
	//	for k1, v1 := range in {
	//		temp[k1] = v1
	//	}
	//	return temp
	//}
	//for _, v := range grid {
	//	help[toKey(v)] += 1
	//}
	//n := len(grid)
	//ans := 0
	//for i := 0; i < n; i++ {
	//	temp := make([]int, n)
	//	for j := 0; j < n; j++ {
	//		temp[j] = grid[j][i]
	//	}
	//	ans += help[toKey(temp)]
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
