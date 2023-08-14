package main

//给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
//
//
// 示例1：
//
//
//
//
//输入：l1 = [7,2,4,3], l2 = [5,6,4]
//输出：[7,8,0,7]
//
//
// 示例2：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[8,0,7]
//
//
// 示例3：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 链表的长度范围为 [1, 100]
// 0 <= node.val <= 9
// 输入数据保证链表代表的数字无前导 0
//
//
//
//
// 进阶：如果输入链表不能翻转该如何解决？
//
// Related Topics 栈 链表 数学 👍 633 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	rl1 := reList(l1)
	rl2 := reList(l2)
	var res *ListNode
	n := 0
	for {
		if rl1 == nil && rl2 == nil && n != 0 {
			break
		}
		n1, n2 := 0, 0
		if rl1 != nil {
			n1 = rl1.Val
			rl1 = rl1.Next
		}
		if rl2 != nil {
			n2 = rl2.Val
			rl2 = rl2.Next
		}
		s := n1 + n2 + n
		res = &ListNode{
			Val:  s % 10,
			Next: res,
		}
		n = s / 10
	}
	return res
}

func reList(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	var head *ListNode

	for l != nil {
		next := l.Next
		l.Next = head
		head = l
		l = next
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
