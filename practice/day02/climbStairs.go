package day02

/*
@Time    : 2021/3/9 9:42 上午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : climbStairs.go
@Software: GoLand
*/

// 爬楼梯：
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 注意：给定 n 是一个正整数。

// 思路：
// 1。爬一级台阶  1种 （1个）
// 2。二级  2种 （1个+1个 2个）
// 3。三级  3种  (1个+1个+1个 2个+1个 1个+2个)   相当于 stairs(2) + stairs(1)
// n级  stairs(n-1)+ stairs(n-2)
// 动态规划

func climbStairs(n int) int {
	f, l := 0, 1
	for i := 0; i < n; i++ {
		f, l = l, f+l
	}
	return l
}
