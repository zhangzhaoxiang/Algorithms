package problems

/*
@Time    : 2021/3/8 6:48 下午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : moveZeros.go
@Software: GoLand
*/

// 移动零：给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

// 思路：双指针 使用两个指针i,j
// i表示数组当前的指向的位置 j表示数组最后一个非零的位置
// 最终结果为j位置的左边都是非零数，右边都是零

func moveZeros(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}

// 时间复杂度 O(n) 其中 n 为序列长度。每个位置至多被遍历两次。
// 空间复杂度 O(1)
