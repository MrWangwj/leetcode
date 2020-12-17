package main
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。 
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 
//
// 注意空字符串可被认为是有效字符串。 
//
// 示例 1: 
//
// 输入: "()"
//输出: true
// 
//
// 示例 2: 
//
// 输入: "()[]{}"
//输出: true
// 
//
// 示例 3: 
//
// 输入: "(]"
//输出: false
// 
//
// 示例 4: 
//
// 输入: "([)]"
//输出: false
// 
//
// 示例 5: 
//
// 输入: "{[]}"
//输出: true 
// Related Topics 栈 字符串 
// 👍 2043 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	s2 := make([]int32, 0, len(s)/2)
	index := 0
	m := map[int32]int32{
		'[':']',
		'(':')',
		'{':'}',
	}
	for _, c := range s {
		switch c {
		case '[', '(', '{':
			s2 = append(s2, c)
			index++
		case ']', ')', '}':
			if index-1 < 0 || m[s2[index-1]] != c {
				return false
			}
			s2 = s2[:index-1]
			index--
		default:
			return false
		}
	}
	if len(s2) > 0 {
		return false
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
