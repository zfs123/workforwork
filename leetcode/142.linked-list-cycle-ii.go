package leetcode

/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	p, q := head, head
	var cycle bool
	for q != nil {
		p = p.Next
		if q.Next != nil {
			q = q.Next.Next
		} else {
			return nil
		}
		if p == q {
			cycle = true
			break
		}
	}
	if !cycle {
		return nil
	}

	p = head
	for p != q {
		p = p.Next
		q = q.Next
	}
	return p
}

// @lc code=end
