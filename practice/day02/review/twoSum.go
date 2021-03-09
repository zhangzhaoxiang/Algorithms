package review

/*
@Time    : 2021/3/9 9:15 上午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : twoSum.go
@Software: GoLand
*/

// 两数之和
// 使用map 记录{num : index}

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
