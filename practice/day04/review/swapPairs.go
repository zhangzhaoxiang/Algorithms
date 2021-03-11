package review

import . "Algorithms/pkg"

/*
@Time    : 2021/3/11 8:39 上午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : swapPairs.go
@Software: GoLand
*/

// 链表反转

// 递归
func swapPairs(head *ListNode) *ListNode {
	// 单个节点或空节点不用反转
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

// 迭代

func swapPairs1(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head} // 记录头节点位置
	temp := newHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return newHead.Next

}
