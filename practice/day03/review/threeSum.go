package review

import "sort"

/*
@Time    : 2021/3/10 8:04 上午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : threeSum.go
@Software: GoLand
*/

// 三数之和等于零
// 固定一个数字 其他两个数字双指针夹逼

func threeSum(nums []int) [][]int {
	// 异常判断 小于三个
	if len(nums) < 3 {
		return nil
	}
	// 排序
	sort.Ints(nums)
	var res [][]int
	length := len(nums)
	for i := 0; i < length-2; i++ {
		// 去重 和前一个相等 直接跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 如果前三个都大于零 证明后面的三数之和一定大于零 直接返回
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			return res
		}
		// 当前数与最后两个和小于零 直接进行下一轮循环
		if nums[i]+nums[length-1]+nums[length-2] < 0 {
			continue
		}
		// 定义两指针
		l := i + 1
		r := length - 1

		for l < r {
			// 求和
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++ // 和小于零 左指针右移
			} else if sum > 0 {
				r-- // 和大于零 右指针左移
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 去除r l重复
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				// 此时左右指针指向还是刚才的值 所以要各自移动一位
				l++
				r--
			}
		}
	}
	return res
}
