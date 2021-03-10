package day03

import . "Algorithms/pkg"

/*
@Time    : 2021/3/10 8:26 下午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : hasCycle.go
@Software: GoLand
*/

// 环形链表
//给定一个链表，判断链表中是否有环。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的
//位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
// 如果链表中存在环，则返回 true 。 否则，返回 false 。

// 进阶：
// 你能用 O(1)（即，常量）内存解决此问题吗？

// 示例 1：
// 输入：head = [3,2,0,-4], pos = 1
//输出：true
//解释：链表中有一个环，其尾部连接到第二个节点。

// 示例 2：
// 输入：head = [1,2], pos = 0
//输出：true
//解释：链表中有一个环，其尾部连接到第一个节点。

// 示例 3：
// 输入：head = [1], pos = -1
//输出：false
//解释：链表中没有环。

// 提示：
// 链表中节点的数目范围是 [0, 104]
// -105 <= Node.val <= 105
// pos 为 -1 或者链表中的一个 有效索引 。

// 哈希表
func hasCycle(head *ListNode) bool {
	has := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := has[head]; ok {
			return true
		}
		has[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 快慢指针
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		// 如果快指针到头了 则表示没有环
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}