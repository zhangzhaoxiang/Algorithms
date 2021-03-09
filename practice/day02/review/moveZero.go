package review

/*
@Time    : 2021/3/9 9:07 上午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : moveZero.go
@Software: GoLand
*/

// 移动零
// 双指针 i j i表示当前值 j表示最后一个非零数的位置
// 结果：j指针左边是非零数 右边是0

func moveZero(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
