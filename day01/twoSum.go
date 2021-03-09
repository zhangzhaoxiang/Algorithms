package problems

/*
@Time    : 2021/3/8 8:26 下午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : twoSum.go
@Software: GoLand
*/

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回它们的数组下标。

func twoSum(nums []int, target int) []int {
	result := make(map[int]int)
	for index, num := range nums {
		if i, ok := result[target-num]; ok {
			return []int{i, index}
		}
		result[num] = index
	}
	return nil
}
