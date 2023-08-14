package main

import (
	"math"
)

//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
//
//
//
// 示例 1:
//
//
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// 示例 2:
//
//
//输入: strs = [""]
//输出: [[""]]
//
//
// 示例 3:
//
//
//输入: strs = ["a"]
//输出: [["a"]]
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 10⁴
// 0 <= strs[i].length <= 100
// strs[i] 仅包含小写字母
//
//
// Related Topics 数组 哈希表 字符串 排序 👍 1502 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	//m := make(map[[26]rune][]string, len(strs))
	//for _, v := range strs {
	//	k := [26]rune{}
	//	for _, s := range v {
	//		k[s-'a']++
	//	}
	//	m[k] = append(m[k], v)
	//}
	//
	//r := make([][]string, 0, len(m))
	//for _, v := range m {
	//	r = append(r, v)
	//}
	//return r

	primes := []int{2, 3, 5, 7, 11, 13, 17, 19,
		23, 29, 31, 37, 41, 43, 47, 53,
		59, 61, 67, 71, 73, 79, 83, 89,
		97, 101}
	m := make(map[int][]string, len(strs))
	for _, v := range strs {
		n := 1
		for _, s := range v {
			n = n * primes[int(s-'a')] % math.MaxInt
		}
		m[n] = append(m[n], v)
	}

	r := make([][]string, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

//leetcode submit region end(Prohibit modification and deletion)
