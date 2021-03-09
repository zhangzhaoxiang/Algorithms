package day02

import "sort"

/*
@Time    : 2021/3/9 1:53 下午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : threeSum.go
@Software: GoLand
*/

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
// 复的三元组。
// 注意：答案中不可以包含重复的三元组。
//
// 示例 1：
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
//
// 示例 2：
// 输入：nums = []
// 输出：[]
//
// 示例 3：
// 输入：nums = [0]
// 输出：[]

// 提示：
// 0 <= nums.length <= 3000
// -105 <= nums[i] <= 105

// 使用一个数组和两个指针
// 1。将给定的数组排序
// 2。最外层有一个循环 之后拿每次循环到的值去和后面的值求和
// 3。后面的值需要两个指针 一个在头部一个在尾部 夹逼
// 4。结果大了 右指针左移 结果小了 左指针右移

func threeSum(nums []int) [][]int {
	// 空数组以及长度小于3的数组直接跳过
	var res [][]int
	length := len(nums)
	if length < 3 {
		return res
	}
	// 排序
	sort.Ints(nums)
	for i, num := range nums {
		// 当前值和前一个相同 直接跳过 避免重复
		if i > 0 && num == nums[i-1] {
			continue
		}
		// 定义左右指针
		l, r := i+1, length-1
		for l < r {
			if num+nums[l]+nums[r] < 0 {
				l++ // 整体小 左指针右移
				continue
			}
			if num+nums[l]+nums[r] > 0 {
				r-- // 整体大 右指针左移
				continue
			}
			if num+nums[l]+nums[r] == 0 {
				res = append(res, []int{num, nums[l], nums[r]}) // 相等 记录结果
				// 再起两个loop 分别过滤掉左右两指针中重复的值
				for l < r {
					if nums[l] == nums[l+1] {
						l++
					} else {
						break
					}
				}
				for l < r {
					if nums[r] == nums[r-1] {
						r--
					} else {
						break
					}
				}
				// 左右指针移位
				l++
				r--
			}
		}
	}
	return res
}
