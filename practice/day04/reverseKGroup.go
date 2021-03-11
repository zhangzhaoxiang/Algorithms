package day04

import . "Algorithms/pkg"

/*
@Time    : 2021/3/11 9:51 上午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : reverseKGroup.go
@Software: GoLand
*/

// k个一组反转链表

// 迭代
// 先找到需要的k个长度的链表的头尾节点
// 将找到的链表反转

func reverseKGroup(head *ListNode, k int) *ListNode {
	newHead := &ListNode{Next: head} // 记录链表的开始节点

	prev := head
	for head != nil {
		end := prev
		for i := 0; i < k; i++ {
			end = end.Next
			if end == nil {
				return newHead.Next
			}
		}
	}
	return nil
}
