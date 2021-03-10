package day03

import . "Algorithms/pkg"

/*
@Time    : 2021/3/10 8:53 下午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : detectCycle.go
@Software: GoLand
*/

// 环形链表II 找出环的起点
// 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos
// 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
// 说明：不允许修改给定的链表。

// 进阶：
// 你是否可以使用 O(1) 空间解决此题？

// 示例 1：
//输入：head = [3,2,0,-4], pos = 1
//输出：返回索引为 1 的链表节点
//解释：链表中有一个环，其尾部连接到第二个节点。

// 示例 2：
//输入：head = [1,2], pos = 0
//输出：返回索引为 0 的链表节点
//解释：链表中有一个环，其尾部连接到第一个节点。

// 示例 3：
//输入：head = [1], pos = -1
//输出：返回 null
//解释：链表中没有环。

// 哈希表  时间复杂度O(n) 空间复杂度O(1)
func detectCycle(head *ListNode) *ListNode {
	has := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := has[head]; ok {
			return head
		}
		has[head] = struct{}{}
		head = head.Next
	}
	return nil
}

// 双指针
func detectCycle1(head *ListNode) *ListNode {
	// 定义快慢双指针
	slow, fast := head, head
	// 快指针到头 表示没环
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		// 如果有环 慢指针和头节点指针一起走，当两指针相遇时 相遇的节点就是环的第一个节点
		if slow == fast {
			s := head
			for slow != s {
				slow = slow.Next
				s = s.Next
			}
			return s
		}
	}
	return nil
}
