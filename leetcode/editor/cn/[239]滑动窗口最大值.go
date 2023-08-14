package main

import (
	"container/heap"
	"sort"
)

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。
//
// 返回 滑动窗口中的最大值 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
//输入：nums = [1], k = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// 1 <= k <= nums.length
//
//
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 2413 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
var a []int

type hp2 struct {
	sort.IntSlice
}

func (h hp2) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}
func (h *hp2) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp2) Pop() any {
	v := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return v
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}
func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)
	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans

	//max, l := 0, 0
	//for i := 0; i < k; i++ {
	//	if nums[i] > max {
	//		max = nums[i]
	//		l = i
	//	}
	//}
	//res := make([]int, 0, len(nums)-k)
	//res = append(res, max)
	//for i := k; i < len(nums); i++ {
	//	if nums[i] > max {
	//		res = append(res, nums[i])
	//		max = nums[i]
	//		l = i
	//		continue
	//	}
	//
	//	if i-k+1 <= l {
	//		res = append(res, max)
	//		continue
	//	}
	//	max = nums[i]
	//	for j := i - k + 1; j <= i; j++ {
	//		if nums[j] > max {
	//			max = nums[j]
	//			l = j
	//		}
	//	}
	//	res = append(res, max)
	//}
	//return res
}

//leetcode submit region end(Prohibit modification and deletion)
