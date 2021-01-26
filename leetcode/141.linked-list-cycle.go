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
	Val int
	Next *ListNode
}
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
		return false
	}
	var p, q = head, head.Next.Next
	for q != nil {
		if p == q {
			return true
		}
		p = p.Next
		q = q.Next
		if q == nil {
			return false
		}
		q = q.Next
	}
	return false
}
// @lc code=end

