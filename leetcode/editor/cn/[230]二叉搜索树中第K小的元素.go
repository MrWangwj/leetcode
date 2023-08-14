package main

//给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,1,4,null,2], k = 1
//输出：1
//
//
// 示例 2：
//
//
//输入：root = [5,3,6,2,4,null,null,1], k = 3
//输出：3
//
//
//
//
//
//
// 提示：
//
//
// 树中的节点数为 n 。
// 1 <= k <= n <= 10⁴
// 0 <= Node.val <= 10⁴
//
//
//
//
// 进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 756 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var m = map[*TreeNode]int{}

func kthSmallest(root *TreeNode, k int) int {
	//stack := []*TreeNode{}
	//for {
	//	for root != nil {
	//		stack = append(stack, root)
	//		root = root.Left
	//	}
	//
	//	root = stack[len(stack)-1]
	//	stack = stack[:len(stack)-1]
	//	k--
	//	if k == 0 {
	//		return root.Val
	//	}
	//	root = root.Right
	//}

	depth(root)
	for {
		l := m[root.Left]
		if l == k-1 {
			return root.Val
		}
		if l < k-1 {
			root = root.Right
			k -= l + 1
			continue
		}
		root = root.Left
	}
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m[root] = depth(root.Left) + depth(root.Right) + 1
	return m[root]
}

//leetcode submit region end(Prohibit modification and deletion)
