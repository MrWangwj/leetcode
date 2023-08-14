package main
//给你两个整数 left 和 right ，在闭区间 [left, right] 范围内，统计并返回 计算置位位数为质数 的整数个数。 
//
// 计算置位位数 就是二进制表示中 1 的个数。 
//
// 
// 例如， 21 的二进制表示 10101 有 3 个计算置位。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：left = 6, right = 10
//输出：4
//解释：
//6 -> 110 (2 个计算置位，2 是质数)
//7 -> 111 (3 个计算置位，3 是质数)
//9 -> 1001 (2 个计算置位，2 是质数)
//10-> 1010 (2 个计算置位，2 是质数)
//共计 4 个计算置位为质数的数字。
// 
//
// 示例 2： 
//
// 
//输入：left = 10, right = 15
//输出：5
//解释：
//10 -> 1010 (2 个计算置位, 2 是质数)
//11 -> 1011 (3 个计算置位, 3 是质数)
//12 -> 1100 (2 个计算置位, 2 是质数)
//13 -> 1101 (3 个计算置位, 3 是质数)
//14 -> 1110 (3 个计算置位, 3 是质数)
//15 -> 1111 (4 个计算置位, 4 不是质数)
//共计 5 个计算置位为质数的数字。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= left <= right <= 10⁶ 
// 0 <= right - left <= 10⁴ 
// 
// Related Topics 位运算 数学 👍 77 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func binaryCount(num int) (c int) {
	for ;num > 0; num &= num-1 {
		c++
	}
	return
}

func isPrimes (num int) bool {
	//不超过20的质数 2,3,5,7,11,13,17,19
	//665772=10100010100010101100
	if 1<<num&665772 != 0 {
		return true
	}
	return false
}

func countPrimeSetBits(left int, right int) (n int) {
	for i := left; i <= right; i++ {
		if isPrimes(binaryCount(i)) {
			n++
		}
	}
	return n;
}
//leetcode submit region end(Prohibit modification and deletion)
