package main

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//
//
// 示例 2：
//
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//
//
//
//
//
// 提示：
//
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10⁶ <= nums1[i], nums2[i] <= 10⁶
//
//
// Related Topics 数组 二分查找 分治 👍 6288 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	s := len(nums1) + len(nums2)
	if s%2 == 1 {
		return float64(findK(s/2+1, nums1, nums2))
	}
	return float64(findK(s/2, nums1, nums2)+findK(s/2+1, nums1, nums2)) / 2.0
}

func findK(k int, nums1, nums2 []int) int {
	for {
		if len(nums1) == 0 {
			return nums2[k-1]
		}
		if len(nums2) == 0 {
			return nums1[k-1]
		}
		if k == 1 {
			return min(nums1[0], nums2[0])
		}

		m := k/2 - 1
		i1 := min(m, len(nums1)-1)
		i2 := min(m, len(nums2)-1)
		if nums1[i1] <= nums2[i2] {
			nums1 = nums1[i1+1:]
			k = k - (i1 + 1)
			continue
		}

		nums2 = nums2[i2+1:]
		k = k - (i2 + 1)
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
