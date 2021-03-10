package day03

import . "Algorithms/pkg"

/*
@Time    : 2021/3/10 6:43 下午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : swapPairs.go
@Software: GoLand
*/

// 交换链表中的节点

// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// 示例 1：
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]

// 示例 2：
// 输入：head = []
// 输出：[]

// 示例 3：
// 输入：head = [1]
// 输出：[1]

// 提示：
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100

// 迭代
func swapPairs1(head *ListNode) *ListNode {
	// 存储链表的头部节点
	listNodeHead := &ListNode{
		Val:  0,
		Next: head,
	}
	// 中间节点 用于节点交换
	temp := listNodeHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return listNodeHead.Next
}

// TODO 递归
func swapPairs(head *ListNode) *ListNode {
	// 空节点或只剩一个节点 直接跳过
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head

	return newHead
}
