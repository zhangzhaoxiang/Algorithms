package review

/*
@Time    : 2021/3/10 8:57 上午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : reverseList.go
@Software: GoLand
*/

// 链表反转
// 双指针

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 定义一个链表
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next // 将右指针的下一个节点记录下来
		cur.Next = prev  // 右指针的Next节点指向左指针
		prev = cur       // 左指针后移
		cur = next       // 右指针后移
	}
	return prev
}
