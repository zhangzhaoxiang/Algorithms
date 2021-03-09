package pkg

/*
@Time    : 2021/3/9 4:57 下午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : listNode.go
@Software: GoLand
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListConvertListNode(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}
	node := &ListNode{}
	// 存储链表头位置
	res := node
	for _, num := range array {
		temp := &ListNode{}
		temp.Val = num
		node.Next = temp
		node = temp
	}
	return res.Next
}

func ListNodeConvertList(node *ListNode) []int {
	var myArray []int
	for node != nil {
		myArray = append(myArray, node.Val)
		node = node.Next
	}
	return myArray
}
