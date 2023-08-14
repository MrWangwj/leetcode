package main

import "math"

//编写一种方法，对字符串数组进行排序，将所有变位词组合在一起。变位词是指字母相同，但排列不同的字符串。
//
// 注意：本题相对原题稍作修改
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//
// 说明：
//
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
//
// Related Topics 数组 哈希表 字符串 排序 👍 110 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	arr := []int{2, 3, 5, 7, 11, 13, 17, 19,
		23, 29, 31, 37, 41, 43, 47, 53,
		59, 61, 67, 71, 73, 79, 83, 89,
		97, 101}
	m := make(map[int][]string, len(strs))
	for _, str := range strs {
		n := 1
		for _, s := range str {
			n = n * arr[s-'a'] % math.MaxInt
		}
		m[n] = append(m[n], str)
	}
	r := make([][]string, 0, len(m))
	for _, strings := range m {
		r = append(r, strings)
	}
	return r
}

//leetcode submit region end(Prohibit modification and deletion)
