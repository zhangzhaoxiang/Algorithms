package day02

import . "Algorithms/pkg"

/*
@Time    : 2021/3/9 4:20 下午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : reverseList.go
@Software: GoLand
*/

//反转一个单链表。

// 示例:
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//
// 进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

// 思路 双指针

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // 定义一个新链表 nil
	cur := head        // cur指向给定的链表头
	for cur != nil {
		temp := cur.Next // 中间temp 存储cur的下一个节点位置
		cur.Next = prev  // 将cur下一个节点指向prev
		prev = cur       // prev前移
		cur = temp       // cur前移
	}
	return prev
}
