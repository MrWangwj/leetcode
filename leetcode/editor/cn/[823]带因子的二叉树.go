package main

import "sort"

//给出一个含有不重复整数元素的数组 arr ，每个整数 arr[i] 均大于 1。
//
// 用这些整数来构建二叉树，每个整数可以使用任意次数。其中：每个非叶结点的值应等于它的两个子结点的值的乘积。
//
// 满足条件的二叉树一共有多少个？答案可能很大，返回 对 10⁹ + 7 取余 的结果。
//
//
//
// 示例 1:
//
//
//输入: arr = [2, 4]
//输出: 3
//解释: 可以得到这些二叉树: [2], [4], [4, 2, 2]
//
// 示例 2:
//
//
//输入: arr = [2, 4, 5, 10]
//输出: 7
//解释: 可以得到这些二叉树: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 1000
// 2 <= arr[i] <= 10⁹
// arr 中的所有值 互不相同
//
//
// Related Topics 数组 哈希表 动态规划 排序 👍 201 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	res, mod := int64(0), int64(1e9+7)
	dp := make([]int64, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for l, r := 0, i-1; l <= r; l++ {
			for l <= r && int64(arr[l]*arr[r]) > int64(arr[i]) {
				r--
			}

			if l <= r && int64(arr[l]*arr[r]) == int64(arr[i]) {
				if arr[l] == arr[r] {
					dp[i] = (dp[i] + dp[l]*dp[r]) % mod
				} else {
					dp[i] = (dp[i] + dp[l]*dp[r]*2) % mod
				}
			}
		}
		res = (res + dp[i]) % mod
	}
	return int(res)
}

//leetcode submit region end(Prohibit modification and deletion)
