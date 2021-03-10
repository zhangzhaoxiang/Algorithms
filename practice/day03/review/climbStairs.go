package review

/*
@Time    : 2021/3/10 8:52 上午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : climbStairs.go
@Software: GoLand
*/

// 动态规划 斐波那契数列

func climbStairs(n int) int {
	f, s := 0, 1
	for i := 0; i < n; i++ {
		f, s = s, f+s
	}
	return s
}
