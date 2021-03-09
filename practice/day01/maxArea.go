package problems

/*
@Time    : 2021/3/8 7:43 下午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : maxArea.go
@Software: GoLand
*/

// 盛最多水的容器：给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i,
// ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

// 思路： 双指针法 i,j i指向数组头部 j指向数组尾部 每次都移动值小的指针 移动后计算面积 记录 遍历数组一遍 找到最大面积

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		// 头部高度小于尾部 用头部高度计算面积
		area := 0
		if height[i] < height[j] {
			area = height[i] * (j - i)
			i++
		} else {
			area = height[j] * (j - i)
			j--
		}
		if max < area {
			max = area
		}
	}
	return max
}

// 时间复杂度 O(n)
// 空间复杂度 O(1)
