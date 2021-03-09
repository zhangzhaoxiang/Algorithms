package review

/*
@Time    : 2021/3/9 8:45 上午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : maxArea.go
@Software: GoLand
*/

// 盛水的面积最大
// 双指针

func maxArea(height []int) int {
	i, j := 0, len(height)-1 // i j 双指针分别指向头尾
	max := 0                 // 初始化最大max
	for i < j {              // 左指针在右指针左边
		area := 0
		if height[i] < height[j] { // 移动短的
			area = height[i] * (j - i) // 计算面积
			i++
		} else {
			area = height[j] * (j - i)
			j--
		}
		if max < area { // 如果当前的面积大于最大面积 更新记录
			max = area
		}
	}
	return max
}
