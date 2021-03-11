package review

import . "Algorithms/pkg"

/*
@Time    : 2021/3/11 8:24 上午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : hasCycle.go
@Software: GoLand
*/

// 链表中存在环

// 迭代 定义一个map 把所有路过的节点都记录下来 再次路过表示有环
func hasCycle(head *ListNode) bool {
	has := make(map[*ListNode]int)
	for head != nil {
		if _, ok := has[head]; ok {
			return true
		}
		has[head] = 1
		head = head.Next
	}
	return false
}

// 快慢指针
// 定义快慢指针 快指针一轮走两步 慢指针一轮走一步 如果两指针再次相遇 证明有环
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
