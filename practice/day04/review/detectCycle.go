package review

import . "Algorithms/pkg"

/*
@Time    : 2021/3/11 9:03 上午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : detectCycle.go
@Software: GoLand
*/

// 环形链表 找出环入口

// 哈希
func detectCycle(head *ListNode) *ListNode {
	has := make(map[*ListNode]int)
	for head != nil {
		if _, ok := has[head]; ok {
			return head
		}
		has[head] = 1
		head = head.Next
	}
	return nil
}

// 快慢指针
func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			// 快慢指针相遇 证明有环
			// 此时头节点和慢指针同时走 当两者相遇时 相遇的地方就是节点入口
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			return head
		}
	}
	return nil
}
