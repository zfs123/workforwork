package leetcode

/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	p, q := head, head
	for q != nil {
		p = p.Next
		if q.Next != nil {
			q = q.Next.Next
		} else {
			return false
		}
		if p == q {
			return true
		}
	}
	return false
}

// @lc code=end
